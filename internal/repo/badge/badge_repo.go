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
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/service/badge"
	"github.com/apache/incubator-answer/internal/service/unique"
)

type badgeRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

// NewBadgeRepo creates a new badge repository
func NewBadgeRepo(data *data.Data, uniqueIDRepo unique.UniqueIDRepo) badge.BadgeRepo {
	return &badgeRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

func (r badgeRepo) GetByID(ctx context.Context, id string) (badge *entity.Badge, exists bool, err error) {
	badge = &entity.Badge{}
	exists, err = r.data.DB.Context(ctx).Where("id = ?", id).Get(badge)
	return
}

// ListByLevel returns a list of badges by level
func (r *badgeRepo) ListByLevel(ctx context.Context, level entity.BadgeLevel) (badges []*entity.Badge, err error) {
	badges = make([]*entity.Badge, 0)
	err = r.data.DB.Context(ctx).Where("level = ?", level).Find(&badges)
	return
}

// ListByGroup returns a list of badges by group
func (r *badgeRepo) ListByGroup(ctx context.Context, groupID int64) (badges []*entity.Badge, err error) {
	badges = make([]*entity.Badge, 0)
	err = r.data.DB.Context(ctx).Where("group_id = ?", groupID).Find(&badges)
	return
}

// ListByLevelAndGroup returns a list of badges by level and group
func (r *badgeRepo) ListByLevelAndGroup(ctx context.Context, level entity.BadgeLevel, groupID int64) (badges []*entity.Badge, err error) {
	badges = make([]*entity.Badge, 0)
	err = r.data.DB.Context(ctx).Where("level = ? AND group_id = ?", level, groupID).Find(&badges)
	return
}

// ListActivated returns a list of activated badges
func (r *badgeRepo) ListActivated(ctx context.Context) (badges []*entity.Badge, err error) {
	badges = make([]*entity.Badge, 0)
	err = r.data.DB.Context(ctx).Where("status = ?", entity.BadgeStatusActive).Find(&badges)
	return
}

// ListInactivated returns a list of inactivated badges
func (r *badgeRepo) ListInactivated(ctx context.Context) (badges []*entity.Badge, err error) {
	badges = make([]*entity.Badge, 0)
	err = r.data.DB.Context(ctx).Where("status = ?", entity.BadgeStatusInactive).Find(&badges)
	return
}