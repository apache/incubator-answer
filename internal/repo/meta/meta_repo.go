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

package meta

import (
	"context"

	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/service/meta"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/builder"
)

// metaRepo meta repository
type metaRepo struct {
	data *data.Data
}

// NewMetaRepo new repository
func NewMetaRepo(data *data.Data) meta.MetaRepo {
	return &metaRepo{
		data: data,
	}
}

// AddMeta add meta
func (mr *metaRepo) AddMeta(ctx context.Context, meta *entity.Meta) (err error) {
	_, err = mr.data.DB.Context(ctx).Insert(meta)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// RemoveMeta delete meta
func (mr *metaRepo) RemoveMeta(ctx context.Context, id int) (err error) {
	_, err = mr.data.DB.Context(ctx).ID(id).Delete(&entity.Meta{})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateMeta update meta
func (mr *metaRepo) UpdateMeta(ctx context.Context, meta *entity.Meta) (err error) {
	_, err = mr.data.DB.Context(ctx).ID(meta.ID).Update(meta)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetMetaByObjectIdAndKey get meta one
func (mr *metaRepo) GetMetaByObjectIdAndKey(ctx context.Context, objectID, key string) (
	meta *entity.Meta, exist bool, err error) {
	meta = &entity.Meta{}
	exist, err = mr.data.DB.Context(ctx).Where(builder.Eq{"object_id": objectID}.And(builder.Eq{"`key`": key})).Desc("created_at").Get(meta)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetMetaList get meta list all
func (mr *metaRepo) GetMetaList(ctx context.Context, meta *entity.Meta) (metaList []*entity.Meta, err error) {
	metaList = make([]*entity.Meta, 0)
	err = mr.data.DB.Context(ctx).Find(&metaList, meta)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
