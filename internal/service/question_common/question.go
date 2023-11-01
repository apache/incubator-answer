/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package questioncommon

import (
	"context"
	"encoding/json"
	"math"
	"time"

	"github.com/apache/incubator-answer/internal/base/constant"
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/handler"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/service/activity_common"
	"github.com/apache/incubator-answer/internal/service/activity_queue"
	"github.com/apache/incubator-answer/internal/service/config"
	"github.com/apache/incubator-answer/internal/service/meta"
	"github.com/apache/incubator-answer/pkg/checker"
	"github.com/apache/incubator-answer/pkg/htmltext"
	"github.com/apache/incubator-answer/pkg/uid"
	"github.com/segmentfault/pacman/errors"

	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	answercommon "github.com/apache/incubator-answer/internal/service/answer_common"
	collectioncommon "github.com/apache/incubator-answer/internal/service/collection_common"
	tagcommon "github.com/apache/incubator-answer/internal/service/tag_common"
	usercommon "github.com/apache/incubator-answer/internal/service/user_common"
	"github.com/segmentfault/pacman/log"
)

// QuestionRepo question repository
type QuestionRepo interface {
	AddQuestion(ctx context.Context, question *entity.Question) (err error)
	RemoveQuestion(ctx context.Context, id string) (err error)
	UpdateQuestion(ctx context.Context, question *entity.Question, Cols []string) (err error)
	GetQuestion(ctx context.Context, id string) (question *entity.Question, exist bool, err error)
	GetQuestionList(ctx context.Context, question *entity.Question) (questions []*entity.Question, err error)
	GetQuestionPage(ctx context.Context, page, pageSize int, userID, tagID, orderCond string, inDays int) (
		questionList []*entity.Question, total int64, err error)
	UpdateQuestionStatus(ctx context.Context, questionID string, status int) (err error)
	UpdateQuestionStatusWithOutUpdateTime(ctx context.Context, question *entity.Question) (err error)
	RecoverQuestion(ctx context.Context, questionID string) (err error)
	UpdateQuestionOperation(ctx context.Context, question *entity.Question) (err error)
	GetQuestionsByTitle(ctx context.Context, title string, pageSize int) (questionList []*entity.Question, err error)
	UpdatePvCount(ctx context.Context, questionID string) (err error)
	UpdateAnswerCount(ctx context.Context, questionID string, num int) (err error)
	UpdateCollectionCount(ctx context.Context, questionID string, num int) (err error)
	UpdateAccepted(ctx context.Context, question *entity.Question) (err error)
	UpdateLastAnswer(ctx context.Context, question *entity.Question) (err error)
	FindByID(ctx context.Context, id []string) (questionList []*entity.Question, err error)
	AdminQuestionPage(ctx context.Context, search *schema.AdminQuestionPageReq) ([]*entity.Question, int64, error)
	GetQuestionCount(ctx context.Context) (count int64, err error)
	GetUserQuestionCount(ctx context.Context, userID string) (count int64, err error)
	SitemapQuestions(ctx context.Context, page, pageSize int) (questionIDList []*schema.SiteMapQuestionInfo, err error)
	RemoveAllUserQuestion(ctx context.Context, userID string) (err error)
}

// QuestionCommon user service
type QuestionCommon struct {
	questionRepo         QuestionRepo
	answerRepo           answercommon.AnswerRepo
	voteRepo             activity_common.VoteRepo
	followCommon         activity_common.FollowRepo
	tagCommon            *tagcommon.TagCommonService
	userCommon           *usercommon.UserCommon
	collectionCommon     *collectioncommon.CollectionCommon
	AnswerCommon         *answercommon.AnswerCommon
	metaService          *meta.MetaService
	configService        *config.ConfigService
	activityQueueService activity_queue.ActivityQueueService
	data                 *data.Data
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
	configService *config.ConfigService,
	activityQueueService activity_queue.ActivityQueueService,
	data *data.Data,
) *QuestionCommon {
	return &QuestionCommon{
		questionRepo:         questionRepo,
		answerRepo:           answerRepo,
		voteRepo:             voteRepo,
		followCommon:         followCommon,
		tagCommon:            tagCommon,
		userCommon:           userCommon,
		collectionCommon:     collectionCommon,
		AnswerCommon:         answerCommon,
		metaService:          metaService,
		configService:        configService,
		activityQueueService: activityQueueService,
		data:                 data,
	}
}

