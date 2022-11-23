package service

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/object_info"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/revision"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/copier"
)

// RevisionService user service
type RevisionService struct {
	revisionRepo      revision.RevisionRepo
	userCommon        *usercommon.UserCommon
	questionCommon    *questioncommon.QuestionCommon
	answerService     *AnswerService
	objectInfoService *object_info.ObjService
}

func NewRevisionService(
	revisionRepo revision.RevisionRepo,
	userCommon *usercommon.UserCommon,
	questionCommon *questioncommon.QuestionCommon,
	answerService *AnswerService,
	objectInfoService *object_info.ObjService,
) *RevisionService {
	return &RevisionService{
		revisionRepo:      revisionRepo,
		userCommon:        userCommon,
		questionCommon:    questionCommon,
		answerService:     answerService,
		objectInfoService: objectInfoService,
	}
}

// SearchUnreviewedList get unreviewed list
func (rs *RevisionService) GetUnreviewedRevisionList(ctx context.Context, req *schema.RevisionSearch) (resp []*schema.GetUnreviewedRevisionResp, count int64, err error) {
	resp = []*schema.GetUnreviewedRevisionResp{}
	search := &entity.RevisionSearch{}
	_ = copier.Copy(search, req)
	list, count, err := rs.revisionRepo.SearchUnreviewedList(ctx, search)
	for _, revision := range list {
		spew.Dump(revision)
		item := &schema.GetUnreviewedRevisionResp{}
		_, ok := constant.ObjectTypeNumberMapping[revision.ObjectType]
		if !ok {
			continue
		}
		item.Type = constant.ObjectTypeNumberMapping[revision.ObjectType]
		info, err := rs.objectInfoService.GetInfo(ctx, revision.ObjectID)
		spew.Dump(info, err)
		resp = append(resp, item)
	}

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
