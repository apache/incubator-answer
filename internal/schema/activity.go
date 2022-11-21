package schema

import "github.com/answerdev/answer/internal/base/constant"

// ActivityMsg activity message
type ActivityMsg struct {
	UserID          string                   `json:"user_id"`
	TriggerUserID   int64                    `json:"trigger_user_id"`
	ObjectID        string                   `json:"object_id"`
	ActivityTypeKey constant.ActivityTypeKey `json:"activity_type_key"`
}
