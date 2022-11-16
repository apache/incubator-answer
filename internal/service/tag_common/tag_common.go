package tagcommon

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/revision_common"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type TagRepo interface {
	AddTagList(ctx context.Context, tagList []*entity.Tag) (err error)
	GetTagListByIDs(ctx context.Context, ids []string) (tagList []*entity.Tag, err error)
	GetTagBySlugName(ctx context.Context, slugName string) (tagInfo *entity.Tag, exist bool, err error)
	GetTagListByName(ctx context.Context, name string, limit int, hasReserved bool) (tagList []*entity.Tag, err error)
	GetTagListByNames(ctx context.Context, names []string) (tagList []*entity.Tag, err error)
	RemoveTag(ctx context.Context, tagID string) (err error)
	UpdateTag(ctx context.Context, tag *entity.Tag) (err error)
	UpdateTagQuestionCount(ctx context.Context, tagID string, questionCount int) (err error)
	UpdateTagSynonym(ctx context.Context, tagSlugNameList []string, mainTagID int64, mainTagSlugName string) (err error)
	GetTagByID(ctx context.Context, tagID string) (tag *entity.Tag, exist bool, err error)
	GetTagList(ctx context.Context, tag *entity.Tag) (tagList []*entity.Tag, err error)
	GetTagPage(ctx context.Context, page, pageSize int, tag *entity.Tag, queryCond string) (tagList []*entity.Tag, total int64, err error)
	GetRecommendTagList(ctx context.Context) (tagList []*entity.Tag, err error)
	GetReservedTagList(ctx context.Context) (tagList []*entity.Tag, err error)
	UpdateTagsAttribute(ctx context.Context, tags []string, attribute string, value bool) (err error)
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
	tagRepo         TagRepo
	tagRelRepo      TagRelRepo
	siteInfoService *siteinfo_common.SiteInfoCommonService
}

// NewTagCommonService new tag service
func NewTagCommonService(tagRepo TagRepo, tagRelRepo TagRelRepo,
	revisionService *revision_common.RevisionService,
	siteInfoService *siteinfo_common.SiteInfoCommonService,
) *TagCommonService {
	return &TagCommonService{
		tagRepo:         tagRepo,
		tagRelRepo:      tagRelRepo,
		revisionService: revisionService,
		siteInfoService: siteInfoService,
	}
}

func (ts *TagCommonService) GetSiteWriteRecommendTag(ctx context.Context) (tags []string, err error) {
	tags = make([]string, 0)
	list, err := ts.tagRepo.GetRecommendTagList(ctx)
	for _, item := range list {
		tags = append(tags, item.SlugName)
	}
	return tags, nil
}

