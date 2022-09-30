package questioncommon

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/service/activity_common"
	"github.com/segmentfault/answer/internal/service/config"
	"github.com/segmentfault/answer/internal/service/meta"
	"github.com/segmentfault/pacman/errors"

	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	answercommon "github.com/segmentfault/answer/internal/service/answer_common"
	collectioncommon "github.com/segmentfault/answer/internal/service/collection_common"
	tagcommon "github.com/segmentfault/answer/internal/service/tag_common"
	usercommon "github.com/segmentfault/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/log"
)

// QuestionRepo question repository
type QuestionRepo interface {
	AddQuestion(ctx context.Context, question *entity.Question) (err error)
	RemoveQuestion(ctx context.Context, id string) (err error)
	UpdateQuestion(ctx context.Context, question *entity.Question, Cols []string) (err error)
	GetQuestion(ctx context.Context, id string) (question *entity.Question, exist bool, err error)
	GetQuestionList(ctx context.Context, question *entity.Question) (questions []*entity.Question, err error)
	GetQuestionPage(ctx context.Context, page, pageSize int, question *entity.Question) (questions []*entity.Question, total int64, err error)
	SearchList(ctx context.Context, search *schema.QuestionSearch) ([]*entity.QuestionTag, int64, error)
	UpdateQuestionStatus(ctx context.Context, question *entity.Question) (err error)
	SearchByTitleLike(ctx context.Context, title string) (questionList []*entity.Question, err error)
	UpdatePvCount(ctx context.Context, questionId string) (err error)
	UpdateAnswerCount(ctx context.Context, questionId string, num int) (err error)
	UpdateCollectionCount(ctx context.Context, questionId string, num int) (err error)
	UpdateAccepted(ctx context.Context, question *entity.Question) (err error)
	UpdateLastAnswer(ctx context.Context, question *entity.Question) (err error)
	FindByID(ctx context.Context, id []string) (questionList []*entity.Question, err error)
	CmsSearchList(ctx context.Context, search *schema.CmsQuestionSearch) ([]*entity.Question, int64, error)
}

// QuestionCommon user service
type QuestionCommon struct {
	questionRepo     QuestionRepo
	answerRepo       answercommon.AnswerRepo
	voteRepo         activity_common.VoteRepo
	followCommon     activity_common.FollowRepo
	tagCommon        *tagcommon.TagCommonService
	userCommon       *usercommon.UserCommon
	collectionCommon *collectioncommon.CollectionCommon
	AnswerCommon     *answercommon.AnswerCommon
	metaService      *meta.MetaService
	configRepo       config.ConfigRepo
}

func NewQuestionCommon(questionRepo QuestionRepo,
	answerRepo answercommon.AnswerRepo,
	voteRepo activity_common.VoteRepo,
	followCommon activity_common.FollowRepo,
	tagCommon *tagcommon.TagCommonService,
	userCommon *usercommon.UserCommon,
	collectionCommon *collectioncommon.CollectionCommon,
	answerCommon *answercommon.AnswerCommon,
	metaService *meta.MetaService,
	configRepo config.ConfigRepo,

) *QuestionCommon {
	return &QuestionCommon{
		questionRepo:     questionRepo,
		answerRepo:       answerRepo,
		voteRepo:         voteRepo,
		followCommon:     followCommon,
		tagCommon:        tagCommon,
		userCommon:       userCommon,
		collectionCommon: collectionCommon,
		AnswerCommon:     answerCommon,
		metaService:      metaService,
		configRepo:       configRepo,
	}
}

func (qs *QuestionCommon) UpdataPv(ctx context.Context, questionId string) error {
	return qs.questionRepo.UpdatePvCount(ctx, questionId)
}
func (qs *QuestionCommon) UpdateAnswerCount(ctx context.Context, questionId string, num int) error {
	return qs.questionRepo.UpdateAnswerCount(ctx, questionId, num)
}
func (qs *QuestionCommon) UpdateCollectionCount(ctx context.Context, questionId string, num int) error {
	return qs.questionRepo.UpdateCollectionCount(ctx, questionId, num)
}

func (qs *QuestionCommon) UpdateAccepted(ctx context.Context, questionId, AnswerId string) error {
	question := &entity.Question{}
	question.ID = questionId
	question.AcceptedAnswerID = AnswerId
	return qs.questionRepo.UpdateAccepted(ctx, question)
}

func (qs *QuestionCommon) UpdateLastAnswer(ctx context.Context, questionId, AnswerId string) error {
	question := &entity.Question{}
	question.ID = questionId
	question.LastAnswerID = AnswerId
	return qs.questionRepo.UpdateLastAnswer(ctx, question)
}

