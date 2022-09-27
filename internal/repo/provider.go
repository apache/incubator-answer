package repo

import (
	"github.com/google/wire"
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/repo/activity"
	"github.com/segmentfault/answer/internal/repo/activity_common"
	"github.com/segmentfault/answer/internal/repo/auth"
	"github.com/segmentfault/answer/internal/repo/captcha"
	"github.com/segmentfault/answer/internal/repo/collection"
	"github.com/segmentfault/answer/internal/repo/comment"
	"github.com/segmentfault/answer/internal/repo/common"
	"github.com/segmentfault/answer/internal/repo/config"
	"github.com/segmentfault/answer/internal/repo/export"
	"github.com/segmentfault/answer/internal/repo/meta"
	"github.com/segmentfault/answer/internal/repo/notification"
	"github.com/segmentfault/answer/internal/repo/rank"
	"github.com/segmentfault/answer/internal/repo/reason"
	"github.com/segmentfault/answer/internal/repo/report"
	"github.com/segmentfault/answer/internal/repo/revision"
	"github.com/segmentfault/answer/internal/repo/tag"
	"github.com/segmentfault/answer/internal/repo/unique"
	"github.com/segmentfault/answer/internal/repo/user"
)

// ProviderSetRepo is data providers.
var ProviderSetRepo = wire.NewSet(
	common.NewCommonRepo,
	data.NewData,
	data.NewDB,
	data.NewCache,
	comment.NewCommentRepo,
	comment.NewCommentCommonRepo,
	captcha.NewCaptchaRepo,
	unique.NewUniqueIDRepo,
	report.NewReportRepo,
	activity_common.NewFollowRepo,
	activity_common.NewVoteRepo,
	config.NewConfigRepo,
	user.NewUserRepo,
	user.NewUserBackyardRepo,
	rank.NewUserRankRepo,
	NewQuestionRepo,
	NewAnswerRepo,
	NewActivityRepo,
	activity.NewVoteRepo,
	activity.NewFollowRepo,
	activity.NewAnswerActivityRepo,
	activity.NewQuestionActivityRepo,
	activity.NewUserActiveActivityRepo,
	tag.NewTagRepo,
	tag.NewTagListRepo,
	collection.NewCollectionRepo,
	collection.NewCollectionGroupRepo,
	auth.NewAuthRepo,
	revision.NewRevisionRepo,
	NewSearchRepo,
	meta.NewMetaRepo,
	export.NewEmailRepo,
	reason.NewReasonRepo,
	NewSiteInfo,
	notification.NewNotificationRepo,
)
