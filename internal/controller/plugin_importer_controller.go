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
	"bytes"
	"io"

	"github.com/apache/incubator-answer/internal/base/handler"
	"github.com/apache/incubator-answer/internal/base/middleware"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/base/translator"
	"github.com/apache/incubator-answer/internal/base/validator"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service/content"
	"github.com/apache/incubator-answer/internal/service/permission"
	"github.com/apache/incubator-answer/internal/service/rank"
	usercommon "github.com/apache/incubator-answer/internal/service/user_common"
	"github.com/apache/incubator-answer/plugin"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type ImporterController struct {
	// userRepo            *usercommon.UserRepo
	questionService     *content.QuestionService
	rankService         *rank.RankService
	rateLimitMiddleware *middleware.RateLimitMiddleware
	userCommon          *usercommon.UserCommon
}

func NewImporterController(
	questionService *content.QuestionService,
	rankService *rank.RankService,
	rateLimitMiddleware *middleware.RateLimitMiddleware,
	userCommon *usercommon.UserCommon,
) *ImporterController {
	return &ImporterController{
		questionService:     questionService,
		rankService:         rankService,
		rateLimitMiddleware: rateLimitMiddleware,
		userCommon:          userCommon,
	}
}

// ImportCommand godoc
// @Summary ImportCommand
// @Description ImportCommand
// @Tags PluginImporter
// @Accept json
// @Produce json
// @Router /answer/api/v1/importer/command [post]
// @Success 200 {object} handler.RespBody{data=[]plugin.importStatus}
func (ipc *ImporterController) ImportCommand(ctx *gin.Context) {
	questionInfo := &plugin.QuestionImporterInfo{}
	if !plugin.ImporterEnabled() {
		log.Errorf("error: plugin is not enabled")
		ctx.JSON(200, gin.H{"text": "plugin is not enabled"})
		return
	}
	body, err := io.ReadAll(ctx.Request.Body)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	cmd := ctx.PostForm("command")
	if cmd != "/ask" {
		log.Errorf("error: Invalid command")
		ctx.JSON(200, gin.H{"text": "Invalid command"})
		return
	}
	_ = plugin.CallImporter(func(importer plugin.Importer) error {
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		questionInfo, err = importer.GetQuestion(ctx)
		return nil
	})
	if err != nil {
		log.Errorf("error: %v", err)
		ctx.JSON(200, gin.H{"text": err.Error()})
		return
	}

	req := &schema.QuestionAdd{}
	errFields := make([]*validator.FormErrorField, 0)
	// errFields := handler.BindAndCheckReturnErr(ctx, req)
	if ctx.IsAborted() {
		return
	}

	// reject, rejectKey := ipc.rateLimitMiddleware.DuplicateRequestRejection(ctx, req)
	// if reject {
	// 	return
	// }
	user_info, exist, err := ipc.userCommon.GetByEmail(ctx, questionInfo.UserEmail)
	if err != nil {
		log.Errorf("error: %v", err)
		ctx.JSON(200, gin.H{"text": err.Error()})
		return
	}
	if !exist {
		log.Errorf("error: User Email not found")
		ctx.JSON(200, gin.H{"text": "User Email not found"})
		return
	}

	// defer func() {
	// 	// If status is not 200 means that the bad request has been returned, so the record should be cleared
	// 	if ctx.Writer.Status() != http.StatusOK {
	// 		ipc.rateLimitMiddleware.DuplicateRequestClear(ctx, rejectKey)
	// 	}
	// }()
	req.UserID = user_info.ID
	req.Title = questionInfo.Title
	req.Content = questionInfo.Content
	req.HTML = "<p>" + questionInfo.Content + "</p>"
	req.Tags = make([]*schema.TagItem, len(questionInfo.Tags))
	for i, tag := range questionInfo.Tags {
		req.Tags[i] = &schema.TagItem{
			SlugName:    tag,
			DisplayName: tag,
		}
	}
	canList, requireRanks, err := ipc.rankService.CheckOperationPermissionsForRanks(ctx, req.UserID, []string{
		permission.QuestionAdd,
		permission.QuestionEdit,
		permission.QuestionDelete,
		permission.QuestionClose,
		permission.QuestionReopen,
		permission.TagUseReservedTag,
		permission.TagAdd,
		permission.LinkUrlLimit,
	})
	if err != nil {
		log.Errorf("error: %v", err)
		ctx.JSON(200, gin.H{"text": err.Error()})
		return
	}
	req.CanAdd = canList[0]
	req.CanEdit = canList[1]
	req.CanDelete = canList[2]
	req.CanClose = canList[3]
	req.CanReopen = canList[4]
	req.CanUseReservedTag = canList[5]
	req.CanAddTag = canList[6]
	if !req.CanAdd {
		ctx.JSON(200, gin.H{"text": "No permission to add question"})
		log.Errorf("error: %v", err)
		return
	}
	hasNewTag, err := ipc.questionService.HasNewTag(ctx, req.Tags)
	if err != nil {
		log.Errorf("error: %v", err)
		ctx.JSON(200, gin.H{"text": err.Error()})
		return
	}
	if !req.CanAddTag && hasNewTag {
		lang := handler.GetLang(ctx)
		msg := translator.TrWithData(lang, reason.NoEnoughRankToOperate, &schema.PermissionTrTplData{Rank: requireRanks[6]})
		log.Errorf("error: %v", msg)
		ctx.JSON(200, gin.H{"text": msg})
		return
	}

	errList, err := ipc.questionService.CheckAddQuestion(ctx, req)
	if err != nil {
		errlist, ok := errList.([]*validator.FormErrorField)
		if ok {
			errFields = append(errFields, errlist...)
		}
	}
	if len(errFields) > 0 {
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), errFields)
		log.Errorf("error: RequestFormatError")
		ctx.JSON(200, gin.H{"text": "RequestFormatError"})
		return
	}
	req.UserAgent = ctx.GetHeader("User-Agent")
	req.IP = ctx.ClientIP()
	resp, err := ipc.questionService.AddQuestion(ctx, req)
	if err != nil {
		errlist, ok := resp.([]*validator.FormErrorField)
		if ok {
			errFields = append(errFields, errlist...)
		}
	}

	if len(errFields) > 0 {
		log.Errorf("error: RequestFormatError")
		ctx.JSON(200, gin.H{"text": "RequestFormatError"})
		return
	}

	ctx.JSON(200, gin.H{"text": "Add Question Successfully"})
}
