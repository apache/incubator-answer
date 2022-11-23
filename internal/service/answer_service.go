package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/activity_queue"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	collectioncommon "github.com/answerdev/answer/internal/service/collection_common"
	"github.com/answerdev/answer/internal/service/notice_queue"
	"github.com/answerdev/answer/internal/service/permission"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/revision_common"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// AnswerService user service
type AnswerService struct {
	answerRepo            answercommon.AnswerRepo
	questionRepo          questioncommon.QuestionRepo
	questionCommon        *questioncommon.QuestionCommon
	answerActivityService *activity.AnswerActivityService
	userCommon            *usercommon.UserCommon
	collectionCommon      *collectioncommon.CollectionCommon
	userRepo              usercommon.UserRepo
	revisionService       *revision_common.RevisionService
	AnswerCommon          *answercommon.AnswerCommon
	voteRepo              activity_common.VoteRepo
}

func NewAnswerService(
	answerRepo answercommon.AnswerRepo,
	questionRepo questioncommon.QuestionRepo,
	questionCommon *questioncommon.QuestionCommon,
	userCommon *usercommon.UserCommon,
	collectionCommon *collectioncommon.CollectionCommon,
	userRepo usercommon.UserRepo,
	revisionService *revision_common.RevisionService,
	answerAcceptActivityRepo *activity.AnswerActivityService,
	answerCommon *answercommon.AnswerCommon,
	voteRepo activity_common.VoteRepo,
) *AnswerService {
	return &AnswerService{
		answerRepo:            answerRepo,
		questionRepo:          questionRepo,
		userCommon:            userCommon,
		collectionCommon:      collectionCommon,
		questionCommon:        questionCommon,
		userRepo:              userRepo,
		revisionService:       revisionService,
		answerActivityService: answerAcceptActivityRepo,
		AnswerCommon:          answerCommon,
		voteRepo:              voteRepo,
	}
}

// RemoveAnswer delete answer
func (as *AnswerService) RemoveAnswer(ctx context.Context, req *schema.RemoveAnswerReq) (err error) {
	answerInfo, exist, err := as.answerRepo.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}
	if answerInfo.UserID != req.UserID {
		return errors.BadRequest(reason.UnauthorizedError)
	}
	if answerInfo.VoteCount > 0 {
		return errors.BadRequest(reason.UnauthorizedError)
	}
	if answerInfo.Adopted == schema.AnswerAdoptedEnable {
		return errors.BadRequest(reason.UnauthorizedError)
	}
	questionInfo, exist, err := as.questionRepo.GetQuestion(ctx, answerInfo.QuestionID)
	if err != nil {
		return errors.BadRequest(reason.UnauthorizedError)
	}
	if !exist {
		return errors.BadRequest(reason.UnauthorizedError)
	}
	if questionInfo.AnswerCount > 1 {
		return errors.BadRequest(reason.UnauthorizedError)
	}
	if questionInfo.AcceptedAnswerID != "" {
		return errors.BadRequest(reason.UnauthorizedError)
	}

	// user add question count
	err = as.questionCommon.UpdateAnswerCount(ctx, answerInfo.QuestionID, -1)
	if err != nil {
		log.Error("IncreaseAnswerCount error", err.Error())
	}

	err = as.userCommon.UpdateAnswerCount(ctx, answerInfo.UserID, -1)
	if err != nil {
		log.Error("user IncreaseAnswerCount error", err.Error())
	}

	err = as.answerRepo.RemoveAnswer(ctx, req.ID)
	if err != nil {
		return err
	}
	err = as.answerActivityService.DeleteAnswer(ctx, answerInfo.ID, answerInfo.CreatedAt, answerInfo.VoteCount)
	if err != nil {
		log.Errorf("delete answer activity change failed: %s", err.Error())
	}
	activity_queue.AddActivity(&schema.ActivityMsg{
		UserID:           req.UserID,
		ObjectID:         answerInfo.ID,
		OriginalObjectID: answerInfo.ID,
		ActivityTypeKey:  constant.ActAnswerDeleted,
	})
	return
}

