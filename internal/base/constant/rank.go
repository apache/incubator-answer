package constant

type Privilege struct {
	Label string `json:"label"`
	Value int    `json:"value"`
	Key   string `json:"-"`
}

const (
	RankQuestionAddKey               = "rank.question.add"
	RankQuestionEditKey              = "rank.question.edit"
	RankQuestionDeleteKey            = "rank.question.delete"
	RankQuestionVoteUpKey            = "rank.question.vote_up"
	RankQuestionVoteDownKey          = "rank.question.vote_down"
	RankAnswerAddKey                 = "rank.answer.add"
	RankAnswerEditKey                = "rank.answer.edit"
	RankAnswerDeleteKey              = "rank.answer.delete"
	RankAnswerAcceptKey              = "rank.answer.accept"
	RankAnswerVoteUpKey              = "rank.answer.vote_up"
	RankAnswerVoteDownKey            = "rank.answer.vote_down"
	RankCommentAddKey                = "rank.comment.add"
	RankCommentEditKey               = "rank.comment.edit"
	RankCommentDeleteKey             = "rank.comment.delete"
	RankReportAddKey                 = "rank.report.add"
	RankTagAddKey                    = "rank.tag.add"
	RankTagEditKey                   = "rank.tag.edit"
	RankTagDeleteKey                 = "rank.tag.delete"
	RankTagSynonymKey                = "rank.tag.synonym"
	RankLinkUrlLimitKey              = "rank.link.url_limit"
	RankVoteDetailKey                = "rank.vote.detail"
	RankCommentVoteUpKey             = "rank.comment.vote_up"
	RankCommentVoteDownKey           = "rank.comment.vote_down"
	RankQuestionEditWithoutReviewKey = "rank.question.edit_without_review"
	RankAnswerEditWithoutReviewKey   = "rank.answer.edit_without_review"
	RankTagEditWithoutReviewKey      = "rank.tag.edit_without_review"
	RankAnswerAuditKey               = "rank.answer.audit"
	RankQuestionAuditKey             = "rank.question.audit"
	RankTagAuditKey                  = "rank.tag.audit"
	RankQuestionCloseKey             = "rank.question.close"
	RankQuestionReopenKey            = "rank.question.reopen"
	RankTagUseReservedTagKey         = "rank.tag.use_reserved_tag"
)

//| Permission                             | Level 1                                          | Level 2                                       | Level 3                                       | Custom Level |
//| -------------------------------------- | ------------------------------------------------ | --------------------------------------------- | --------------------------------------------- | ------------ |
//| Description                            | less reputation required for private team, group | low reputation required for startup community | high reputation required for mature community |              |
//| Ask question                           | 1                                                | 1                                             | 1                                             |              |
//| Write answer                           | 1                                                | 1                                             | 1                                             |              |
//| Write comment                          | 1                                                | 1                                             | 1                                             |              |
//| Accept answer                          | 1                                                | 1                                             | 1                                             |              |
//| Flag                                   | 1                                                | 1                                             | 1                                             |              |
//| Upvote comment                         | 1                                                | 1                                             | 1                                             |              |
//| Post more than 2 links at a time       | 1                                                | 10                                            | 10                                            |              |
//| Upvote question                        | 1                                                | 1                                             | 15                                            |              |
//| Upvote answer                          | 1                                                | 1                                             | 15                                            |              |
//| Downvote question                      | 125                                              | 125                                           | 125                                           |              |
//| Downvote answer                        | 125                                              | 125                                           | 125                                           |              |
//| Create new tag                         | 1                                                | 750                                           | 1500                                          |              |
//| Edit tag description (need to review)  | 1                                                | 50                                            | 100                                           |              |
//| Edit other's question (need to review) | 1                                                | 100                                           | 200                                           |              |
//| Edit other's answer (need to review)   | 1                                                | 100                                           | 200                                           |              |
//| Edit other's question without review   | 1                                                | 1000                                          | 2000                                          |              |
//| Edit other's answer without review     | 1                                                | 1000                                          | 2000                                          |              |
//| Revew question edits                   | 1                                                | 1000                                          | 2000                                          |              |
//| Review answer edits                    | 1                                                | 1000                                          | 2000                                          |              |
//| Review tag edits                       | 1                                                | 2500                                          | 5000                                          |              |
//| Edit tag description without review    | 1                                                | 10000                                         | 20000                                         |              |
//| Manage tag synonyms                    | 1                                                | 10000                                         | 20000                                         |              |

