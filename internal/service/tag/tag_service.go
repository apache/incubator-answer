package tag

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/service/activity_queue"
	"github.com/answerdev/answer/internal/service/revision_common"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/answerdev/answer/internal/service/tag_common"
	"github.com/answerdev/answer/pkg/htmltext"

	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/permission"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type TagRepo interface {
	RemoveTag(ctx context.Context, tagID string) (err error)
	UpdateTag(ctx context.Context, tag *entity.Tag) (err error)
	UpdateTagSynonym(ctx context.Context, tagSlugNameList []string, mainTagID int64, mainTagSlugName string) (err error)
	GetTagList(ctx context.Context, tag *entity.Tag) (tagList []*entity.Tag, err error)
}

// TagService user service
type TagService struct {
	tagRepo          TagRepo
	tagCommonService *tag_common.TagCommonService
	revisionService  *revision_common.RevisionService
	followCommon     activity_common.FollowRepo
	siteInfoService  *siteinfo_common.SiteInfoCommonService
}

// NewTagService new tag service
func NewTagService(
	tagRepo TagRepo,
	tagCommonService *tag_common.TagCommonService,
	revisionService *revision_common.RevisionService,
	followCommon activity_common.FollowRepo,
	siteInfoService *siteinfo_common.SiteInfoCommonService) *TagService {
	return &TagService{
		tagRepo:          tagRepo,
		tagCommonService: tagCommonService,
		revisionService:  revisionService,
		followCommon:     followCommon,
		siteInfoService:  siteInfoService,
	}
}

// RemoveTag delete tag
func (ts *TagService) RemoveTag(ctx context.Context, tagID string) (err error) {
	// TODO permission

	err = ts.tagRepo.RemoveTag(ctx, tagID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTag update tag
func (ts *TagService) UpdateTag(ctx context.Context, req *schema.UpdateTagReq) (err error) {
	tag := &entity.Tag{}
	_ = copier.Copy(tag, req)
	tag.ID = req.TagID
	err = ts.tagRepo.UpdateTag(ctx, tag)
	if err != nil {
		return err
	}

	tagInfo, exist, err := ts.tagCommonService.GetTagByID(ctx, req.TagID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.TagNotFound)
	}
	if tagInfo.MainTagID == 0 && len(req.SlugName) > 0 {
		log.Debugf("tag %s update slug_name", tagInfo.SlugName)
		tagList, err := ts.tagRepo.GetTagList(ctx, &entity.Tag{MainTagID: converter.StringToInt64(tagInfo.ID)})
		if err != nil {
			return err
		}
		updateTagSlugNames := make([]string, 0)
		for _, tag := range tagList {
			updateTagSlugNames = append(updateTagSlugNames, tag.SlugName)
		}
		err = ts.tagRepo.UpdateTagSynonym(ctx, updateTagSlugNames, converter.StringToInt64(tagInfo.ID), tagInfo.MainTagSlugName)
		if err != nil {
			return err
		}
	}

	revisionDTO := &schema.AddRevisionDTO{
		UserID:   req.UserID,
		ObjectID: tag.ID,
		Title:    tag.SlugName,
		Log:      req.EditSummary,
	}
	tagInfoJson, _ := json.Marshal(tagInfo)
	revisionDTO.Content = string(tagInfoJson)
	revisionID, err := ts.revisionService.AddRevision(ctx, revisionDTO, true)
	if err != nil {
		return err
	}
	activity_queue.AddActivity(&schema.ActivityMsg{
		UserID:           req.UserID,
		ObjectID:         tag.ID,
		OriginalObjectID: tag.ID,
		ActivityTypeKey:  constant.ActTagEdited,
		RevisionID:       revisionID,
	})
	return
}

// GetTagInfo get tag one
func (ts *TagService) GetTagInfo(ctx context.Context, req *schema.GetTagInfoReq) (resp *schema.GetTagResp, err error) {
	var (
		tagInfo *entity.Tag
		exist   bool
	)
	if len(req.ID) > 0 {
		tagInfo, exist, err = ts.tagCommonService.GetTagByID(ctx, req.ID)
	} else {
		tagInfo, exist, err = ts.tagCommonService.GetTagBySlugName(ctx, req.Name)
	}
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.BadRequest(reason.TagNotFound)
	}

	resp = &schema.GetTagResp{}
	// if tag is synonyms get original tag info
	if tagInfo.MainTagID > 0 {
		tagInfo, exist, err = ts.tagCommonService.GetTagByID(ctx, converter.IntToString(tagInfo.MainTagID))
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, errors.BadRequest(reason.TagNotFound)
		}
		resp.MainTagSlugName = tagInfo.SlugName
	}
	resp.TagID = tagInfo.ID
	resp.CreatedAt = tagInfo.CreatedAt.Unix()
	resp.UpdatedAt = tagInfo.UpdatedAt.Unix()
	resp.SlugName = tagInfo.SlugName
	resp.DisplayName = tagInfo.DisplayName
	resp.OriginalText = tagInfo.OriginalText
	resp.ParsedText = tagInfo.ParsedText
	resp.FollowCount = tagInfo.FollowCount
	resp.QuestionCount = tagInfo.QuestionCount
	resp.Recommend = tagInfo.Recommend
	resp.Reserved = tagInfo.Reserved
	resp.IsFollower = ts.checkTagIsFollow(ctx, req.UserID, tagInfo.ID)
	resp.MemberActions = permission.GetTagPermission(ctx, req.UserID, req.UserID)
	resp.GetExcerpt()
	return resp, nil
}

