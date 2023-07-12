package object_info

import (
	"context"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	"github.com/answerdev/answer/internal/service/comment_common"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	tagcommon "github.com/answerdev/answer/internal/service/tag_common"
	"github.com/answerdev/answer/pkg/obj"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/segmentfault/pacman/errors"
)

// ObjService user service
type ObjService struct {
	answerRepo   answercommon.AnswerRepo
	questionRepo questioncommon.QuestionRepo
	commentRepo  comment_common.CommentCommonRepo
	tagRepo      tagcommon.TagCommonRepo
	tagCommon    *tagcommon.TagCommonService
}

// NewObjService new object service
func NewObjService(
	answerRepo answercommon.AnswerRepo,
	questionRepo questioncommon.QuestionRepo,
	commentRepo comment_common.CommentCommonRepo,
	tagRepo tagcommon.TagCommonRepo,
	tagCommon *tagcommon.TagCommonService,
) *ObjService {
	return &ObjService{
		answerRepo:   answerRepo,
		questionRepo: questionRepo,
		commentRepo:  commentRepo,
		tagRepo:      tagRepo,
		tagCommon:    tagCommon,
	}
}
func (os *ObjService) GetUnreviewedRevisionInfo(ctx context.Context, objectID string) (objInfo *schema.UnreviewedRevisionInfoInfo, err error) {
	objectType, err := obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		return nil, err
	}
	switch objectType {
	case constant.QuestionObjectType:
		questionInfo, exist, err := os.questionRepo.GetQuestion(ctx, objectID)
		if err != nil {
			return nil, err
		}
		questionInfo.ID = uid.EnShortID(questionInfo.ID)
		if !exist {
			break
		}
		taglist, err := os.tagCommon.GetObjectEntityTag(ctx, objectID)
		if err != nil {
			return nil, err
		}
		os.tagCommon.TagsFormatRecommendAndReserved(ctx, taglist)
		tags, err := os.tagCommon.TagFormat(ctx, taglist)
		if err != nil {
			return nil, err
		}
		objInfo = &schema.UnreviewedRevisionInfoInfo{
			ObjectID: questionInfo.ID,
			Title:    questionInfo.Title,
			Content:  questionInfo.OriginalText,
			Html:     questionInfo.ParsedText,
			Tags:     tags,
		}
	case constant.AnswerObjectType:
		answerInfo, exist, err := os.answerRepo.GetAnswer(ctx, objectID)
		if err != nil {
			return nil, err
		}
		if !exist {
			break
		}

		questionInfo, exist, err := os.questionRepo.GetQuestion(ctx, answerInfo.QuestionID)
		if err != nil {
			return nil, err
		}
		if !exist {
			break
		}
		questionInfo.ID = uid.EnShortID(questionInfo.ID)
		objInfo = &schema.UnreviewedRevisionInfoInfo{
			ObjectID: answerInfo.ID,
			Title:    questionInfo.Title,
			Content:  answerInfo.OriginalText,
			Html:     answerInfo.ParsedText,
		}

	case constant.TagObjectType:
		tagInfo, exist, err := os.tagRepo.GetTagByID(ctx, objectID, true)
		if err != nil {
			return nil, err
		}
		if !exist {
			break
		}
		objInfo = &schema.UnreviewedRevisionInfoInfo{
			ObjectID: tagInfo.ID,
			Title:    tagInfo.SlugName,
			Content:  tagInfo.OriginalText,
			Html:     tagInfo.ParsedText,
		}
	}
	if objInfo == nil {
		err = errors.BadRequest(reason.ObjectNotFound)
	}
	return objInfo, err
}

// GetInfo get object simple information
func (os *ObjService) GetInfo(ctx context.Context, objectID string) (objInfo *schema.SimpleObjectInfo, err error) {
	objectType, err := obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		return nil, err
	}
	switch objectType {
	case constant.QuestionObjectType:
		questionInfo, exist, err := os.questionRepo.GetQuestion(ctx, objectID)
		if err != nil {
			return nil, err
		}
		if !exist {
			break
		}
		objInfo = &schema.SimpleObjectInfo{
			ObjectID:            questionInfo.ID,
			ObjectCreatorUserID: questionInfo.UserID,
			QuestionID:          questionInfo.ID,
			QuestionStatus:      questionInfo.Status,
			ObjectType:          objectType,
			Title:               questionInfo.Title,
			Content:             questionInfo.ParsedText, // todo trim
		}
	case constant.AnswerObjectType:
		answerInfo, exist, err := os.answerRepo.GetAnswer(ctx, objectID)
		if err != nil {
			return nil, err
		}
		if !exist {
			break
		}
		questionInfo, exist, err := os.questionRepo.GetQuestion(ctx, answerInfo.QuestionID)
		if err != nil {
			return nil, err
		}
		if !exist {
			break
		}
		objInfo = &schema.SimpleObjectInfo{
			ObjectID:            answerInfo.ID,
			ObjectCreatorUserID: answerInfo.UserID,
			QuestionID:          answerInfo.QuestionID,
			QuestionStatus:      questionInfo.Status,
			AnswerID:            answerInfo.ID,
			ObjectType:          objectType,
			Title:               questionInfo.Title,    // this should be question title
			Content:             answerInfo.ParsedText, // todo trim
		}
	case constant.CommentObjectType:
		commentInfo, exist, err := os.commentRepo.GetComment(ctx, objectID)
		if err != nil {
			return nil, err
		}
		if !exist {
			break
		}
		objInfo = &schema.SimpleObjectInfo{
			ObjectID:            commentInfo.ID,
			ObjectCreatorUserID: commentInfo.UserID,
			ObjectType:          objectType,
			Content:             commentInfo.ParsedText, // todo trim
			CommentID:           commentInfo.ID,
		}
		if len(commentInfo.QuestionID) > 0 {
			questionInfo, exist, err := os.questionRepo.GetQuestion(ctx, commentInfo.QuestionID)
			if err != nil {
				return nil, err
			}
			if exist {
				objInfo.QuestionID = questionInfo.ID
				objInfo.QuestionStatus = questionInfo.Status
				objInfo.Title = questionInfo.Title
			}
			answerInfo, exist, err := os.answerRepo.GetAnswer(ctx, commentInfo.ObjectID)
			if err != nil {
				return nil, err
			}
			if exist {
				objInfo.AnswerID = answerInfo.ID
			}
		}
	case constant.TagObjectType:
		tagInfo, exist, err := os.tagRepo.GetTagByID(ctx, objectID, true)
		if err != nil {
			return nil, err
		}
		if !exist {
			break
		}
		objInfo = &schema.SimpleObjectInfo{
			ObjectID:   tagInfo.ID,
			TagID:      tagInfo.ID,
			ObjectType: objectType,
			Title:      tagInfo.ParsedText,
			Content:    tagInfo.ParsedText, // todo trim
		}
	}
	if objInfo == nil {
		err = errors.BadRequest(reason.ObjectNotFound)
	}
	return objInfo, err
}
