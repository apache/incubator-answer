package constant

type ReviewingType string

const (
	QueuedPost        ReviewingType = "queued_post"
	QueuedUser        ReviewingType = "queued_user"
	FlaggedPost       ReviewingType = "flagged_post"
	FlaggedUser       ReviewingType = "flagged_user"
	SuggestedPostEdit ReviewingType = "suggested_post_edit"
)

const (
	ReportOperationEditPost     = "edit_post"
	ReportOperationClosePost    = "close_post"
	ReportOperationDeletePost   = "delete_post"
	ReportOperationUnlistPost   = "unlist_post"
	ReportOperationIgnoreReport = "ignore_report"
)
