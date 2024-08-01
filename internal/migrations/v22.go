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

package migrations

import (
	"context"
	"github.com/apache/incubator-answer/internal/entity"
	"time"
	"xorm.io/xorm"
)

var (
	defaultBadgeGroupTable = []*entity.BadgeGroup{
		{ID: "1", Name: "Getting Started"},
		{ID: "2", Name: "Community"},
		{ID: "3", Name: "Posting"},
	}

	defaultBadgeTable = []*entity.Badge{
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.badges.autobiographer.name",
			Icon:         "",
			AwardCount:   0,
			Description:  "badge.badges.autobiographer.desc",
			Status:       entity.BadgeStatusAvailable,
			BadgeGroupId: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.badges.editor.name",
			Icon:         "",
			AwardCount:   0,
			Description:  "badge.badges.editor.desc",
			Status:       entity.BadgeStatusAvailable,
			BadgeGroupId: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "question",
			Handler:      "FirstQuestion",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.badges.first_flag.name",
			Icon:         "",
			AwardCount:   0,
			Description:  "badge.badges.first_flag.desc",
			Status:       entity.BadgeStatusAvailable,
			BadgeGroupId: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.badges.first_upvote.name",
			Icon:         "",
			AwardCount:   0,
			Description:  "badge.badges.first_upvote.desc",
			Status:       entity.BadgeStatusAvailable,
			BadgeGroupId: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.badges.first_reaction.name",
			Icon:         "",
			AwardCount:   0,
			Description:  "badge.badges.first_reaction.desc",
			Status:       entity.BadgeStatusAvailable,
			BadgeGroupId: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.badges.first_share.name",
			Icon:         "",
			AwardCount:   0,
			Description:  "badge.badges.first_share.desc",
			Status:       entity.BadgeStatusAvailable,
			BadgeGroupId: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.badges.scholar.name",
			Icon:         "",
			AwardCount:   0,
			Description:  "badge.badges.scholar.desc",
			Status:       entity.BadgeStatusAvailable,
			BadgeGroupId: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.badges.solved.name",
			Icon:         "",
			AwardCount:   0,
			Description:  "badge.badges.solved.desc",
			Status:       entity.BadgeStatusAvailable,
			BadgeGroupId: 2,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
	}
)

func addBadges(ctx context.Context, x *xorm.Engine) (err error) {
	// create table
	err = x.Context(ctx).Sync(new(entity.Badge))
	if err != nil {
		return
	}

	err = x.Context(ctx).Sync(new(entity.BadgeGroup))
	if err != nil {
		return
	}

	err = x.Context(ctx).Sync(new(entity.BadgeAward))
	if err != nil {
		return
	}

	// insert default data
	_, err = x.Context(ctx).Insert(defaultBadgeGroupTable)
	if err != nil {
		return
	}
	_, err = x.Context(ctx).Insert(defaultBadgeTable)
	return
}