func (qs *QuestionCommon) GetUserQuestionCount(ctx context.Context, userID string) (count int64, err error) {
	return qs.questionRepo.GetUserQuestionCount(ctx, userID)
}

func (qs *QuestionCommon) UpdatePv(ctx context.Context, questionID string) error {
	return qs.questionRepo.UpdatePvCount(ctx, questionID)
}

func (qs *QuestionCommon) UpdateAnswerCount(ctx context.Context, questionID string) error {
	count, err := qs.answerRepo.GetCountByQuestionID(ctx, questionID)
	if err != nil {
		return err
	}
	return qs.questionRepo.UpdateAnswerCount(ctx, questionID, int(count))
}

func (qs *QuestionCommon) UpdateCollectionCount(ctx context.Context, questionID string, num int) error {
	return qs.questionRepo.UpdateCollectionCount(ctx, questionID, num)
}

func (qs *QuestionCommon) UpdateAccepted(ctx context.Context, questionID, AnswerID string) error {
	question := &entity.Question{}
	question.ID = questionID
	question.AcceptedAnswerID = AnswerID
	return qs.questionRepo.UpdateAccepted(ctx, question)
}

func (qs *QuestionCommon) UpdateLastAnswer(ctx context.Context, questionID, AnswerID string) error {
	question := &entity.Question{}
	question.ID = questionID
	question.LastAnswerID = AnswerID
	return qs.questionRepo.UpdateLastAnswer(ctx, question)
}

func (qs *QuestionCommon) UpdatePostTime(ctx context.Context, questionID string) error {
	questioninfo := &entity.Question{}
	now := time.Now()
	questioninfo.ID = questionID
	questioninfo.PostUpdateTime = now
	return qs.questionRepo.UpdateQuestion(ctx, questioninfo, []string{"post_update_time"})
}
func (qs *QuestionCommon) UpdatePostSetTime(ctx context.Context, questionID string, setTime time.Time) error {
	questioninfo := &entity.Question{}
	questioninfo.ID = questionID
	questioninfo.PostUpdateTime = setTime
	return qs.questionRepo.UpdateQuestion(ctx, questioninfo, []string{"post_update_time"})
}

func (qs *QuestionCommon) FindInfoByID(ctx context.Context, questionIDs []string, loginUserID string) (map[string]*schema.QuestionInfo, error) {
	list := make(map[string]*schema.QuestionInfo)
	questionList, err := qs.questionRepo.FindByID(ctx, questionIDs)
	if err != nil {
		return list, err
	}
	questions, err := qs.FormatQuestions(ctx, questionList, loginUserID)
	if err != nil {
		return list, err
	}
	for _, item := range questions {
		list[item.ID] = item
	}
	return list, nil
}

func (qs *QuestionCommon) InviteUserInfo(ctx context.Context, questionID string) (inviteList []*schema.UserBasicInfo, err error) {
	InviteUserInfo := make([]*schema.UserBasicInfo, 0)
	dbinfo, has, err := qs.questionRepo.GetQuestion(ctx, questionID)
	if err != nil {
		return InviteUserInfo, err
	}
	if !has {
		return InviteUserInfo, errors.NotFound(reason.QuestionNotFound)
	}
	//InviteUser
	if dbinfo.InviteUserID != "" {
		InviteUserIDs := make([]string, 0)
		err := json.Unmarshal([]byte(dbinfo.InviteUserID), &InviteUserIDs)
		if err == nil {
			inviteUserInfoMap, err := qs.userCommon.BatchUserBasicInfoByID(ctx, InviteUserIDs)
			if err == nil {
				for _, userid := range InviteUserIDs {
					_, ok := inviteUserInfoMap[userid]
					if ok {
						InviteUserInfo = append(InviteUserInfo, inviteUserInfoMap[userid])
					}
				}
			}
		}
	}
	return InviteUserInfo, nil
}