func (qs *QuestionCommon) UpdataPostTime(ctx context.Context, questionId string) error {
	questioninfo := &entity.Question{}
	now := time.Now()
	questioninfo.ID = questionId
	questioninfo.PostUpdateTime = now
	return qs.questionRepo.UpdateQuestion(ctx, questioninfo, []string{"post_update_time"})
}

func (qs *QuestionCommon) FindInfoByID(ctx context.Context, questionIds []string, loginUserID string) (map[string]*schema.QuestionInfo, error) {
	list := make(map[string]*schema.QuestionInfo)
	listAddTag := make([]*entity.QuestionTag, 0)
	questionList, err := qs.questionRepo.FindByID(ctx, questionIds)
	if err != nil {
		return list, err
	}
	for _, item := range questionList {
		itemAddTag := &entity.QuestionTag{}
		itemAddTag.Question = *item
		listAddTag = append(listAddTag, itemAddTag)
	}
	QuestionInfo, err := qs.ListFormat(ctx, listAddTag, loginUserID)
	if err != nil {
		return list, err
	}
	for _, item := range QuestionInfo {
		list[item.ID] = item
	}
	return list, nil
}

func (qs *QuestionCommon) Info(ctx context.Context, questionId string, loginUserID string) (showinfo *schema.QuestionInfo, err error) {
	dbinfo, has, err := qs.questionRepo.GetQuestion(ctx, questionId)
	if err != nil {
		return showinfo, err
	}
	if !has {
		return showinfo, errors.BadRequest(reason.QuestionNotFound)
	}
	showinfo = qs.ShowFormat(ctx, dbinfo)

	if showinfo.Status == 2 {
		metainfo, err := qs.metaService.GetMetaByObjectIdAndKey(ctx, dbinfo.ID, entity.QuestionCloseReasonKey)
		if err != nil {
			log.Error(err)
		} else {
			//metainfo.Value
			closemsg := &schema.CloseQuestionMeta{}
			err := json.Unmarshal([]byte(metainfo.Value), closemsg)
			if err != nil {
				log.Error("json.Unmarshal CloseQuestionMeta error", err.Error())
			} else {
				closeinfo := &schema.GetReportTypeResp{}
				err = qs.configRepo.GetConfigById(closemsg.CloseType, closeinfo)
				if err != nil {
					log.Error("json.Unmarshal QuestionCloseJson error", err.Error())
				} else {
					operation := &schema.Operation{}
					operation.Operation_Type = closeinfo.Name
					operation.Operation_Description = closeinfo.Description
					operation.Operation_Msg = closemsg.CloseMsg
					operation.Operation_Time = metainfo.CreatedAt.Unix()
					showinfo.Operation = operation
				}

			}

		}
	}

	tagmap, err := qs.tagCommon.GetObjectTag(ctx, questionId)
	if err != nil {
		return showinfo, err
	}
	showinfo.Tags = tagmap

	userinfo, has, err := qs.userCommon.GetUserBasicInfoByID(ctx, dbinfo.UserID)
	if err != nil {
		return showinfo, err
	}
	if has {
		showinfo.UserInfo = userinfo
		showinfo.UpdateUserInfo = userinfo
		showinfo.LastAnsweredUserInfo = userinfo
	}

	if loginUserID == "" {
		return showinfo, nil
	}

	showinfo.VoteStatus = qs.voteRepo.GetVoteStatus(ctx, questionId, loginUserID)

	// // check is followed
	isFollowed, _ := qs.followCommon.IsFollowed(loginUserID, questionId)
	showinfo.IsFollowed = isFollowed

	has, err = qs.AnswerCommon.SearchAnswered(ctx, loginUserID, dbinfo.ID)
	if err != nil {
		log.Error("AnswerFunc.SearchAnswered", err)
	}
	showinfo.Answered = has

	//login user  Collected information

	CollectedMap, err := qs.collectionCommon.SearchObjectCollected(ctx, loginUserID, []string{dbinfo.ID})
	if err != nil {
		log.Error("CollectionFunc.SearchObjectCollected", err)
	}
	_, ok := CollectedMap[dbinfo.ID]
	if ok {
		showinfo.Collected = true
	}

	return showinfo, nil
}

