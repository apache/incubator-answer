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

package badge

import (
	"context"
	"github.com/apache/incubator-answer/internal/base/handler"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/base/translator"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service/badge_award"
	"github.com/apache/incubator-answer/internal/service/badge_group"
	"github.com/apache/incubator-answer/pkg/converter"
	"github.com/apache/incubator-answer/pkg/uid"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

type BadgeRepo interface {
	GetByID(ctx context.Context, id string) (badge *entity.Badge, exists bool, err error)
	GetByIDs(ctx context.Context, ids []string) (badges []*entity.Badge, err error)

	ListByLevel(ctx context.Context, level entity.BadgeLevel) ([]*entity.Badge, error)
	ListByGroup(ctx context.Context, groupID int64) ([]*entity.Badge, error)
	ListByLevelAndGroup(ctx context.Context, level entity.BadgeLevel, groupID int64) ([]*entity.Badge, error)
	ListActivated(ctx context.Context) ([]*entity.Badge, error)
	ListInactivated(ctx context.Context) ([]*entity.Badge, error)

	UpdateAwardCount(ctx context.Context, id string, count int64) error
}

type BadgeService struct {
	badgeRepo         BadgeRepo
	badgeGroupRepo    badge_group.BadgeGroupRepo
	badgeAwardRepo    badge_award.BadgeAwardRepo
	badgeEventService *BadgeEventService
}

func NewBadgeService(
	badgeRepo BadgeRepo,
	badgeGroupRepo badge_group.BadgeGroupRepo,
	badgeAwardRepo badge_award.BadgeAwardRepo,
	badgeEventService *BadgeEventService,
) *BadgeService {
	return &BadgeService{
		badgeRepo:         badgeRepo,
		badgeGroupRepo:    badgeGroupRepo,
		badgeAwardRepo:    badgeAwardRepo,
		badgeEventService: badgeEventService,
	}
}

// ListByGroup list all badges group by group
func (b *BadgeService) ListByGroup(ctx context.Context, userID string) (resp []*schema.GetBadgeListResp, err error) {
	var (
		groups       []*entity.BadgeGroup
		badges       []*entity.Badge
		earnedCounts []*entity.BadgeEarnedCount

		groupMap  = make(map[int64]string, 0)
		badgesMap = make(map[int64][]*schema.BadgeListInfo, 0)
	)
	resp = make([]*schema.GetBadgeListResp, 0)

	groups, err = b.badgeGroupRepo.ListGroups(ctx)
	if err != nil {
		return
	}
	badges, err = b.badgeRepo.ListActivated(ctx)
	if err != nil {
		return
	}

	if len(userID) > 0 {
		earnedCounts, err = b.badgeAwardRepo.SumUserEarnedGroupByBadgeID(ctx, userID)
		if err != nil {
			return
		}
	}

	for _, group := range groups {
		groupMap[converter.StringToInt64(group.ID)] = group.Name
	}

	for _, badge := range badges {
		// check is earned
		earned := false
		if len(earnedCounts) > 0 {
			for _, earnedCount := range earnedCounts {
				if badge.ID == earnedCount.BadgeID {
					earned = true
					break
				}
			}
		}

		badgesMap[badge.BadgeGroupID] = append(badgesMap[badge.BadgeGroupID], &schema.BadgeListInfo{
			ID:         uid.EnShortID(badge.ID),
			Name:       translator.Tr(handler.GetLangByCtx(ctx), badge.Name),
			Icon:       badge.Icon,
			AwardCount: badge.AwardCount,
			Earned:     earned,
			Level:      badge.Level,
		})
	}

	for _, group := range groups {
		resp = append(resp, &schema.GetBadgeListResp{
			GroupName: group.Name,
			Badges:    badgesMap[converter.StringToInt64(group.ID)],
		})
	}

	return
}

// GetBadgeInfo get badge info
func (b *BadgeService) GetBadgeInfo(ctx *gin.Context, id string, userID string) (info *schema.GetBadgeInfoResp, err error) {
	var (
		badge       *entity.Badge
		earnedTotal int64 = 0
		exists            = false
	)

	badge, exists, err = b.badgeRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	if !exists || badge.Status == entity.BadgeStatusInactive {
		err = errors.BadRequest(reason.BadgeObjectNotFound)
		return
	}

	if len(userID) > 0 {
		earnedTotal = b.badgeAwardRepo.CountByUserIdAndBadgeId(ctx, userID, badge.ID)
	}

	info = &schema.GetBadgeInfoResp{
		ID:          uid.EnShortID(badge.ID),
		Name:        translator.Tr(handler.GetLangByCtx(ctx), badge.Name),
		Description: translator.Tr(handler.GetLangByCtx(ctx), badge.Description),
		Icon:        badge.Icon,
		AwardCount:  badge.AwardCount,
		EarnedCount: earnedTotal,
		IsSingle:    badge.Single == entity.BadgeSingleAward,
		Level:       badge.Level,
	}
	return
}
