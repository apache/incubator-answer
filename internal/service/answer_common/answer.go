package answercommon

import (
	"context"

	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
)

type AnswerRepo interface {
	AddAnswer(ctx context.Context, answer *entity.Answer) (err error)
	RemoveAnswer(ctx context.Context, id string) (err error)
	UpdateAnswer(ctx context.Context, answer *entity.Answer, Colar []string) (err error)
	GetAnswer(ctx context.Context, id string) (answer *entity.Answer, exist bool, err error)
	GetAnswerList(ctx context.Context, answer *entity.Answer) (answerList []*entity.Answer, err error)
	GetAnswerPage(ctx context.Context, page, pageSize int, answer *entity.Answer) (answerList []*entity.Answer, total int64, err error)
	UpdateAdopted(ctx context.Context, id string, questionId string) error
	GetByID(ctx context.Context, id string) (*entity.Answer, bool, error)
	GetByUserIdQuestionId(ctx context.Context, userId string, questionId string) (*entity.Answer, bool, error)
	SearchList(ctx context.Context, search *entity.AnswerSearch) ([]*entity.Answer, int64, error)
	CmsSearchList(ctx context.Context, search *entity.CmsAnswerSearch) ([]*entity.Answer, int64, error)
	UpdateAnswerStatus(ctx context.Context, answer *entity.Answer) (err error)
}

// AnswerCommon user service
type AnswerCommon struct {
	answerRepo AnswerRepo
}

func NewAnswerCommon(answerRepo AnswerRepo) *AnswerCommon {
	return &AnswerCommon{
		answerRepo: answerRepo,
	}
}

func (as *AnswerCommon) SearchAnswered(ctx context.Context, userId, questionId string) (bool, error) {
	_, has, err := as.answerRepo.GetByUserIdQuestionId(ctx, userId, questionId)
	if err != nil {
		return has, err
	}
	return has, nil
}

func (as *AnswerCommon) CmsSearchList(ctx context.Context, search *entity.CmsAnswerSearch) ([]*entity.Answer, int64, error) {
	return as.answerRepo.CmsSearchList(ctx, search)
}

func (as *AnswerCommon) Search(ctx context.Context, search *entity.AnswerSearch) ([]*entity.Answer, int64, error) {
	list, count, err := as.answerRepo.SearchList(ctx, search)
	if err != nil {
		return list, count, err
	}
	return list, count, err
}

func (as *AnswerCommon) ShowFormat(ctx context.Context, data *entity.Answer) *schema.AnswerInfo {
	info := schema.AnswerInfo{}
	info.ID = data.ID
	info.QuestionId = data.QuestionID
	info.Content = data.OriginalText
	info.Html = data.ParsedText
	info.Adopted = data.Adopted
	info.VoteCount = data.VoteCount
	info.CreateTime = data.CreatedAt.Unix()
	info.UpdateTime = data.UpdatedAt.Unix()
	info.UserId = data.UserID
	return &info
}

func (as *AnswerCommon) AdminShowFormat(ctx context.Context, data *entity.Answer) *schema.AdminAnswerInfo {
	info := schema.AdminAnswerInfo{}
	info.ID = data.ID
	info.QuestionId = data.QuestionID
	info.Description = data.ParsedText
	info.Adopted = data.Adopted
	info.VoteCount = data.VoteCount
	info.CreateTime = data.CreatedAt.Unix()
	info.UpdateTime = data.UpdatedAt.Unix()
	info.UserId = data.UserID
	return &info
}
