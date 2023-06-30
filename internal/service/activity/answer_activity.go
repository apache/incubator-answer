package activity

import (
	"context"
)

// AnswerActivityRepo answer activity
type AnswerActivityRepo interface {
	AcceptAnswer(ctx context.Context,
		answerObjID, questionObjID, questionUserID, answerUserID string, isSelf bool) (err error)
	CancelAcceptAnswer(ctx context.Context,
		answerObjID, questionObjID, questionUserID, answerUserID string) (err error)
}

// AnswerActivityService user service
type AnswerActivityService struct {
	answerActivityRepo AnswerActivityRepo
}

// NewAnswerActivityService new comment service
func NewAnswerActivityService(
	answerActivityRepo AnswerActivityRepo) *AnswerActivityService {
	return &AnswerActivityService{
		answerActivityRepo: answerActivityRepo,
	}
}

// AcceptAnswer accept answer change activity
func (as *AnswerActivityService) AcceptAnswer(ctx context.Context,
	answerObjID, questionObjID, questionUserID, answerUserID string, isSelf bool) (err error) {
	return as.answerActivityRepo.AcceptAnswer(ctx, answerObjID, questionObjID, questionUserID, answerUserID, isSelf)
}

// CancelAcceptAnswer cancel accept answer change activity
func (as *AnswerActivityService) CancelAcceptAnswer(ctx context.Context,
	answerObjID, questionObjID, questionUserID, answerUserID string) (err error) {
	return as.answerActivityRepo.CancelAcceptAnswer(ctx, answerObjID, questionObjID, questionUserID, answerUserID)
}
