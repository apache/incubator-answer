package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/constant"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/base/translator"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/activity"
	collectioncommon "github.com/segmentfault/answer/internal/service/collection_common"
	"github.com/segmentfault/answer/internal/service/meta"
	"github.com/segmentfault/answer/internal/service/notice_queue"
	"github.com/segmentfault/answer/internal/service/permission"
	questioncommon "github.com/segmentfault/answer/internal/service/question_common"
	"github.com/segmentfault/answer/internal/service/revision_common"
	tagcommon "github.com/segmentfault/answer/internal/service/tag_common"
	usercommon "github.com/segmentfault/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/i18n"
	"github.com/segmentfault/pacman/log"
	"golang.org/x/net/context"
)

// QuestionRepo question repository

// QuestionService user service
type QuestionService struct {
	questionRepo          questioncommon.QuestionRepo
	tagCommon             *tagcommon.TagCommonService
	questioncommon        *questioncommon.QuestionCommon
	userCommon            *usercommon.UserCommon
	revisionService       *revision_common.RevisionService
	metaService           *meta.MetaService
	collectionCommon      *collectioncommon.CollectionCommon
	answerActivityService *activity.AnswerActivityService
}

func NewQuestionService(
	questionRepo questioncommon.QuestionRepo,
	tagCommon *tagcommon.TagCommonService,
	questioncommon *questioncommon.QuestionCommon,
	userCommon *usercommon.UserCommon,
	revisionService *revision_common.RevisionService,
	metaService *meta.MetaService,
	collectionCommon *collectioncommon.CollectionCommon,
	answerActivityService *activity.AnswerActivityService,
) *QuestionService {
	return &QuestionService{
		questionRepo:          questionRepo,
		tagCommon:             tagCommon,
		questioncommon:        questioncommon,
		userCommon:            userCommon,
		revisionService:       revisionService,
		metaService:           metaService,
		collectionCommon:      collectionCommon,
		answerActivityService: answerActivityService,
	}
}

