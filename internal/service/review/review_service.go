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

package review

import (
	"context"

	"github.com/apache/incubator-answer/internal/base/constant"
	"github.com/apache/incubator-answer/internal/base/pager"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	answercommon "github.com/apache/incubator-answer/internal/service/answer_common"
	"github.com/apache/incubator-answer/internal/service/notice_queue"
	"github.com/apache/incubator-answer/internal/service/object_info"
	questioncommon "github.com/apache/incubator-answer/internal/service/question_common"
	"github.com/apache/incubator-answer/internal/service/role"
	tagcommon "github.com/apache/incubator-answer/internal/service/tag_common"
	usercommon "github.com/apache/incubator-answer/internal/service/user_common"
	"github.com/apache/incubator-answer/pkg/token"
	"github.com/apache/incubator-answer/pkg/uid"
	"github.com/apache/incubator-answer/plugin"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// ReviewRepo review repository
type ReviewRepo interface {
	AddReview(ctx context.Context, review *entity.Review) (err error)
	UpdateReviewStatus(ctx context.Context, reviewID int, reviewerUserID string, status int) (err error)
	GetReview(ctx context.Context, reviewID int) (review *entity.Review, exist bool, err error)
	GetReviewCount(ctx context.Context, status int) (count int64, err error)
	GetReviewPage(ctx context.Context, page, pageSize int, cond *entity.Review) (reviewList []*entity.Review, total int64, err error)
}

// ReviewService user service
type ReviewService struct {
	reviewRepo                       ReviewRepo
	objectInfoService                *object_info.ObjService
	userCommon                       *usercommon.UserCommon
	userRepo                         usercommon.UserRepo
	questionRepo                     questioncommon.QuestionRepo
	answerRepo                       answercommon.AnswerRepo
	userRoleService                  *role.UserRoleRelService
	tagCommon                        *tagcommon.TagCommonService
	externalNotificationQueueService notice_queue.ExternalNotificationQueueService
	notificationQueueService         notice_queue.NotificationQueueService
}

// NewReviewService new review service
func NewReviewService(
	reviewRepo ReviewRepo,
	objectInfoService *object_info.ObjService,
	userCommon *usercommon.UserCommon,
	questionRepo questioncommon.QuestionRepo,
	answerRepo answercommon.AnswerRepo,
	userRoleService *role.UserRoleRelService,
	externalNotificationQueueService notice_queue.ExternalNotificationQueueService,
	tagCommon *tagcommon.TagCommonService,
	notificationQueueService notice_queue.NotificationQueueService,
) *ReviewService {
	return &ReviewService{
		reviewRepo:                       reviewRepo,
		objectInfoService:                objectInfoService,
		userCommon:                       userCommon,
		questionRepo:                     questionRepo,
		answerRepo:                       answerRepo,
		userRoleService:                  userRoleService,
		externalNotificationQueueService: externalNotificationQueueService,
		tagCommon:                        tagCommon,
		notificationQueueService:         notificationQueueService,
	}
}

// AddQuestionReview add review for question if needed
func (cs *ReviewService) AddQuestionReview(ctx context.Context,
	question *entity.Question, tags []*schema.TagItem) (needReview bool) {
	reviewContent := &plugin.ReviewContent{
		ObjectType: constant.QuestionObjectType,
		Title:      question.Title,
		Content:    question.ParsedText,
	}
	for _, tag := range tags {
		reviewContent.Tags = append(reviewContent.Tags, tag.SlugName)
	}
	reviewContent.Author = cs.getReviewContentAuthorInfo(ctx, question.UserID)
	return cs.callPluginToReview(ctx, question.UserID, question.ID, reviewContent)
}

// AddAnswerReview add review for answer if needed
func (cs *ReviewService) AddAnswerReview(ctx context.Context,
	answer *entity.Answer) (needReview bool) {
	reviewContent := &plugin.ReviewContent{
		ObjectType: constant.AnswerObjectType,
		Content:    answer.ParsedText,
	}
	reviewContent.Author = cs.getReviewContentAuthorInfo(ctx, answer.UserID)
	return cs.callPluginToReview(ctx, answer.UserID, answer.ID, reviewContent)
}

// get review content author info
func (cs *ReviewService) getReviewContentAuthorInfo(ctx context.Context, userID string) (author plugin.ReviewContentAuthor) {
	user, exist, err := cs.userCommon.GetUserBasicInfoByID(ctx, userID)
	if err != nil {
		log.Errorf("get user info failed, err: %v", err)
		return
	}
	if !exist {
		log.Errorf("user not found by id: %s", userID)
		return
	}
	author.Rank = user.Rank
	author.ApprovedQuestionAmount, _ = cs.questionRepo.GetUserQuestionCount(ctx, userID)
	author.ApprovedAnswerAmount, _ = cs.answerRepo.GetCountByUserID(ctx, userID)
	author.Role, _ = cs.userRoleService.GetUserRole(ctx, userID)
	return
}

// call plugin to review
func (cs *ReviewService) callPluginToReview(ctx context.Context, userID, objectID string,
	reviewContent *plugin.ReviewContent) (approved bool) {
	// As default, no need review
	approved = true
	objectID = uid.DeShortID(objectID)

	r := &entity.Review{
		UserID:         userID,
		ObjectID:       objectID,
		ObjectType:     constant.ObjectTypeStrMapping[reviewContent.ObjectType],
		ReviewerUserID: "0",
		Status:         entity.ReviewStatusPending,
	}

	_ = plugin.CallReviewer(func(reviewer plugin.Reviewer) error {
		// If one of the reviewer plugin return false, then the review is not approved
		if !approved {
			return nil
		}
		if result := reviewer.Review(reviewContent); !result.Approved {
			approved = false
			r.Reason = result.Reason
			r.Submitter = reviewer.Info().SlugName
		}
		return nil
	})

	if !approved {
		if err := cs.reviewRepo.AddReview(ctx, r); err != nil {
			log.Errorf("add review failed, err: %v", err)
		}
	}
	return approved
}

// UpdateReview update review
func (cs *ReviewService) UpdateReview(ctx context.Context, req *schema.UpdateReviewReq) (err error) {
	review, exist, err := cs.reviewRepo.GetReview(ctx, req.ReviewID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.ObjectNotFound)
	}
	if review.Status != entity.ReviewStatusPending {
		return nil
	}

	if err = cs.updateObjectStatus(ctx, review, req.IsApprove()); err != nil {
		return err
	}

	if req.IsApprove() {
		err = cs.reviewRepo.UpdateReviewStatus(ctx, req.ReviewID, req.UserID, entity.ReviewStatusApproved)
	} else {
		err = cs.reviewRepo.UpdateReviewStatus(ctx, req.ReviewID, req.UserID, entity.ReviewStatusRejected)
	}
	return
}

