package schema

import (
	"github.com/answerdev/answer/internal/entity"
	"github.com/jinzhu/copier"
)

// AddCommentReq add comment request
type AddCommentReq struct {
	// object id
	ObjectID string `validate:"required" json:"object_id"`
	// reply comment id
	ReplyCommentID string `validate:"omitempty" json:"reply_comment_id"`
	// original comment content
	OriginalText string `validate:"required" json:"original_text"`
	// parsed comment content
	ParsedText string `validate:"required" json:"parsed_text"`
	// @ user id list
	MentionUsernameList []string `validate:"omitempty" json:"mention_username_list"`
	// user id
	UserID string `json:"-"`
}

// RemoveCommentReq remove comment
type RemoveCommentReq struct {
	// comment id
	CommentID string `validate:"required" json:"comment_id"`
	// user id
	UserID string `json:"-"`
}

// UpdateCommentReq update comment request
type UpdateCommentReq struct {
	// comment id
	CommentID string `validate:"required" json:"comment_id"`
	// original comment content
	OriginalText string `validate:"omitempty" json:"original_text"`
	// parsed comment content
	ParsedText string `validate:"omitempty" json:"parsed_text"`
	// user id
	UserID string `json:"-"`
}

// GetCommentListReq get comment list all request
type GetCommentListReq struct {
	// user id
	UserID int64 `validate:"omitempty" comment:"user id" form:"user_id"`
	// reply user id
	ReplyUserID int64 `validate:"omitempty" comment:"reply user id" form:"reply_user_id"`
	// reply comment id
	ReplyCommentID int64 `validate:"omitempty" comment:"reply comment id" form:"reply_comment_id"`
	// object id
	ObjectID int64 `validate:"omitempty" comment:"object id" form:"object_id"`
	// user vote amount
	VoteCount int `validate:"omitempty" comment:"user vote amount" form:"vote_count"`
	// comment status(available: 0; deleted: 10)
	Status int `validate:"omitempty" comment:"comment status(available: 0; deleted: 10)" form:"status"`
	// original comment content
	OriginalText string `validate:"omitempty" comment:"original comment content" form:"original_text"`
	// parsed comment content
	ParsedText string `validate:"omitempty" comment:"parsed comment content" form:"parsed_text"`
}

// GetCommentWithPageReq get comment list page request
type GetCommentWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// object id
	ObjectID string `validate:"required" form:"object_id"`
	// comment id
	CommentID string `validate:"omitempty" form:"comment_id"`
	// query condition
	QueryCond string `validate:"omitempty,oneof=vote" form:"query_cond"`
	// user id
	UserID string `json:"-"`
}

// GetCommentReq get comment list page request
type GetCommentReq struct {
	// object id
	ID string `validate:"required" form:"id"`
	// user id
	UserID string `json:"-"`
}

// GetCommentResp comment response
type GetCommentResp struct {
	// comment id
	CommentID string `json:"comment_id"`
	// create time
	CreatedAt int64 `json:"created_at"`

	// object id
	ObjectID string `json:"object_id"`
	// user vote amount
	VoteCount int `json:"vote_count"`
	// current user if already vote this comment
	IsVote bool `json:"is_vote"`
	// original comment content
	OriginalText string `json:"original_text"`
	// parsed comment content
	ParsedText string `json:"parsed_text"`

	// user id
	UserID string `json:"user_id"`
	// username
	Username string `json:"username"`
	// user display name
	UserDisplayName string `json:"user_display_name"`
	// user avatar
	UserAvatar string `json:"user_avatar"`
	// user status
	UserStatus string `json:"user_status"`

	// reply user id
	ReplyUserID string `json:"reply_user_id"`
	// reply user username
	ReplyUsername string `json:"reply_username"`
	// reply user display name
	ReplyUserDisplayName string `json:"reply_user_display_name"`
	// reply comment id
	ReplyCommentID string `json:"reply_comment_id"`
	// reply user status
	ReplyUserStatus string `json:"reply_user_status"`

	// MemberActions
	MemberActions []*PermissionMemberAction `json:"member_actions"`
}

func (r *GetCommentResp) SetFromComment(comment *entity.Comment) {
	_ = copier.Copy(r, comment)
	r.CommentID = comment.ID
	r.CreatedAt = comment.CreatedAt.Unix()
	r.ReplyUserID = comment.GetReplyUserID()
	r.ReplyCommentID = comment.GetReplyCommentID()
}

// GetCommentPersonalWithPageReq get comment list page request
type GetCommentPersonalWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// username
	Username string `validate:"omitempty,gt=0,lte=100" form:"username"`
	// user id
	UserID string `json:"-"`
}

// GetCommentPersonalWithPageResp comment response
type GetCommentPersonalWithPageResp struct {
	// comment id
	CommentID string `json:"comment_id"`
	// create time
	CreatedAt int64 `json:"created_at"`
	// object id
	ObjectID string `json:"object_id"`
	// question id
	QuestionID string `json:"question_id"`
	// answer id
	AnswerID string `json:"answer_id"`
	// object type
	ObjectType string `json:"object_type" enums:"question,answer,tag,comment"`
	// title
	Title string `json:"title"`
	// content
	Content string `json:"content"`
}
