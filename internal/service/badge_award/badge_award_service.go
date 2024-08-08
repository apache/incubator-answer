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

package badge_award

import (
	"context"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	answercommon "github.com/apache/incubator-answer/internal/service/answer_common"
	"github.com/apache/incubator-answer/internal/service/object_info"
	questioncommon "github.com/apache/incubator-answer/internal/service/question_common"
	usercommon "github.com/apache/incubator-answer/internal/service/user_common"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/log"
	"time"
)

type BadgeAwardRepo interface {
	Award(ctx context.Context, badgeID string, userID string, awardKey string, force bool, createdAt time.Time)
	CheckIsAward(ctx context.Context, badgeID string, userID string, awardKey string) bool

	CountByUserIdAndBadgeLevel(ctx context.Context, userID string, badgeLevel entity.BadgeLevel) (awardCount int64)
	CountByUserId(ctx context.Context, userID string) (awardCount int64)
	CountByUserIdAndBadgeId(ctx context.Context, userID string, badgeID string) (awardCount int64)
	CountByObjectId(ctx context.Context, objectID string) (awardCount int64)
	CountByObjectIdAndBadgeId(ctx context.Context, objectID string, badgeID string) (awardCount int64)
	CountBadgesByUserIdAndObjectId(ctx context.Context, userID string, objectID string, badgeID string) (awardCount int64)

	SumUserEarnedGroupByBadgeID(ctx context.Context, userID string) (earnedCounts []*entity.BadgeEarnedCount, err error)

	ListAllByUserId(ctx context.Context, userID string) (badgeAwards []*entity.BadgeAward)
	ListPagedByBadgeId(ctx context.Context, badgeID string, page int, pageSize int) (badgeAwardList []*entity.BadgeAward, total int64, err error)
	ListPagedByBadgeIdAndUserId(ctx context.Context, badgeID string, userID string, page int, pageSize int) (badgeAwards []*entity.BadgeAward, total int64, err error)
	ListPagedByObjectId(ctx context.Context, badgeID string, objectID string, page int, pageSize int) (badgeAwards []*entity.BadgeAward, total int64, err error)
	ListPagedByObjectIdAndUserId(ctx context.Context, badgeID string, objectID string, userID string, page int, pageSize int) (badgeAwards []*entity.BadgeAward, total int64, err error)
	ListTagPagedByBadgeId(ctx context.Context, badgeIDs []string, page int, pageSize int, filterUserID string) (badgeAwards []*entity.BadgeAward, total int64, err error)
	ListTagPagedByBadgeIdAndUserId(ctx context.Context, badgeIDs []string, userID string, page int, pageSize int) (badgeAwards []*entity.BadgeAward, total int64, err error)
	ListPagedLatest(ctx context.Context, page int, pageSize int) (badgeAwards []*entity.BadgeAward, total int64, err error)
	ListNewestEarnedByLevel(ctx context.Context, userID string, level entity.BadgeLevel, num int) (badgeAwards []*entity.BadgeAward, total int64, err error)
	ListNewestByUserIdAndLevel(ctx context.Context, userID string, level int, page int, pageSize int) (badgeAwards []*entity.BadgeAward, total int64, err error)

	GetByUserIdAndBadgeId(ctx context.Context, userID string, badgeID string) (badgeAward *entity.BadgeAward)
	GetByUserIdAndBadgeIdAndObjectId(ctx context.Context, userID string, badgeID string, objectID string) (badgeAward *entity.BadgeAward)
}

type BadgeAwardService struct {
	badgeAwardRepo    BadgeAwardRepo
	userCommon        *usercommon.UserCommon
	objectInfoService *object_info.ObjService
	questionRepo      questioncommon.QuestionRepo
	answerRepo        answercommon.AnswerRepo
}

func NewBadgeAwardService(
	badgeAwardRepo BadgeAwardRepo,
	userCommon *usercommon.UserCommon,
	objectInfoService *object_info.ObjService,
	questionRepo questioncommon.QuestionRepo,
	answerRepo answercommon.AnswerRepo,
) *BadgeAwardService {
	return &BadgeAwardService{
		badgeAwardRepo:    badgeAwardRepo,
		userCommon:        userCommon,
		objectInfoService: objectInfoService,
		questionRepo:      questionRepo,
		answerRepo:        answerRepo,
	}
}

// GetBadgeAwardList get badge award list
func (b *BadgeAwardService) GetBadgeAwardList(
	ctx context.Context, req *schema.GetBadgeAwardWithPageReq,
) (resp []*schema.GetBadgeAwardWithPageResp, total int64, err error) {
	var (
		badgeAwardList []*entity.BadgeAward
	)

	badgeAwardList, total, err = b.badgeAwardRepo.ListPagedByBadgeId(ctx, req.BadgeID, req.Page, req.PageSize)
	if err != nil {
		return
	}

	resp = make([]*schema.GetBadgeAwardWithPageResp, 0, len(badgeAwardList))

	for i, badgeAward := range badgeAwardList {
		var (
			objectID, questionID, answerID, commentID, objectType, urlTitle string
		)

		// if exist object info
		objInfo, e := b.objectInfoService.GetInfo(ctx, badgeAward.AwardKey)
		if e == nil && !objInfo.IsDeleted() {
			objectID = objInfo.ObjectID
			questionID = objInfo.QuestionID
			answerID = objInfo.AnswerID
			commentID = objInfo.CommentID
			objectType = objInfo.ObjectType
			urlTitle = objInfo.Title
		}

		row := &schema.GetBadgeAwardWithPageResp{
			CreatedAt:      badgeAward.CreatedAt.Unix(),
			ObjectID:       objectID,
			QuestionID:     questionID,
			AnswerID:       answerID,
			CommentID:      commentID,
			ObjectType:     objectType,
			UrlTitle:       urlTitle,
			AuthorUserInfo: schema.UserBasicInfo{},
		}

		// get user info
		userInfo, exists, e := b.userCommon.GetUserBasicInfoByID(ctx, badgeAward.UserID)
		if e != nil {
			log.Errorf("user not found by id: %s, err: %v", badgeAward.UserID, e)
		}
		if exists {
			_ = copier.Copy(&row.AuthorUserInfo, userInfo)
		}

		resp[i] = row
	}

	return
}
