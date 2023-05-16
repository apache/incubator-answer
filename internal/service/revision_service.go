package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_queue"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	"github.com/answerdev/answer/internal/service/notice_queue"
	"github.com/answerdev/answer/internal/service/object_info"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/revision"
	"github.com/answerdev/answer/internal/service/tag_common"
	tagcommon "github.com/answerdev/answer/internal/service/tag_common"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/answerdev/answer/pkg/obj"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// RevisionService user service
type RevisionService struct {
	revisionRepo      revision.RevisionRepo
	userCommon        *usercommon.UserCommon
	questionCommon    *questioncommon.QuestionCommon
	answerService     *AnswerService
	objectInfoService *object_info.ObjService
	questionRepo      questioncommon.QuestionRepo
	answerRepo        answercommon.AnswerRepo
	tagRepo           tag_common.TagRepo
	tagCommon         *tagcommon.TagCommonService
}

func NewRevisionService(
	revisionRepo revision.RevisionRepo,
	userCommon *usercommon.UserCommon,
	questionCommon *questioncommon.QuestionCommon,
	answerService *AnswerService,
	objectInfoService *object_info.ObjService,
	questionRepo questioncommon.QuestionRepo,
	answerRepo answercommon.AnswerRepo,
	tagRepo tag_common.TagRepo,
	tagCommon *tagcommon.TagCommonService,
) *RevisionService {
	return &RevisionService{
		revisionRepo:      revisionRepo,
		userCommon:        userCommon,
		questionCommon:    questionCommon,
		answerService:     answerService,
		objectInfoService: objectInfoService,
		questionRepo:      questionRepo,
		answerRepo:        answerRepo,
		tagRepo:           tagRepo,
		tagCommon:         tagCommon,
	}
}

func (rs *RevisionService) RevisionAudit(ctx context.Context, req *schema.RevisionAuditReq) (err error) {
	revisioninfo, exist, err := rs.revisionRepo.GetRevisionByID(ctx, req.ID)
	if err != nil {
		return
	}
	if !exist {
		return
	}
	if revisioninfo.Status != entity.RevisionUnreviewedStatus {
		return
	}
	if req.Operation == schema.RevisionAuditReject {
		err = rs.revisionRepo.UpdateStatus(ctx, req.ID, entity.RevisionReviewRejectStatus, req.UserID)
		return
	}
	if req.Operation == schema.RevisionAuditApprove {
		objectType, objectTypeerr := obj.GetObjectTypeStrByObjectID(revisioninfo.ObjectID)
		if objectTypeerr != nil {
			return objectTypeerr
		}
		revisionitem := &schema.GetRevisionResp{}
		_ = copier.Copy(revisionitem, revisioninfo)
		rs.parseItem(ctx, revisionitem)
		var saveErr error
		switch objectType {
		case constant.QuestionObjectType:
			if !req.CanReviewQuestion {
				saveErr = errors.BadRequest(reason.RevisionNoPermission)
			} else {
				saveErr = rs.revisionAuditQuestion(ctx, revisionitem)
			}
		case constant.AnswerObjectType:
			if !req.CanReviewAnswer {
				saveErr = errors.BadRequest(reason.RevisionNoPermission)
			} else {
				saveErr = rs.revisionAuditAnswer(ctx, revisionitem)
			}
		case constant.TagObjectType:
			if !req.CanReviewTag {
				saveErr = errors.BadRequest(reason.RevisionNoPermission)
			} else {
				saveErr = rs.revisionAuditTag(ctx, revisionitem)
			}
		}
		if saveErr != nil {
			return saveErr
		}
		err = rs.revisionRepo.UpdateStatus(ctx, req.ID, entity.RevisionReviewPassStatus, req.UserID)
		return
	}

	return nil
}

