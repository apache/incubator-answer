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

	"github.com/apache/incubator-answer/internal/base/pager"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service/object_info"
	usercommon "github.com/apache/incubator-answer/internal/service/user_common"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// ReviewRepo review repository
type ReviewRepo interface {
	AddReview(ctx context.Context, review *entity.Review) (err error)
	UpdateReviewStatus(ctx context.Context, reviewID, reviewerUserID string, status int) (err error)
	GetReview(ctx context.Context, reviewID string) (review *entity.Review, exist bool, err error)
	GetReviewCount(ctx context.Context, status int) (count int64, err error)
	GetReviewPage(ctx context.Context, page, pageSize int, cond *entity.Review) (reviewList []*entity.Review, total int64, err error)
}

// ReviewService user service
type ReviewService struct {
	reviewRepo        ReviewRepo
	objectInfoService *object_info.ObjService
	userCommon        *usercommon.UserCommon
}

// NewReviewService new review service
func NewReviewService(
	reviewRepo ReviewRepo,
	objectInfoService *object_info.ObjService,
	userCommon *usercommon.UserCommon,
) *ReviewService {
	return &ReviewService{
		reviewRepo:        reviewRepo,
		objectInfoService: objectInfoService,
		userCommon:        userCommon,
	}
}

// AddReview add review
func (cs *ReviewService) AddReview(ctx context.Context, review *entity.Review) (err error) {
	return cs.reviewRepo.AddReview(ctx, review)
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

	if req.IsApprove() {
		// TODO update object status
	} else {
		// TODO update object status
	}

	if req.IsApprove() {
		err = cs.reviewRepo.UpdateReviewStatus(ctx, req.ReviewID, req.UserID, entity.ReviewStatusApproved)
	} else {
		err = cs.reviewRepo.UpdateReviewStatus(ctx, req.ReviewID, req.UserID, entity.ReviewStatusRejected)
	}
	return
}

// GetReviewPendingCount get review pending count
func (cs *ReviewService) GetReviewPendingCount(ctx context.Context) (count int64, err error) {
	return cs.reviewRepo.GetReviewCount(ctx, entity.ReviewStatusPending)
}

// GetUnreviewedPostPage get review page
func (cs *ReviewService) GetUnreviewedPostPage(ctx context.Context, req *schema.GetUnreviewedPostPageReq) (
	pageModel *pager.PageModel, err error) {
	reviewList, total, err := cs.reviewRepo.GetReviewPage(ctx, req.Page, 1, &entity.Review{})
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
			ObjectType:           info.ObjectID,
			Title:                info.Title,
			ParsedText:           info.Html,
			Tags:                 info.Tags,
			AuthorUserID:         info.ObjectCreatorUserID,
			SubmitAt:             review.CreatedAt.Unix(),
			SubmitterDisplayName: review.SubmitterDisplayName,
		}

		// get user info
		userInfo, exists, e := cs.userCommon.GetUserBasicInfoByID(ctx, info.ObjectCreatorUserID)
		if e != nil {
			log.Errorf("user not found by id: %s, err: %v", info.ObjectCreatorUserID, e)
		}
		if exists {
			_ = copier.Copy(&r.UserInfo, userInfo)
		}
	}
	return pager.NewPageModel(total, resp), nil
}
