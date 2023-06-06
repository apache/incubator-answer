package question

import (
	"context"
	"fmt"
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
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/answerdev/answer/pkg/htmltext"
	"github.com/answerdev/answer/pkg/uid"

	"github.com/segmentfault/pacman/errors"
)

// questionRepo question repository
type questionRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

// NewQuestionRepo new repository
func NewQuestionRepo(
	data *data.Data,
	uniqueIDRepo unique.UniqueIDRepo,
) questioncommon.QuestionRepo {
	return &questionRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

// AddQuestion add question
func (qr *questionRepo) AddQuestion(ctx context.Context, question *entity.Question) (err error) {
	question.ID, err = qr.uniqueIDRepo.GenUniqueIDStr(ctx, question.TableName())
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	_, err = qr.data.DB.Context(ctx).Insert(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	question.ID = uid.EnShortID(question.ID)
	return
}

// RemoveQuestion delete question
func (qr *questionRepo) RemoveQuestion(ctx context.Context, id string) (err error) {
	id = uid.DeShortID(id)
	_, err = qr.data.DB.Context(ctx).Where("id =?", id).Delete(&entity.Question{})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateQuestion update question
func (qr *questionRepo) UpdateQuestion(ctx context.Context, question *entity.Question, Cols []string) (err error) {
	question.ID = uid.DeShortID(question.ID)
	_, err = qr.data.DB.Context(ctx).Where("id =?", question.ID).Cols(Cols...).Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	question.ID = uid.EnShortID(question.ID)
	return
}

func (qr *questionRepo) UpdatePvCount(ctx context.Context, questionID string) (err error) {
	questionID = uid.DeShortID(questionID)
	question := &entity.Question{}
	_, err = qr.data.DB.Context(ctx).Where("id =?", questionID).Incr("view_count", 1).Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateAnswerCount(ctx context.Context, questionID string, num int) (err error) {
	questionID = uid.DeShortID(questionID)
	question := &entity.Question{}
	question.AnswerCount = num
	_, err = qr.data.DB.Context(ctx).Where("id =?", questionID).Cols("answer_count").Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateCollectionCount(ctx context.Context, questionID string, num int) (err error) {
	questionID = uid.DeShortID(questionID)
	question := &entity.Question{}
	_, err = qr.data.DB.Context(ctx).Where("id =?", questionID).Incr("collection_count", num).Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateQuestionStatus(ctx context.Context, question *entity.Question) (err error) {
	question.ID = uid.DeShortID(question.ID)
	now := time.Now()
	question.UpdatedAt = now
	_, err = qr.data.DB.Context(ctx).Where("id =?", question.ID).Cols("status", "updated_at").Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateQuestionStatusWithOutUpdateTime(ctx context.Context, question *entity.Question) (err error) {
	question.ID = uid.DeShortID(question.ID)
	_, err = qr.data.DB.Context(ctx).Where("id =?", question.ID).Cols("status").Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateQuestionOperation(ctx context.Context, question *entity.Question) (err error) {
	question.ID = uid.DeShortID(question.ID)
	_, err = qr.data.DB.Context(ctx).Where("id =?", question.ID).Cols("pin", "show").Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateAccepted(ctx context.Context, question *entity.Question) (err error) {
	question.ID = uid.DeShortID(question.ID)
	_, err = qr.data.DB.Context(ctx).Where("id =?", question.ID).Cols("accepted_answer_id").Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateLastAnswer(ctx context.Context, question *entity.Question) (err error) {
	question.ID = uid.DeShortID(question.ID)
	_, err = qr.data.DB.Context(ctx).Where("id =?", question.ID).Cols("last_answer_id").Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// GetQuestion get question one
func (qr *questionRepo) GetQuestion(ctx context.Context, id string) (
	question *entity.Question, exist bool, err error,
) {
	id = uid.DeShortID(id)
	question = &entity.Question{}
	question.ID = id
	exist, err = qr.data.DB.Context(ctx).Where("id = ?", id).Get(question)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	question.ID = uid.EnShortID(question.ID)
	return
}

// GetTagBySlugName get tag by slug name
func (qr *questionRepo) SearchByTitleLike(ctx context.Context, title string) (questionList []*entity.Question, err error) {
	questionList = make([]*entity.Question, 0)
	err = qr.data.DB.Context(ctx).Table("question").Where("title like ?", "%"+title+"%").Limit(10, 0).Find(&questionList)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	for _, item := range questionList {
		item.ID = uid.EnShortID(item.ID)
	}
	return
}

func (qr *questionRepo) FindByID(ctx context.Context, id []string) (questionList []*entity.Question, err error) {
	for key, itemID := range id {
		id[key] = uid.DeShortID(itemID)
	}
	questionList = make([]*entity.Question, 0)
	err = qr.data.DB.Context(ctx).Table("question").In("id", id).Find(&questionList)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	for _, item := range questionList {
		item.ID = uid.EnShortID(item.ID)
	}
	return
}

// GetQuestionList get question list all
func (qr *questionRepo) GetQuestionList(ctx context.Context, question *entity.Question) (questionList []*entity.Question, err error) {
	question.ID = uid.DeShortID(question.ID)
	questionList = make([]*entity.Question, 0)
	err = qr.data.DB.Context(ctx).Find(questionList, question)
	if err != nil {
		return questionList, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	for _, item := range questionList {
		item.ID = uid.DeShortID(item.ID)
	}
	return
}

func (qr *questionRepo) GetQuestionCount(ctx context.Context) (count int64, err error) {
	questionList := make([]*entity.Question, 0)

	count, err = qr.data.DB.Context(ctx).In("question.status", []int{entity.QuestionStatusAvailable, entity.QuestionStatusClosed}).FindAndCount(&questionList)
	if err != nil {
		return count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (qr *questionRepo) GetUserQuestionCount(ctx context.Context, userID string) (count int64, err error) {
	questionList := make([]*entity.Question, 0)
	count, err = qr.data.DB.In("question.status", []int{entity.QuestionStatusAvailable, entity.QuestionStatusClosed}).And("question.user_id = ?", userID).FindAndCount(&questionList)
	if err != nil {
		return count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (qr *questionRepo) GetQuestionCountByIDs(ctx context.Context, ids []string) (count int64, err error) {
	questionList := make([]*entity.Question, 0)
	count, err = qr.data.DB.In("question.status", []int{entity.QuestionStatusAvailable, entity.QuestionStatusClosed}).In("id = ?", ids).Count(&questionList)
	if err != nil {
		return count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (qr *questionRepo) GetQuestionIDsPage(ctx context.Context, page, pageSize int) (questionIDList []*schema.SiteMapQuestionInfo, err error) {
	questionIDList = make([]*schema.SiteMapQuestionInfo, 0)
	rows := make([]*entity.Question, 0)
	if page > 0 {
		page = page - 1
	} else {
		page = 0
	}
	if pageSize == 0 {
		pageSize = constant.DefaultPageSize
	}
	offset := page * pageSize
	session := qr.data.DB.Context(ctx).Table("question")
	session = session.In("question.status", []int{entity.QuestionStatusAvailable, entity.QuestionStatusClosed})
	session.And("question.show = ?", entity.QuestionShow)
	session = session.Limit(pageSize, offset)
	session = session.OrderBy("question.created_at asc")
	err = session.Select("id,title,created_at,post_update_time").Find(&rows)
	if err != nil {
		return questionIDList, err
	}
	for _, question := range rows {
		item := &schema.SiteMapQuestionInfo{}
		item.ID = uid.EnShortID(question.ID)
		item.Title = htmltext.UrlTitle(question.Title)
		updateTime := fmt.Sprintf("%v", question.PostUpdateTime.Format(time.RFC3339))
		if question.PostUpdateTime.Unix() < 1 {
			updateTime = fmt.Sprintf("%v", question.CreatedAt.Format(time.RFC3339))
		}
		item.UpdateTime = updateTime
		questionIDList = append(questionIDList, item)
	}
	return questionIDList, nil
}

// GetQuestionPage query question page
func (qr *questionRepo) GetQuestionPage(ctx context.Context, page, pageSize int, userID, tagID, orderCond string, inDays int) (
	questionList []*entity.Question, total int64, err error) {
	questionList = make([]*entity.Question, 0)

	session := qr.data.DB.Context(ctx).Where("question.status = ? OR question.status = ?",
		entity.QuestionStatusAvailable, entity.QuestionStatusClosed)
	if len(tagID) > 0 {
		session.Join("LEFT", "tag_rel", "question.id = tag_rel.object_id")
		session.And("tag_rel.tag_id = ?", tagID)
		session.And("tag_rel.status = ?", entity.TagRelStatusAvailable)
	}
	if len(userID) > 0 {
		session.And("question.user_id = ?", userID)
	} else {
		session.And("question.show = ?", entity.QuestionShow)
	}
	if inDays > 0 {
		session.And("question.created_at > ?", time.Now().AddDate(0, 0, -inDays))
	}

	switch orderCond {
	case "newest":
		session.OrderBy("question.pin desc,question.created_at DESC")
	case "active":
		session.OrderBy("question.pin desc,question.post_update_time DESC, question.updated_at DESC")
	case "frequent":
		session.OrderBy("question.pin desc,question.view_count DESC")
	case "score":
		session.OrderBy("question.pin desc,question.vote_count DESC, question.view_count DESC")
	case "unanswered":
		session.Where("question.last_answer_id = 0")
		session.OrderBy("question.pin desc,question.created_at DESC")
	}

	total, err = pager.Help(page, pageSize, &questionList, &entity.Question{}, session)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	for _, item := range questionList {
		item.ID = uid.EnShortID(item.ID)
	}
	return questionList, total, err
}

func (qr *questionRepo) AdminSearchList(ctx context.Context, search *schema.AdminQuestionSearch) ([]*entity.Question, int64, error) {
	var (
		count   int64
		err     error
		session = qr.data.DB.Context(ctx).Table("question")
	)

	session.Where(builder.Eq{
		"status": search.Status,
	})

	rows := make([]*entity.Question, 0)
	if search.Page > 0 {
		search.Page = search.Page - 1
	} else {
		search.Page = 0
	}
	if search.PageSize == 0 {
		search.PageSize = constant.DefaultPageSize
	}

	// search by question title like or question id
	if len(search.Query) > 0 {
		// check id search
		var (
			idSearch = false
			id       = ""
		)

		if strings.Contains(search.Query, "question:") {
			idSearch = true
			id = strings.TrimSpace(strings.TrimPrefix(search.Query, "question:"))
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
			session.And(builder.Like{
				"title", search.Query,
			})
		}
	}

	offset := search.Page * search.PageSize

	session.OrderBy("created_at desc").
		Limit(search.PageSize, offset)
	count, err = session.FindAndCount(&rows)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return rows, count, err
	}
	for _, item := range rows {
		item.ID = uid.EnShortID(item.ID)
	}
	return rows, count, nil
}
