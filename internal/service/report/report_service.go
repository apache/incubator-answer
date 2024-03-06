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

package report

import (
	"encoding/json"

	"github.com/apache/incubator-answer/internal/base/constant"
	"github.com/apache/incubator-answer/internal/base/handler"
	"github.com/apache/incubator-answer/internal/base/pager"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	answercommon "github.com/apache/incubator-answer/internal/service/answer_common"
	"github.com/apache/incubator-answer/internal/service/comment_common"
	"github.com/apache/incubator-answer/internal/service/config"
	"github.com/apache/incubator-answer/internal/service/object_info"
	questioncommon "github.com/apache/incubator-answer/internal/service/question_common"
	"github.com/apache/incubator-answer/internal/service/report_common"
	"github.com/apache/incubator-answer/internal/service/report_handle"
	usercommon "github.com/apache/incubator-answer/internal/service/user_common"
	"github.com/apache/incubator-answer/pkg/htmltext"
	"github.com/apache/incubator-answer/pkg/obj"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
	"golang.org/x/net/context"
)

// ReportService user service
type ReportService struct {
	reportRepo        report_common.ReportRepo
	objectInfoService *object_info.ObjService
	commonUser        *usercommon.UserCommon
	answerRepo        answercommon.AnswerRepo
	questionRepo      questioncommon.QuestionRepo
	commentCommonRepo comment_common.CommentCommonRepo
	reportHandle      *report_handle.ReportHandle
	configService     *config.ConfigService
}

// NewReportService new report service
func NewReportService(
	reportRepo report_common.ReportRepo,
	objectInfoService *object_info.ObjService,
	commonUser *usercommon.UserCommon,
	answerRepo answercommon.AnswerRepo,
	questionRepo questioncommon.QuestionRepo,
	commentCommonRepo comment_common.CommentCommonRepo,
	reportHandle *report_handle.ReportHandle,
	configService *config.ConfigService,
) *ReportService {
	return &ReportService{
		reportRepo:        reportRepo,
		objectInfoService: objectInfoService,
		commonUser:        commonUser,
		answerRepo:        answerRepo,
		questionRepo:      questionRepo,
		commentCommonRepo: commentCommonRepo,
		reportHandle:      reportHandle,
		configService:     configService,
	}
}

// AddReport add report
func (rs *ReportService) AddReport(ctx context.Context, req *schema.AddReportReq) (err error) {
	objectTypeNumber, err := obj.GetObjectTypeNumberByObjectID(req.ObjectID)
	if err != nil {
		return err
	}

	// TODO this reported user id should be get by revision
	objInfo, err := rs.objectInfoService.GetInfo(ctx, req.ObjectID)
	if err != nil {
		return err
	}

	report := &entity.Report{
		UserID:         req.UserID,
		ReportedUserID: objInfo.ObjectCreatorUserID,
		ObjectID:       req.ObjectID,
		ObjectType:     objectTypeNumber,
		ReportType:     req.ReportType,
		Content:        req.Content,
		Status:         entity.ReportStatusPending,
	}
	return rs.reportRepo.AddReport(ctx, report)
}

// GetUnreviewedReportPostPage get unreviewed report post page
func (rs *ReportService) GetUnreviewedReportPostPage(ctx context.Context, req *schema.GetUnreviewedReportPostPageReq) (
	pageModel *pager.PageModel, err error) {
	reports, total, err := rs.reportRepo.GetReportListPage(ctx, &schema.GetReportListPageDTO{
		Page:     req.Page,
		PageSize: 1,
		Status:   entity.ReportStatusPending,
	})
	if err != nil {
		return
	}

	resp := make([]*schema.GetReportListPageResp, 0)
	_ = copier.Copy(&resp, reports)
	var flaggedUserIds, userIds []string
	for _, r := range resp {
		flaggedUserIds = append(flaggedUserIds, r.ReportedUserID)
		userIds = append(userIds, r.UserID)
		r.Format()
	}

	// flagged users
	flaggedUsers, err := rs.commonUser.BatchUserBasicInfoByID(ctx, flaggedUserIds)
	if err != nil {
		return nil, err
	}

	// flag users
	users, err := rs.commonUser.BatchUserBasicInfoByID(ctx, userIds)
	if err != nil {
		return nil, err
	}
	for _, r := range resp {
		r.ReportedUser = flaggedUsers[r.ReportedUserID]
		r.ReportUser = users[r.UserID]
		rs.decorateReportResp(ctx, r)
	}
	return pager.NewPageModel(total, resp), nil
}

// ReviewReport review report
func (rs *ReportService) ReviewReport(ctx context.Context, req *schema.ReviewReportReq) (err error) {
	report, exist, err := rs.reportRepo.GetByID(ctx, req.ReportID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.NotFound(reason.ReportNotFound)
	}
	// check if handle or not
	if report.Status != entity.ReportStatusPending {
		return nil
	}

	// ignore this report
	if req.OperationType == constant.ReportOperationIgnoreReport {
		return rs.reportRepo.UpdateStatus(ctx, report.ID, entity.ReportStatusIgnore)
	}

	if err = rs.reportHandle.UpdateReportedObject(ctx, report, req); err != nil {
		return
	}

	return rs.reportRepo.UpdateStatus(ctx, report.ID, entity.ReportStatusCompleted)
}

func (rs *ReportService) decorateReportResp(ctx context.Context, resp *schema.GetReportListPageResp) {
	lang := handler.GetLangByCtx(ctx)
	objectInfo, err := rs.objectInfoService.GetInfo(ctx, resp.ObjectID)
	if err != nil {
		log.Error(err)
		return
	}

	resp.QuestionID = objectInfo.QuestionID
	resp.AnswerID = objectInfo.AnswerID
	resp.CommentID = objectInfo.CommentID
	resp.Title = objectInfo.Title
	resp.Excerpt = htmltext.FetchExcerpt(objectInfo.Content, "...", 240)

	if resp.ReportType > 0 {
		resp.Reason = &schema.ReasonItem{ReasonType: resp.ReportType}
		cf, err := rs.configService.GetConfigByID(ctx, resp.ReportType)
		if err != nil {
			log.Error(err)
		} else {
			_ = json.Unmarshal([]byte(cf.Value), resp.Reason)
			resp.Reason.Translate(cf.Key, lang)
		}
	}
	if resp.FlaggedType > 0 {
		resp.FlaggedReason = &schema.ReasonItem{ReasonType: resp.FlaggedType}
		cf, err := rs.configService.GetConfigByID(ctx, resp.FlaggedType)
		if err != nil {
			log.Error(err)
		} else {
			_ = json.Unmarshal([]byte(cf.Value), resp.Reason)
			resp.Reason.Translate(cf.Key, lang)
		}
	}
}
