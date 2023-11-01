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

package collection

import (
	"context"

	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/pager"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service"
	"github.com/segmentfault/pacman/errors"
)

// collectionGroupRepo collectionGroup repository
type collectionGroupRepo struct {
	data *data.Data
}

// NewCollectionGroupRepo new repository
func NewCollectionGroupRepo(data *data.Data) service.CollectionGroupRepo {
	return &collectionGroupRepo{
		data: data,
	}
}

// AddCollectionGroup add collection group
func (cr *collectionGroupRepo) AddCollectionGroup(ctx context.Context, collectionGroup *entity.CollectionGroup) (err error) {
	_, err = cr.data.DB.Context(ctx).Insert(collectionGroup)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// AddCollectionDefaultGroup add collection group
func (cr *collectionGroupRepo) AddCollectionDefaultGroup(ctx context.Context, userID string) (collectionGroup *entity.CollectionGroup, err error) {
	defaultGroup := &entity.CollectionGroup{
		Name:         "default",
		DefaultGroup: schema.CGDefault,
		UserID:       userID,
	}
	_, err = cr.data.DB.Context(ctx).Insert(defaultGroup)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	}
	collectionGroup = defaultGroup
	return
}

// UpdateCollectionGroup update collection group
func (cr *collectionGroupRepo) UpdateCollectionGroup(ctx context.Context, collectionGroup *entity.CollectionGroup, cols []string) (err error) {
	_, err = cr.data.DB.Context(ctx).ID(collectionGroup.ID).Cols(cols...).Update(collectionGroup)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetCollectionGroup get collection group one
func (cr *collectionGroupRepo) GetCollectionGroup(ctx context.Context, id string) (
	collectionGroup *entity.CollectionGroup, exist bool, err error,
) {
	collectionGroup = &entity.CollectionGroup{}
	exist, err = cr.data.DB.Context(ctx).ID(id).Get(collectionGroup)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetCollectionGroupPage get collection group page
func (cr *collectionGroupRepo) GetCollectionGroupPage(ctx context.Context, page, pageSize int, collectionGroup *entity.CollectionGroup) (collectionGroupList []*entity.CollectionGroup, total int64, err error) {
	collectionGroupList = make([]*entity.CollectionGroup, 0)

	session := cr.data.DB.Context(ctx)
	if collectionGroup.UserID != "" && collectionGroup.UserID != "0" {
		session = session.Where("user_id = ?", collectionGroup.UserID)
	}
	session = session.OrderBy("update_time desc")

	total, err = pager.Help(page, pageSize, collectionGroupList, collectionGroup, session)
	err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	return
}

func (cr *collectionGroupRepo) GetDefaultID(ctx context.Context, userID string) (collectionGroup *entity.CollectionGroup, has bool, err error) {
	collectionGroup = &entity.CollectionGroup{}
	has, err = cr.data.DB.Context(ctx).Where("user_id =? and  default_group = ?", userID, schema.CGDefault).Get(collectionGroup)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	}
	return
}