// GetFollowingTags get following tags
func (ts *TagService) GetFollowingTags(ctx context.Context, userID string) (
	resp []*schema.GetFollowingTagsResp, err error) {
	resp = make([]*schema.GetFollowingTagsResp, 0)
	if len(userID) == 0 {
		return resp, nil
	}
	objIDs, err := ts.followCommon.GetFollowIDs(ctx, userID, entity.Tag{}.TableName())
	if err != nil {
		return nil, err
	}
	tagList, err := ts.tagCommonService.GetTagListByIDs(ctx, objIDs)
	if err != nil {
		return nil, err
	}
	for _, t := range tagList {
		tagInfo := &schema.GetFollowingTagsResp{
			TagID:       t.ID,
			SlugName:    t.SlugName,
			DisplayName: t.DisplayName,
			Recommend:   t.Recommend,
			Reserved:    t.Reserved,
		}
		if t.MainTagID > 0 {
			mainTag, exist, err := ts.tagCommonService.GetTagByID(ctx, converter.IntToString(t.MainTagID))
			if err != nil {
				return nil, err
			}
			if exist {
				tagInfo.MainTagSlugName = mainTag.SlugName
			}
		}
		resp = append(resp, tagInfo)
	}
	return resp, nil
}

// GetTagSynonyms get tag synonyms
func (ts *TagService) GetTagSynonyms(ctx context.Context, req *schema.GetTagSynonymsReq) (
	resp []*schema.GetTagSynonymsResp, err error) {
	tag, exist, err := ts.tagCommonService.GetTagByID(ctx, req.TagID)
	if err != nil {
		return
	}
	if !exist {
		return nil, errors.BadRequest(reason.TagNotFound)
	}

	var tagList []*entity.Tag
	var mainTagSlugName string
	if tag.MainTagID > 0 {
		tagList, err = ts.tagRepo.GetTagList(ctx, &entity.Tag{MainTagID: tag.MainTagID})
	} else {
		tagList, err = ts.tagRepo.GetTagList(ctx, &entity.Tag{MainTagID: converter.StringToInt64(tag.ID)})
	}
	if err != nil {
		return
	}

	// get main tag slug name
	if tag.MainTagID > 0 {
		for _, tagInfo := range tagList {
			if tag.MainTagID == 0 {
				mainTagSlugName = tagInfo.SlugName
				break
			}
		}
	} else {
		mainTagSlugName = tag.SlugName
	}

	resp = make([]*schema.GetTagSynonymsResp, 0)
	for _, t := range tagList {
		resp = append(resp, &schema.GetTagSynonymsResp{
			TagID:           t.ID,
			SlugName:        t.SlugName,
			DisplayName:     t.DisplayName,
			MainTagSlugName: mainTagSlugName,
		})
	}
	return
}