func (rs *RevisionService) revisionAuditQuestion(ctx context.Context, revisionitem *schema.GetRevisionResp) (err error) {
	questioninfo, ok := revisionitem.ContentParsed.(*schema.QuestionInfo)
	if ok {
		var PostUpdateTime time.Time
		dbquestion, exist, dberr := rs.questionRepo.GetQuestion(ctx, questioninfo.ID)
		if dberr != nil || !exist {
			return
		}

		PostUpdateTime = time.Unix(questioninfo.UpdateTime, 0)
		if dbquestion.PostUpdateTime.Unix() > PostUpdateTime.Unix() {
			PostUpdateTime = dbquestion.PostUpdateTime
		}
		question := &entity.Question{}
		question.ID = questioninfo.ID
		question.Title = questioninfo.Title
		question.OriginalText = questioninfo.Content
		question.ParsedText = questioninfo.HTML
		question.UpdatedAt = time.Unix(questioninfo.UpdateTime, 0)
		question.PostUpdateTime = PostUpdateTime
		question.LastEditUserID = revisionitem.UserID
		saveerr := rs.questionRepo.UpdateQuestion(ctx, question, []string{"title", "original_text", "parsed_text", "updated_at", "post_update_time", "last_edit_user_id"})
		if saveerr != nil {
			return saveerr
		}
		objectTagTags := make([]*schema.TagItem, 0)
		for _, tag := range questioninfo.Tags {
			item := &schema.TagItem{}
			item.SlugName = tag.SlugName
			objectTagTags = append(objectTagTags, item)
		}
		objectTagData := schema.TagChange{}
		objectTagData.ObjectID = question.ID
		objectTagData.Tags = objectTagTags
		saveerr = rs.tagCommon.ObjectChangeTag(ctx, &objectTagData)
		if saveerr != nil {
			return saveerr
		}
		activity_queue.AddActivity(&schema.ActivityMsg{
			UserID:           revisionitem.UserID,
			ObjectID:         revisionitem.ObjectID,
			ActivityTypeKey:  constant.ActQuestionEdited,
			RevisionID:       revisionitem.ID,
			OriginalObjectID: revisionitem.ObjectID,
		})
	}
	return nil
}

func (rs *RevisionService) revisionAuditAnswer(ctx context.Context, revisionitem *schema.GetRevisionResp) (err error) {
	answerinfo, ok := revisionitem.ContentParsed.(*schema.AnswerInfo)
	if ok {

		var PostUpdateTime time.Time
		dbquestion, exist, dberr := rs.questionRepo.GetQuestion(ctx, answerinfo.QuestionID)
		if dberr != nil || !exist {
			return
		}

		PostUpdateTime = time.Unix(answerinfo.UpdateTime, 0)
		if dbquestion.PostUpdateTime.Unix() > PostUpdateTime.Unix() {
			PostUpdateTime = dbquestion.PostUpdateTime
		}

		insertData := new(entity.Answer)
		insertData.ID = answerinfo.ID
		insertData.OriginalText = answerinfo.Content
		insertData.ParsedText = answerinfo.HTML
		insertData.UpdatedAt = time.Unix(answerinfo.UpdateTime, 0)
		insertData.LastEditUserID = revisionitem.UserID
		saveerr := rs.answerRepo.UpdateAnswer(ctx, insertData, []string{"original_text", "parsed_text", "updated_at", "last_edit_user_id"})
		if saveerr != nil {
			return saveerr
		}
		saveerr = rs.questionCommon.UpdatePostSetTime(ctx, answerinfo.QuestionID, PostUpdateTime)
		if saveerr != nil {
			return saveerr
		}
		questionInfo, exist, err := rs.questionRepo.GetQuestion(ctx, answerinfo.QuestionID)
		if err != nil {
			return err
		}
		if !exist {
			return errors.BadRequest(reason.QuestionNotFound)
		}
		msg := &schema.NotificationMsg{
			TriggerUserID:  revisionitem.UserID,
			ReceiverUserID: questionInfo.UserID,
			Type:           schema.NotificationTypeInbox,
			ObjectID:       answerinfo.ID,
		}
		msg.ObjectType = constant.AnswerObjectType
		msg.NotificationAction = constant.NotificationUpdateAnswer
		notice_queue.AddNotification(msg)

		activity_queue.AddActivity(&schema.ActivityMsg{
			UserID:           revisionitem.UserID,
			ObjectID:         insertData.ID,
			OriginalObjectID: insertData.ID,
			ActivityTypeKey:  constant.ActAnswerEdited,
			RevisionID:       revisionitem.ID,
		})
	}
	return nil
}