func (qs *QuestionCommon) Info(ctx context.Context, questionID string, loginUserID string) (showinfo *schema.QuestionInfo, err error) {
	dbinfo, has, err := qs.questionRepo.GetQuestion(ctx, questionID)
	if err != nil {
		return showinfo, err
	}
	dbinfo.ID = uid.DeShortID(dbinfo.ID)
	if !has {
		return showinfo, errors.NotFound(reason.QuestionNotFound)
	}
	showinfo = qs.ShowFormat(ctx, dbinfo)

	if showinfo.Status == 2 {
		var metainfo *entity.Meta
		metainfo, err = qs.metaService.GetMetaByObjectIdAndKey(ctx, dbinfo.ID, entity.QuestionCloseReasonKey)
		if err != nil {
			log.Error(err)
		} else {
			// metainfo.Value
			closemsg := &schema.CloseQuestionMeta{}
			err = json.Unmarshal([]byte(metainfo.Value), closemsg)
			if err != nil {
				log.Error("json.Unmarshal CloseQuestionMeta error", err.Error())
			} else {
				cfg, err := qs.configService.GetConfigByID(ctx, closemsg.CloseType)
				if err != nil {
					log.Error("json.Unmarshal QuestionCloseJson error", err.Error())
				} else {
					reasonItem := &schema.ReasonItem{}
					_ = json.Unmarshal(cfg.GetByteValue(), reasonItem)
					reasonItem.Translate(cfg.Key, handler.GetLangByCtx(ctx))
					operation := &schema.Operation{}
					operation.Type = reasonItem.Name
					operation.Description = reasonItem.Description
					operation.Msg = closemsg.CloseMsg
					operation.Time = metainfo.CreatedAt.Unix()
					operation.Level = schema.OperationLevelInfo
					showinfo.Operation = operation
				}
			}
		}
	}

	tagmap, err := qs.tagCommon.GetObjectTag(ctx, questionID)
	if err != nil {
		return showinfo, err
	}
	showinfo.Tags = tagmap

	userIds := make([]string, 0)
	if checker.IsNotZeroString(dbinfo.UserID) {
		userIds = append(userIds, dbinfo.UserID)
	}
	if checker.IsNotZeroString(dbinfo.LastEditUserID) {
		userIds = append(userIds, dbinfo.LastEditUserID)
	}
	if checker.IsNotZeroString(showinfo.LastAnsweredUserID) {
		userIds = append(userIds, showinfo.LastAnsweredUserID)
	}
	userInfoMap, err := qs.userCommon.BatchUserBasicInfoByID(ctx, userIds)
	if err != nil {
		return showinfo, err
	}

	_, ok := userInfoMap[dbinfo.UserID]
	if ok {
		showinfo.UserInfo = userInfoMap[dbinfo.UserID]
	}
	_, ok = userInfoMap[dbinfo.LastEditUserID]
	if ok {
		showinfo.UpdateUserInfo = userInfoMap[dbinfo.LastEditUserID]
	}
	_, ok = userInfoMap[showinfo.LastAnsweredUserID]
	if ok {
		showinfo.LastAnsweredUserInfo = userInfoMap[showinfo.LastAnsweredUserID]
	}

	if loginUserID == "" {
		return showinfo, nil
	}

	showinfo.VoteStatus = qs.voteRepo.GetVoteStatus(ctx, questionID, loginUserID)

	// // check is followed
	isFollowed, _ := qs.followCommon.IsFollowed(ctx, loginUserID, questionID)
	showinfo.IsFollowed = isFollowed

	has, err = qs.AnswerCommon.SearchAnswered(ctx, loginUserID, dbinfo.ID)
	if err != nil {
		log.Error("AnswerFunc.SearchAnswered", err)
	}
	showinfo.Answered = has

	collectedMap, err := qs.collectionCommon.SearchObjectCollected(ctx, loginUserID, []string{dbinfo.ID})
	if err != nil {
		return nil, err
	}
	if len(collectedMap) > 0 {
		showinfo.Collected = true
	}
	return showinfo, nil
}

