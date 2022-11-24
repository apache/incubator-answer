package tag_common

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/validator"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_queue"
	"github.com/answerdev/answer/internal/service/revision_common"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type TagCommonRepo interface {
	AddTagList(ctx context.Context, tagList []*entity.Tag) (err error)
	GetTagListByIDs(ctx context.Context, ids []string) (tagList []*entity.Tag, err error)
	GetTagBySlugName(ctx context.Context, slugName string) (tagInfo *entity.Tag, exist bool, err error)
	GetTagListByName(ctx context.Context, name string, limit int, hasReserved bool) (tagList []*entity.Tag, err error)
	GetTagListByNames(ctx context.Context, names []string) (tagList []*entity.Tag, err error)
	GetTagByID(ctx context.Context, tagID string, includeDeleted bool) (tag *entity.Tag, exist bool, err error)
	GetTagPage(ctx context.Context, page, pageSize int, tag *entity.Tag, queryCond string) (tagList []*entity.Tag, total int64, err error)
	GetRecommendTagList(ctx context.Context) (tagList []*entity.Tag, err error)
	GetReservedTagList(ctx context.Context) (tagList []*entity.Tag, err error)
	UpdateTagsAttribute(ctx context.Context, tags []string, attribute string, value bool) (err error)
	UpdateTagQuestionCount(ctx context.Context, tagID string, questionCount int) (err error)
}

type TagRepo interface {
	RemoveTag(ctx context.Context, tagID string) (err error)
	UpdateTag(ctx context.Context, tag *entity.Tag) (err error)
	UpdateTagSynonym(ctx context.Context, tagSlugNameList []string, mainTagID int64, mainTagSlugName string) (err error)
	GetTagList(ctx context.Context, tag *entity.Tag) (tagList []*entity.Tag, err error)
}

type TagRelRepo interface {
	AddTagRelList(ctx context.Context, tagList []*entity.TagRel) (err error)
	RemoveTagRelListByIDs(ctx context.Context, ids []int64) (err error)
	EnableTagRelByIDs(ctx context.Context, ids []int64) (err error)
	GetObjectTagRelWithoutStatus(ctx context.Context, objectId, tagID string) (tagRel *entity.TagRel, exist bool, err error)
	GetObjectTagRelList(ctx context.Context, objectId string) (tagListList []*entity.TagRel, err error)
	BatchGetObjectTagRelList(ctx context.Context, objectIds []string) (tagListList []*entity.TagRel, err error)
	CountTagRelByTagID(ctx context.Context, tagID string) (count int64, err error)
}

// TagCommonService user service
type TagCommonService struct {
	revisionService *revision_common.RevisionService
	tagCommonRepo   TagCommonRepo
	tagRelRepo      TagRelRepo
	tagRepo         TagRepo
	siteInfoService *siteinfo_common.SiteInfoCommonService
}

// NewTagCommonService new tag service
func NewTagCommonService(
	tagCommonRepo TagCommonRepo,
	tagRelRepo TagRelRepo,
	tagRepo TagRepo,
	revisionService *revision_common.RevisionService,
	siteInfoService *siteinfo_common.SiteInfoCommonService,
) *TagCommonService {
	return &TagCommonService{
		tagCommonRepo:   tagCommonRepo,
		tagRelRepo:      tagRelRepo,
		tagRepo:         tagRepo,
		revisionService: revisionService,
		siteInfoService: siteInfoService,
	}
}

// SearchTagLike get tag list all
func (ts *TagCommonService) SearchTagLike(ctx context.Context, req *schema.SearchTagLikeReq) (resp []schema.SearchTagLikeResp, err error) {
	tags, err := ts.tagCommonRepo.GetTagListByName(ctx, req.Tag, 5, req.IsAdmin)
	if err != nil {
		return
	}
	ts.TagsFormatRecommendAndReserved(ctx, tags)
	for _, tag := range tags {
		item := schema.SearchTagLikeResp{}
		item.SlugName = tag.SlugName
		item.Recommend = tag.Recommend
		item.Reserved = tag.Reserved
		resp = append(resp, item)
	}
	return resp, nil
}

func (ts *TagCommonService) GetSiteWriteRecommendTag(ctx context.Context) (tags []string, err error) {
	tags = make([]string, 0)
	list, err := ts.tagCommonRepo.GetRecommendTagList(ctx)
	if err != nil {
		return tags, err
	}
	for _, item := range list {
		tags = append(tags, item.SlugName)
	}
	return tags, nil
}

