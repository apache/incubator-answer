package service

import (
	"github.com/google/wire"
	"github.com/segmentfault/answer/internal/service/action"
	"github.com/segmentfault/answer/internal/service/activity"
	answercommon "github.com/segmentfault/answer/internal/service/answer_common"
	"github.com/segmentfault/answer/internal/service/auth"
	collectioncommon "github.com/segmentfault/answer/internal/service/collection_common"
	"github.com/segmentfault/answer/internal/service/comment"
	"github.com/segmentfault/answer/internal/service/comment_common"
	"github.com/segmentfault/answer/internal/service/export"
	"github.com/segmentfault/answer/internal/service/follow"
	"github.com/segmentfault/answer/internal/service/meta"
	"github.com/segmentfault/answer/internal/service/notification"
	notficationcommon "github.com/segmentfault/answer/internal/service/notification_common"
	"github.com/segmentfault/answer/internal/service/object_info"
	questioncommon "github.com/segmentfault/answer/internal/service/question_common"
	"github.com/segmentfault/answer/internal/service/rank"
	"github.com/segmentfault/answer/internal/service/reason"
	"github.com/segmentfault/answer/internal/service/report"
	"github.com/segmentfault/answer/internal/service/report_backyard"
	"github.com/segmentfault/answer/internal/service/report_handle_backyard"
	"github.com/segmentfault/answer/internal/service/revision_common"
	"github.com/segmentfault/answer/internal/service/tag"
	tagcommon "github.com/segmentfault/answer/internal/service/tag_common"
	"github.com/segmentfault/answer/internal/service/uploader"
	"github.com/segmentfault/answer/internal/service/user_backyard"
	usercommon "github.com/segmentfault/answer/internal/service/user_common"
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
	NewSearchService,
	meta.NewMetaService,
	object_info.NewObjService,
	report_handle_backyard.NewReportHandle,
	report_backyard.NewReportBackyardService,
	user_backyard.NewUserBackyardService,
	reason.NewReasonService,
	NewSiteInfoService,
	notficationcommon.NewNotificationCommon,
	notification.NewNotificationService,
	activity.NewAnswerActivityService,
)
