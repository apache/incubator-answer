package service

import (
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/constant"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	questioncommon "github.com/segmentfault/answer/internal/service/question_common"
	"github.com/segmentfault/answer/internal/service/revision"
	usercommon "github.com/segmentfault/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/errors"
)

// RevisionService user service
type RevisionService struct {
	revisionRepo   revision.RevisionRepo
	userCommon     *usercommon.UserCommon
	questionCommon *questioncommon.QuestionCommon
	answerService  *AnswerService
}

func NewRevisionService(
	revisionRepo revision.RevisionRepo,
	userCommon *usercommon.UserCommon,
	questionCommon *questioncommon.QuestionCommon,
	answerService *AnswerService) *RevisionService {
	return &RevisionService{
		revisionRepo:   revisionRepo,
		userCommon:     userCommon,
		questionCommon: questionCommon,
		answerService:  answerService,
	}
}

// GetRevision get revision one
func (rs *RevisionService) GetRevision(ctx context.Context, id string) (resp schema.GetRevisionResp, err error) {
	var (
		rev    *entity.Revision
		exists bool
	)

	resp = schema.GetRevisionResp{}

	rev, exists, err = rs.revisionRepo.GetRevision(ctx, id)
	if err != nil {
		return
	}

	if !exists {
		err = errors.BadRequest(reason.ObjectNotFound)
		return
	}

	_ = copier.Copy(&resp, rev)
	rs.parseItem(ctx, &resp)

	return
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
		question     entity.Question
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
		questionInfo = rs.questionCommon.ShowFormat(ctx, &question)
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
		}
		tagInfo.GetExcerpt()
		item.ContentParsed = tagInfo
	}

	if err != nil {
		item.ContentParsed = item.Content
	}
	item.CreatedAtParsed = item.CreatedAt.Unix()
}