func (ts *TagCommonService) SetSiteWriteTag(ctx context.Context, recommendTags, reservedTags []string, userID string) (
	errFields []*validator.FormErrorField, err error) {
	recommendErr := ts.CheckTag(ctx, recommendTags, userID)
	reservedErr := ts.CheckTag(ctx, reservedTags, userID)
	if recommendErr != nil {
		errFields = append(errFields, &validator.FormErrorField{
			ErrorField: "recommend_tags",
			ErrorMsg:   recommendErr.Error(),
		})
		err = recommendErr
	}
	if reservedErr != nil {
		errFields = append(errFields, &validator.FormErrorField{
			ErrorField: "reserved_tags",
			ErrorMsg:   reservedErr.Error(),
		})
		err = reservedErr
	}
	if len(errFields) > 0 {
		return errFields, err
	}

	err = ts.SetTagsAttribute(ctx, recommendTags, "recommend")
	if err != nil {
		return nil, err
	}
	err = ts.SetTagsAttribute(ctx, reservedTags, "reserved")
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (ts *TagCommonService) GetSiteWriteReservedTag(ctx context.Context) (tags []string, err error) {
	tags = make([]string, 0)
	list, err := ts.tagCommonRepo.GetReservedTagList(ctx)
	if err != nil {
		return tags, err
	}
	for _, item := range list {
		tags = append(tags, item.SlugName)
	}
	return tags, nil
}

// SetTagsAttribute
func (ts *TagCommonService) SetTagsAttribute(ctx context.Context, tags []string, attribute string) (err error) {
	var tagslist []string
	switch attribute {
	case "recommend":
		tagslist, err = ts.GetSiteWriteRecommendTag(ctx)
	case "reserved":
		tagslist, err = ts.GetSiteWriteReservedTag(ctx)
	default:
		return
	}
	err = ts.tagCommonRepo.UpdateTagsAttribute(ctx, tagslist, attribute, false)
	if err != nil {
		return err
	}
	err = ts.tagCommonRepo.UpdateTagsAttribute(ctx, tags, attribute, true)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TagCommonService) GetTagListByNames(ctx context.Context, tagNames []string) ([]*entity.Tag, error) {
	for k, tagname := range tagNames {
		tagNames[k] = strings.ToLower(tagname)
	}
	tagList, err := ts.tagCommonRepo.GetTagListByNames(ctx, tagNames)
	if err != nil {
		return nil, err
	}
	ts.TagsFormatRecommendAndReserved(ctx, tagList)
	return tagList, nil
}

func (ts *TagCommonService) ExistRecommend(ctx context.Context, tags []*schema.TagItem) (bool, error) {
	taginfo, err := ts.siteInfoService.GetSiteWrite(ctx)
	if err != nil {
		return false, err
	}
	if !taginfo.RequiredTag {
		return true, nil
	}
	tagNames := make([]string, 0)
	for _, item := range tags {
		tagNames = append(tagNames, item.SlugName)
	}
	list, err := ts.GetTagListByNames(ctx, tagNames)
	if err != nil {
		return false, err
	}
	for _, item := range list {
		if item.Recommend {
			return true, nil
		}
	}
	return false, nil
}

// GetObjectTag get object tag
func (ts *TagCommonService) GetObjectTag(ctx context.Context, objectId string) (objTags []*schema.TagResp, err error) {
	tagsInfoList, err := ts.GetObjectEntityTag(ctx, objectId)
	return ts.TagFormat(ctx, tagsInfoList)
}

// AddTagList get object tag
func (ts *TagCommonService) AddTagList(ctx context.Context, tagList []*entity.Tag) (err error) {
	return ts.tagCommonRepo.AddTagList(ctx, tagList)
}

// GetTagByID get object tag
func (ts *TagCommonService) GetTagByID(ctx context.Context, tagID string) (tag *entity.Tag, exist bool, err error) {
	tag, exist, err = ts.tagCommonRepo.GetTagByID(ctx, tagID, false)
	if !exist {
		return
	}
	ts.tagFormatRecommendAndReserved(ctx, tag)
	return
}

// GetTagBySlugName get object tag
func (ts *TagCommonService) GetTagBySlugName(ctx context.Context, slugName string) (tag *entity.Tag, exist bool, err error) {
	tag, exist, err = ts.tagCommonRepo.GetTagBySlugName(ctx, slugName)
	if !exist {
		return
	}
	ts.tagFormatRecommendAndReserved(ctx, tag)
	return
}

// GetTagListByIDs get object tag
func (ts *TagCommonService) GetTagListByIDs(ctx context.Context, ids []string) (tagList []*entity.Tag, err error) {
	tagList, err = ts.tagCommonRepo.GetTagListByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	ts.TagsFormatRecommendAndReserved(ctx, tagList)
	return
}

// GetTagPage get object tag
func (ts *TagCommonService) GetTagPage(ctx context.Context, page, pageSize int, tag *entity.Tag, queryCond string) (
	tagList []*entity.Tag, total int64, err error) {
	tagList, total, err = ts.tagCommonRepo.GetTagPage(ctx, page, pageSize, tag, queryCond)
	if err != nil {
		return nil, 0, err
	}
	ts.TagsFormatRecommendAndReserved(ctx, tagList)
	return
}

func (ts *TagCommonService) GetObjectEntityTag(ctx context.Context, objectId string) (objTags []*entity.Tag, err error) {
	tagIDList := make([]string, 0)
	tagList, err := ts.tagRelRepo.GetObjectTagRelList(ctx, objectId)
	if err != nil {
		return nil, err
	}
	for _, tag := range tagList {
		tagIDList = append(tagIDList, tag.TagID)
	}
	objTags, err = ts.GetTagListByIDs(ctx, tagIDList)
	if err != nil {
		return nil, err
	}
	return objTags, nil
}

func (ts *TagCommonService) TagFormat(ctx context.Context, tags []*entity.Tag) (objTags []*schema.TagResp, err error) {
	objTags = make([]*schema.TagResp, 0)
	for _, tagInfo := range tags {
		objTags = append(objTags, &schema.TagResp{
			SlugName:        tagInfo.SlugName,
			DisplayName:     tagInfo.DisplayName,
			MainTagSlugName: tagInfo.MainTagSlugName,
			Recommend:       tagInfo.Recommend,
			Reserved:        tagInfo.Reserved,
		})
	}
	return objTags, nil
}

func (ts *TagCommonService) TagsFormatRecommendAndReserved(ctx context.Context, tagList []*entity.Tag) {
	if len(tagList) == 0 {
		return
	}
	tagConfig, err := ts.siteInfoService.GetSiteWrite(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	if !tagConfig.RequiredTag {
		for _, tag := range tagList {
			tag.Recommend = false
		}
	}
}

func (ts *TagCommonService) tagFormatRecommendAndReserved(ctx context.Context, tag *entity.Tag) {
	if tag == nil {
		return
	}
	tagConfig, err := ts.siteInfoService.GetSiteWrite(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	if !tagConfig.RequiredTag {
		tag.Recommend = false
	}
}

// BatchGetObjectTag batch get object tag
func (ts *TagCommonService) BatchGetObjectTag(ctx context.Context, objectIds []string) (map[string][]*schema.TagResp, error) {
	objectIDTagMap := make(map[string][]*schema.TagResp)
	tagIDList := make([]string, 0)
	tagsInfoMap := make(map[string]*entity.Tag)

	tagList, err := ts.tagRelRepo.BatchGetObjectTagRelList(ctx, objectIds)
	if err != nil {
		return objectIDTagMap, err
	}
	for _, tag := range tagList {
		tagIDList = append(tagIDList, tag.TagID)
	}
	tagsInfoList, err := ts.GetTagListByIDs(ctx, tagIDList)
	if err != nil {
		return objectIDTagMap, err
	}
	for _, item := range tagsInfoList {
		tagsInfoMap[item.ID] = item
	}
	for _, item := range tagList {
		_, ok := tagsInfoMap[item.TagID]
		if ok {
			tagInfo := tagsInfoMap[item.TagID]
			t := &schema.TagResp{
				SlugName:        tagInfo.SlugName,
				DisplayName:     tagInfo.DisplayName,
				MainTagSlugName: tagInfo.MainTagSlugName,
				Recommend:       tagInfo.Recommend,
				Reserved:        tagInfo.Reserved,
			}
			objectIDTagMap[item.ObjectID] = append(objectIDTagMap[item.ObjectID], t)
		}
	}
	for _, taglist := range objectIDTagMap {
		sort.SliceStable(taglist, func(i, j int) bool {
			return taglist[i].Reserved
		})
		sort.SliceStable(taglist, func(i, j int) bool {
			return taglist[i].Recommend
		})
	}
	return objectIDTagMap, nil
}

func (ts *TagCommonService) CheckTag(ctx context.Context, tags []string, userID string) (err error) {
	if len(tags) == 0 {
		return nil
	}

	// find tags name
	tagListInDb, err := ts.GetTagListByNames(ctx, tags)
	if err != nil {
		return err
	}

	tagInDbMapping := make(map[string]*entity.Tag)
	checktags := make([]string, 0)

	for _, tag := range tagListInDb {
		if tag.MainTagID != 0 {
			checktags = append(checktags, fmt.Sprintf("\"%s\"", tag.SlugName))
		}
		tagInDbMapping[tag.SlugName] = tag
	}
	if len(checktags) > 0 {
		err = errors.BadRequest(reason.TagNotContainSynonym).WithMsg(fmt.Sprintf("Should not contain synonym tags %s", strings.Join(checktags, ",")))
		return err
	}

	addTagList := make([]*entity.Tag, 0)
	addTagMsgList := make([]string, 0)
	for _, tag := range tags {
		_, ok := tagInDbMapping[tag]
		if ok {
			continue
		}
		item := &entity.Tag{}
		item.SlugName = tag
		item.DisplayName = tag
		item.OriginalText = ""
		item.ParsedText = ""
		item.Status = entity.TagStatusAvailable
		addTagList = append(addTagList, item)
		addTagMsgList = append(addTagMsgList, tag)
	}

	if len(addTagList) > 0 {
		err = errors.BadRequest(reason.TagNotFound).WithMsg(fmt.Sprintf("tag [%s] does not exist",
			strings.Join(addTagMsgList, ",")))
		return err

	}

	return nil
}

func (ts *TagCommonService) CheckChangeReservedTag(ctx context.Context, oldobjectTagData, objectTagData []*entity.Tag) (bool, []string) {
	reservedTagsMap := make(map[string]bool)
	needTagsMap := make([]string, 0)
	for _, tag := range objectTagData {
		if tag.Reserved {
			reservedTagsMap[tag.SlugName] = true
		}
	}
	for _, tag := range oldobjectTagData {
		if tag.Reserved {
			_, ok := reservedTagsMap[tag.SlugName]
			if !ok {
				needTagsMap = append(needTagsMap, tag.SlugName)
			}
		}
	}
	if len(needTagsMap) > 0 {
		return false, needTagsMap
	}

	return true, []string{}
}

// ObjectChangeTag change object tag list
func (ts *TagCommonService) ObjectChangeTag(ctx context.Context, objectTagData *schema.TagChange) (err error) {
	if len(objectTagData.Tags) == 0 {
		return nil
	}

	thisObjTagNameList := make([]string, 0)
	thisObjTagIDList := make([]string, 0)
	for _, t := range objectTagData.Tags {
		t.SlugName = strings.ToLower(t.SlugName)
		thisObjTagNameList = append(thisObjTagNameList, t.SlugName)
	}

	// find tags name
	tagListInDb, err := ts.tagCommonRepo.GetTagListByNames(ctx, thisObjTagNameList)
	if err != nil {
		return err
	}

	tagInDbMapping := make(map[string]*entity.Tag)
	for _, tag := range tagListInDb {
		tagInDbMapping[tag.SlugName] = tag
		thisObjTagIDList = append(thisObjTagIDList, tag.ID)
	}

	addTagList := make([]*entity.Tag, 0)
	for _, tag := range objectTagData.Tags {
		_, ok := tagInDbMapping[tag.SlugName]
		if ok {
			continue
		}
		item := &entity.Tag{}
		item.SlugName = tag.SlugName
		item.DisplayName = tag.DisplayName
		item.OriginalText = tag.OriginalText
		item.ParsedText = tag.ParsedText
		item.Status = entity.TagStatusAvailable
		addTagList = append(addTagList, item)
	}

	if len(addTagList) > 0 {
		err = ts.tagCommonRepo.AddTagList(ctx, addTagList)
		if err != nil {
			return err
		}
		for _, tag := range addTagList {
			thisObjTagIDList = append(thisObjTagIDList, tag.ID)
			revisionDTO := &schema.AddRevisionDTO{
				UserID:   objectTagData.UserID,
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
				UserID:           objectTagData.UserID,
				ObjectID:         tag.ID,
				OriginalObjectID: tag.ID,
				ActivityTypeKey:  constant.ActTagCreated,
				RevisionID:       revisionID,
			})
		}
	}

	err = ts.CreateOrUpdateTagRelList(ctx, objectTagData.ObjectID, thisObjTagIDList)
	if err != nil {
		return err
	}
	return nil
}

// RefreshTagQuestionCount refresh tag question count
func (ts *TagCommonService) RefreshTagQuestionCount(ctx context.Context, tagIDs []string) (err error) {
	for _, tagID := range tagIDs {
		count, err := ts.tagRelRepo.CountTagRelByTagID(ctx, tagID)
		if err != nil {
			return err
		}
		err = ts.tagCommonRepo.UpdateTagQuestionCount(ctx, tagID, int(count))
		if err != nil {
			return err
		}
		log.Debugf("tag count updated %s %d", tagID, count)
	}
	return nil
}

// CreateOrUpdateTagRelList if tag relation is exists update status, if not create it
func (ts *TagCommonService) CreateOrUpdateTagRelList(ctx context.Context, objectId string, tagIDs []string) (err error) {
	addTagIDMapping := make(map[string]bool)
	needRefreshTagIDs := make([]string, 0)
	for _, t := range tagIDs {
		addTagIDMapping[t] = true
	}

	// get all old relation
	oldTagRelList, err := ts.tagRelRepo.GetObjectTagRelList(ctx, objectId)
	if err != nil {
		return err
	}
	var deleteTagRel []int64
	for _, rel := range oldTagRelList {
		if !addTagIDMapping[rel.TagID] {
			deleteTagRel = append(deleteTagRel, rel.ID)
			needRefreshTagIDs = append(needRefreshTagIDs, rel.TagID)
		}
	}

	addTagRelList := make([]*entity.TagRel, 0)
	enableTagRelList := make([]int64, 0)
	for _, tagID := range tagIDs {
		needRefreshTagIDs = append(needRefreshTagIDs, tagID)
		rel, exist, err := ts.tagRelRepo.GetObjectTagRelWithoutStatus(ctx, objectId, tagID)
		if err != nil {
			return err
		}
		// if not exist add tag relation
		if !exist {
			addTagRelList = append(addTagRelList, &entity.TagRel{
				TagID: tagID, ObjectID: objectId, Status: entity.TagStatusAvailable,
			})
		}
		// if exist and has been removed, that should be enabled
		if exist && rel.Status != entity.TagStatusAvailable {
			enableTagRelList = append(enableTagRelList, rel.ID)
		}
	}

	if len(deleteTagRel) > 0 {
		if err = ts.tagRelRepo.RemoveTagRelListByIDs(ctx, deleteTagRel); err != nil {
			return err
		}
	}
	if len(addTagRelList) > 0 {
		if err = ts.tagRelRepo.AddTagRelList(ctx, addTagRelList); err != nil {
			return err
		}
	}
	if len(enableTagRelList) > 0 {
		if err = ts.tagRelRepo.EnableTagRelByIDs(ctx, enableTagRelList); err != nil {
			return err
		}
	}

	err = ts.RefreshTagQuestionCount(ctx, needRefreshTagIDs)
	if err != nil {
		log.Error(err)
	}
	return nil
}

func (ts *TagCommonService) UpdateTag(ctx context.Context, req *schema.UpdateTagReq) (err error) {
	var canUpdate bool
	_, existUnreviewed, err := ts.revisionService.ExistUnreviewedByObjectID(ctx, req.TagID)
	if err != nil {
		return err
	}
	if existUnreviewed {
		err = errors.BadRequest(reason.AnswerCannotUpdate)
		return err
	}

	tagInfo, exist, err := ts.GetTagByID(ctx, req.TagID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.TagNotFound)
	}

	tagInfo.SlugName = req.SlugName
	tagInfo.DisplayName = req.DisplayName
	tagInfo.OriginalText = req.OriginalText
	tagInfo.ParsedText = req.ParsedText

	revisionDTO := &schema.AddRevisionDTO{
		UserID:   req.UserID,
		ObjectID: tagInfo.ID,
		Title:    tagInfo.SlugName,
		Log:      req.EditSummary,
	}

	if !req.IsAdmin {
		revisionDTO.Status = entity.RevisionUnreviewedStatus
	} else {
		canUpdate = true
		err = ts.tagRepo.UpdateTag(ctx, tagInfo)
		if err != nil {
			return err
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
		revisionDTO.Status = entity.RevisionReviewPassStatus
	}

	tagInfoJson, _ := json.Marshal(tagInfo)
	revisionDTO.Content = string(tagInfoJson)
	revisionID, err := ts.revisionService.AddRevision(ctx, revisionDTO, true)
	if err != nil {
		return err
	}
	if canUpdate {
		activity_queue.AddActivity(&schema.ActivityMsg{
			UserID:           req.UserID,
			ObjectID:         tagInfo.ID,
			OriginalObjectID: tagInfo.ID,
			ActivityTypeKey:  constant.ActTagEdited,
			RevisionID:       revisionID,
		})
	}

	return
}
