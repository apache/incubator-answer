package repo

import (
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/repo/activity"
	"github.com/answerdev/answer/internal/repo/activity_common"
	"github.com/answerdev/answer/internal/repo/answer"
	"github.com/answerdev/answer/internal/repo/auth"
	"github.com/answerdev/answer/internal/repo/captcha"
	"github.com/answerdev/answer/internal/repo/collection"
	"github.com/answerdev/answer/internal/repo/comment"
	"github.com/answerdev/answer/internal/repo/common"
	"github.com/answerdev/answer/internal/repo/config"
	"github.com/answerdev/answer/internal/repo/export"
	"github.com/answerdev/answer/internal/repo/meta"
	"github.com/answerdev/answer/internal/repo/notification"
	"github.com/answerdev/answer/internal/repo/question"
	"github.com/answerdev/answer/internal/repo/rank"
	"github.com/answerdev/answer/internal/repo/reason"
	"github.com/answerdev/answer/internal/repo/report"
	"github.com/answerdev/answer/internal/repo/revision"
	"github.com/answerdev/answer/internal/repo/search_common"
	"github.com/answerdev/answer/internal/repo/site_info"
	"github.com/answerdev/answer/internal/repo/tag"
	"github.com/answerdev/answer/internal/repo/unique"
	"github.com/answerdev/answer/internal/repo/user"
	"github.com/google/wire"
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
	question.NewQuestionRepo,
	answer.NewAnswerRepo,
	activity_common.NewActivityRepo,
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
	search_common.NewSearchRepo,
	meta.NewMetaRepo,
	export.NewEmailRepo,
	reason.NewReasonRepo,
	site_info.NewSiteInfo,
	notification.NewNotificationRepo,
)
