package constant

const (
	ReportSpamName             = "report.spam.name"
	ReportSpamDescription      = "report.spam.description"
	ReportRudeName             = "report.rude.name"
	ReportRudeDescription      = "report.rude.description"
	ReportDuplicateName        = "report.duplicate.name"
	ReportDuplicateDescription = "report.duplicate.description"
	ReportOtherName            = "report.other.name"
	ReportOtherDescription     = "report.other.description"
	ReportNotAnswerName        = "report.not_answer.name"
	ReportNotAnswerDescription = "report.not_answer.description"
	ReportNotNeedName          = "report.not_need.name"
	ReportNotNeedDescription   = "report.not_need.description"
	// question close
	QuestionCloseDuplicateName        = "question.close.duplicate.name"
	QuestionCloseDuplicateDescription = "question.close.duplicate.description"
	QuestionCloseGuidelineName        = "question.close.guideline.name"
	QuestionCloseGuidelineDescription = "question.close.guideline.description"
	QuestionCloseMultipleName         = "question.close.multiple.name"
	QuestionCloseMultipleDescription  = "question.close.multiple.description"
	QuestionCloseOtherName            = "question.close.other.name"
	QuestionCloseOtherDescription     = "question.close.other.description"
)

const (
	// TODO put this in database
	// TODO need reason controller to resolve
	QuestionCloseJSON  = `[{"name":"question.close.duplicate.name","description":"question.close.duplicate.description","source":"question","type":1,"have_content":false,"content_type":""},{"name":"question.close.guideline.name","description":"question.close.guideline.description","source":"question","type":2,"have_content":false,"content_type":""},{"name":"question.close.multiple.name","description":"question.close.multiple.description","source":"question","type":3,"have_content":true,"content_type":"text"},{"name":"question.close.other.name","description":"question.close.other.description","source":"question","type":4,"have_content":true,"content_type":"textarea"}]`
	QuestionReportJSON = `[{"name":"report.spam.name","description":"report.spam.description","source":"question","type":1,"have_content":false,"content_type":""},{"name":"report.rude.name","description":"report.rude.description","source":"question","type":2,"have_content":false,"content_type":""},{"name":"report.duplicate.name","description":"report.duplicate.description","source":"question","type":3,"have_content":true,"content_type":"text"},{"name":"report.other.name","description":"report.other.description","source":"question","type":4,"have_content":true,"content_type":"textarea"}]`
	AnswerReportJSON   = `[{"name":"report.spam.name","description":"report.spam.description","source":"answer","type":1,"have_content":false,"content_type":""},{"name":"report.rude.name","description":"report.rude.description","source":"answer","type":2,"have_content":false,"content_type":""},{"name":"report.not_answer.name","description":"report.not_answer.description","source":"answer","type":3,"have_content":false,"content_type":""},{"name":"report.other.name","description":"report.other.description","source":"answer","type":4,"have_content":true,"content_type":"textarea"}]`
	CommentReportJSON  = `[{"name":"report.spam.name","description":"report.spam.description","source":"comment","type":1,"have_content":false,"content_type":""},{"name":"report.rude.name","description":"report.rude.description","source":"comment","type":2,"have_content":false,"content_type":""},{"name":"report.not_need.name","description":"report.not_need.description","source":"comment","type":3,"have_content":true,"content_type":"text"},{"name":"report.other.name","description":"report.other.description","source":"comment","type":4,"have_content":true,"content_type":"textarea"}]`
)