// update object status
func (cs *ReviewService) updateObjectStatus(ctx context.Context, review *entity.Review, isApprove bool) (err error) {
	objectType := constant.ObjectTypeNumberMapping[review.ObjectType]
	switch objectType {
	case constant.QuestionObjectType:
		question, exist, err := cs.questionRepo.GetQuestion(ctx, review.ObjectID)
		if err != nil {
			return err
		}
		if !exist {
			return errors.BadRequest(reason.ObjectNotFound)
		}
		if isApprove {
			question.Status = entity.QuestionStatusAvailable
		} else {
			question.Status = entity.QuestionStatusDeleted
		}
		if err := cs.questionRepo.UpdateQuestionStatus(ctx, question.ID, question.Status); err != nil {
			return err
		}
		if isApprove {
			tags, err := cs.tagCommon.GetObjectEntityTag(ctx, question.ID)
			if err != nil {
				log.Errorf("get question tags failed, err: %v", err)
			}
			cs.externalNotificationQueueService.Send(ctx,
				schema.CreateNewQuestionNotificationMsg(question.ID, question.Title, question.UserID, tags))
		}
	case constant.AnswerObjectType:
		answerInfo, exist, err := cs.answerRepo.GetAnswer(ctx, review.ObjectID)
		if err != nil {
			return err
		}
		if !exist {
			return errors.BadRequest(reason.ObjectNotFound)
		}
		if isApprove {
			answerInfo.Status = entity.AnswerStatusAvailable
		} else {
			answerInfo.Status = entity.AnswerStatusDeleted
		}
		if err := cs.answerRepo.UpdateAnswerStatus(ctx, answerInfo.ID, answerInfo.Status); err != nil {
			return err
		}
		questionInfo, exist, err := cs.questionRepo.GetQuestion(ctx, answerInfo.QuestionID)
		if err != nil {
			return err
		}
		if !exist {
			return errors.BadRequest(reason.ObjectNotFound)
		}
		if isApprove {
			cs.notificationAnswerTheQuestion(ctx, questionInfo.UserID, questionInfo.ID, answerInfo.ID,
				answerInfo.UserID, questionInfo.Title, answerInfo.OriginalText)
		}
	}
	return
}

