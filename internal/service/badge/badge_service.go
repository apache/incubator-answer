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
	"github.com/apache/incubator-answer/internal/base/translator"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service/badge_award"
	"github.com/apache/incubator-answer/internal/service/badge_group"
	"github.com/apache/incubator-answer/pkg/converter"
)

type BadgeRepo interface {
	ListByLevel(ctx context.Context, level entity.BadgeLevel) ([]*entity.Badge, error)
	ListByGroup(ctx context.Context, groupID int64) ([]*entity.Badge, error)
	ListByLevelAndGroup(ctx context.Context, level entity.BadgeLevel, groupID int64) ([]*entity.Badge, error)
	ListActivated(ctx context.Context) ([]*entity.Badge, error)
	ListInactivated(ctx context.Context) ([]*entity.Badge, error)
}

type BadgeService struct {
	badgeRepo      BadgeRepo
	badgeGroupRepo badge_group.BadgeGroupRepo
	badgeAwardRepo badge_award.BadgeAwardRepo
}

func NewBadgeService(
	badgeRepo BadgeRepo,
	badgeGroupRepo badge_group.BadgeGroupRepo,
	badgeAwardRepo badge_award.BadgeAwardRepo) *BadgeService {
	return &BadgeService{
		badgeRepo:      badgeRepo,
		badgeGroupRepo: badgeGroupRepo,
		badgeAwardRepo: badgeAwardRepo,
	}
}

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
	earnedCounts, err = b.badgeAwardRepo.SumUserEarnedGroupByBadgeID(ctx, userID)

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

		badgesMap[badge.BadgeGroupId] = append(badgesMap[badge.BadgeGroupId], &schema.BadgeListInfo{
			ID:         badge.ID,
			Name:       translator.Tr(handler.GetLangByCtx(ctx), badge.Name),
			Icon:       badge.Icon,
			AwardCount: badge.AwardCount,
			Earned:     earned,
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
