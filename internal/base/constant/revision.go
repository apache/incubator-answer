package constant

type ReviewingType string

const (
	QueuedPost        ReviewingType = "queued_post"
	QueuedUser        ReviewingType = "queued_user"
	FlaggedPost       ReviewingType = "flagged_post"
	FlaggedUser       ReviewingType = "flagged_user"
	SuggestedPostEdit ReviewingType = "suggested_post_edit"
)
