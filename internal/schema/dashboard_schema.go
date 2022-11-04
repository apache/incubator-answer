package schema

type DashboardInfo struct {
	QuestionCount         int64  `json:"question_count"`
	AnswerCount           int64  `json:"answer_count"`
	CommentCount          int64  `json:"comment_count"`
	VoteCount             int64  `json:"vote_count"`
	UserCount             int64  `json:"user_count"`
	ReportCount           int64  `json:"report_count"`
	UploadingFiles        bool   `json:"uploading_files"`
	SMTP                  bool   `json:"smtp"`
	HTTPS                 bool   `json:"https"`
	TimeZone              string `json:"time_zone"`
	OccupyingStorageSpace string `json:"occupying_storage_space"`
	AppStartTime          string `json:"app_start_time"`
}