func (ts *TagCommonService) SetSiteWriteRecommendTag(ctx context.Context, tags []string, required bool, userID string) (msg string, err error) {
	err = ts.UpdateTag(ctx, tags, userID)
	if err != nil {
		return err.Error(), err
	}
	err = ts.SetTagsAttribute(ctx, tags, "recommend", true)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (ts *TagCommonService) GetSiteWriteReservedTag(ctx context.Context) (tags []string, err error) {
	tags = make([]string, 0)
	list, err := ts.tagRepo.GetReservedTagList(ctx)
	for _, item := range list {
		tags = append(tags, item.SlugName)
	}
	return tags, nil
}

func (ts *TagCommonService) SetSiteWriteReservedTag(ctx context.Context, tags []string, userID string) (msg string, err error) {
	err = ts.UpdateTag(ctx, tags, userID)
	if err != nil {
		return err.Error(), err
	}
	err = ts.SetTagsAttribute(ctx, tags, "reserved", true)
	if err != nil {
		return "", err
	}
	return "", nil
}

// SetTagsAttribute
func (ts *TagCommonService) SetTagsAttribute(ctx context.Context, tags []string, attribute string, value bool) (err error) {
	var tagslist []string
	switch attribute {
	case "recommend":
		tagslist, err = ts.GetSiteWriteRecommendTag(ctx)
	case "reserved":
		tagslist, err = ts.GetSiteWriteReservedTag(ctx)
	default:
		return
	}
	err = ts.tagRepo.UpdateTagsAttribute(ctx, tagslist, attribute, false)
	if err != nil {
		return err
	}
	err = ts.tagRepo.UpdateTagsAttribute(ctx, tags, attribute, value)
	if err != nil {
		return err
	}
	return nil
}

// GetTagListByName
func (ts *TagCommonService) GetTagListByName(ctx context.Context, tagName string) (tagInfo *entity.Tag, exist bool, err error) {
	tagName = strings.ToLower(tagName)
	return ts.tagRepo.GetTagBySlugName(ctx, tagName)
}

func (ts *TagCommonService) GetTagListByNames(ctx context.Context, tagNames []string) ([]*entity.Tag, error) {
	for k, tagname := range tagNames {
		tagNames[k] = strings.ToLower(tagname)
	}
	return ts.tagRepo.GetTagListByNames(ctx, tagNames)
}

func (ts *TagCommonService) ExistRecommend(ctx context.Context, tags []*schema.TagItem) (bool, error) {
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

//

// GetObjectTag get object tag
func (ts *TagCommonService) GetObjectTag(ctx context.Context, objectId string) (objTags []*schema.TagResp, err error) {
	objTags = make([]*schema.TagResp, 0)
	tagIDList := make([]string, 0)
	tagList, err := ts.tagRelRepo.GetObjectTagRelList(ctx, objectId)
	if err != nil {
		return nil, err
	}
	for _, tag := range tagList {
		tagIDList = append(tagIDList, tag.TagID)
	}
	tagsInfoList, err := ts.tagRepo.GetTagListByIDs(ctx, tagIDList)
	if err != nil {
		return nil, err
	}
	for _, tagInfo := range tagsInfoList {
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
	tagsInfoList, err := ts.tagRepo.GetTagListByIDs(ctx, tagIDList)
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
	return objectIDTagMap, nil
}

func (ts *TagCommonService) UpdateTag(ctx context.Context, tags []string, userID string) (err error) {
	if len(tags) == 0 {
		return nil
	}

	thisTagNameList := make([]string, 0)
	thisTagIDList := make([]string, 0)
	for _, t := range tags {
		t = strings.ToLower(t)
		thisTagNameList = append(thisTagNameList, t)
	}

	// find tags name
	tagListInDb, err := ts.tagRepo.GetTagListByNames(ctx, thisTagNameList)
	if err != nil {
		return err
	}

	tagInDbMapping := make(map[string]*entity.Tag)
	for _, tag := range tagListInDb {
		if tag.MainTagID != 0 {
			err = errors.BadRequest(reason.TagNotContainSynonym).WithMsg(fmt.Sprintf("tag name:%s", tag.SlugName))
			return err
		}
		tagInDbMapping[tag.SlugName] = tag
		thisTagIDList = append(thisTagIDList, tag.ID)
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
		// todo if need add
		// err = ts.tagRepo.AddTagList(ctx, addTagList)
		// if err != nil {
		// 	return err
		// }
		// for _, tag := range addTagList {
		// 	thisTagIDList = append(thisTagIDList, tag.ID)
		// 	revisionDTO := &schema.AddRevisionDTO{
		// 		UserID:   userID,
		// 		ObjectID: tag.ID,
		// 		Title:    tag.SlugName,
		// 	}
		// 	tagInfoJson, _ := json.Marshal(tag)
		// 	revisionDTO.Content = string(tagInfoJson)
		// 	err = ts.revisionService.AddRevision(ctx, revisionDTO, true)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
	}

	return nil
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
	tagListInDb, err := ts.tagRepo.GetTagListByNames(ctx, thisObjTagNameList)
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
		err = ts.tagRepo.AddTagList(ctx, addTagList)
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
			err = ts.revisionService.AddRevision(ctx, revisionDTO, true)
			if err != nil {
				return err
			}
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
		err = ts.tagRepo.UpdateTagQuestionCount(ctx, tagID, int(count))
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