func (qs *QuestionService) CloseQuestion(ctx context.Context, req *schema.CloseQuestionReq) error {
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

// CloseMsgList list close question condition
func (qs *QuestionService) CloseMsgList(ctx context.Context, lang i18n.Language) (
	resp []*schema.GetCloseTypeResp, err error) {
	resp = make([]*schema.GetCloseTypeResp, 0)
	err = json.Unmarshal([]byte(constant.QuestionCloseJson), &resp)
	if err != nil {
		return nil, errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	for _, t := range resp {
		t.Name = translator.GlobalTrans.Tr(lang, t.Name)
		t.Description = translator.GlobalTrans.Tr(lang, t.Description)
	}
	return resp, err
}

// AddQuestion add question
func (qs *QuestionService) AddQuestion(ctx context.Context, req *schema.QuestionAdd) (questionInfo *schema.QuestionInfo, err error) {
	questionInfo = &schema.QuestionInfo{}
	question := &entity.Question{}
	now := time.Now()
	question.UserID = req.UserID
	question.Title = req.Title
	question.OriginalText = req.Content
	question.ParsedText = req.Html
	question.AcceptedAnswerID = "0"
	question.LastAnswerID = "0"
	question.PostUpdateTime = now
	question.Status = entity.QuestionStatusAvailable
	question.RevisionID = "0"
	question.CreatedAt = now
	question.UpdatedAt = now
	err = qs.questionRepo.AddQuestion(ctx, question)
	if err != nil {
		return
	}
	objectTagData := schema.TagChange{}
	objectTagData.ObjectId = question.ID
	objectTagData.Tags = req.Tags
	objectTagData.UserID = req.UserID
	err = qs.ChangeTag(ctx, &objectTagData)
	if err != nil {
		return
	}

	revisionDTO := &schema.AddRevisionDTO{
		UserID:   question.UserID,
		ObjectID: question.ID,
		Title:    "",
	}
	InfoJson, _ := json.Marshal(question)
	revisionDTO.Content = string(InfoJson)
	err = qs.revisionService.AddRevision(ctx, revisionDTO, true)
	if err != nil {
		return
	}

	//user add question count
	err = qs.userCommon.UpdateQuestionCount(ctx, question.UserID, 1)
	if err != nil {
		log.Error("user IncreaseQuestionCount error", err.Error())
	}

	questionInfo, err = qs.GetQuestion(ctx, question.ID, question.UserID, false)
	return
}

// RemoveQuestion delete question
func (qs *QuestionService) RemoveQuestion(ctx context.Context, req *schema.RemoveQuestionReq) (err error) {
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
		log.Error("user IncreaseQuestionCount error", err.Error())
	}

	err = qs.answerActivityService.DeleteQuestion(ctx, questionInfo.ID, questionInfo.CreatedAt, questionInfo.VoteCount)
	if err != nil {
		log.Errorf("user DeleteQuestion rank rollback error %s", err.Error())
	}

	return nil
}

// UpdateQuestion update question
func (qs *QuestionService) UpdateQuestion(ctx context.Context, req *schema.QuestionUpdate) (questionInfo *schema.QuestionInfo, err error) {
	questionInfo = &schema.QuestionInfo{}
	now := time.Now()
	question := &entity.Question{}
	question.UserID = req.UserID
	question.Title = req.Title
	question.OriginalText = req.Content
	question.ParsedText = req.Html
	question.ID = req.ID
	question.UpdatedAt = now
	dbinfo, has, err := qs.questionRepo.GetQuestion(ctx, question.ID)
	if err != nil {
		return
	}
	if !has {
		return
	}
	if dbinfo.UserID != req.UserID {
		return
	}
	err = qs.questionRepo.UpdateQuestion(ctx, question, []string{"title", "original_text", "parsed_text", "updated_at"})
	if err != nil {
		return
	}
	objectTagData := schema.TagChange{}
	objectTagData.ObjectId = question.ID
	objectTagData.Tags = req.Tags
	objectTagData.UserID = req.UserID
	err = qs.ChangeTag(ctx, &objectTagData)
	if err != nil {
		return
	}

	revisionDTO := &schema.AddRevisionDTO{
		UserID:   question.UserID,
		ObjectID: question.ID,
		Title:    "",
		Log:      req.EditSummary,
	}
	InfoJson, _ := json.Marshal(question)
	revisionDTO.Content = string(InfoJson)
	err = qs.revisionService.AddRevision(ctx, revisionDTO, true)
	if err != nil {
		return
	}

	questionInfo, err = qs.GetQuestion(ctx, question.ID, question.UserID, false)
	return
}

// GetQuestion get question one
func (qs *QuestionService) GetQuestion(ctx context.Context, id, loginUserID string, addpv bool) (resp *schema.QuestionInfo, err error) {
	question, err := qs.questioncommon.Info(ctx, id, loginUserID)
	if err != nil {
		return
	}
	if addpv {
		err = qs.questioncommon.UpdataPv(ctx, id)
		if err != nil {
			log.Error("UpdataPv", err)
		}
	}

	question.MemberActions = permission.GetQuestionPermission(loginUserID, question.UserId)
	return question, nil
}

func (qs *QuestionService) ChangeTag(ctx context.Context, objectTagData *schema.TagChange) error {
	return qs.tagCommon.ObjectChangeTag(ctx, objectTagData)
}

func (qs *QuestionService) SearchUserList(ctx context.Context, userName, order string, page, pageSize int, loginUserID string) ([]*schema.UserQuestionInfo, int64, error) {
	userlist := make([]*schema.UserQuestionInfo, 0)

	userinfo, Exist, err := qs.userCommon.GetUserBasicInfoByUserName(ctx, userName)
	if err != nil {
		return userlist, 0, err
	}
	if !Exist {
		return userlist, 0, nil
	}
	search := &schema.QuestionSearch{}
	search.Order = order
	search.Page = page
	search.PageSize = pageSize
	search.UserID = userinfo.ID
	questionlist, count, err := qs.SearchList(ctx, search, loginUserID)
	if err != nil {
		return userlist, 0, err
	}
	for _, item := range questionlist {
		info := &schema.UserQuestionInfo{}
		_ = copier.Copy(info, item)
		status, ok := entity.CmsQuestionSearchStatusIntToString[item.Status]
		if ok {
			info.Status = status
		}
		userlist = append(userlist, info)
	}
	return userlist, count, nil
}

func (qs *QuestionService) SearchUserAnswerList(ctx context.Context, userName, order string, page, pageSize int, loginUserID string) ([]*schema.UserAnswerInfo, int64, error) {
	answerlist := make([]*schema.AnswerInfo, 0)
	userAnswerlist := make([]*schema.UserAnswerInfo, 0)
	userinfo, Exist, err := qs.userCommon.GetUserBasicInfoByUserName(ctx, userName)
	if err != nil {
		return userAnswerlist, 0, err
	}
	if !Exist {
		return userAnswerlist, 0, nil
	}
	answersearch := &entity.AnswerSearch{}
	answersearch.UserID = userinfo.ID
	answersearch.PageSize = pageSize
	answersearch.Page = page
	if order == "newest" {
		answersearch.Order = entity.Answer_Search_OrderBy_Time
	} else {
		answersearch.Order = entity.Answer_Search_OrderBy_Default
	}
	questionIDs := make([]string, 0)
	answerList, count, err := qs.questioncommon.AnswerCommon.Search(ctx, answersearch)
	if err != nil {
		return userAnswerlist, count, err
	}
	for _, item := range answerList {
		answerinfo := qs.questioncommon.AnswerCommon.ShowFormat(ctx, item)
		answerlist = append(answerlist, answerinfo)
		questionIDs = append(questionIDs, item.QuestionID)
	}
	questionMaps, err := qs.questioncommon.FindInfoByID(ctx, questionIDs, loginUserID)
	if err != nil {
		return userAnswerlist, count, err
	}
	for _, item := range answerlist {
		_, ok := questionMaps[item.QuestionId]
		if ok {
			item.QuestionInfo = questionMaps[item.QuestionId]
		}
	}
	for _, item := range answerlist {
		info := &schema.UserAnswerInfo{}
		_ = copier.Copy(info, item)
		info.AnswerID = item.ID
		info.QuestionID = item.QuestionId
		userAnswerlist = append(userAnswerlist, info)
	}
	return userAnswerlist, count, nil
}

func (qs *QuestionService) SearchUserCollectionList(ctx context.Context, page, pageSize int, loginUserID string) ([]*schema.QuestionInfo, int64, error) {
	list := make([]*schema.QuestionInfo, 0)
	userinfo, Exist, err := qs.userCommon.GetUserBasicInfoByID(ctx, loginUserID)
	if err != nil {
		return list, 0, err
	}
	if !Exist {
		return list, 0, nil
	}
	collectionSearch := &entity.CollectionSearch{}
	collectionSearch.UserID = userinfo.ID
	collectionSearch.Page = page
	collectionSearch.PageSize = pageSize
	collectionlist, count, err := qs.collectionCommon.SearchList(ctx, collectionSearch)
	if err != nil {
		return list, 0, err
	}
	questionIDs := make([]string, 0)
	for _, item := range collectionlist {
		questionIDs = append(questionIDs, item.ObjectID)
	}

	questionMaps, err := qs.questioncommon.FindInfoByID(ctx, questionIDs, loginUserID)
	if err != nil {
		return list, count, err
	}
	for _, id := range questionIDs {
		_, ok := questionMaps[id]
		if ok {
			questionMaps[id].LastAnsweredUserInfo = nil
			questionMaps[id].UpdateUserInfo = nil
			questionMaps[id].Content = ""
			questionMaps[id].Html = ""
			list = append(list, questionMaps[id])
		}
	}

	return list, count, nil
}

func (qs *QuestionService) SearchUserTopList(ctx context.Context, userName string, loginUserID string) ([]*schema.UserQuestionInfo, []*schema.UserAnswerInfo, error) {
	answerlist := make([]*schema.AnswerInfo, 0)

	userAnswerlist := make([]*schema.UserAnswerInfo, 0)
	userQuestionlist := make([]*schema.UserQuestionInfo, 0)

	userinfo, Exist, err := qs.userCommon.GetUserBasicInfoByUserName(ctx, userName)
	if err != nil {
		return userQuestionlist, userAnswerlist, err
	}
	if !Exist {
		return userQuestionlist, userAnswerlist, nil
	}
	search := &schema.QuestionSearch{}
	search.Order = "score"
	search.Page = 0
	search.PageSize = 5
	search.UserID = userinfo.ID
	questionlist, _, err := qs.SearchList(ctx, search, loginUserID)
	if err != nil {
		return userQuestionlist, userAnswerlist, err
	}
	answersearch := &entity.AnswerSearch{}
	answersearch.UserID = userinfo.ID
	answersearch.PageSize = 5
	answersearch.Order = entity.Answer_Search_OrderBy_Vote
	questionIDs := make([]string, 0)
	answerList, _, err := qs.questioncommon.AnswerCommon.Search(ctx, answersearch)
	if err != nil {
		return userQuestionlist, userAnswerlist, err
	}
	for _, item := range answerList {
		answerinfo := qs.questioncommon.AnswerCommon.ShowFormat(ctx, item)
		answerlist = append(answerlist, answerinfo)
		questionIDs = append(questionIDs, item.QuestionID)
	}
	questionMaps, err := qs.questioncommon.FindInfoByID(ctx, questionIDs, loginUserID)
	if err != nil {
		return userQuestionlist, userAnswerlist, err
	}
	for _, item := range answerlist {
		_, ok := questionMaps[item.QuestionId]
		if ok {
			item.QuestionInfo = questionMaps[item.QuestionId]
		}
	}

	for _, item := range questionlist {
		info := &schema.UserQuestionInfo{}
		_ = copier.Copy(info, item)
		userQuestionlist = append(userQuestionlist, info)
	}

	for _, item := range answerlist {
		info := &schema.UserAnswerInfo{}
		_ = copier.Copy(info, item)
		info.AnswerID = item.ID
		info.QuestionID = item.QuestionId
		userAnswerlist = append(userAnswerlist, info)
	}

	return userQuestionlist, userAnswerlist, nil
}

// SearchByTitleLike
func (qs *QuestionService) SearchByTitleLike(ctx context.Context, title string, loginUserID string) ([]*schema.QuestionBaseInfo, error) {
	list := make([]*schema.QuestionBaseInfo, 0)
	dblist, err := qs.questionRepo.SearchByTitleLike(ctx, title)
	if err != nil {
		return list, err
	}
	for _, question := range dblist {
		item := &schema.QuestionBaseInfo{}
		item.ID = question.ID
		item.Title = question.Title
		item.ViewCount = question.ViewCount
		item.AnswerCount = question.AnswerCount
		item.CollectionCount = question.CollectionCount
		item.FollowCount = question.FollowCount
		status, ok := entity.CmsQuestionSearchStatusIntToString[question.Status]
		if ok {
			item.Status = status
		}
		if question.AcceptedAnswerID != "0" {
			item.AcceptedAnswer = true
		}
		list = append(list, item)
	}

	return list, nil
}

// SimilarQuestion
func (qs *QuestionService) SimilarQuestion(ctx context.Context, questionID string, loginUserID string) ([]*schema.QuestionInfo, int64, error) {
	list := make([]*schema.QuestionInfo, 0)
	questionInfo, err := qs.GetQuestion(ctx, questionID, loginUserID, false)
	if err != nil {
		return list, 0, err
	}
	tagNames := make([]string, 0, len(questionInfo.Tags))
	for _, tag := range questionInfo.Tags {
		tagNames = append(tagNames, tag.SlugName)
	}
	search := &schema.QuestionSearch{}
	search.Order = "frequent"
	search.Page = 0
	search.PageSize = 6
	search.Tags = tagNames
	return qs.SearchList(ctx, search, loginUserID)
}

// SearchList
func (qs *QuestionService) SearchList(ctx context.Context, req *schema.QuestionSearch, loginUserID string) ([]*schema.QuestionInfo, int64, error) {
	if len(req.Tags) > 0 {
		taginfo, err := qs.tagCommon.GetTagListByNames(ctx, req.Tags)
		if err != nil {
			log.Error("tagCommon.GetTagListByNames error", err)
		}
		for _, tag := range taginfo {
			req.TagIDs = append(req.TagIDs, tag.ID)
		}
	}
	list := make([]*schema.QuestionInfo, 0)
	if req.UserName != "" {
		userinfo, exist, err := qs.userCommon.GetUserBasicInfoByUserName(ctx, req.UserName)
		if err != nil {
			return list, 0, err
		}
		if !exist {
			return list, 0, err
		}
		req.UserID = userinfo.ID
	}
	questionList, count, err := qs.questionRepo.SearchList(ctx, req)
	if err != nil {
		return list, count, err
	}
	list, err = qs.questioncommon.ListFormat(ctx, questionList, loginUserID)
	if err != nil {
		return list, count, err
	}
	return list, count, nil
}

func (qs *QuestionService) AdminSetQuestionStatus(ctx context.Context, questionID string, setStatusStr string) error {
	setStatus, ok := entity.CmsQuestionSearchStatus[setStatusStr]
	if !ok {
		return fmt.Errorf("question status does not exist")
	}
	questionInfo, exist, err := qs.questionRepo.GetQuestion(ctx, questionID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.QuestionNotFound)
	}
	questionInfo.Status = setStatus
	err = qs.questionRepo.UpdateQuestionStatus(ctx, questionInfo)
	if err != nil {
		return err
	}

	if setStatus == entity.QuestionStatusDeleted {
		err = qs.answerActivityService.DeleteQuestion(ctx, questionInfo.ID, questionInfo.CreatedAt, questionInfo.VoteCount)
		if err != nil {
			log.Errorf("admin delete question then rank rollback error %s", err.Error())
		}
	}
	msg := &schema.NotificationMsg{}
	msg.ObjectID = questionInfo.ID
	msg.Type = schema.NotificationTypeInbox
	msg.ReceiverUserID = questionInfo.UserID
	msg.TriggerUserID = questionInfo.UserID
	msg.ObjectType = constant.QuestionObjectType
	msg.NotificationAction = constant.YourQuestionWasDeleted
	notice_queue.AddNotification(msg)
	return nil
}

