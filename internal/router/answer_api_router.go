package router

import (
	"github.com/answerdev/answer/internal/controller"
	"github.com/answerdev/answer/internal/controller_backyard"
	"github.com/gin-gonic/gin"
)

type AnswerAPIRouter struct {
	langController           *controller.LangController
	userController           *controller.UserController
	commentController        *controller.CommentController
	reportController         *controller.ReportController
	voteController           *controller.VoteController
	tagController            *controller.TagController
	followController         *controller.FollowController
	collectionController     *controller.CollectionController
	questionController       *controller.QuestionController
	answerController         *controller.AnswerController
	searchController         *controller.SearchController
	revisionController       *controller.RevisionController
	rankController           *controller.RankController
	backyardReportController *controller_backyard.ReportController
	backyardUserController   *controller_backyard.UserBackyardController
	reasonController         *controller.ReasonController
	themeController          *controller_backyard.ThemeController
	siteInfoController       *controller_backyard.SiteInfoController
	siteinfoController       *controller.SiteinfoController
	notificationController   *controller.NotificationController
	dashboardController      *controller.DashboardController
	uploadController         *controller.UploadController
	roleController           *controller_backyard.RoleController
}

func NewAnswerAPIRouter(
	langController *controller.LangController,
	userController *controller.UserController,
	commentController *controller.CommentController,
	reportController *controller.ReportController,
	voteController *controller.VoteController,
	tagController *controller.TagController,
	followController *controller.FollowController,
	collectionController *controller.CollectionController,
	questionController *controller.QuestionController,
	answerController *controller.AnswerController,
	searchController *controller.SearchController,
	revisionController *controller.RevisionController,
	rankController *controller.RankController,
	backyardReportController *controller_backyard.ReportController,
	backyardUserController *controller_backyard.UserBackyardController,
	reasonController *controller.ReasonController,
	themeController *controller_backyard.ThemeController,
	siteInfoController *controller_backyard.SiteInfoController,
	siteinfoController *controller.SiteinfoController,
	notificationController *controller.NotificationController,
	dashboardController *controller.DashboardController,
	uploadController *controller.UploadController,
	roleController *controller_backyard.RoleController,
) *AnswerAPIRouter {
	return &AnswerAPIRouter{
		langController:           langController,
		userController:           userController,
		commentController:        commentController,
		reportController:         reportController,
		voteController:           voteController,
		tagController:            tagController,
		followController:         followController,
		collectionController:     collectionController,
		questionController:       questionController,
		answerController:         answerController,
		searchController:         searchController,
		revisionController:       revisionController,
		rankController:           rankController,
		backyardReportController: backyardReportController,
		backyardUserController:   backyardUserController,
		reasonController:         reasonController,
		themeController:          themeController,
		siteInfoController:       siteInfoController,
		notificationController:   notificationController,
		siteinfoController:       siteinfoController,
		dashboardController:      dashboardController,
		uploadController:         uploadController,
		roleController:           roleController,
	}
}

func (a *AnswerAPIRouter) RegisterUnAuthAnswerAPIRouter(r *gin.RouterGroup) {
	// i18n
	r.GET("/language/config", a.langController.GetLangMapping)
	r.GET("/language/options", a.langController.GetUserLangOptions)

	// comment
	r.GET("/comment/page", a.commentController.GetCommentWithPage)
	r.GET("/personal/comment/page", a.commentController.GetCommentPersonalWithPage)
	r.GET("/comment", a.commentController.GetComment)

	// user
	r.GET("/user/info", a.userController.GetUserInfoByUserID)
	r.GET("/user/action/record", a.userController.ActionRecord)
	r.POST("/user/login/email", a.userController.UserEmailLogin)
	r.POST("/user/register/email", a.userController.UserRegisterByEmail)
	r.POST("/user/email/verification", a.userController.UserVerifyEmail)
	r.POST("/user/password/reset", a.userController.RetrievePassWord)
	r.POST("/user/password/replacement", a.userController.UseRePassWord)
	r.GET("/personal/user/info", a.userController.GetOtherUserInfoByUsername)
	r.POST("/user/email/verification/send", a.userController.UserVerifyEmailSend)
	r.GET("/user/logout", a.userController.UserLogout)
	r.PUT("/user/email", a.userController.UserChangeEmailVerify)
	r.POST("/user/email/change/code", a.userController.UserChangeEmailSendCode)

	//answer
	r.GET("/answer/info", a.answerController.Get)
	r.GET("/answer/page", a.answerController.AnswerList)
	r.GET("/personal/answer/page", a.questionController.UserAnswerList)

	//question
	r.GET("/question/info", a.questionController.GetQuestion)
	r.POST("/question/search", a.questionController.SearchList)
	r.GET("/question/page", a.questionController.Index)
	r.GET("/question/similar/tag", a.questionController.SimilarQuestion)
	r.GET("/personal/qa/top", a.questionController.UserTop)
	r.GET("/personal/question/page", a.questionController.UserList)

	//revision
	r.GET("/revisions", a.revisionController.GetRevisionList)

	// tag
	r.GET("/tags/page", a.tagController.GetTagWithPage)
	r.GET("/tags/following", a.tagController.GetFollowingTags)
	r.GET("/tag", a.tagController.GetTagInfo)
	r.GET("/tag/synonyms", a.tagController.GetTagSynonyms)
	r.GET("/question/index", a.questionController.Index)

	//search
	r.GET("/search", a.searchController.Search)

	//rank
	r.GET("/personal/rank/page", a.rankController.GetRankPersonalWithPage)

	//siteinfo
	r.GET("/siteinfo", a.siteinfoController.GetSiteInfo)
	r.GET("/siteinfo/legal", a.siteinfoController.GetSiteLegalInfo)
}

