package answercommon

import (
	"context"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
)

type AnswerRepo interface {
	AddAnswer(ctx context.Context, answer *entity.Answer) (err error)
	RemoveAnswer(ctx context.Context, id string) (err error)
	UpdateAnswer(ctx context.Context, answer *entity.Answer, Colar []string) (err error)
	GetAnswer(ctx context.Context, id string) (answer *entity.Answer, exist bool, err error)
	GetAnswerList(ctx context.Context, answer *entity.Answer) (answerList []*entity.Answer, err error)
	GetAnswerPage(ctx context.Context, page, pageSize int, answer *entity.Answer) (answerList []*entity.Answer, total int64, err error)
	UpdateAdopted(ctx context.Context, id string, questionID string) error
	GetByID(ctx context.Context, id string) (*entity.Answer, bool, error)
	GetByUserIDQuestionID(ctx context.Context, userID string, questionID string) (*entity.Answer, bool, error)
	SearchList(ctx context.Context, search *entity.AnswerSearch) ([]*entity.Answer, int64, error)
	CmsSearchList(ctx context.Context, search *entity.CmsAnswerSearch) ([]*entity.Answer, int64, error)
	UpdateAnswerStatus(ctx context.Context, answer *entity.Answer) (err error)
	GetAnswerCount(ctx context.Context) (count int64, err error)
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

func (as *AnswerCommon) SearchAnswered(ctx context.Context, userID, questionID string) (bool, error) {
	_, has, err := as.answerRepo.GetByUserIDQuestionID(ctx, userID, questionID)
	if err != nil {
		return has, err
	}
	return has, nil
}

func (as *AnswerCommon) CmsSearchList(ctx context.Context, search *entity.CmsAnswerSearch) ([]*entity.Answer, int64, error) {
	if search.Status == 0 {
		search.Status = 1
	}
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
	info.QuestionID = data.QuestionID
	info.Content = data.OriginalText
	info.HTML = data.ParsedText
	info.Adopted = data.Adopted
	info.VoteCount = data.VoteCount
	info.CreateTime = data.CreatedAt.Unix()
	info.UpdateTime = data.UpdatedAt.Unix()
	info.UserID = data.UserID
	return &info
}

func (as *AnswerCommon) AdminShowFormat(ctx context.Context, data *entity.Answer) *schema.AdminAnswerInfo {
	info := schema.AdminAnswerInfo{}
	info.ID = data.ID
	info.QuestionID = data.QuestionID
	info.Description = data.ParsedText
	info.Adopted = data.Adopted
	info.VoteCount = data.VoteCount
	info.CreateTime = data.CreatedAt.Unix()
	info.UpdateTime = data.UpdatedAt.Unix()
	info.UserID = data.UserID
	return &info
}
