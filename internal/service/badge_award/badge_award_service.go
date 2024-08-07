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
	"time"
)

type BadgeAwardRepo interface {
	Award(ctx context.Context, badgeID string, userID string, objectID string, force bool, createdAt time.Time)
	CheckIsAward(ctx context.Context, badgeID string, userID string, objectID string) bool

	CountByUserIdAndBadgeLevel(ctx context.Context, userID string, badgeLevel entity.BadgeLevel) (awardCount int64)
	CountByUserId(ctx context.Context, userID string) (awardCount int64)
	CountByUserIdAndBadgeId(ctx context.Context, userID string, badgeID string) (awardCount int64)
	CountByObjectId(ctx context.Context, objectID string) (awardCount int64)
	CountByObjectIdAndBadgeId(ctx context.Context, objectID string, badgeID string) int64
	CountBadgesByUserIdAndObjectId(ctx context.Context, userID string, objectID string, badgeID string) (awardCount int64)

	SumUserEarnedGroupByBadgeID(ctx context.Context, userID string) (earnedCounts []*entity.BadgeEarnedCount, err error)

	ListAllByUserId(ctx context.Context, userID string) (badgeAwards []*entity.BadgeAward)
	ListPagedByBadgeId(ctx context.Context, badgeID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64)
	ListPagedByBadgeIdAndUserId(ctx context.Context, badgeID string, userID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64)
	ListPagedByObjectId(ctx context.Context, badgeID string, objectID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64)
	ListPagedByObjectIdAndUserId(ctx context.Context, badgeID string, objectID string, userID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64)
	ListTagPagedByBadgeId(ctx context.Context, badgeIDs []int64, page int64, pageSize int64, filterUserID int64) (badgeAwards []*entity.BadgeAward, total int64)
	ListTagPagedByBadgeIdAndUserId(ctx context.Context, badgeIDs []int64, userID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64)
	ListPagedLatest(ctx context.Context, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64)
	ListNewestEarnedByLevel(ctx context.Context, userID string, level entity.BadgeLevel, num int64) (badgeAwards []*entity.BadgeAward, total int64)
	ListNewestByUserIdAndLevel(ctx context.Context, userID string, level int64, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64)

	GetByUserIdAndBadgeId(ctx context.Context, userID string, badgeID string) (badgeAward *entity.BadgeAward)
	GetByUserIdAndBadgeIdAndObjectId(ctx context.Context, userID string, badgeID string, objectID string) (badgeAward *entity.BadgeAward)
}

type BadgeAwardService struct {
	badgeAwardRepo BadgeAwardRepo
}

func NewBadgeAwardService(badgeAwardRepo BadgeAwardRepo) *BadgeAwardService {
	return &BadgeAwardService{
		badgeAwardRepo: badgeAwardRepo,
	}
}
