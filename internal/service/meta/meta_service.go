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
	"errors"

	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	usercommon "github.com/apache/incubator-answer/internal/service/user_common"
	myErrors "github.com/segmentfault/pacman/errors"
)

// MetaRepo meta repository
type MetaRepo interface {
	AddMeta(ctx context.Context, meta *entity.Meta) (err error)
	RemoveMeta(ctx context.Context, id int) (err error)
	UpdateMeta(ctx context.Context, meta *entity.Meta) (err error)
	GetMetaByObjectIdAndKey(ctx context.Context, objectId, key string) (meta *entity.Meta, exist bool, err error)
	GetMetaList(ctx context.Context, meta *entity.Meta) (metas []*entity.Meta, err error)
}

// MetaService user service
type MetaService struct {
	metaRepo   MetaRepo
	userCommon *usercommon.UserCommon
}

func NewMetaService(metaRepo MetaRepo, userCommon *usercommon.UserCommon) *MetaService {
	return &MetaService{
		metaRepo:   metaRepo,
		userCommon: userCommon,
	}
}

// AddMeta add meta
func (ms *MetaService) AddMeta(ctx context.Context, objID, key, value string) (err error) {
	meta := &entity.Meta{
		ObjectID: objID,
		Key:      key,
		Value:    value,
	}
	return ms.metaRepo.AddMeta(ctx, meta)
}

// RemoveMeta delete meta
func (ms *MetaService) RemoveMeta(ctx context.Context, id int) (err error) {
	return ms.metaRepo.RemoveMeta(ctx, id)
}

// UpdateMeta update meta
func (ms *MetaService) UpdateMeta(ctx context.Context, metaID int, key, value string) (err error) {
	meta := &entity.Meta{
		ID:    metaID,
		Key:   key,
		Value: value,
	}
	return ms.metaRepo.UpdateMeta(ctx, meta)
}

// GetMetaByObjectIdAndKey get meta one
func (ms *MetaService) GetMetaByObjectIdAndKey(ctx context.Context, objectID, key string) (meta *entity.Meta, err error) {
	meta, exist, err := ms.metaRepo.GetMetaByObjectIdAndKey(ctx, objectID, key)
	if err != nil {
		return
	}
	if !exist {
		return nil, myErrors.BadRequest(reason.MetaObjectNotFound)
	}
	return meta, nil
}

// GetMetaList get meta list all
func (ms *MetaService) GetMetaList(ctx context.Context, objID string) (metas []*entity.Meta, err error) {
	metas, err = ms.metaRepo.GetMetaList(ctx, &entity.Meta{ObjectID: objID})
	if err != nil {
		return nil, err
	}
	return metas, err
}

// GetReactionByObjectId get reaction
func (ms *MetaService) GetReactionByObjectId(ctx context.Context, objectID string) (resp schema.ReactionResp, err error) {
	reactionMeta, err := ms.GetMetaByObjectIdAndKey(ctx, objectID, entity.ObjectReactSummaryKey)

	// if not exist, return nil
	if err != nil {
		var pacmanErr *myErrors.Error
		if errors.As(err, &pacmanErr) && pacmanErr.Reason == reason.MetaObjectNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}

	var reaction schema.ReactSummaryMeta
	err = json.Unmarshal([]byte(reactionMeta.Value), &reaction)
	if err != nil {
		return nil, err
	}
	return ms.convertToReactionResp(ctx, reaction)
}

// AddOrUpdateReaction add or update reaction
func (ms *MetaService) AddOrUpdateReaction(ctx context.Context, req *schema.UpdateReactionReq) (resp schema.ReactionResp, err error) {
	// get reaction for this object
	reactionMeta, err := ms.GetMetaByObjectIdAndKey(ctx, req.ObjectID, entity.ObjectReactSummaryKey)

	var reaction schema.ReactSummaryMeta
	if err != nil {
		var pacmanErr *myErrors.Error
		if errors.As(err, &pacmanErr) && pacmanErr.Reason == reason.MetaObjectNotFound {
			// create new reaction summary
			reaction = schema.ReactSummaryMeta{}
		} else {
			return nil, err
		}
	} else {
		// json unmarshal reactionMeta.Value to reaction
		err = json.Unmarshal([]byte(reactionMeta.Value), &reaction)
		if err != nil {
			return nil, err
		}
	}

	// update reaction
	ms.updateReaction(req, reaction)

	// write back to meta repo
	reactSumBytes, err := json.Marshal(reaction)
	if err != nil {
		return nil, err
	}

	if reactionMeta == nil {
		err = ms.AddMeta(ctx, req.ObjectID, entity.ObjectReactSummaryKey, string(reactSumBytes))
	} else {
		err = ms.UpdateMeta(ctx, reactionMeta.ID, entity.ObjectReactSummaryKey, string(reactSumBytes))
	}

	if err != nil {
		return nil, err
	}

	resp, err = ms.convertToReactionResp(ctx, reaction)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// updateReaction update reaction
func (ms *MetaService) updateReaction(req *schema.UpdateReactionReq, reaction schema.ReactSummaryMeta) {
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

	if req.Type == "activate" && !found {
		emojiUserIds = append(emojiUserIds, req.UserID)
	} else if req.Type == "deactivate" && found {
		emojiUserIds = removeItem(emojiUserIds, req.UserID)
	} else if req.Type == "toggle" {
		if found {
			emojiUserIds = removeItem(emojiUserIds, req.UserID)
		} else {
			emojiUserIds = append(emojiUserIds, req.UserID)
		}
	}

	reaction[req.Emoji] = emojiUserIds
}

func (ms *MetaService) convertToReactionResp(ctx context.Context, reaction schema.ReactSummaryMeta) (schema.ReactionResp, error) {
	resp := schema.ReactionResp{}
	// traverse map and convert to username
	for emoji, userIds := range reaction {
		userNames := make([]string, 0)
		userBasicInfos, err := ms.userCommon.BatchUserBasicInfoByID(ctx, userIds)
		if err != nil {
			return nil, err
		}
		// get username
		for _, userBasicInfo := range userBasicInfos {
			userNames = append(userNames, userBasicInfo.Username)
		}
		resp[emoji] = userNames
	}

	return resp, nil
}