func (cs *ReviewService) notificationAnswerTheQuestion(ctx context.Context,
	questionUserID, questionID, answerID, answerUserID, questionTitle, answerSummary string) {
	// If the question is answered by me, there is no notification for myself.
	if questionUserID == answerUserID {
		return
	}
	msg := &schema.NotificationMsg{
		TriggerUserID:  answerUserID,
		ReceiverUserID: questionUserID,
		Type:           schema.NotificationTypeInbox,
		ObjectID:       answerID,
	}
	msg.ObjectType = constant.AnswerObjectType
	msg.NotificationAction = constant.NotificationAnswerTheQuestion
	cs.notificationQueueService.Send(ctx, msg)

	receiverUserInfo, exist, err := cs.userRepo.GetByUserID(ctx, questionUserID)
	if err != nil {
		log.Error(err)
		return
	}
	if !exist {
		log.Warnf("user %s not found", questionUserID)
		return
	}

	externalNotificationMsg := &schema.ExternalNotificationMsg{
		ReceiverUserID: receiverUserInfo.ID,
		ReceiverEmail:  receiverUserInfo.EMail,
		ReceiverLang:   receiverUserInfo.Language,
	}
	rawData := &schema.NewAnswerTemplateRawData{
		QuestionTitle:   questionTitle,
		QuestionID:      questionID,
		AnswerID:        answerID,
		AnswerSummary:   answerSummary,
		UnsubscribeCode: token.GenerateToken(),
	}
	answerUser, _, _ := cs.userCommon.GetUserBasicInfoByID(ctx, answerUserID)
	if answerUser != nil {
		rawData.AnswerUserDisplayName = answerUser.DisplayName
	}
	externalNotificationMsg.NewAnswerTemplateRawData = rawData
	cs.externalNotificationQueueService.Send(ctx, externalNotificationMsg)
}

// GetReviewPendingCount get review pending count
func (cs *ReviewService) GetReviewPendingCount(ctx context.Context) (count int64, err error) {
	return cs.reviewRepo.GetReviewCount(ctx, entity.ReviewStatusPending)
}

// GetUnreviewedPostPage get review page
func (cs *ReviewService) GetUnreviewedPostPage(ctx context.Context, req *schema.GetUnreviewedPostPageReq) (
	pageModel *pager.PageModel, err error) {
	cond := &entity.Review{
		ObjectID: req.ObjectID,
		Status:   entity.ReviewStatusPending,
	}
	reviewList, total, err := cs.reviewRepo.GetReviewPage(ctx, req.Page, 1, cond)
	if err != nil {
		return
	}

	resp := make([]*schema.GetUnreviewedPostPageResp, 0)
	for _, review := range reviewList {
		info, err := cs.objectInfoService.GetUnreviewedRevisionInfo(ctx, review.ObjectID)
		if err != nil {
			log.Errorf("GetUnreviewedRevisionInfo failed, err: %v", err)
			continue
		}

		r := &schema.GetUnreviewedPostPageResp{
			ReviewID:             review.ID,
			CreatedAt:            info.CreatedAt,
			ObjectID:             info.ObjectID,
			QuestionID:           info.QuestionID,
			AnswerID:             info.AnswerID,
			CommentID:            info.CommentID,
			ObjectType:           info.ObjectType,
			Title:                info.Title,
			OriginalText:         info.Content,
			Tags:                 info.Tags,
			ObjectStatus:         info.Status,
			ObjectShowStatus:     info.ShowStatus,
			SubmitAt:             review.CreatedAt.Unix(),
			SubmitterDisplayName: req.ReviewerMapping[review.Submitter],
			Reason:               review.Reason,
		}

		// get user info
		userInfo, exists, e := cs.userCommon.GetUserBasicInfoByID(ctx, info.ObjectCreatorUserID)
		if e != nil {
			log.Errorf("user not found by id: %s, err: %v", info.ObjectCreatorUserID, e)
		}
		if exists {
			_ = copier.Copy(&r.AuthorUserInfo, userInfo)
		}
		resp = append(resp, r)
	}
	return pager.NewPageModel(total, resp), nil
}
