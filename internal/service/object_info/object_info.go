package object_info

import (
	"context"

	"github.com/segmentfault/answer/internal/base/constant"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/schema"
	answercommon "github.com/segmentfault/answer/internal/service/answer_common"
	"github.com/segmentfault/answer/internal/service/comment_common"
	questioncommon "github.com/segmentfault/answer/internal/service/question_common"
	tagcommon "github.com/segmentfault/answer/internal/service/tag_common"
	"github.com/segmentfault/answer/pkg/obj"
	"github.com/segmentfault/pacman/errors"
)

// ObjService user service
type ObjService struct {
	answerRepo   answercommon.AnswerRepo
	questionRepo questioncommon.QuestionRepo
	commentRepo  comment_common.CommentCommonRepo
	tagRepo      tagcommon.TagRepo
}

// NewObjService new object service
func NewObjService(
	answerRepo answercommon.AnswerRepo,
	questionRepo questioncommon.QuestionRepo,
	commentRepo comment_common.CommentCommonRepo,
	tagRepo tagcommon.TagRepo) *ObjService {
	return &ObjService{
		answerRepo:   answerRepo,
		questionRepo: questionRepo,
		commentRepo:  commentRepo,
		tagRepo:      tagRepo,
	}
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
			ObjectID:      questionInfo.ID,
			ObjectCreator: questionInfo.UserID,
			QuestionID:    questionInfo.ID,
			ObjectType:    objectType,
			Title:         questionInfo.Title,
			Content:       questionInfo.ParsedText, // todo trim
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
		objInfo = &schema.SimpleObjectInfo{
			ObjectID:      answerInfo.ID,
			ObjectCreator: answerInfo.UserID,
			QuestionID:    answerInfo.QuestionID,
			AnswerID:      answerInfo.ID,
			ObjectType:    objectType,
			Title:         questionInfo.Title,    // this should be question title
			Content:       answerInfo.ParsedText, // todo trim
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
			ObjectID:      commentInfo.ID,
			ObjectCreator: commentInfo.UserID,
			ObjectType:    objectType,
			Content:       commentInfo.ParsedText, // todo trim
			CommentID:     commentInfo.ID,
		}
		if len(commentInfo.QuestionID) > 0 {
			questionInfo, exist, err := os.questionRepo.GetQuestion(ctx, commentInfo.QuestionID)
			if err != nil {
				return nil, err
			}
			if exist {
				objInfo.QuestionID = questionInfo.ID
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
		tagInfo, exist, err := os.tagRepo.GetTagByID(ctx, objectID)
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