func (qs *QuestionCommon) FormatQuestionsPage(
	ctx context.Context, questionList []*entity.Question, loginUserID string, orderCond string) (
	formattedQuestions []*schema.QuestionPageResp, err error) {
	formattedQuestions = make([]*schema.QuestionPageResp, 0)
	questionIDs := make([]string, 0)
	userIDs := make([]string, 0)
	for _, questionInfo := range questionList {
		t := &schema.QuestionPageResp{
			ID:               questionInfo.ID,
			CreatedAt:        questionInfo.CreatedAt.Unix(),
			Title:            questionInfo.Title,
			UrlTitle:         htmltext.UrlTitle(questionInfo.Title),
			Description:      htmltext.FetchExcerpt(questionInfo.ParsedText, "...", 240),
			Status:           questionInfo.Status,
			ViewCount:        questionInfo.ViewCount,
			UniqueViewCount:  questionInfo.UniqueViewCount,
			VoteCount:        questionInfo.VoteCount,
			AnswerCount:      questionInfo.AnswerCount,
			CollectionCount:  questionInfo.CollectionCount,
			FollowCount:      questionInfo.FollowCount,
			AcceptedAnswerID: questionInfo.AcceptedAnswerID,
			LastAnswerID:     questionInfo.LastAnswerID,
			Pin:              questionInfo.Pin,
			Show:             questionInfo.Show,
		}

		questionIDs = append(questionIDs, questionInfo.ID)
		userIDs = append(userIDs, questionInfo.UserID)
		haveEdited, haveAnswered := false, false
		if checker.IsNotZeroString(questionInfo.LastEditUserID) {
			haveEdited = true
			userIDs = append(userIDs, questionInfo.LastEditUserID)
		}
		if checker.IsNotZeroString(questionInfo.LastAnswerID) {
			haveAnswered = true

			answerInfo, exist, err := qs.answerRepo.GetAnswer(ctx, questionInfo.LastAnswerID)
			if err == nil && exist {
				if answerInfo.LastEditUserID != "0" {
					t.LastAnsweredUserID = answerInfo.LastEditUserID
				} else {
					t.LastAnsweredUserID = answerInfo.UserID
				}
				t.LastAnsweredAt = answerInfo.CreatedAt
				userIDs = append(userIDs, t.LastAnsweredUserID)
			}
		}

		// if order condition is newest or nobody edited or nobody answered, only show question author
		if orderCond == schema.QuestionOrderCondNewest || (!haveEdited && !haveAnswered) {
			t.OperationType = schema.QuestionPageRespOperationTypeAsked
			t.OperatedAt = questionInfo.CreatedAt.Unix()
			t.Operator = &schema.QuestionPageRespOperator{ID: questionInfo.UserID}
		} else {
			// if no one
			if haveEdited {
				t.OperationType = schema.QuestionPageRespOperationTypeModified
				t.OperatedAt = questionInfo.UpdatedAt.Unix()
				t.Operator = &schema.QuestionPageRespOperator{ID: questionInfo.LastEditUserID}
			}

			if haveAnswered {
				if t.LastAnsweredAt.Unix() > t.OperatedAt {
					t.OperationType = schema.QuestionPageRespOperationTypeAnswered
					t.OperatedAt = t.LastAnsweredAt.Unix()
					t.Operator = &schema.QuestionPageRespOperator{ID: t.LastAnsweredUserID}
				}
			}
		}
		formattedQuestions = append(formattedQuestions, t)
	}

	tagsMap, err := qs.tagCommon.BatchGetObjectTag(ctx, questionIDs)
	if err != nil {
		return formattedQuestions, err
	}
	userInfoMap, err := qs.userCommon.BatchUserBasicInfoByID(ctx, userIDs)
	if err != nil {
		return formattedQuestions, err
	}

	for _, item := range formattedQuestions {
		tags, ok := tagsMap[item.ID]
		if ok {
			item.Tags = tags
		} else {
			item.Tags = make([]*schema.TagResp, 0)
		}
		userInfo, ok := userInfoMap[item.Operator.ID]
		if ok {
			if userInfo != nil {
				item.Operator.DisplayName = userInfo.DisplayName
				item.Operator.Username = userInfo.Username
				item.Operator.Rank = userInfo.Rank
				item.Operator.Status = userInfo.Status
			}
		}

	}
	return formattedQuestions, nil
}