func (qs *QuestionService) CmsSearchList(ctx context.Context, search *schema.CmsQuestionSearch, loginUserID string) ([]*schema.AdminQuestionInfo, int64, error) {
	list := make([]*schema.AdminQuestionInfo, 0)

	status, ok := entity.CmsQuestionSearchStatus[search.StatusStr]
	if ok {
		search.Status = status
	}

	if search.Status == 0 {
		search.Status = 1
	}
	dblist, count, err := qs.questionRepo.CmsSearchList(ctx, search)
	if err != nil {
		return list, count, err
	}
	userIds := make([]string, 0)
	for _, dbitem := range dblist {
		item := &schema.AdminQuestionInfo{}
		_ = copier.Copy(item, dbitem)
		item.CreateTime = dbitem.CreatedAt.Unix()
		item.UpdateTime = dbitem.PostUpdateTime.Unix()
		item.EditTime = dbitem.UpdatedAt.Unix()
		list = append(list, item)
		userIds = append(userIds, dbitem.UserID)
	}
	userInfoMap, err := qs.userCommon.BatchUserBasicInfoByID(ctx, userIds)
	if err != nil {
		return list, count, err
	}
	for _, item := range list {
		_, ok = userInfoMap[item.UserID]
		if ok {
			item.UserInfo = userInfoMap[item.UserID]
		}
	}

	return list, count, nil
}

