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
	"encoding/json"

	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service/meta"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/builder"
	"xorm.io/xorm"
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

// AddOrUpdateMetaByObjectIdAndKey if exist record with same objectID and key, update it. Or create a new one
func (mr *metaRepo) AddOrUpdateMetaByObjectIdAndKey(ctx context.Context, req *schema.UpdateReactionReq) (schema.ReactSummaryMeta, error) {
	result, err := mr.data.DB.Transaction(func(session *xorm.Session) (interface{}, error) {
		session = session.Context(ctx)

		// 1. acquire meta entity with target object id and key
		metaEntity := &entity.Meta{}
		exist, err := mr.data.DB.Context(ctx).Where(builder.Eq{"object_id": req.ObjectID}.And(builder.Eq{"`key`": entity.ObjectReactSummaryKey})).ForUpdate().Get(metaEntity)
		if err != nil {
			return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		}

		var reaction schema.ReactSummaryMeta
		// if not exist, create new one
		if !exist {
			reaction = schema.ReactSummaryMeta{}
			return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		} else {
			err = json.Unmarshal([]byte(metaEntity.Value), &reaction)
			if err != nil {
				return nil, err
			}
		}

		// update reaction
		mr.updateReaction(req, reaction)

		// write back to meta repo
		reactSumBytes, err := json.Marshal(reaction)
		if err != nil {
			return nil, err
		}

		metaObj := &entity.Meta{
			ObjectID: req.ObjectID,
			Key:      entity.ObjectReactSummaryKey,
			Value:    string(reactSumBytes),
		}
		if exist {
			_, err = session.Update(metaObj)
		} else {
			_, err = session.Insert(metaObj)
		}

		return reaction, err
	})

	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err)
	}

	if ret, ok := result.(schema.ReactSummaryMeta); ok {
		return ret, nil
	} else {
		return nil, errors.InternalServer(reason.UnknownError).WithMsg("Unable to cast to schema.ReactSummaryMeta.")
	}
}

// updateReaction update reaction
func (mr *metaRepo) updateReaction(req *schema.UpdateReactionReq, reaction schema.ReactSummaryMeta) {
	emojiUserIds, ok := reaction[req.Emoji]

	if !ok {
		emojiUserIds = make([]string, 0)
	}

	found := false
	for _, item := range emojiUserIds {
		if item == req.UserID {
			found = true
			break
		}
	}

	removeItem := func(arr []string, target string) []string {
		result := make([]string, 0, len(arr))

		for _, item := range arr {
			if item != target {
				result = append(result, item)
			}
		}

		return result
	}

	if req.Reaction == "activate" && !found {
		emojiUserIds = append(emojiUserIds, req.UserID)
	} else if req.Reaction == "deactivate" && found {
		emojiUserIds = removeItem(emojiUserIds, req.UserID)
	}

	reaction[req.Emoji] = emojiUserIds
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
