package question

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
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/answerdev/answer/pkg/htmltext"

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
	_, err = qr.data.DB.Insert(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// RemoveQuestion delete question
func (qr *questionRepo) RemoveQuestion(ctx context.Context, id string) (err error) {
	_, err = qr.data.DB.Where("id =?", id).Delete(&entity.Question{})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateQuestion update question
func (qr *questionRepo) UpdateQuestion(ctx context.Context, question *entity.Question, Cols []string) (err error) {
	_, err = qr.data.DB.Where("id =?", question.ID).Cols(Cols...).Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (qr *questionRepo) UpdatePvCount(ctx context.Context, questionID string) (err error) {
	question := &entity.Question{}
	_, err = qr.data.DB.Where("id =?", questionID).Incr("view_count", 1).Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateAnswerCount(ctx context.Context, questionID string, num int) (err error) {
	question := &entity.Question{}
	_, err = qr.data.DB.Where("id =?", questionID).Incr("answer_count", num).Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateCollectionCount(ctx context.Context, questionID string, num int) (err error) {
	question := &entity.Question{}
	_, err = qr.data.DB.Where("id =?", questionID).Incr("collection_count", num).Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateQuestionStatus(ctx context.Context, question *entity.Question) (err error) {
	now := time.Now()
	question.UpdatedAt = now
	_, err = qr.data.DB.Where("id =?", question.ID).Cols("status", "updated_at").Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateAccepted(ctx context.Context, question *entity.Question) (err error) {
	_, err = qr.data.DB.Where("id =?", question.ID).Cols("accepted_answer_id").Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (qr *questionRepo) UpdateLastAnswer(ctx context.Context, question *entity.Question) (err error) {
	_, err = qr.data.DB.Where("id =?", question.ID).Cols("last_answer_id").Update(question)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// GetQuestion get question one
func (qr *questionRepo) GetQuestion(ctx context.Context, id string) (
	question *entity.Question, exist bool, err error,
) {
	question = &entity.Question{}
	question.ID = id
	exist, err = qr.data.DB.Where("id = ?", id).Get(question)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetTagBySlugName get tag by slug name
func (qr *questionRepo) SearchByTitleLike(ctx context.Context, title string) (questionList []*entity.Question, err error) {
	questionList = make([]*entity.Question, 0)
	err = qr.data.DB.Table("question").Where("title like ?", "%"+title+"%").Limit(10, 0).Find(&questionList)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (qr *questionRepo) FindByID(ctx context.Context, id []string) (questionList []*entity.Question, err error) {
	questionList = make([]*entity.Question, 0)
	err = qr.data.DB.Table("question").In("id", id).Find(&questionList)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetQuestionList get question list all
func (qr *questionRepo) GetQuestionList(ctx context.Context, question *entity.Question) (questionList []*entity.Question, err error) {
	questionList = make([]*entity.Question, 0)
	err = qr.data.DB.Find(questionList, question)
	if err != nil {
		return questionList, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (qr *questionRepo) GetQuestionCount(ctx context.Context) (count int64, err error) {
	questionList := make([]*entity.Question, 0)

	count, err = qr.data.DB.In("question.status", []int{entity.QuestionStatusAvailable, entity.QuestionStatusClosed}).FindAndCount(&questionList)
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
	session := qr.data.DB.Table("question")
	session = session.In("question.status", []int{entity.QuestionStatusAvailable, entity.QuestionStatusClosed})
	session = session.Limit(pageSize, offset)
	session = session.OrderBy("question.created_at asc")
	err = session.Select("id,title,post_update_time").Find(&rows)
	if err != nil {
		return questionIDList, err
	}
	for _, question := range rows {
		item := &schema.SiteMapQuestionInfo{}
		item.ID = question.ID
		item.Title = htmltext.UrlTitle(question.Title)
		item.UpdateTime = question.PostUpdateTime.Format("2006-01-02 15:04:05")
		questionIDList = append(questionIDList, item)
	}
	return questionIDList, nil
}

// GetQuestionPage get question page
func (qr *questionRepo) GetQuestionPage(ctx context.Context, page, pageSize int, question *entity.Question) (questionList []*entity.Question, total int64, err error) {
	questionList = make([]*entity.Question, 0)
	total, err = pager.Help(page, pageSize, questionList, question, qr.data.DB.NewSession())
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// SearchList
func (qr *questionRepo) SearchList(ctx context.Context, search *schema.QuestionSearch) ([]*entity.QuestionTag, int64, error) {
	var count int64
	var err error
	rows := make([]*entity.QuestionTag, 0)
	if search.Page > 0 {
		search.Page = search.Page - 1
	} else {
		search.Page = 0
	}
	if search.PageSize == 0 {
		search.PageSize = constant.DefaultPageSize
	}
	offset := search.Page * search.PageSize
	session := qr.data.DB.Table("question")

	if len(search.TagIDs) > 0 {
		session = session.Join("LEFT", "tag_rel", "question.id = tag_rel.object_id")
		session = session.And("tag_rel.tag_id =?", search.TagIDs[0])
		// session = session.In("tag_rel.tag_id ", search.TagIDs)
		session = session.And("tag_rel.status =?", entity.TagRelStatusAvailable)
	}

	if len(search.UserID) > 0 {
		session = session.And("question.user_id = ?", search.UserID)
	}

	session = session.In("question.status", []int{entity.QuestionStatusAvailable, entity.QuestionStatusClosed})
	// if search.Status > 0 {
	// 	session = session.And("question.status = ?", search.Status)
	// }
	// switch
	// newest, active,frequent,score,unanswered
	switch search.Order {
	case "newest":
		session = session.OrderBy("question.created_at desc")
	case "active":
		session = session.OrderBy("question.post_update_time desc,question.updated_at desc")
	case "frequent":
		session = session.OrderBy("question.view_count desc")
	case "score":
		session = session.OrderBy("question.vote_count desc,question.view_count desc")
	case "unanswered":
		session = session.And("question.last_answer_id = 0")
		session = session.OrderBy("question.created_at desc")
	}
	session = session.Limit(search.PageSize, offset)
	session = session.Select("question.id,question.user_id,last_edit_user_id,question.title,question.original_text,question.parsed_text,question.status,question.view_count,question.unique_view_count,question.vote_count,question.answer_count,question.collection_count,question.follow_count,question.accepted_answer_id,question.last_answer_id,question.created_at,question.updated_at,question.post_update_time,question.revision_id")
	count, err = session.FindAndCount(&rows)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return rows, count, err
	}
	return rows, count, nil
}

func (qr *questionRepo) CmsSearchList(ctx context.Context, search *schema.CmsQuestionSearch) ([]*entity.Question, int64, error) {
	var (
		count   int64
		err     error
		session = qr.data.DB.Table("question")
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

	session.OrderBy("updated_at desc").
		Limit(search.PageSize, offset)
	count, err = session.FindAndCount(&rows)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return rows, count, err
	}
	return rows, count, nil
}
