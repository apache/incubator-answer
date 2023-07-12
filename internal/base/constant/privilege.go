package constant

import "github.com/answerdev/answer/internal/base/reason"

type Privilege struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Value int    `json:"value"`
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
	RankInviteSomeoneToAnswerKey     = "rank.answer.invite_someone_to_answer"
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

var (
	RankAllPrivileges = []*Privilege{
		{Label: reason.RankQuestionAddLabel, Key: RankQuestionAddKey},
		{Label: reason.RankAnswerAddLabel, Key: RankAnswerAddKey},
		{Label: reason.RankCommentAddLabel, Key: RankCommentAddKey},
		{Label: reason.RankReportAddLabel, Key: RankReportAddKey},
		{Label: reason.RankCommentVoteUpLabel, Key: RankCommentVoteUpKey},
		{Label: reason.RankLinkUrlLimitLabel, Key: RankLinkUrlLimitKey},
		{Label: reason.RankQuestionVoteUpLabel, Key: RankQuestionVoteUpKey},
		{Label: reason.RankAnswerVoteUpLabel, Key: RankAnswerVoteUpKey},
		{Label: reason.RankQuestionVoteDownLabel, Key: RankQuestionVoteDownKey},
		{Label: reason.RankAnswerVoteDownLabel, Key: RankAnswerVoteDownKey},
		{Label: reason.RankInviteSomeoneToAnswerLabel, Key: RankInviteSomeoneToAnswerKey},
		{Label: reason.RankTagAddLabel, Key: RankTagAddKey},
		{Label: reason.RankTagEditLabel, Key: RankTagEditKey},
		{Label: reason.RankQuestionEditLabel, Key: RankQuestionEditKey},
		{Label: reason.RankAnswerEditLabel, Key: RankAnswerEditKey},
		{Label: reason.RankQuestionEditWithoutReviewLabel, Key: RankQuestionEditWithoutReviewKey},
		{Label: reason.RankAnswerEditWithoutReviewLabel, Key: RankAnswerEditWithoutReviewKey},
		{Label: reason.RankQuestionAuditLabel, Key: RankQuestionAuditKey},
		{Label: reason.RankAnswerAuditLabel, Key: RankAnswerAuditKey},
		{Label: reason.RankTagAuditLabel, Key: RankTagAuditKey},
		{Label: reason.RankTagEditWithoutReviewLabel, Key: RankTagEditWithoutReviewKey},
		{Label: reason.RankTagSynonymLabel, Key: RankTagSynonymKey},
	}
)
