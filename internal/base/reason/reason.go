package reason

const (
	// Success .
	Success = "base.success"
	// UnknownError unknown error
	UnknownError = "base.unknown"
	// RequestFormatError request format error
	RequestFormatError = "base.request_format_error"
	// UnauthorizedError unauthorized error
	UnauthorizedError = "base.unauthorized_error"
	// DatabaseError database error
	DatabaseError = "base.database_error"
)

const (
	EmailOrPasswordWrong         = "error.user.email_or_password_wrong"
	CommentNotFound              = "error.comment.not_found"
	QuestionNotFound             = "error.question.not_found"
	AnswerNotFound               = "error.answer.not_found"
	CommentEditWithoutPermission = "error.comment.edit_without_permission"
	DisallowVote                 = "error.object.disallow_vote"
	DisallowFollow               = "error.object.disallow_follow"
	DisallowVoteYourSelf         = "error.object.disallow_vote_your_self"
	CaptchaVerificationFailed    = "error.object.captcha_verification_failed"
	UserNotFound                 = "error.user.not_found"
	UsernameInvalid              = "error.user.username_invalid"
	UsernameDuplicate            = "error.user.username_duplicate"
	EmailDuplicate               = "error.email.duplicate"
	EmailVerifyUrlExpired        = "error.email.verify_url_expired"
	EmailNeedToBeVerified        = "error.email.need_to_be_verified"
	UserSuspended                = "error.user.suspended"
	ObjectNotFound               = "error.object.not_found"
	TagNotFound                  = "error.tag.not_found"
	RankFailToMeetTheCondition   = "error.rank.fail_to_meet_the_condition"
	ThemeNotFound                = "error.theme.not_found"
	LangNotFound                 = "error.lang.not_found"
	ReportHandleFailed           = "error.report.handle_failed"
	ReportNotFound               = "error.report.not_found"
)