const (
	RankQuestionAddLabel               = "Ask question"
	RankAnswerAddLabel                 = "Write answer"
	RankCommentAddLabel                = "Write comment"
	RankAnswerAcceptLabel              = "Accept answer"
	RankReportAddLabel                 = "Flag"
	RankCommentVoteUpLabel             = "Upvote comment"
	RankLinkUrlLimitLabel              = "Post more than 2 links at a time"
	RankQuestionVoteUpLabel            = "Upvote question"
	RankAnswerVoteUpLabel              = "Upvote answer"
	RankQuestionVoteDownLabel          = "Downvote question"
	RankAnswerVoteDownLabel            = "Downvote answer"
	RankTagAddLabel                    = "Create new tag"
	RankTagEditLabel                   = "Edit tag description (need to review)"
	RankQuestionEditLabel              = "Edit other's question (need to review)"
	RankAnswerEditLabel                = "Edit other's answer (need to review)"
	RankQuestionEditWithoutReviewLabel = "Edit other's question without review"
	RankAnswerEditWithoutReviewLabel   = "Edit other's answer without review"
	RankQuestionAuditLabel             = "Review question edits"
	RankAnswerAuditLabel               = "Review answer edits"
	RankTagAuditLabel                  = "Review tag edits"
	RankTagEditWithoutReviewLabel      = "Edit tag description without review"
	RankTagSynonymLabel                = "Manage tag synonyms"
)

var (
	RankAllPrivileges = []*Privilege{
		{Label: RankQuestionAddLabel, Key: RankQuestionAddKey},
		{Label: RankAnswerAddLabel, Key: RankAnswerAddKey},
		{Label: RankCommentAddLabel, Key: RankCommentAddKey},
		{Label: RankAnswerAcceptLabel, Key: RankAnswerAcceptKey},
		{Label: RankReportAddLabel, Key: RankReportAddKey},
		{Label: RankCommentVoteUpLabel, Key: RankCommentVoteUpKey},
		{Label: RankLinkUrlLimitLabel, Key: RankLinkUrlLimitKey},
		{Label: RankQuestionVoteUpLabel, Key: RankQuestionVoteUpKey},
		{Label: RankAnswerVoteUpLabel, Key: RankAnswerVoteUpKey},
		{Label: RankQuestionVoteDownLabel, Key: RankQuestionVoteDownKey},
		{Label: RankAnswerVoteDownLabel, Key: RankAnswerVoteDownKey},
		{Label: RankTagAddLabel, Key: RankTagAddKey},
		{Label: RankTagEditLabel, Key: RankTagEditKey},
		{Label: RankQuestionEditLabel, Key: RankQuestionEditKey},
		{Label: RankAnswerEditLabel, Key: RankAnswerEditKey},
		{Label: RankQuestionEditWithoutReviewLabel, Key: RankQuestionEditWithoutReviewKey},
		{Label: RankAnswerEditWithoutReviewLabel, Key: RankAnswerEditWithoutReviewKey},
		{Label: RankQuestionAuditLabel, Key: RankQuestionAuditKey},
		{Label: RankAnswerAuditLabel, Key: RankAnswerAuditKey},
		{Label: RankTagAuditLabel, Key: RankTagAuditKey},
		{Label: RankTagEditWithoutReviewLabel, Key: RankTagEditWithoutReviewKey},
		{Label: RankTagSynonymLabel, Key: RankTagSynonymKey},
	}
)