func (as *AnswerService) Insert(ctx context.Context, req *schema.AnswerAddReq) (string, error) {
	questionInfo, exist, err := as.questionRepo.GetQuestion(ctx, req.QuestionID)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errors.BadRequest(reason.QuestionNotFound)
	}
	now := time.Now()
	insertData := new(entity.Answer)
	insertData.UserID = req.UserID
	insertData.OriginalText = req.Content
	insertData.ParsedText = req.HTML
	insertData.Adopted = schema.AnswerAdoptedFailed
	insertData.QuestionID = req.QuestionID
	insertData.RevisionID = "0"
	insertData.Status = entity.AnswerStatusAvailable
	insertData.UpdatedAt = now
	if err = as.answerRepo.AddAnswer(ctx, insertData); err != nil {
		return "", err
	}
	err = as.questionCommon.UpdateAnswerCount(ctx, req.QuestionID, 1)
	if err != nil {
		log.Error("IncreaseAnswerCount error", err.Error())
	}
	err = as.questionCommon.UpdateLastAnswer(ctx, req.QuestionID, insertData.ID)
	if err != nil {
		log.Error("UpdateLastAnswer error", err.Error())
	}
	err = as.questionCommon.UpdataPostTime(ctx, req.QuestionID)
	if err != nil {
		return insertData.ID, err
	}

	err = as.userCommon.UpdateAnswerCount(ctx, req.UserID, 1)
	if err != nil {
		log.Error("user IncreaseAnswerCount error", err.Error())
	}

	revisionDTO := &schema.AddRevisionDTO{
		UserID:   insertData.UserID,
		ObjectID: insertData.ID,
		Title:    "",
	}
	infoJSON, _ := json.Marshal(insertData)
	revisionDTO.Content = string(infoJSON)
	revisionID, err := as.revisionService.AddRevision(ctx, revisionDTO, true)
	if err != nil {
		return insertData.ID, err
	}
	as.notificationAnswerTheQuestion(ctx, questionInfo.UserID, insertData.ID, req.UserID)

	activity_queue.AddActivity(&schema.ActivityMsg{
		UserID:           insertData.UserID,
		ObjectID:         insertData.ID,
		OriginalObjectID: insertData.ID,
		ActivityTypeKey:  constant.ActAnswerAnswered,
		RevisionID:       revisionID,
	})
	activity_queue.AddActivity(&schema.ActivityMsg{
		UserID:           insertData.UserID,
		ObjectID:         insertData.ID,
		OriginalObjectID: questionInfo.ID,
		ActivityTypeKey:  constant.ActQuestionAnswered,
	})
	return insertData.ID, nil
}

func (as *AnswerService) Update(ctx context.Context, req *schema.AnswerUpdateReq) (string, error) {
	questionInfo, exist, err := as.questionRepo.GetQuestion(ctx, req.QuestionID)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errors.BadRequest(reason.QuestionNotFound)
	}
	now := time.Now()
	insertData := new(entity.Answer)
	insertData.ID = req.ID
	insertData.QuestionID = req.QuestionID
	insertData.UserID = req.UserID
	insertData.OriginalText = req.Content
	insertData.ParsedText = req.HTML
	insertData.UpdatedAt = now
	if err = as.answerRepo.UpdateAnswer(ctx, insertData, []string{"original_text", "parsed_text", "update_time"}); err != nil {
		return "", err
	}
	err = as.questionCommon.UpdataPostTime(ctx, req.QuestionID)
	if err != nil {
		return insertData.ID, err
	}
	revisionDTO := &schema.AddRevisionDTO{
		UserID:   req.UserID,
		ObjectID: req.ID,
		Title:    "",
		Log:      req.EditSummary,
	}
	infoJSON, _ := json.Marshal(insertData)
	revisionDTO.Content = string(infoJSON)
	revisionID, err := as.revisionService.AddRevision(ctx, revisionDTO, true)
	if err != nil {
		return insertData.ID, err
	}
	as.notificationUpdateAnswer(ctx, questionInfo.UserID, insertData.ID, req.UserID)

	activity_queue.AddActivity(&schema.ActivityMsg{
		UserID:           insertData.UserID,
		ObjectID:         insertData.ID,
		OriginalObjectID: insertData.ID,
		ActivityTypeKey:  constant.ActAnswerEdited,
		RevisionID:       revisionID,
	})
	return insertData.ID, nil
}

