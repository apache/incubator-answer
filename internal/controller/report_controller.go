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
	"github.com/apache/incubator-answer/internal/base/handler"
	"github.com/apache/incubator-answer/internal/base/middleware"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/base/translator"
	"github.com/apache/incubator-answer/internal/base/validator"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service/action"
	"github.com/apache/incubator-answer/internal/service/permission"
	"github.com/apache/incubator-answer/internal/service/rank"
	"github.com/apache/incubator-answer/internal/service/report"
	"github.com/apache/incubator-answer/pkg/converter"
	"github.com/apache/incubator-answer/pkg/uid"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// ReportController report controller
type ReportController struct {
	reportService *report.ReportService
	rankService   *rank.RankService
	actionService *action.CaptchaService
}

// NewReportController new controller
func NewReportController(
	reportService *report.ReportService,
	rankService *rank.RankService,
	actionService *action.CaptchaService,
) *ReportController {
	return &ReportController{
		reportService: reportService,
		rankService:   rankService,
		actionService: actionService,
	}
}

// AddReport add report
// @Summary add report
// @Description add report <br> source (question, answer, comment, user)
// @Security ApiKeyAuth
// @Tags Report
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.AddReportReq true "report"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/report [post]
func (rc *ReportController) AddReport(ctx *gin.Context) {
	req := &schema.AddReportReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.ObjectID = uid.DeShortID(req.ObjectID)
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	isAdmin := middleware.GetUserIsAdminModerator(ctx)
	if !isAdmin {
		captchaPass := rc.actionService.ActionRecordVerifyCaptcha(ctx, entity.CaptchaActionReport, req.UserID, req.CaptchaID, req.CaptchaCode)
		if !captchaPass {
			errFields := append([]*validator.FormErrorField{}, &validator.FormErrorField{
				ErrorField: "captcha_code",
				ErrorMsg:   translator.Tr(handler.GetLang(ctx), reason.CaptchaVerificationFailed),
			})
			handler.HandleResponse(ctx, errors.BadRequest(reason.CaptchaVerificationFailed), errFields)
			return
		}
	}

	can, err := rc.rankService.CheckOperationPermission(ctx, req.UserID, permission.ReportAdd, "")
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if !can {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), nil)
		return
	}

	err = rc.reportService.AddReport(ctx, req)
	if !isAdmin {
		rc.actionService.ActionRecordAdd(ctx, entity.CaptchaActionReport, req.UserID)
	}
	handler.HandleResponse(ctx, err, nil)
}

// ListReportPage godoc
// @Summary list report page
// @Description list report records
// @Security ApiKeyAuth
// @Tags Report
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param status query string true "status" Enums(pending, completed)
// @Param object_type query string true "object_type" Enums(all, question,answer,comment)
// @Param page query int false "page size"
// @Param page_size query int false "page size"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/reports/page [get]
func (rc *ReportController) ListReportPage(ctx *gin.Context) {
	var (
		objectType = ctx.Query("object_type")
		status     = ctx.Query("status")
		page       = converter.StringToInt(ctx.DefaultQuery("page", "1"))
		pageSize   = converter.StringToInt(ctx.DefaultQuery("page_size", "20"))
	)

	dto := schema.GetReportListPageDTO{
		ObjectType: objectType,
		Status:     status,
		Page:       page,
		PageSize:   pageSize,
	}

	resp, err := rc.reportService.ListReportPage(ctx, dto)
	if err != nil {
		handler.HandleResponse(ctx, err, schema.ErrTypeModal)
	} else {
		handler.HandleResponse(ctx, err, resp)
	}
}

// Handle godoc
// @Summary handle flag
// @Description handle flag
// @Security ApiKeyAuth
// @Tags Report
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.ReportHandleReq true "flag"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/report/ [put]
func (rc *ReportController) Handle(ctx *gin.Context) {
	req := schema.ReportHandleReq{}
	if handler.BindAndCheck(ctx, &req) {
		return
	}

	err := rc.reportService.HandleReported(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}
