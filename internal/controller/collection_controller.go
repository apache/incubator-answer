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

package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
)

// CollectionController collection controller
type CollectionController struct {
	collectionService *service.CollectionService
}

// NewCollectionController new controller
func NewCollectionController(collectionService *service.CollectionService) *CollectionController {
	return &CollectionController{collectionService: collectionService}
}

// CollectionSwitch add collection
// @Summary add collection
// @Description add collection
// @Tags Collection
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.CollectionSwitchReq true "collection"
// @Success 200 {object} handler.RespBody{data=schema.CollectionSwitchResp}
// @Router /answer/api/v1/collection/switch [post]
func (cc *CollectionController) CollectionSwitch(ctx *gin.Context) {
	req := &schema.CollectionSwitchReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.ObjectID = uid.DeShortID(req.ObjectID)

	dto := &schema.CollectionSwitchDTO{}
	_ = copier.Copy(dto, req)

	dto.UserID = middleware.GetLoginUserIDFromContext(ctx)

	if converter.StringToInt64(req.ObjectID) < 1 {
		return
	}
	if converter.StringToInt64(dto.UserID) < 1 {
		handler.HandleResponse(ctx, errors.Unauthorized(reason.UnauthorizedError), nil)
		return
	}

	resp, err := cc.collectionService.CollectionSwitch(ctx, dto)
	handler.HandleResponse(ctx, err, resp)
}
