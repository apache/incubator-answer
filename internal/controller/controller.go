package controller

import "github.com/google/wire"

// ProviderSetController is controller providers.
var ProviderSetController = wire.NewSet(
	NewLangController,
	NewCommentController,
	NewReportController,
	NewVoteController,
	NewTagController,
	NewFollowController,
	NewCollectionController,
	NewUserController,
	NewQuestionController,
	NewAnswerController,
	NewSearchController,
	NewRevisionController,
	NewRankController,
	NewReasonController,
	NewNotificationController,
	NewSiteInfoController,
	NewDashboardController,
	NewUploadController,
	NewActivityController,
	NewTemplateController,
	NewConnectorController,
	NewUserCenterController,
	NewPermissionController,
)