func (qs *QuestionCommon) FormatQuestions(ctx context.Context, questionList []*entity.Question, loginUserID string) ([]*schema.QuestionInfo, error) {
	list := make([]*schema.QuestionInfo, 0)
	objectIds := make([]string, 0)
	userIds := make([]string, 0)

	for _, questionInfo := range questionList {
		item := qs.ShowFormat(ctx, questionInfo)
		list = append(list, item)
		objectIds = append(objectIds, item.ID)
		userIds = append(userIds, item.UserID, item.LastEditUserID, item.LastAnsweredUserID)
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
		item.Tags = tagsMap[item.ID]
		item.UserInfo = userInfoMap[item.UserID]
		item.UpdateUserInfo = userInfoMap[item.LastEditUserID]
		item.LastAnsweredUserInfo = userInfoMap[item.LastAnsweredUserID]
	}
	if loginUserID == "" {
		return list, nil
	}

	collectedMap, err := qs.collectionCommon.SearchObjectCollected(ctx, loginUserID, objectIds)
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		item.Collected = collectedMap[item.ID]
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

	if questionInfo.Status == entity.QuestionStatusDeleted {
		return nil
	}

	questionInfo.Status = entity.QuestionStatusDeleted
	err = qs.questionRepo.UpdateQuestionStatus(ctx, questionInfo.ID, questionInfo.Status)
	if err != nil {
		return err
	}

	userQuestionCount, err := qs.GetUserQuestionCount(ctx, questionInfo.UserID)
	if err != nil {
		log.Error("user GetUserQuestionCount error", err.Error())
	} else {
		err = qs.userCommon.UpdateQuestionCount(ctx, questionInfo.UserID, userQuestionCount)
		if err != nil {
			log.Error("user IncreaseQuestionCount error", err.Error())
		}
	}

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
	questionInfo.Status = entity.QuestionStatusClosed
	err = qs.questionRepo.UpdateQuestionStatus(ctx, questionInfo.ID, questionInfo.Status)
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

	qs.activityQueueService.Send(ctx, &schema.ActivityMsg{
		UserID:           questionInfo.UserID,
		ObjectID:         questionInfo.ID,
		OriginalObjectID: questionInfo.ID,
		ActivityTypeKey:  constant.ActQuestionClosed,
	})
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

	// user add question count

	err = as.UpdateAnswerCount(ctx, answerinfo.QuestionID)
	if err != nil {
		log.Error("UpdateAnswerCount error", err.Error())
	}
	userAnswerCount, err := as.answerRepo.GetCountByUserID(ctx, answerinfo.UserID)
	if err != nil {
		log.Error("GetCountByUserID error", err.Error())
	}
	err = as.userCommon.UpdateAnswerCount(ctx, answerinfo.UserID, int(userAnswerCount))
	if err != nil {
		log.Error("user UpdateAnswerCount error", err.Error())
	}

	return as.answerRepo.RemoveAnswer(ctx, id)
}

