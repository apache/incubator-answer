package activity

import (
	"context"
	"time"

	"github.com/segmentfault/pacman/log"
)

// AnswerActivityRepo answer activity
type AnswerActivityRepo interface {
	AcceptAnswer(ctx context.Context,
		answerObjID, questionObjID, questionUserID, answerUserID string, isSelf bool) (err error)
	CancelAcceptAnswer(ctx context.Context,
		answerObjID, questionObjID, questionUserID, answerUserID string) (err error)
	DeleteAnswer(ctx context.Context, answerID string) (err error)
}

// QuestionActivityRepo answer activity
type QuestionActivityRepo interface {
	DeleteQuestion(ctx context.Context, questionID string) (err error)
}

// AnswerActivityService user service
type AnswerActivityService struct {
	answerActivityRepo   AnswerActivityRepo
	questionActivityRepo QuestionActivityRepo
}

// NewAnswerActivityService new comment service
func NewAnswerActivityService(
	answerActivityRepo AnswerActivityRepo, questionActivityRepo QuestionActivityRepo) *AnswerActivityService {
	return &AnswerActivityService{
		answerActivityRepo:   answerActivityRepo,
		questionActivityRepo: questionActivityRepo,
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

// DeleteAnswer delete answer change activity
func (as *AnswerActivityService) DeleteAnswer(ctx context.Context, answerID string, createdAt time.Time,
	voteCount int) (err error) {
	if voteCount >= 3 {
		log.Infof("There is no need to roll back the reputation by answering likes above the target value. %s %d", answerID, voteCount)
		return nil
	}
	if createdAt.Before(time.Now().AddDate(0, 0, -60)) {
		log.Infof("There is no need to roll back the reputation by answer's existence time meets the target. %s %s", answerID, createdAt.String())
		return nil
	}
	return as.answerActivityRepo.DeleteAnswer(ctx, answerID)
}

// DeleteQuestion delete question change activity
func (as *AnswerActivityService) DeleteQuestion(ctx context.Context, questionID string, createdAt time.Time,
	voteCount int) (err error) {
	if voteCount >= 3 {
		log.Infof("There is no need to roll back the reputation by answering likes above the target value. %s %d", questionID, voteCount)
		return nil
	}
	if createdAt.Before(time.Now().AddDate(0, 0, -60)) {
		log.Infof("There is no need to roll back the reputation by answer's existence time meets the target. %s %s", questionID, createdAt.String())
		return nil
	}
	return as.questionActivityRepo.DeleteQuestion(ctx, questionID)
}