// CmsSearchList
func (qs *QuestionService) CmsSearchAnswerList(ctx context.Context, search *entity.CmsAnswerSearch, loginUserID string) ([]*schema.AdminAnswerInfo, int64, error) {
	answerlist := make([]*schema.AdminAnswerInfo, 0)

	status, ok := entity.CmsAnswerSearchStatus[search.StatusStr]
	if ok {
		search.Status = status
	}

	if search.Status == 0 {
		search.Status = 1
	}
	dblist, count, err := qs.questioncommon.AnswerCommon.CmsSearchList(ctx, search)
	if err != nil {
		return answerlist, count, err
	}
	questionIDs := make([]string, 0)
	userIds := make([]string, 0)
	for _, item := range dblist {
		answerinfo := qs.questioncommon.AnswerCommon.AdminShowFormat(ctx, item)
		answerlist = append(answerlist, answerinfo)
		questionIDs = append(questionIDs, item.QuestionID)
		userIds = append(userIds, item.UserID)
	}
	userInfoMap, err := qs.userCommon.BatchUserBasicInfoByID(ctx, userIds)
	if err != nil {
		return answerlist, count, err
	}

	questionMaps, err := qs.questioncommon.FindInfoByID(ctx, questionIDs, loginUserID)
	if err != nil {
		return answerlist, count, err
	}
	for _, item := range answerlist {
		_, ok := questionMaps[item.QuestionId]
		if ok {
			item.QuestionInfo.Title = questionMaps[item.QuestionId].Title
		}
		_, ok = userInfoMap[item.UserId]
		if ok {
			item.UserInfo = userInfoMap[item.UserId]
		}
	}
	return answerlist, count, nil
}