// UpdateAdopted
func (as *AnswerService) UpdateAdopted(ctx context.Context, req *schema.AnswerAdoptedReq) error {
	if req.AnswerID == "" {
		req.AnswerID = "0"
	}
	if req.UserID == "" {
		return nil
	}

	newAnswerInfo := &entity.Answer{}
	newAnswerInfoexist := false
	var err error

	if req.AnswerID != "0" {
		newAnswerInfo, newAnswerInfoexist, err = as.answerRepo.GetByID(ctx, req.AnswerID)
		if err != nil {
			return err
		}
		if !newAnswerInfoexist {
			return errors.BadRequest(reason.AnswerNotFound)
		}
	}

	questionInfo, exist, err := as.questionRepo.GetQuestion(ctx, req.QuestionID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.QuestionNotFound)
	}
	if questionInfo.UserID != req.UserID {
		return fmt.Errorf("no permission to set answer")
	}
	if questionInfo.AcceptedAnswerID == req.AnswerID {
		return nil
	}

	var oldAnswerInfo *entity.Answer
	if len(questionInfo.AcceptedAnswerID) > 0 && questionInfo.AcceptedAnswerID != "0" {
		oldAnswerInfo, _, err = as.answerRepo.GetByID(ctx, questionInfo.AcceptedAnswerID)
		if err != nil {
			return err
		}
	}

	err = as.answerRepo.UpdateAdopted(ctx, req.AnswerID, req.QuestionID)
	if err != nil {
		return err
	}

	err = as.questionCommon.UpdateAccepted(ctx, req.QuestionID, req.AnswerID)
	if err != nil {
		log.Error("UpdateLastAnswer error", err.Error())
	}

	as.updateAnswerRank(ctx, req.UserID, questionInfo, newAnswerInfo, oldAnswerInfo)
	return nil
}

func (as *AnswerService) updateAnswerRank(ctx context.Context, userID string,
	questionInfo *entity.Question, newAnswerInfo *entity.Answer, oldAnswerInfo *entity.Answer,
) {
	// if this question is already been answered, should cancel old answer rank
	if oldAnswerInfo != nil {
		err := as.answerActivityService.CancelAcceptAnswer(
			ctx, questionInfo.AcceptedAnswerID, questionInfo.ID, questionInfo.UserID, oldAnswerInfo.UserID)
		if err != nil {
			log.Error(err)
		}
	}
	if newAnswerInfo.ID != "" {
		err := as.answerActivityService.AcceptAnswer(
			ctx, newAnswerInfo.ID, questionInfo.ID, questionInfo.UserID, newAnswerInfo.UserID, newAnswerInfo.UserID == userID)
		if err != nil {
			log.Error(err)
		}
	}
}

func (as *AnswerService) Get(ctx context.Context, answerID, loginUserID string) (*schema.AnswerInfo, *schema.QuestionInfo, bool, error) {
	answerInfo, has, err := as.answerRepo.GetByID(ctx, answerID)
	if err != nil {
		return nil, nil, has, err
	}
	info := as.ShowFormat(ctx, answerInfo)
	// todo questionFunc
	questionInfo, err := as.questionCommon.Info(ctx, answerInfo.QuestionID, loginUserID)
	if err != nil {
		return nil, nil, has, err
	}
	// todo UserFunc
	userinfo, has, err := as.userCommon.GetUserBasicInfoByID(ctx, answerInfo.UserID)
	if err != nil {
		return nil, nil, has, err
	}
	if has {
		info.UserInfo = userinfo
		info.UpdateUserInfo = userinfo
	}

	if loginUserID == "" {
		return info, questionInfo, has, nil
	}

	info.VoteStatus = as.voteRepo.GetVoteStatus(ctx, answerID, loginUserID)

	CollectedMap, err := as.collectionCommon.SearchObjectCollected(ctx, loginUserID, []string{answerInfo.ID})
	if err != nil {
		log.Error("CollectionFunc.SearchObjectCollected error", err)
	}
	_, ok := CollectedMap[answerInfo.ID]
	if ok {
		info.Collected = true
	}

	return info, questionInfo, has, nil
}

func (as *AnswerService) AdminSetAnswerStatus(ctx context.Context, req *schema.AdminSetAnswerStatusRequest) error {
	setStatus, ok := entity.CmsAnswerSearchStatus[req.StatusStr]
	if !ok {
		return fmt.Errorf("question status does not exist")
	}
	answerInfo, exist, err := as.answerRepo.GetAnswer(ctx, req.AnswerID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("answer does not exist")
	}
	answerInfo.Status = setStatus
	err = as.answerRepo.UpdateAnswerStatus(ctx, answerInfo)
	if err != nil {
		return err
	}

	if setStatus == entity.AnswerStatusDeleted {
		err = as.answerActivityService.DeleteAnswer(ctx, answerInfo.ID, answerInfo.CreatedAt, answerInfo.VoteCount)
		if err != nil {
			log.Errorf("admin delete question then rank rollback error %s", err.Error())
		} else {
			activity_queue.AddActivity(&schema.ActivityMsg{
				UserID:           req.UserID,
				ObjectID:         answerInfo.ID,
				OriginalObjectID: answerInfo.ID,
				ActivityTypeKey:  constant.ActAnswerDeleted,
			})
		}
	}

	msg := &schema.NotificationMsg{}
	msg.ObjectID = answerInfo.ID
	msg.Type = schema.NotificationTypeInbox
	msg.ReceiverUserID = answerInfo.UserID
	msg.TriggerUserID = answerInfo.UserID
	msg.ObjectType = constant.AnswerObjectType
	msg.NotificationAction = constant.YourAnswerWasDeleted
	notice_queue.AddNotification(msg)

	return nil
}