func (rs *RevisionService) revisionAuditTag(ctx context.Context, revisionitem *schema.GetRevisionResp) (err error) {
	taginfo, ok := revisionitem.ContentParsed.(*schema.GetTagResp)
	if ok {
		tag := &entity.Tag{}
		tag.ID = taginfo.TagID
		tag.OriginalText = taginfo.OriginalText
		tag.ParsedText = taginfo.ParsedText
		saveerr := rs.tagRepo.UpdateTag(ctx, tag)
		if saveerr != nil {
			return saveerr
		}

		tagInfo, exist, err := rs.tagCommon.GetTagByID(ctx, taginfo.TagID)
		if err != nil {
			return err
		}
		if !exist {
			return errors.BadRequest(reason.TagNotFound)
		}
		if tagInfo.MainTagID == 0 && len(tagInfo.SlugName) > 0 {
			log.Debugf("tag %s update slug_name", tagInfo.SlugName)
			tagList, err := rs.tagRepo.GetTagList(ctx, &entity.Tag{MainTagID: converter.StringToInt64(tagInfo.ID)})
			if err != nil {
				return err
			}
			updateTagSlugNames := make([]string, 0)
			for _, tag := range tagList {
				updateTagSlugNames = append(updateTagSlugNames, tag.SlugName)
			}
			err = rs.tagRepo.UpdateTagSynonym(ctx, updateTagSlugNames, converter.StringToInt64(tagInfo.ID), tagInfo.MainTagSlugName)
			if err != nil {
				return err
			}
		}

		activity_queue.AddActivity(&schema.ActivityMsg{
			UserID:           revisionitem.UserID,
			ObjectID:         taginfo.TagID,
			OriginalObjectID: taginfo.TagID,
			ActivityTypeKey:  constant.ActTagEdited,
			RevisionID:       revisionitem.ID,
		})
	}
	return nil
}

// GetUnreviewedRevisionPage get unreviewed list
func (rs *RevisionService) GetUnreviewedRevisionPage(ctx context.Context, req *schema.RevisionSearch) (
	resp *pager.PageModel, err error) {
	revisionResp := make([]*schema.GetUnreviewedRevisionResp, 0)
	if len(req.GetCanReviewObjectTypes()) == 0 {
		return pager.NewPageModel(0, revisionResp), nil
	}
	revisionPage, total, err := rs.revisionRepo.GetUnreviewedRevisionPage(
		ctx, req.Page, 1, req.GetCanReviewObjectTypes())
	if err != nil {
		return nil, err
	}
	for _, rev := range revisionPage {
		item := &schema.GetUnreviewedRevisionResp{}
		_, ok := constant.ObjectTypeNumberMapping[rev.ObjectType]
		if !ok {
			continue
		}
		item.Type = constant.ObjectTypeNumberMapping[rev.ObjectType]
		info, err := rs.objectInfoService.GetUnreviewedRevisionInfo(ctx, rev.ObjectID)
		if err != nil {
			return nil, err
		}
		item.Info = info
		revisionitem := &schema.GetRevisionResp{}
		_ = copier.Copy(revisionitem, rev)
		rs.parseItem(ctx, revisionitem)
		item.UnreviewedInfo = revisionitem

		// get user info
		userInfo, exists, e := rs.userCommon.GetUserBasicInfoByID(ctx, revisionitem.UserID)
		if e != nil {
			return nil, e
		}
		if exists {
			var uinfo schema.UserBasicInfo
			_ = copier.Copy(&uinfo, userInfo)
			item.UnreviewedInfo.UserInfo = uinfo
		}
		revisionResp = append(revisionResp, item)
	}
	return pager.NewPageModel(total, revisionResp), nil
}

