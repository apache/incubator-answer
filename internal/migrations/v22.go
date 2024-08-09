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
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/repo/unique"
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
			Name:         "badge.default_badges.autobiographer.name",
			Icon:         "person-badge-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.autobiographer.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.editor.name",
			Icon:         "pencil-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.editor.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "question",
			Handler:      "FirstQuestion",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.first_flag.name",
			Icon:         "flag-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.first_flag.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.first_upvote.name",
			Icon:         "hand-thumbs-up-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.first_upvote.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.first_reaction.name",
			Icon:         "emoji-smile-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.first_reaction.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.first_share.name",
			Icon:         "share-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.first_share.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.scholar.name",
			Icon:         "check-circle-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.scholar.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 1,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.solved.name",
			Icon:         "check-square-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.solved.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 2,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.nice_answer.name",
			Icon:         "chat-square-text-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.nice_answer.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 3,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeMultiAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.good_answer.name",
			Icon:         "chat-square-text-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.good_answer.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 3,
			Level:        entity.BadgeLevelSilver,
			Single:       entity.BadgeMultiAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.great_answer.name",
			Icon:         "chat-square-text-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.great_answer.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 3,
			Level:        entity.BadgeLevelGold,
			Single:       entity.BadgeMultiAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.nice_question.name",
			Icon:         "question-circle-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.nice_question.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 3,
			Level:        entity.BadgeLevelBronze,
			Single:       entity.BadgeMultiAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.good_question.name",
			Icon:         "question-circle-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.good_question.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 3,
			Level:        entity.BadgeLevelSilver,
			Single:       entity.BadgeSingleAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
		{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Name:         "badge.default_badges.great_question.name",
			Icon:         "question-circle-fill",
			AwardCount:   0,
			Description:  "badge.default_badges.great_question.desc",
			Status:       entity.BadgeStatusActive,
			BadgeGroupID: 3,
			Level:        entity.BadgeLevelGold,
			Single:       entity.BadgeMultiAward,
			Collect:      "",
			Handler:      "",
			Param:        "",
		},
	}
)

func addBadges(ctx context.Context, x *xorm.Engine) (err error) {
	uniqueIDRepo := unique.NewUniqueIDRepo(&data.Data{DB: x})
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
	for _, badge := range defaultBadgeTable {
		badge.ID, err = uniqueIDRepo.GenUniqueIDStr(ctx, entity.Badge{}.TableName())
		if err != nil {
			return
		}
		_, err = x.Context(ctx).Insert(badge)
		if err != nil {
			return
		}
	}
	return
}