// UpdateTagSynonym add tag synonym
func (ts *TagService) UpdateTagSynonym(ctx context.Context, req *schema.UpdateTagSynonymReq) (err error) {
	// format tag slug name
	req.Format()
	addSynonymTagList := make([]string, 0)
	removeSynonymTagList := make([]string, 0)
	mainTagInfo, exist, err := ts.tagCommonService.GetTagByID(ctx, req.TagID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.TagNotFound)
	}

	// find all exist tag
	for _, item := range req.SynonymTagList {
		addSynonymTagList = append(addSynonymTagList, item.SlugName)
	}
	tagListInDB, err := ts.tagCommonService.GetTagListByNames(ctx, addSynonymTagList)
	if err != nil {
		return err
	}
	existTagMapping := make(map[string]*entity.Tag, 0)
	for _, tag := range tagListInDB {
		existTagMapping[tag.SlugName] = tag
	}

	// add tag list
	needAddTagList := make([]*entity.Tag, 0)
	for _, tag := range req.SynonymTagList {
		if existTagMapping[tag.SlugName] != nil {
			continue
		}
		item := &entity.Tag{}
		item.SlugName = tag.SlugName
		item.DisplayName = tag.DisplayName
		item.OriginalText = tag.OriginalText
		item.ParsedText = tag.ParsedText
		item.Status = entity.TagStatusAvailable
		needAddTagList = append(needAddTagList, item)
	}

	if len(needAddTagList) > 0 {
		err = ts.tagCommonService.AddTagList(ctx, needAddTagList)
		if err != nil {
			return err
		}
		// update tag revision
		for _, tag := range needAddTagList {
			existTagMapping[tag.SlugName] = tag
			revisionDTO := &schema.AddRevisionDTO{
				UserID:   req.UserID,
				ObjectID: tag.ID,
				Title:    tag.SlugName,
			}
			tagInfoJson, _ := json.Marshal(tag)
			revisionDTO.Content = string(tagInfoJson)
			revisionID, err := ts.revisionService.AddRevision(ctx, revisionDTO, true)
			if err != nil {
				return err
			}
			activity_queue.AddActivity(&schema.ActivityMsg{
				UserID:           req.UserID,
				ObjectID:         tag.ID,
				OriginalObjectID: tag.ID,
				ActivityTypeKey:  constant.ActTagCreated,
				RevisionID:       revisionID,
			})
		}
	}

	// get all old synonyms list
	oldSynonymList, err := ts.tagRepo.GetTagList(ctx, &entity.Tag{MainTagID: converter.StringToInt64(mainTagInfo.ID)})
	if err != nil {
		return err
	}
	for _, oldSynonym := range oldSynonymList {
		if existTagMapping[oldSynonym.SlugName] == nil {
			removeSynonymTagList = append(removeSynonymTagList, oldSynonym.SlugName)
		}
	}

	// remove old synonyms
	if len(removeSynonymTagList) > 0 {
		err = ts.tagRepo.UpdateTagSynonym(ctx, removeSynonymTagList, 0, "")
		if err != nil {
			return err
		}
	}

	// update new synonyms
	if len(addSynonymTagList) > 0 {
		err = ts.tagRepo.UpdateTagSynonym(ctx, addSynonymTagList, converter.StringToInt64(req.TagID), mainTagInfo.SlugName)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetTagWithPage get tag list page
func (ts *TagService) GetTagWithPage(ctx context.Context, req *schema.GetTagWithPageReq) (pageModel *pager.PageModel, err error) {
	tag := &entity.Tag{}
	_ = copier.Copy(tag, req)

	page := req.Page
	pageSize := req.PageSize

	tags, total, err := ts.tagCommonService.GetTagPage(ctx, page, pageSize, tag, req.QueryCond)
	if err != nil {
		return
	}

	resp := make([]*schema.GetTagPageResp, 0)
	for _, tag := range tags {
		excerpt := htmltext.FetchExcerpt(tag.ParsedText, "...", 240)
		resp = append(resp, &schema.GetTagPageResp{
			TagID:         tag.ID,
			SlugName:      tag.SlugName,
			DisplayName:   tag.DisplayName,
			OriginalText:  excerpt,
			ParsedText:    excerpt,
			FollowCount:   tag.FollowCount,
			QuestionCount: tag.QuestionCount,
			IsFollower:    ts.checkTagIsFollow(ctx, req.UserID, tag.ID),
			CreatedAt:     tag.CreatedAt.Unix(),
			UpdatedAt:     tag.UpdatedAt.Unix(),
			Recommend:     tag.Recommend,
			Reserved:      tag.Reserved,
		})
	}
	return pager.NewPageModel(total, resp), nil
}

// checkTagIsFollow get tag list page
func (ts *TagService) checkTagIsFollow(ctx context.Context, userID, tagID string) bool {
	if len(userID) == 0 {
		return false
	}
	followed, err := ts.followCommon.IsFollowed(userID, tagID)
	if err != nil {
		log.Error(err)
	}
	return followed
}