// GetRevisionList get revision list all
func (rs *RevisionService) GetRevisionList(ctx context.Context, req *schema.GetRevisionListReq) (resp []schema.GetRevisionResp, err error) {
	var (
		rev  entity.Revision
		revs []entity.Revision
	)

	resp = []schema.GetRevisionResp{}
	_ = copier.Copy(&rev, req)

	revs, err = rs.revisionRepo.GetRevisionList(ctx, &rev)
	if err != nil {
		return
	}

	for _, r := range revs {
		var (
			uinfo schema.UserBasicInfo
			item  schema.GetRevisionResp
		)

		_ = copier.Copy(&item, r)
		rs.parseItem(ctx, &item)

		// get user info
		userInfo, exists, e := rs.userCommon.GetUserBasicInfoByID(ctx, item.UserID)
		if e != nil {
			return nil, e
		}
		if exists {
			err = copier.Copy(&uinfo, userInfo)
			item.UserInfo = uinfo
		}
		resp = append(resp, item)
	}
	return
}

func (rs *RevisionService) parseItem(ctx context.Context, item *schema.GetRevisionResp) {
	var (
		err          error
		question     entity.QuestionWithTagsRevision
		questionInfo *schema.QuestionInfo
		answer       entity.Answer
		answerInfo   *schema.AnswerInfo
		tag          entity.Tag
		tagInfo      *schema.GetTagResp
	)

	switch item.ObjectType {
	case constant.ObjectTypeStrMapping["question"]:
		err = json.Unmarshal([]byte(item.Content), &question)
		if err != nil {
			break
		}
		questionInfo = rs.questionCommon.ShowFormatWithTag(ctx, &question)
		item.ContentParsed = questionInfo
	case constant.ObjectTypeStrMapping["answer"]:
		err = json.Unmarshal([]byte(item.Content), &answer)
		if err != nil {
			break
		}
		answerInfo = rs.answerService.ShowFormat(ctx, &answer)
		item.ContentParsed = answerInfo
	case constant.ObjectTypeStrMapping["tag"]:
		err = json.Unmarshal([]byte(item.Content), &tag)
		if err != nil {
			break
		}
		tagInfo = &schema.GetTagResp{
			TagID:         tag.ID,
			CreatedAt:     tag.CreatedAt.Unix(),
			UpdatedAt:     tag.UpdatedAt.Unix(),
			SlugName:      tag.SlugName,
			DisplayName:   tag.DisplayName,
			OriginalText:  tag.OriginalText,
			ParsedText:    tag.ParsedText,
			FollowCount:   tag.FollowCount,
			QuestionCount: tag.QuestionCount,
			Recommend:     tag.Recommend,
			Reserved:      tag.Reserved,
		}
		tagInfo.GetExcerpt()
		item.ContentParsed = tagInfo
	}

	if err != nil {
		item.ContentParsed = item.Content
	}
	item.CreatedAtParsed = item.CreatedAt.Unix()
}

// CheckCanUpdateRevision can check revision
func (rs *RevisionService) CheckCanUpdateRevision(ctx context.Context, req *schema.CheckCanQuestionUpdate) (
	resp *schema.ErrTypeData, err error) {
	_, exist, err := rs.revisionRepo.ExistUnreviewedByObjectID(ctx, req.ID)
	if err != nil {
		return nil, nil
	}
	if exist {
		return &schema.ErrTypeToast, errors.BadRequest(reason.RevisionReviewUnderway)
	}
	return nil, nil
}