func (a *AnswerAPIRouter) RegisterAnswerAPIRouter(r *gin.RouterGroup) {
	// comment
	r.POST("/comment", a.commentController.AddComment)
	r.DELETE("/comment", a.commentController.RemoveComment)
	r.PUT("/comment", a.commentController.UpdateComment)

	// report
	r.POST("/report", a.reportController.AddReport)

	// vote
	r.POST("/vote/up", a.voteController.VoteUp)
	r.POST("/vote/down", a.voteController.VoteDown)

	// follow
	r.POST("/follow", a.followController.Follow)
	r.PUT("/follow/tags", a.followController.UpdateFollowTags)

	// tag
	r.GET("/question/tags", a.tagController.SearchTagLike)
	r.PUT("/tag", a.tagController.UpdateTag)
	r.DELETE("/tag", a.tagController.RemoveTag)
	r.PUT("/tag/synonym", a.tagController.UpdateTagSynonym)

	// collection
	r.POST("/collection/switch", a.collectionController.CollectionSwitch)
	r.GET("/personal/collection/page", a.questionController.UserCollectionList)

	// question
	r.POST("/question", a.questionController.AddQuestion)
	r.PUT("/question", a.questionController.UpdateQuestion)
	r.DELETE("/question", a.questionController.RemoveQuestion)
	r.PUT("/question/status", a.questionController.CloseQuestion)
	r.GET("/question/similar", a.questionController.SearchByTitleLike)

	// answer
	r.POST("/answer", a.answerController.Add)
	r.PUT("/answer", a.answerController.Update)
	r.POST("/answer/acceptance", a.answerController.Adopted)
	r.DELETE("/answer", a.answerController.RemoveAnswer)

	// user
	r.PUT("/user/password", a.userController.UserModifyPassWord)
	r.PUT("/user/info", a.userController.UserUpdateInfo)
	r.PUT("/user/interface", a.userController.UserUpdateInterface)
	r.POST("/user/notice/set", a.userController.UserNoticeSet)

	// vote
	r.GET("/personal/vote/page", a.voteController.UserVotes)

	// reason
	r.GET("/reasons", a.reasonController.Reasons)

	// notification
	r.GET("/notification/status", a.notificationController.GetRedDot)
	r.PUT("/notification/status", a.notificationController.ClearRedDot)
	r.GET("/notification/page", a.notificationController.GetList)
	r.PUT("/notification/read/state/all", a.notificationController.ClearUnRead)
	r.PUT("/notification/read/state", a.notificationController.ClearIDUnRead)

	// upload file
	r.POST("/file", a.uploadController.UploadFile)
}

func (a *AnswerAPIRouter) RegisterAnswerCmsAPIRouter(r *gin.RouterGroup) {
	r.GET("/question/page", a.questionController.CmsSearchList)
	r.PUT("/question/status", a.questionController.AdminSetQuestionStatus)
	r.GET("/answer/page", a.questionController.CmsSearchAnswerList)
	r.PUT("/answer/status", a.answerController.AdminSetAnswerStatus)

	// report
	r.GET("/reports/page", a.backyardReportController.ListReportPage)
	r.PUT("/report", a.backyardReportController.Handle)

	// user
	r.GET("/users/page", a.backyardUserController.GetUserPage)
	r.PUT("/user/status", a.backyardUserController.UpdateUserStatus)
	r.PUT("/user/role", a.backyardUserController.UpdateUserRole)

	// reason
	r.GET("/reasons", a.reasonController.Reasons)

	// language
	r.GET("/language/options", a.langController.GetAdminLangOptions)

	// theme
	r.GET("/theme/options", a.themeController.GetThemeOptions)

	// siteinfo
	r.GET("/siteinfo/general", a.siteInfoController.GetGeneral)
	r.GET("/siteinfo/interface", a.siteInfoController.GetInterface)
	r.GET("/siteinfo/branding", a.siteInfoController.GetSiteBranding)
	r.GET("/siteinfo/write", a.siteInfoController.GetSiteWrite)
	r.GET("/siteinfo/legal", a.siteInfoController.GetSiteLegal)
	r.PUT("/siteinfo/general", a.siteInfoController.UpdateGeneral)
	r.PUT("/siteinfo/interface", a.siteInfoController.UpdateInterface)
	r.PUT("/siteinfo/branding", a.siteInfoController.UpdateBranding)
	r.PUT("/siteinfo/write", a.siteInfoController.UpdateSiteWrite)
	r.PUT("/siteinfo/legal", a.siteInfoController.UpdateSiteLegal)
	r.GET("/setting/smtp", a.siteInfoController.GetSMTPConfig)
	r.PUT("/setting/smtp", a.siteInfoController.UpdateSMTPConfig)

	// dashboard
	r.GET("/dashboard", a.dashboardController.DashboardInfo)

	// roles
	r.GET("/roles", a.roleController.GetRoleList)
}
