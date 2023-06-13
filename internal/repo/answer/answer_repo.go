package answer

import (
	"context"
	"strings"
	"time"
	"unicode"

	"xorm.io/builder"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/answerdev/answer/pkg/uid"
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
	answer.QuestionID = uid.DeShortID(answer.QuestionID)
	ID, err := ar.uniqueIDRepo.GenUniqueIDStr(ctx, answer.TableName())
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	answer.ID = ID
	_, err = ar.data.DB.Context(ctx).Insert(answer)

	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	answer.ID = uid.EnShortID(answer.ID)
	answer.QuestionID = uid.EnShortID(answer.QuestionID)
	return nil
}

// RemoveAnswer delete answer
func (ar *answerRepo) RemoveAnswer(ctx context.Context, id string) (err error) {
	id = uid.DeShortID(id)
	answer := &entity.Answer{
		ID:     id,
		Status: entity.AnswerStatusDeleted,
	}
	_, err = ar.data.DB.Context(ctx).Where("id = ?", id).Cols("status").Update(answer)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// UpdateAnswer update answer
func (ar *answerRepo) UpdateAnswer(ctx context.Context, answer *entity.Answer, Colar []string) (err error) {
	answer.ID = uid.DeShortID(answer.ID)
	answer.QuestionID = uid.DeShortID(answer.QuestionID)
	_, err = ar.data.DB.Context(ctx).ID(answer.ID).Cols(Colar...).Update(answer)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return err
}

func (ar *answerRepo) UpdateAnswerStatus(ctx context.Context, answer *entity.Answer) (err error) {
	now := time.Now()
	answer.ID = uid.DeShortID(answer.ID)
	answer.UpdatedAt = now
	_, err = ar.data.DB.Context(ctx).Where("id =?", answer.ID).Cols("status", "updated_at").Update(answer)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetAnswer get answer one
func (ar *answerRepo) GetAnswer(ctx context.Context, id string) (
	answer *entity.Answer, exist bool, err error,
) {
	id = uid.DeShortID(id)
	answer = &entity.Answer{}
	exist, err = ar.data.DB.Context(ctx).ID(id).Get(answer)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	answer.ID = uid.EnShortID(answer.ID)
	answer.QuestionID = uid.EnShortID(answer.QuestionID)

	return
}

// GetQuestionCount
func (ar *answerRepo) GetAnswerCount(ctx context.Context) (count int64, err error) {
	list := make([]*entity.Answer, 0)
	count, err = ar.data.DB.Context(ctx).Where("status = ?", entity.AnswerStatusAvailable).FindAndCount(&list)
	if err != nil {
		return count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetAnswerList get answer list all
func (ar *answerRepo) GetAnswerList(ctx context.Context, answer *entity.Answer) (answerList []*entity.Answer, err error) {
	answerList = make([]*entity.Answer, 0)
	answer.ID = uid.DeShortID(answer.ID)
	answer.QuestionID = uid.DeShortID(answer.QuestionID)
	err = ar.data.DB.Context(ctx).Find(answerList, answer)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	for _, item := range answerList {
		item.ID = uid.EnShortID(item.ID)
		item.QuestionID = uid.EnShortID(item.QuestionID)
	}
	return
}

// GetAnswerPage get answer page
func (ar *answerRepo) GetAnswerPage(ctx context.Context, page, pageSize int, answer *entity.Answer) (answerList []*entity.Answer, total int64, err error) {
	answer.ID = uid.DeShortID(answer.ID)
	answer.QuestionID = uid.DeShortID(answer.QuestionID)
	answerList = make([]*entity.Answer, 0)
	total, err = pager.Help(page, pageSize, answerList, answer, ar.data.DB.Context(ctx))
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	for _, item := range answerList {
		item.ID = uid.EnShortID(item.ID)
		item.QuestionID = uid.EnShortID(item.QuestionID)
	}
	return
}

// UpdateAccepted
// If no answer is selected, the answer id can be 0
func (ar *answerRepo) UpdateAccepted(ctx context.Context, id string, questionID string) error {
	if questionID == "" {
		return nil
	}
	id = uid.DeShortID(id)
	questionID = uid.DeShortID(questionID)
	var data entity.Answer
	data.ID = id

	data.Accepted = schema.AnswerAcceptedFailed
	_, err := ar.data.DB.Context(ctx).Where("question_id =?", questionID).Cols("adopted").Update(&data)
	if err != nil {
		return err
	}
	if id != "0" {
		data.Accepted = schema.AnswerAcceptedEnable
		_, err = ar.data.DB.Context(ctx).Where("id = ?", id).Cols("adopted").Update(&data)
		if err != nil {
			return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		}
	}
	return nil
}

// GetByID
func (ar *answerRepo) GetByID(ctx context.Context, id string) (*entity.Answer, bool, error) {
	var resp entity.Answer
	id = uid.DeShortID(id)
	has, err := ar.data.DB.Context(ctx).Where("id =? ", id).Get(&resp)
	if err != nil {
		return &resp, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	resp.ID = uid.EnShortID(resp.ID)
	resp.QuestionID = uid.EnShortID(resp.QuestionID)
	return &resp, has, nil
}

func (ar *answerRepo) GetCountByQuestionID(ctx context.Context, questionID string) (int64, error) {
	questionID = uid.DeShortID(questionID)
	rows := make([]*entity.Answer, 0)
	count, err := ar.data.DB.Context(ctx).Where("question_id =? and  status = ?", questionID, entity.AnswerStatusAvailable).FindAndCount(&rows)
	if err != nil {
		return count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return count, nil
}

func (ar *answerRepo) GetCountByUserID(ctx context.Context, userID string) (int64, error) {
	rows := make([]*entity.Answer, 0)
	count, err := ar.data.DB.Context(ctx).Where(" user_id = ?  and  status = ?", userID, entity.AnswerStatusAvailable).FindAndCount(&rows)
	if err != nil {
		return count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return count, nil
}

func (ar *answerRepo) GetByUserIDQuestionID(ctx context.Context, userID string, questionID string) (*entity.Answer, bool, error) {
	questionID = uid.DeShortID(questionID)
	var resp entity.Answer
	has, err := ar.data.DB.Context(ctx).Where("question_id =? and  user_id = ? and status = ?", questionID, userID, entity.AnswerStatusAvailable).Get(&resp)
	if err != nil {
		return &resp, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	resp.ID = uid.EnShortID(resp.ID)
	resp.QuestionID = uid.EnShortID(resp.QuestionID)
	return &resp, has, nil
}

// SearchList
func (ar *answerRepo) SearchList(ctx context.Context, search *entity.AnswerSearch) ([]*entity.Answer, int64, error) {
	if search.QuestionID != "" {
		search.QuestionID = uid.DeShortID(search.QuestionID)
	}
	search.ID = uid.DeShortID(search.ID)
	var count int64
	var err error
	rows := make([]*entity.Answer, 0)
	if search.Page > 0 {
		search.Page = search.Page - 1
	} else {
		search.Page = 0
	}
	if search.PageSize == 0 {
		search.PageSize = constant.DefaultPageSize
	}
	offset := search.Page * search.PageSize
	session := ar.data.DB.Context(ctx).Where("")

	if search.QuestionID != "" {
		session = session.And("question_id = ?", search.QuestionID)
	}
	if len(search.UserID) > 0 {
		session = session.And("user_id = ?", search.UserID)
	}
	switch search.Order {
	case entity.AnswerSearchOrderByTime:
		session = session.OrderBy("created_at desc")
	case entity.AnswerSearchOrderByVote:
		session = session.OrderBy("vote_count desc")
	default:
		session = session.OrderBy("adopted desc,vote_count desc,created_at asc")
	}
	if !search.IncludeDeleted {
		if search.LoginUserID == "" {
			session = session.And("status = ? ", entity.AnswerStatusAvailable)
		} else {
			session = session.And("status = ? OR user_id = ?", entity.AnswerStatusAvailable, search.LoginUserID)
		}
	}

	session = session.Limit(search.PageSize, offset)
	count, err = session.FindAndCount(&rows)
	if err != nil {
		return rows, count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	for _, item := range rows {
		item.ID = uid.EnShortID(item.ID)
		item.QuestionID = uid.EnShortID(item.QuestionID)
	}
	return rows, count, nil
}

func (ar *answerRepo) AdminSearchList(ctx context.Context, search *entity.AdminAnswerSearch) ([]*entity.Answer, int64, error) {
	var (
		count   int64
		err     error
		session = ar.data.DB.Context(ctx).Table([]string{entity.Answer{}.TableName(), "a"}).Select("a.*")
	)
	if search.QuestionID != "" {
		search.QuestionID = uid.DeShortID(search.QuestionID)
	}

	session.Where(builder.Eq{
		"a.status": search.Status,
	})

	rows := make([]*entity.Answer, 0)
	if search.Page > 0 {
		search.Page = search.Page - 1
	} else {
		search.Page = 0
	}
	if search.PageSize == 0 {
		search.PageSize = constant.DefaultPageSize
	}

	// search by question title like or answer id
	if len(search.Query) > 0 {
		// check id search
		var (
			idSearch = false
			id       = ""
		)

		if strings.Contains(search.Query, "answer:") {
			idSearch = true
			id = strings.TrimSpace(strings.TrimPrefix(search.Query, "answer:"))
			id = uid.DeShortID(id)
			for _, r := range id {
				if !unicode.IsDigit(r) {
					idSearch = false
					break
				}
			}
		}

		if idSearch {
			session.And(builder.Eq{
				"id": id,
			})
		} else {
			session.Join("LEFT", []string{entity.Question{}.TableName(), "q"}, "q.id = a.question_id")
			session.And(builder.Like{
				"q.title", search.Query,
			})
		}
	}

	// check search by question id
	if len(search.QuestionID) > 0 {
		session.And(builder.Eq{
			"question_id": search.QuestionID,
		})
	}

	offset := search.Page * search.PageSize
	session.
		OrderBy("a.created_at desc").
		Limit(search.PageSize, offset)
	count, err = session.FindAndCount(&rows)
	if err != nil {
		return rows, count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	for _, item := range rows {
		item.ID = uid.EnShortID(item.ID)
		item.QuestionID = uid.EnShortID(item.QuestionID)
	}
	return rows, count, nil
}