func (as *AnswerService) SearchList(ctx context.Context, search *schema.AnswerList) ([]*schema.AnswerInfo, int64, error) {
	list := make([]*schema.AnswerInfo, 0)
	dbSearch := entity.AnswerSearch{}
	dbSearch.QuestionID = search.QuestionID
	dbSearch.Page = search.Page
	dbSearch.PageSize = search.PageSize
	dbSearch.Order = search.Order
	dblist, count, err := as.answerRepo.SearchList(ctx, &dbSearch)
	if err != nil {
		return list, count, err
	}
	AnswerList, err := as.SearchFormatInfo(ctx, dblist, search.LoginUserID)
	if err != nil {
		return AnswerList, count, err
	}
	return AnswerList, count, nil
}

func (as *AnswerService) SearchFormatInfo(ctx context.Context, dblist []*entity.Answer, loginUserID string) ([]*schema.AnswerInfo, error) {
	list := make([]*schema.AnswerInfo, 0)
	objectIds := make([]string, 0)
	userIds := make([]string, 0)
	for _, dbitem := range dblist {
		item := as.ShowFormat(ctx, dbitem)
		list = append(list, item)
		objectIds = append(objectIds, dbitem.ID)
		userIds = append(userIds, dbitem.UserID)
		if loginUserID != "" {
			// item.VoteStatus = as.activityFunc.GetVoteStatus(ctx, item.TagID, loginUserId)
			item.VoteStatus = as.voteRepo.GetVoteStatus(ctx, item.ID, loginUserID)
		}
	}
	userInfoMap, err := as.userCommon.BatchUserBasicInfoByID(ctx, userIds)
	if err != nil {
		return list, err
	}
	for _, item := range list {
		_, ok := userInfoMap[item.UserID]
		if ok {
			item.UserInfo = userInfoMap[item.UserID]
			item.UpdateUserInfo = userInfoMap[item.UserID]
		}
	}

	if loginUserID == "" {
		return list, nil
	}

	CollectedMap, err := as.collectionCommon.SearchObjectCollected(ctx, loginUserID, objectIds)
	if err != nil {
		log.Error("CollectionFunc.SearchObjectCollected error", err)
	}

	for _, item := range list {
		_, ok := CollectedMap[item.ID]
		if ok {
			item.Collected = true
		}
	}

	for _, item := range list {
		item.MemberActions = permission.GetAnswerPermission(loginUserID, item.UserID)
	}

	return list, nil
}

func (as *AnswerService) ShowFormat(ctx context.Context, data *entity.Answer) *schema.AnswerInfo {
	return as.AnswerCommon.ShowFormat(ctx, data)
}

func (as *AnswerService) notificationUpdateAnswer(ctx context.Context, questionUserID, answerID, answerUserID string) {
	msg := &schema.NotificationMsg{
		TriggerUserID:  answerUserID,
		ReceiverUserID: questionUserID,
		Type:           schema.NotificationTypeInbox,
		ObjectID:       answerID,
	}
	msg.ObjectType = constant.AnswerObjectType
	msg.NotificationAction = constant.UpdateAnswer
	notice_queue.AddNotification(msg)
}

func (as *AnswerService) notificationAnswerTheQuestion(ctx context.Context, questionUserID, answerID, answerUserID string) {
	msg := &schema.NotificationMsg{
		TriggerUserID:  answerUserID,
		ReceiverUserID: questionUserID,
		Type:           schema.NotificationTypeInbox,
		ObjectID:       answerID,
	}
	msg.ObjectType = constant.AnswerObjectType
	msg.NotificationAction = constant.AnswerTheQuestion
	notice_queue.AddNotification(msg)
}
