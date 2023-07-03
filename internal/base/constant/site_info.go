package constant

const (
	DefaultGravatarBaseURL = "https://www.gravatar.com/avatar/"
	DefaultAvatar          = "system"
	AvatarTypeDefault      = "default"
	AvatarTypeGravatar     = "gravatar"
	AvatarTypeCustom       = "custom"
)

const (
	// PermaLinkQuestionIDAndTitle /questions/10010000000000001/post-title
	PermaLinkQuestionIDAndTitle = iota + 1
	// PermaLinkQuestionID /questions/10010000000000001
	PermaLinkQuestionID
	// PermaLinkQuestionIDAndTitleByShortID /questions/11/post-title
	PermaLinkQuestionIDAndTitleByShortID
	// PermaLinkQuestionIDByShortID /questions/11
	PermaLinkQuestionIDByShortID
)
