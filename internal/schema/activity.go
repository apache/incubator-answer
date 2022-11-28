package schema

import "github.com/answerdev/answer/internal/base/constant"

// ActivityMsg activity message
type ActivityMsg struct {
	UserID           string                   `json:"user_id"`
	TriggerUserID    int64                    `json:"trigger_user_id"`
	ObjectID         string                   `json:"object_id"`
	OriginalObjectID string                   `json:"original_object_id"`
	ActivityTypeKey  constant.ActivityTypeKey `json:"activity_type_key"`
	RevisionID       string                   `json:"revision_id"`
}

// GetObjectTimelineReq get object timeline request
type GetObjectTimelineReq struct {
	ObjectID string `validate:"omitempty,gt=0,lte=100" form:"object_id"`
	ShowVote bool   `validate:"omitempty" form:"show_vote"`
	UserID   string `json:"-"`
	IsAdmin  bool   `json:"-"`
}

// GetObjectTimelineResp get object timeline response
type GetObjectTimelineResp struct {
	ObjectInfo *ActObjectInfo       `json:"object_info"`
	Timeline   []*ActObjectTimeline `json:"timeline"`
}

// ActObjectTimeline act object timeline
type ActObjectTimeline struct {
	ActivityID      string `json:"activity_id"`
	RevisionID      string `json:"revision_id"`
	CreatedAt       int64  `json:"created_at"`
	ActivityType    string `json:"activity_type"`
	Username        string `json:"username"`
	UserDisplayName string `json:"user_display_name"`
	Comment         string `json:"comment"`
	ObjectID        string `json:"object_id"`
	ObjectType      string `json:"object_type"`
	Cancelled       bool   `json:"cancelled"`
	CancelledAt     int64  `json:"cancelled_at"`
	UserID          string `json:"-"`
}

// ActObjectInfo act object info
type ActObjectInfo struct {
	ObjectType      string `json:"object_type"`
	Title           string `json:"title"`
	QuestionID      string `json:"question_id"`
	AnswerID        string `json:"answer_id"`
	MainTagSlugName string `json:"main_tag_slug_name"`
	Username        string `json:"username"`
}

// GetObjectTimelineDetailReq get object timeline detail request
type GetObjectTimelineDetailReq struct {
	NewRevisionID string `validate:"required,gt=0,lte=100" form:"new_revision_id"`
	OldRevisionID string `validate:"required,gt=0,lte=100" form:"old_revision_id"`
	UserID        string `json:"-"`
}

// GetObjectTimelineDetailResp get object timeline detail response
type GetObjectTimelineDetailResp struct {
	NewRevision *ObjectTimelineDetail `json:"new_revision"`
	OldRevision *ObjectTimelineDetail `json:"old_revision"`
}

// ObjectTimelineDetail object timeline detail
type ObjectTimelineDetail struct {
	Title           string               `json:"title"`
	Tags            []*ObjectTimelineTag `json:"tags"`
	OriginalText    string               `json:"original_text"`
	SlugName        string               `json:"slug_name"`
	MainTagSlugName string               `json:"main_tag_slug_name"`
}

// ObjectTimelineTag object timeline tags
type ObjectTimelineTag struct {
	SlugName        string `json:"slug_name"`
	DisplayName     string `json:"display_name"`
	MainTagSlugName string `json:"main_tag_slug_name"`
	Recommend       bool   `json:"recommend"`
	Reserved        bool   `json:"reserved"`
}
