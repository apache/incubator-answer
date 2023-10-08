package constant

const (
	DefaultGravatarBaseURL = "https://www.gravatar.com/avatar/"
	DefaultAvatar          = "system"
	AvatarTypeDefault      = "default"
	AvatarTypeGravatar     = "gravatar"
	AvatarTypeCustom       = "custom"
)

const (
	// PermalinkQuestionIDAndTitle /questions/10010000000000001/post-title
	PermalinkQuestionIDAndTitle = iota + 1
	// PermalinkQuestionID /questions/10010000000000001
	PermalinkQuestionID
	// PermalinkQuestionIDAndTitleByShortID /questions/11/post-title
	PermalinkQuestionIDAndTitleByShortID
	// PermalinkQuestionIDByShortID /questions/11
	PermalinkQuestionIDByShortID
)
