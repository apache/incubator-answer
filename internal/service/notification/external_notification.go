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

package notification

import (
	"context"
	"github.com/apache/incubator-answer/internal/base/constant"
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service/activity_common"
	"github.com/apache/incubator-answer/internal/service/export"
	"github.com/apache/incubator-answer/internal/service/notice_queue"
	"github.com/apache/incubator-answer/internal/service/siteinfo_common"
	usercommon "github.com/apache/incubator-answer/internal/service/user_common"
	"github.com/apache/incubator-answer/internal/service/user_external_login"
	"github.com/apache/incubator-answer/internal/service/user_notification_config"
	"github.com/apache/incubator-answer/pkg/display"
	"github.com/apache/incubator-answer/plugin"
	"github.com/segmentfault/pacman/log"
	"strings"
)

type ExternalNotificationService struct {
	data                       *data.Data
	userNotificationConfigRepo user_notification_config.UserNotificationConfigRepo
	followRepo                 activity_common.FollowRepo
	emailService               *export.EmailService
	userRepo                   usercommon.UserRepo
	notificationQueueService   notice_queue.ExternalNotificationQueueService
	userExternalLoginRepo      user_external_login.UserExternalLoginRepo
	siteInfoService            siteinfo_common.SiteInfoCommonService
}

func NewExternalNotificationService(
	data *data.Data,
	userNotificationConfigRepo user_notification_config.UserNotificationConfigRepo,
	followRepo activity_common.FollowRepo,
	emailService *export.EmailService,
	userRepo usercommon.UserRepo,
	notificationQueueService notice_queue.ExternalNotificationQueueService,
	userExternalLoginRepo user_external_login.UserExternalLoginRepo,
	siteInfoService siteinfo_common.SiteInfoCommonService,
) *ExternalNotificationService {
	n := &ExternalNotificationService{
		data:                       data,
		userNotificationConfigRepo: userNotificationConfigRepo,
		followRepo:                 followRepo,
		emailService:               emailService,
		userRepo:                   userRepo,
		notificationQueueService:   notificationQueueService,
		userExternalLoginRepo:      userExternalLoginRepo,
		siteInfoService:            siteInfoService,
	}
	notificationQueueService.RegisterHandler(n.Handler)
	return n
}

func (ns *ExternalNotificationService) Handler(ctx context.Context, msg *schema.ExternalNotificationMsg) error {
	log.Debugf("try to send external notification %+v", msg)

	if msg.NewQuestionTemplateRawData != nil {
		return ns.handleNewQuestionNotification(ctx, msg)
	}
	if msg.NewCommentTemplateRawData != nil {
		return ns.handleNewCommentNotification(ctx, msg)
	}
	if msg.NewAnswerTemplateRawData != nil {
		return ns.handleNewAnswerNotification(ctx, msg)
	}
	if msg.NewInviteAnswerTemplateRawData != nil {
		return ns.handleInviteAnswerNotification(ctx, msg)
	}
	log.Errorf("unknown notification message: %+v", msg)
	return nil
}

func (ns *ExternalNotificationService) syncNotificationToPlugin(ctx context.Context,
	source constant.NotificationSource, msg *schema.ExternalNotificationMsg) {
	pluginNotificationMsg := &plugin.NotificationMessage{
		ReceiverUserID: msg.ReceiverUserID,
		ReceiverLang:   msg.ReceiverLang,
	}

	if len(msg.ReceiverLang) == 0 {
		userInfo, _, _ := ns.userRepo.GetByUserID(ctx, msg.ReceiverUserID)
		if userInfo != nil && len(userInfo.Language) > 0 {
			pluginNotificationMsg.ReceiverLang = userInfo.Language
		}
	}

	switch source {
	case constant.InboxSource:
		if msg.NewCommentTemplateRawData != nil {
			pluginNotificationMsg.Type = plugin.NewComment
			pluginNotificationMsg.NewCommentNoticeData = plugin.NewCommentNoticeData{
				CommentUserDisplayName: msg.NewCommentTemplateRawData.CommentUserDisplayName,
				QuestionTitle:          msg.NewCommentTemplateRawData.QuestionTitle,
				QuestionID:             msg.NewCommentTemplateRawData.QuestionID,
				AnswerID:               msg.NewCommentTemplateRawData.AnswerID,
				CommentID:              msg.NewCommentTemplateRawData.CommentID,
				CommentSummary:         msg.NewCommentTemplateRawData.CommentSummary,
			}
		} else if msg.NewAnswerTemplateRawData != nil {
			pluginNotificationMsg.Type = plugin.NewAnswer
			pluginNotificationMsg.NewAnswerNoticeData = plugin.NewAnswerNoticeData{
				AnswerUserDisplayName: msg.NewAnswerTemplateRawData.AnswerUserDisplayName,
				QuestionTitle:         msg.NewAnswerTemplateRawData.QuestionTitle,
				QuestionID:            msg.NewAnswerTemplateRawData.QuestionID,
				AnswerID:              msg.NewAnswerTemplateRawData.AnswerID,
				AnswerSummary:         msg.NewAnswerTemplateRawData.AnswerSummary,
			}
		} else if msg.NewInviteAnswerTemplateRawData != nil {
			pluginNotificationMsg.Type = plugin.NewInviteAnswer
			pluginNotificationMsg.NewInviteAnswerNoticeData = plugin.NewInviteAnswerNoticeData{
				InviterDisplayName: msg.NewInviteAnswerTemplateRawData.InviterDisplayName,
				QuestionTitle:      msg.NewInviteAnswerTemplateRawData.QuestionTitle,
				QuestionID:         msg.NewInviteAnswerTemplateRawData.QuestionID,
			}
		}
	case constant.AllNewQuestionSource:
		pluginNotificationMsg.Type = plugin.NewQuestion
		pluginNotificationMsg.NewQuestionNoticeData = ns.newPluginQuestionNotification(ctx, msg)
	case constant.AllNewQuestionForFollowingTagsSource:
		pluginNotificationMsg.Type = plugin.NewQuestionFollowedTag
		pluginNotificationMsg.NewQuestionNoticeData = ns.newPluginQuestionNotification(ctx, msg)
	}

	_ = plugin.CallNotification(func(fn plugin.Notification) error {
		userInfo, exist, err := ns.userExternalLoginRepo.GetByUserID(ctx, fn.Info().SlugName, msg.ReceiverUserID)
		if err != nil {
			log.Errorf("get user external login info failed: %v", err)
			return nil
		}
		if exist {
			pluginNotificationMsg.ExternalID = userInfo.ExternalID
		}
		fn.Notify(pluginNotificationMsg)
		return nil
	})
}

func (ns *ExternalNotificationService) newPluginQuestionNotification(
	ctx context.Context, msg *schema.ExternalNotificationMsg) (raw plugin.NewQuestionNoticeData) {
	raw = plugin.NewQuestionNoticeData{
		QuestionTitle: msg.NewQuestionTemplateRawData.QuestionTitle,
		Tags:          strings.Join(msg.NewQuestionTemplateRawData.Tags, ","),
	}
	siteInfo, err := ns.siteInfoService.GetSiteGeneral(ctx)
	if err != nil {
		return raw
	}
	seoInfo, err := ns.siteInfoService.GetSiteSeo(ctx)
	if err != nil {
		return raw
	}
	raw.QuestionUrl = display.QuestionURL(
		seoInfo.Permalink, siteInfo.SiteUrl,
		msg.NewQuestionTemplateRawData.QuestionID, msg.NewQuestionTemplateRawData.QuestionTitle)
	return raw
}