func (qs *QuestionCommon) SitemapCron(ctx context.Context) {
	questionNum, err := qs.questionRepo.GetQuestionCount(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	if questionNum <= constant.SitemapMaxSize {
		_, err = qs.questionRepo.SitemapQuestions(ctx, 1, int(questionNum))
		if err != nil {
			log.Errorf("get site map question error: %v", err)
		}
		return
	}

	totalPages := int(math.Ceil(float64(questionNum) / float64(constant.SitemapMaxSize)))
	for i := 1; i <= totalPages; i++ {
		_, err = qs.questionRepo.SitemapQuestions(ctx, i, constant.SitemapMaxSize)
		if err != nil {
			log.Errorf("get site map question error: %v", err)
			return
		}
	}
}

func (qs *QuestionCommon) SetCache(ctx context.Context, cachekey string, info interface{}) error {
	infoStr, err := json.Marshal(info)
	if err != nil {
		return errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}

	err = qs.data.Cache.SetString(ctx, cachekey, string(infoStr), schema.DashboardCacheTime)
	if err != nil {
		return errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	return nil
}

func (qs *QuestionCommon) ShowListFormat(ctx context.Context, data *entity.Question) *schema.QuestionInfo {
	return qs.ShowFormat(ctx, data)
}

func (qs *QuestionCommon) ShowFormat(ctx context.Context, data *entity.Question) *schema.QuestionInfo {
	info := schema.QuestionInfo{}
	info.ID = data.ID
	info.Title = data.Title
	info.UrlTitle = htmltext.UrlTitle(data.Title)
	info.Content = data.OriginalText
	info.HTML = data.ParsedText
	info.ViewCount = data.ViewCount
	info.UniqueViewCount = data.UniqueViewCount
	info.VoteCount = data.VoteCount
	info.AnswerCount = data.AnswerCount
	info.CollectionCount = data.CollectionCount
	info.FollowCount = data.FollowCount
	info.AcceptedAnswerID = data.AcceptedAnswerID
	info.LastAnswerID = data.LastAnswerID
	info.CreateTime = data.CreatedAt.Unix()
	info.UpdateTime = data.UpdatedAt.Unix()
	info.PostUpdateTime = data.PostUpdateTime.Unix()
	if data.PostUpdateTime.Unix() < 1 {
		info.PostUpdateTime = 0
	}
	info.QuestionUpdateTime = data.UpdatedAt.Unix()
	if data.UpdatedAt.Unix() < 1 {
		info.QuestionUpdateTime = 0
	}
	info.Status = data.Status
	info.Pin = data.Pin
	info.Show = data.Show
	info.UserID = data.UserID
	info.LastEditUserID = data.LastEditUserID
	if data.LastAnswerID != "0" {
		answerInfo, exist, err := qs.answerRepo.GetAnswer(ctx, data.LastAnswerID)
		if err == nil && exist {
			if answerInfo.LastEditUserID != "0" {
				info.LastAnsweredUserID = answerInfo.LastEditUserID
			} else {
				info.LastAnsweredUserID = answerInfo.UserID
			}
		}

	}
	info.Tags = make([]*schema.TagResp, 0)
	return &info
}
func (qs *QuestionCommon) ShowFormatWithTag(ctx context.Context, data *entity.QuestionWithTagsRevision) *schema.QuestionInfo {
	info := qs.ShowFormat(ctx, &data.Question)
	Tags := make([]*schema.TagResp, 0)
	for _, tag := range data.Tags {
		item := &schema.TagResp{}
		item.SlugName = tag.SlugName
		item.DisplayName = tag.DisplayName
		item.Recommend = tag.Recommend
		item.Reserved = tag.Reserved
		Tags = append(Tags, item)
	}
	info.Tags = Tags
	return info
}
