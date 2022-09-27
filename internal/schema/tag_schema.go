package schema

import (
	"strings"

	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/base/validator"
	"github.com/segmentfault/pacman/errors"
)

// SearchTagLikeReq get tag list all request
type SearchTagLikeReq struct {
	// tag
	Tag string `validate:"required,gt=0,lte=50" form:"tag"`
}

// GetTagInfoReq get tag info request
type GetTagInfoReq struct {
	// tag id
	ID string `validate:"omitempty" form:"id"`
	// tag slug name
	Name string `validate:"omitempty,gt=0,lte=50" form:"name"`
	// user id
	UserID string `json:"-"`
}

func (r *GetTagInfoReq) Check() (errField *validator.ErrorField, err error) {
	if len(r.ID) == 0 && len(r.Name) == 0 {
		return nil, errors.BadRequest(reason.RequestFormatError)
	}
	r.Name = strings.ToLower(r.Name)
	return nil, nil
}

// GetTagResp get tag response
type GetTagResp struct {
	// tag id
	TagID string `json:"tag_id"`
	// created time
	CreatedAt int64 `json:"created_at"`
	// updated time
	UpdatedAt int64 `json:"updated_at"`
	// slug name
	SlugName string `json:"slug_name"`
	// display name
	DisplayName string `json:"display_name"`
	// excerpt
	Excerpt string `json:"excerpt"`
	// original text
	OriginalText string `json:"original_text"`
	// parsed text
	ParsedText string `json:"parsed_text"`
	// follower amount
	FollowCount int `json:"follow_count"`
	// question amount
	QuestionCount int `json:"question_count"`
	// is follower
	IsFollower bool `json:"is_follower"`
	// MemberActions
	MemberActions []*PermissionMemberAction `json:"member_actions"`
	// if main tag slug name is not empty, this tag is synonymous with the main tag
	MainTagSlugName string `json:"main_tag_slug_name"`
}

func (tr *GetTagResp) GetExcerpt() {
	excerpt := strings.TrimSpace(tr.OriginalText)
	idx := strings.Index(excerpt, "\n")
	if idx >= 0 {
		excerpt = excerpt[0:idx]
	}
	tr.Excerpt = excerpt
}

// GetTagPageResp get tag response
type GetTagPageResp struct {
	// tag_id
	TagID string `json:"tag_id"`
	// slug_name
	SlugName string `json:"slug_name"`
	// display_name
	DisplayName string `json:"display_name"`
	// excerpt
	Excerpt string `json:"excerpt"`
	// original text
	OriginalText string `json:"original_text"`
	// parsed_text
	ParsedText string `json:"parsed_text"`
	// follower amount
	FollowCount int `json:"follow_count"`
	// question amount
	QuestionCount int `json:"question_count"`
	// is follower
	IsFollower bool `json:"is_follower"`
	// created time
	CreatedAt int64 `json:"created_at"`
	// updated time
	UpdatedAt int64 `json:"updated_at"`
}

func (tr *GetTagPageResp) GetExcerpt() {
	excerpt := strings.TrimSpace(tr.OriginalText)
	idx := strings.Index(excerpt, "\n")
	if idx >= 0 {
		excerpt = excerpt[0:idx]
	}
	tr.Excerpt = excerpt
}

type TagChange struct {
	ObjectId string     `json:"object_id"` // object_id
	Tags     []*TagItem `json:"tags"`      // tags name
	// user id
	UserID string `json:"-"`
}

type TagItem struct {
	// slug_name
	SlugName string `validate:"omitempty,gt=0,lte=50" json:"slug_name"`
	// display_name
	DisplayName string `validate:"omitempty,gt=0,lte=50" json:"display_name"`
	// original text
	OriginalText string `validate:"omitempty" json:"original_text"`
	// parsed text
	ParsedText string `validate:"omitempty" json:"parsed_text"`
}

// RemoveTagReq delete tag request
type RemoveTagReq struct {
	// tag_id
	TagID string `validate:"required" json:"tag_id"`
	// user id
	UserID string `json:"-"`
}

// UpdateTagReq update tag request
type UpdateTagReq struct {
	// tag_id
	TagID string `validate:"required" json:"tag_id"`
	// slug_name
	SlugName string `validate:"omitempty,gt=0,lte=50" json:"slug_name"`
	// display_name
	DisplayName string `validate:"omitempty,gt=0,lte=50" json:"display_name"`
	// original text
	OriginalText string `validate:"omitempty" json:"original_text"`
	// parsed text
	ParsedText string `validate:"omitempty" json:"parsed_text"`
	// edit summary
	EditSummary string `validate:"omitempty" json:"edit_summary"`
	// user id
	UserID string `json:"-"`
}

func (r *UpdateTagReq) Check() (errField *validator.ErrorField, err error) {
	if len(r.EditSummary) == 0 {
		r.EditSummary = "tag.edit.summary" // to i18n
	}
	return nil, nil
}

// GetTagWithPageReq get tag list page request
type GetTagWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// slug_name
	SlugName string `validate:"omitempty,gt=0,lte=50" form:"slug_name"`
	// display_name
	DisplayName string `validate:"omitempty,gt=0,lte=50" form:"display_name"`
	// query condition
	QueryCond string `validate:"omitempty,oneof=popular name newest" form:"query_cond"`
	// user id
	UserID string `json:"-"`
}

// GetTagSynonymsReq get tag synonyms request
type GetTagSynonymsReq struct {
	// tag_id
	TagID string `validate:"required" form:"tag_id"`
}

// GetTagSynonymsResp get tag synonyms response
type GetTagSynonymsResp struct {
	// tag id
	TagID string `json:"tag_id"`
	// slug name
	SlugName string `json:"slug_name"`
	// display name
	DisplayName string `json:"display_name"`
	// if main tag slug name is not empty, this tag is synonymous with the main tag
	MainTagSlugName string `json:"main_tag_slug_name"`
}

// UpdateTagSynonymReq update tag request
type UpdateTagSynonymReq struct {
	// tag_id
	TagID string `validate:"required" json:"tag_id"`
	// synonym tag list
	SynonymTagList []*TagItem `validate:"required,dive" json:"synonym_tag_list"`
	// user id
	UserID string `json:"-"`
}

func (req *UpdateTagSynonymReq) Format() {
	for _, item := range req.SynonymTagList {
		item.SlugName = strings.ToLower(item.SlugName)
	}
}

// GetFollowingTagsResp get following tags response
type GetFollowingTagsResp struct {
	// tag id
	TagID string `json:"tag_id"`
	// slug name
	SlugName string `json:"slug_name"`
	// display name
	DisplayName string `json:"display_name"`
	// if main tag slug name is not empty, this tag is synonymous with the main tag
	MainTagSlugName string `json:"main_tag_slug_name"`
}
