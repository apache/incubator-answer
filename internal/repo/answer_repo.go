package repo

import (
	"context"
	"time"

	"github.com/segmentfault/answer/internal/base/constant"
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/pager"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/activity_common"
	answercommon "github.com/segmentfault/answer/internal/service/answer_common"
	"github.com/segmentfault/answer/internal/service/rank"
	"github.com/segmentfault/answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
)

// answerRepo answer repository
type answerRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
	userRankRepo rank.UserRankRepo
	activityRepo activity_common.ActivityRepo
}

// NewAnswerRepo new repository
func NewAnswerRepo(
	data *data.Data,
	uniqueIDRepo unique.UniqueIDRepo,
	userRankRepo rank.UserRankRepo,
	activityRepo activity_common.ActivityRepo,
) answercommon.AnswerRepo {
	return &answerRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
		userRankRepo: userRankRepo,
		activityRepo: activityRepo,
	}
}

// AddAnswer add answer
func (ar *answerRepo) AddAnswer(ctx context.Context, answer *entity.Answer) (err error) {
	ID, err := ar.uniqueIDRepo.GenUniqueIDStr(ctx, answer.TableName())
	if err != nil {
		errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	answer.ID = ID
	_, err = ar.data.DB.Insert(answer)

	if err != nil {
		errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// RemoveAnswer delete answer
func (ar *answerRepo) RemoveAnswer(ctx context.Context, id string) (err error) {
	answer := &entity.Answer{
		ID:     id,
		Status: entity.AnswerStatusDeleted,
	}
	_, err = ar.data.DB.Where("id = ?", id).Cols("status").Update(answer)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// UpdateAnswer update answer
func (ar *answerRepo) UpdateAnswer(ctx context.Context, answer *entity.Answer, Colar []string) (err error) {
	_, err = ar.data.DB.ID(answer.ID).Cols(Colar...).Update(answer)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return err
}

func (ar *answerRepo) UpdateAnswerStatus(ctx context.Context, answer *entity.Answer) (err error) {
	now := time.Now()
	answer.UpdatedAt = now
	_, err = ar.data.DB.Where("id =?", answer.ID).Cols("status", "updated_at").Update(answer)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetAnswer get answer one
func (ar *answerRepo) GetAnswer(ctx context.Context, id string) (
	answer *entity.Answer, exist bool, err error) {
	answer = &entity.Answer{}
	exist, err = ar.data.DB.ID(id).Get(answer)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetAnswerList get answer list all
func (ar *answerRepo) GetAnswerList(ctx context.Context, answer *entity.Answer) (answerList []*entity.Answer, err error) {
	answerList = make([]*entity.Answer, 0)
	err = ar.data.DB.Find(answerList, answer)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetAnswerPage get answer page
func (ar *answerRepo) GetAnswerPage(ctx context.Context, page, pageSize int, answer *entity.Answer) (answerList []*entity.Answer, total int64, err error) {
	answerList = make([]*entity.Answer, 0)
	total, err = pager.Help(page, pageSize, answerList, answer, ar.data.DB.NewSession())
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateAdopted
// If no answer is selected, the answer id can be 0
func (ar *answerRepo) UpdateAdopted(ctx context.Context, id string, questionId string) error {
	if questionId == "" {
		return nil
	}
	var data entity.Answer
	data.ID = id

	data.Adopted = schema.Answer_Adopted_Failed
	_, err := ar.data.DB.Where("question_id =?", questionId).Cols("adopted").Update(&data)
	if err != nil {
		return err
	}
	if id != "0" {
		data.Adopted = schema.Answer_Adopted_Enable
		_, err = ar.data.DB.Where("id = ?", id).Cols("adopted").Update(&data)
		if err != nil {
			return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		}
	}
	return nil
}

// GetByID
func (ar *answerRepo) GetByID(ctx context.Context, id string) (*entity.Answer, bool, error) {
	var resp entity.Answer
	has, err := ar.data.DB.Where("id =? ", id).Get(&resp)
	if err != nil {
		return &resp, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return &resp, has, nil
}

func (ar *answerRepo) GetByUserIdQuestionId(ctx context.Context, userId string, questionId string) (*entity.Answer, bool, error) {
	var resp entity.Answer
	has, err := ar.data.DB.Where("question_id =? and  user_id = ?", questionId, userId).Get(&resp)
	if err != nil {
		return &resp, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return &resp, has, nil
}

// SearchList
func (ar *answerRepo) SearchList(ctx context.Context, search *entity.AnswerSearch) ([]*entity.Answer, int64, error) {
	var count int64
	var err error
	rows := make([]*entity.Answer, 0)
	if search.Page > 0 {
		search.Page = search.Page - 1
	} else {
		search.Page = 0
	}
	if search.PageSize == 0 {
		search.PageSize = constant.Default_PageSize
	}
	offset := search.Page * search.PageSize
	session := ar.data.DB.Where("")

	if search.QuestionID != "" {
		session = session.And("question_id = ?", search.QuestionID)
	}
	if len(search.UserID) > 0 {
		session = session.And("user_id = ?", search.UserID)
	}
	if search.Order == entity.Answer_Search_OrderBy_Time {
		session = session.OrderBy("created_at desc")
	} else if search.Order == entity.Answer_Search_OrderBy_Vote {
		session = session.OrderBy("vote_count desc")
	} else {
		session = session.OrderBy("adopted desc,vote_count desc")
	}

	session = session.And("status = ?", entity.AnswerStatusAvailable)

	session = session.Limit(search.PageSize, offset)
	count, err = session.FindAndCount(&rows)
	if err != nil {
		return rows, count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return rows, count, nil
}

func (ar *answerRepo) CmsSearchList(ctx context.Context, search *entity.CmsAnswerSearch) ([]*entity.Answer, int64, error) {
	var count int64
	var err error
	if search.Status == 0 {
		search.Status = 1
	}
	rows := make([]*entity.Answer, 0)
	if search.Page > 0 {
		search.Page = search.Page - 1
	} else {
		search.Page = 0
	}
	if search.PageSize == 0 {
		search.PageSize = constant.Default_PageSize
	}
	offset := search.Page * search.PageSize
	session := ar.data.DB.Where("")
	session = session.And("status =?", search.Status)
	session = session.OrderBy("updated_at desc")
	session = session.Limit(search.PageSize, offset)
	count, err = session.FindAndCount(&rows)
	if err != nil {
		return rows, count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return rows, count, nil
}
