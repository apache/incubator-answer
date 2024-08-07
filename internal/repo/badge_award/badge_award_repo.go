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
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/service/badge_award"
	"github.com/apache/incubator-answer/internal/service/unique"
	"time"
)

type badgeAwardRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

func NewBadgeAwardRepo(data *data.Data, uniqueIDRepo unique.UniqueIDRepo) badge_award.BadgeAwardRepo {
	return &badgeAwardRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

func (r *badgeAwardRepo) Award(ctx context.Context, badgeID string, userID string, objectID string, force bool, createdAt time.Time) {
	return
}
func (r *badgeAwardRepo) CheckIsAward(ctx context.Context, badgeID string, userID string, objectID string) (isAward bool) {
	return
}
func (r *badgeAwardRepo) CountByUserIdAndBadgeLevel(ctx context.Context, userID string, badgeLevel entity.BadgeLevel) (awardCount int64) {
	return
}
func (r *badgeAwardRepo) CountByUserId(ctx context.Context, userID string) (awardCount int64) {
	return
}
func (r *badgeAwardRepo) CountByUserIdAndBadgeId(ctx context.Context, userID string, badgeID string) (awardCount int64) {
	return
}
func (r *badgeAwardRepo) CountByObjectId(ctx context.Context, objectID string) (awardCount int64) {
	return
}
func (r *badgeAwardRepo) CountByObjectIdAndBadgeId(ctx context.Context, objectID string, badgeID string) (awardCount int64) {
	return
}
func (r *badgeAwardRepo) CountBadgesByUserIdAndObjectId(ctx context.Context, userID string, objectID string, badgeID string) (awardCount int64) {
	return
}
func (r *badgeAwardRepo) SumUserEarnedGroupByBadgeID(ctx context.Context, userID string) (earnedCounts []*entity.BadgeEarnedCount, err error) {
	err = r.data.DB.Context(ctx).Select("badge_id, count(`id`) AS earned_count").Where("user_id = ?", userID).GroupBy("badge_id").Find(&earnedCounts)
	return
}
func (r *badgeAwardRepo) ListAllByUserId(ctx context.Context, userID string) (badgeAwards []*entity.BadgeAward) {
	return
}
func (r *badgeAwardRepo) ListPagedByBadgeId(ctx context.Context, badgeID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64) {
	return
}
func (r *badgeAwardRepo) ListPagedByBadgeIdAndUserId(ctx context.Context, badgeID string, userID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64) {
	return
}
func (r *badgeAwardRepo) ListPagedByObjectId(ctx context.Context, badgeID string, objectID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64) {
	return
}
func (r *badgeAwardRepo) ListPagedByObjectIdAndUserId(ctx context.Context, badgeID string, objectID string, userID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64) {
	return
}
func (r *badgeAwardRepo) ListTagPagedByBadgeId(ctx context.Context, badgeIDs []int64, page int64, pageSize int64, filterUserID int64) (badgeAwards []*entity.BadgeAward, total int64) {
	return
}
func (r *badgeAwardRepo) ListTagPagedByBadgeIdAndUserId(ctx context.Context, badgeIDs []int64, userID string, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64) {
	return
}
func (r *badgeAwardRepo) ListPagedLatest(ctx context.Context, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64) {
	return
}
func (r *badgeAwardRepo) ListNewestEarnedByLevel(ctx context.Context, userID string, level entity.BadgeLevel, num int64) (badgeAwards []*entity.BadgeAward, total int64) {
	return
}
func (r *badgeAwardRepo) ListNewestByUserIdAndLevel(ctx context.Context, userID string, level int64, page int64, pageSize int64) (badgeAwards []*entity.BadgeAward, total int64) {
	return
}
func (r *badgeAwardRepo) GetByUserIdAndBadgeId(ctx context.Context, userID string, badgeID string) (badgeAward *entity.BadgeAward) {
	return
}
func (r *badgeAwardRepo) GetByUserIdAndBadgeIdAndObjectId(ctx context.Context, userID string, badgeID string, objectID string) (badgeAward *entity.BadgeAward) {
	return
}
