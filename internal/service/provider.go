package service

import (
	"github.com/answerdev/answer/internal/service/action"
	"github.com/answerdev/answer/internal/service/activity"
	"github.com/answerdev/answer/internal/service/activity_common"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	"github.com/answerdev/answer/internal/service/auth"
	collectioncommon "github.com/answerdev/answer/internal/service/collection_common"
	"github.com/answerdev/answer/internal/service/comment"
	"github.com/answerdev/answer/internal/service/comment_common"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/dashboard"
	"github.com/answerdev/answer/internal/service/export"
	"github.com/answerdev/answer/internal/service/follow"
	"github.com/answerdev/answer/internal/service/meta"
	"github.com/answerdev/answer/internal/service/notification"
	notficationcommon "github.com/answerdev/answer/internal/service/notification_common"
	"github.com/answerdev/answer/internal/service/object_info"
	"github.com/answerdev/answer/internal/service/plugin_common"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/answerdev/answer/internal/service/reason"
	"github.com/answerdev/answer/internal/service/report"
	"github.com/answerdev/answer/internal/service/report_admin"
	"github.com/answerdev/answer/internal/service/report_handle_admin"
	"github.com/answerdev/answer/internal/service/revision_common"
	"github.com/answerdev/answer/internal/service/role"
	"github.com/answerdev/answer/internal/service/search_parser"
	"github.com/answerdev/answer/internal/service/siteinfo"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/answerdev/answer/internal/service/tag"
	tagcommon "github.com/answerdev/answer/internal/service/tag_common"
	"github.com/answerdev/answer/internal/service/uploader"
	"github.com/answerdev/answer/internal/service/user_admin"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/internal/service/user_external_login"
	"github.com/google/wire"
)

// ProviderSetService is providers.
var ProviderSetService = wire.NewSet(
	comment.NewCommentService,
	comment_common.NewCommentCommonService,
	report.NewReportService,
	NewVoteService,
	tag.NewTagService,
	follow.NewFollowService,
	NewCollectionGroupService,
	NewCollectionService,
	action.NewCaptchaService,
	auth.NewAuthService,
	NewUserService,
	NewQuestionService,
	NewAnswerService,
	export.NewEmailService,
	tagcommon.NewTagCommonService,
	usercommon.NewUserCommon,
	questioncommon.NewQuestionCommon,
	answercommon.NewAnswerCommon,
	uploader.NewUploaderService,
	collectioncommon.NewCollectionCommon,
	revision_common.NewRevisionService,
	NewRevisionService,
	rank.NewRankService,
	search_parser.NewSearchParser,
	NewSearchService,
	meta.NewMetaService,
	object_info.NewObjService,
	report_handle_admin.NewReportHandle,
	report_admin.NewReportAdminService,
	user_admin.NewUserAdminService,
	reason.NewReasonService,
	siteinfo_common.NewSiteInfoCommonService,
	siteinfo.NewSiteInfoService,
	notficationcommon.NewNotificationCommon,
	notification.NewNotificationService,
	activity.NewAnswerActivityService,
	dashboard.NewDashboardService,
	activity_common.NewActivityCommon,
	activity.NewActivityService,
	role.NewRoleService,
	role.NewUserRoleRelService,
	role.NewRolePowerRelService,
	user_external_login.NewUserExternalLoginService,
	user_external_login.NewUserCenterLoginService,
	plugin_common.NewPluginCommonService,
	config.NewConfigService,
)
