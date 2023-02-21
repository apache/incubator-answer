package schema

import (
	"strings"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/validator"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/segmentfault/pacman/errors"
)

// SearchTagLikeReq get tag list all request
type SearchTagLikeReq struct {
	// tag
	Tag     string `validate:"omitempty" form:"tag"`
	IsAdmin bool   `json:"-"`
}

type SearchTagsBySlugName struct {
	Tags    string   `json:"tags" form:"tags"`
	TagList []string `json:"-"`
	IsAdmin bool     `json:"-"`
}

// GetTagInfoReq get tag info request
type GetTagInfoReq struct {
	// tag id
	ID string `validate:"omitempty" form:"id"`
	// tag slug name
	Name string `validate:"omitempty,gt=0,lte=35" form:"name"`
	// user id
	UserID string `json:"-"`
	// whether user can edit it
	CanEdit bool `json:"-"`
	// whether user can delete it
	CanDelete bool `json:"-"`
}

type GetTamplateTagInfoReq struct {
	// tag id
	ID string `validate:"omitempty" form:"id"`
	// tag slug name
	Name string `validate:"omitempty" form:"name"`
	// user id
	UserID   string `json:"-"`
	Page     int    `validate:"omitempty,min=1" form:"page"`
	PageSize int    `validate:"omitempty,min=1" form:"page_size"`
}

func (r *GetTagInfoReq) Check() (errFields []*validator.FormErrorField, err error) {
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
	// description text
	Description string `json:"description"`
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
	Recommend       bool   `json:"recommend"`
	Reserved        bool   `json:"reserved"`
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
	Recommend bool  `json:"recommend"`
	Reserved  bool  `json:"reserved"`
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
	ObjectID string     `json:"object_id"` // object_id
	Tags     []*TagItem `json:"tags"`      // tags name
	// user id
	UserID string `json:"-"`
}

type TagItem struct {
	// slug_name
	SlugName string `validate:"omitempty,gt=0,lte=35" json:"slug_name"`
	// display_name
	DisplayName string `validate:"omitempty,gt=0,lte=35" json:"display_name"`
	// original text
	OriginalText string `validate:"omitempty" json:"original_text"`
	// parsed text
	ParsedText string `json:"-"`
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
	SlugName string `validate:"omitempty,gt=0,lte=35" json:"slug_name"`
	// display_name
	DisplayName string `validate:"omitempty,gt=0,lte=35" json:"display_name"`
	// original text
	OriginalText string `validate:"omitempty" json:"original_text"`
	// parsed text
	ParsedText string `json:"-"`
	// edit summary
	EditSummary string `validate:"omitempty" json:"edit_summary"`
	// user id
	UserID       string `json:"-"`
	NoNeedReview bool   `json:"-"`
}

func (r *UpdateTagReq) Check() (errFields []*validator.FormErrorField, err error) {
	if len(r.EditSummary) == 0 {
		r.EditSummary = "tag.edit.summary"
	}
	r.ParsedText = converter.Markdown2HTML(r.OriginalText)
	return nil, nil
}

// UpdateTagResp update tag response
type UpdateTagResp struct {
	WaitForReview bool `json:"wait_for_review"`
}

// GetTagWithPageReq get tag list page request
type GetTagWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// slug_name
	SlugName string `validate:"omitempty,gt=0,lte=35" form:"slug_name"`
	// display_name
	DisplayName string `validate:"omitempty,gt=0,lte=35" form:"display_name"`
	// query condition
	QueryCond string `validate:"omitempty,oneof=popular name newest" form:"query_cond"`
	// user id
	UserID string `json:"-"`
}

// GetTagSynonymsReq get tag synonyms request
type GetTagSynonymsReq struct {
	// tag_id
	TagID string `validate:"required" form:"tag_id"`
	// user id
	UserID string `json:"-"`
	// whether user can edit it
	CanEdit bool `json:"-"`
}

// GetTagSynonymsResp get tag synonyms response
type GetTagSynonymsResp struct {
	// synonyms
	Synonyms []*TagSynonym `json:"synonyms"`
	// MemberActions
	MemberActions []*PermissionMemberAction `json:"member_actions"`
}

type TagSynonym struct {
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
	Recommend       bool   `json:"recommend"`
	Reserved        bool   `json:"reserved"`
}

type SearchTagLikeResp struct {
	SlugName  string `json:"slug_name"`
	Recommend bool   `json:"recommend"`
	Reserved  bool   `json:"reserved"`
}