func (qs *QuestionCommon) ListFormat(ctx context.Context, questionList []*entity.QuestionTag, loginUserID string) ([]*schema.QuestionInfo, error) {
	list := make([]*schema.QuestionInfo, 0)
	objectIds := make([]string, 0)
	userIds := make([]string, 0)

	for _, questionInfo := range questionList {
		item := qs.ShowListFormat(ctx, questionInfo)
		list = append(list, item)
		objectIds = append(objectIds, item.ID)
		userIds = append(userIds, questionInfo.UserID)
	}
	tagsMap, err := qs.tagCommon.BatchGetObjectTag(ctx, objectIds)
	if err != nil {
		return list, err
	}

	userInfoMap, err := qs.userCommon.BatchUserBasicInfoByID(ctx, userIds)
	if err != nil {
		return list, err
	}

	for _, item := range list {
		_, ok := tagsMap[item.ID]
		if ok {
			item.Tags = tagsMap[item.ID]
		}
		_, ok = userInfoMap[item.UserId]
		if ok {
			item.UserInfo = userInfoMap[item.UserId]
			item.UpdateUserInfo = userInfoMap[item.UserId]
			item.LastAnsweredUserInfo = userInfoMap[item.UserId]
		}
	}

	if loginUserID == "" {
		return list, nil
	}
	// //login user  Collected information
	CollectedMap, err := qs.collectionCommon.SearchObjectCollected(ctx, loginUserID, objectIds)
	if err != nil {
		log.Error("CollectionFunc.SearchObjectCollected", err)
	}

	for _, item := range list {
		_, ok := CollectedMap[item.ID]
		if ok {
			item.Collected = true
		}
	}
	return list, nil
}

// RemoveQuestion delete question
func (qs *QuestionCommon) RemoveQuestion(ctx context.Context, req *schema.RemoveQuestionReq) (err error) {
	questionInfo, has, err := qs.questionRepo.GetQuestion(ctx, req.ID)
	if err != nil {
		return err
	}
	if !has {
		return nil
	}
	questionInfo.Status = entity.QuestionStatusDeleted
	err = qs.questionRepo.UpdateQuestionStatus(ctx, questionInfo)
	if err != nil {
		return err
	}

	//user add question count
	err = qs.userCommon.UpdateQuestionCount(ctx, questionInfo.UserID, -1)
	if err != nil {
		log.Error("user UpdateQuestionCount error", err.Error())
	}

	// todo rank remove

	return nil
}

func (qs *QuestionCommon) CloseQuestion(ctx context.Context, req *schema.CloseQuestionReq) error {
	questionInfo, has, err := qs.questionRepo.GetQuestion(ctx, req.ID)
	if err != nil {
		return err
	}
	if !has {
		return nil
	}
	questionInfo.Status = entity.QuestionStatusclosed
	err = qs.questionRepo.UpdateQuestionStatus(ctx, questionInfo)
	if err != nil {
		return err
	}

	closeMeta, _ := json.Marshal(schema.CloseQuestionMeta{
		CloseType: req.CloseType,
		CloseMsg:  req.CloseMsg,
	})
	err = qs.metaService.AddMeta(ctx, req.ID, entity.QuestionCloseReasonKey, string(closeMeta))
	if err != nil {
		return err
	}
	return nil
}

// RemoveAnswer delete answer
func (as *QuestionCommon) RemoveAnswer(ctx context.Context, id string) (err error) {
	answerinfo, has, err := as.answerRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if !has {
		return nil
	}

	//user add question count

	err = as.UpdateAnswerCount(ctx, answerinfo.QuestionID, -1)
	if err != nil {
		log.Error("UpdateAnswerCount error", err.Error())
	}

	err = as.userCommon.UpdateAnswerCount(ctx, answerinfo.UserID, -1)
	if err != nil {
		log.Error("user UpdateAnswerCount error", err.Error())
	}

	return as.answerRepo.RemoveAnswer(ctx, id)
}

func (qs *QuestionCommon) ShowListFormat(ctx context.Context, data *entity.QuestionTag) *schema.QuestionInfo {
	return qs.ShowFormat(ctx, &data.Question)
}

func (qs *QuestionCommon) ShowFormat(ctx context.Context, data *entity.Question) *schema.QuestionInfo {
	info := schema.QuestionInfo{}
	info.ID = data.ID
	info.Title = data.Title
	info.Content = data.OriginalText
	info.Html = data.ParsedText
	info.ViewCount = data.ViewCount
	info.UniqueViewCount = data.UniqueViewCount
	info.VoteCount = data.VoteCount
	info.AnswerCount = data.AnswerCount
	info.CollectionCount = data.CollectionCount
	info.FollowCount = data.FollowCount
	info.AcceptedAnswerId = data.AcceptedAnswerID
	info.LastAnswerId = data.LastAnswerID
	info.CreateTime = data.CreatedAt.Unix()
	info.UpdateTime = data.UpdatedAt.Unix()
	info.PostUpdateTime = data.PostUpdateTime.Unix()
	info.QuestionUpdateTime = data.UpdatedAt.Unix()
	info.Status = data.Status
	info.UserId = data.UserID
	info.Tags = make([]*schema.TagResp, 0)
	return &info
}
