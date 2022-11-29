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
}

// ActObjectInfo act object info
type ActObjectInfo struct {
	Title       string `json:"title"`
	ObjectType  string `json:"object_type"`
	QuestionID  string `json:"question_id"`
	AnswerID    string `json:"answer_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
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
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	OriginalText string   `json:"original_text"`
}
