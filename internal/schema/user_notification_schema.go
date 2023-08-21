package schema

import (
	"encoding/json"
	"github.com/answerdev/answer/internal/base/constant"
)

type NotificationChannelConfig struct {
	Key    constant.NotificationChannel `json:"key"`
	Enable bool                         `json:"enable"`
}

type NotificationChannelConfigList []*NotificationChannelConfig

func (n *NotificationChannelConfigList) Format(sequences []constant.NotificationChannel) {
	if n == nil {
		*n = make([]*NotificationChannelConfig, 0)
		return
	}
	newList := make([]*NotificationChannelConfig, 0)
	mapping := make(map[constant.NotificationChannel]*NotificationChannelConfig)
	for _, item := range *n {
		mapping[item.Key] = &NotificationChannelConfig{
			Key:    item.Key,
			Enable: item.Enable,
		}
	}
	for _, ch := range sequences {
		if c, ok := mapping[ch]; ok {
			newList = append(newList, c)
		} else {
			newList = append(newList, &NotificationChannelConfig{
				Key: ch,
			})
		}
	}
	*n = newList
}

func (n *NotificationChannelConfigList) CheckEnable(ch constant.NotificationChannel) bool {
	if n == nil {
		return false
	}
	for _, item := range *n {
		if item.Key == ch {
			return item.Enable
		}
	}
	return false
}

type NotificationConfig struct {
	Inbox                          NotificationChannelConfigList `json:"inbox"`
	AllNewQuestion                 NotificationChannelConfigList `json:"all_new_question"`
	AllNewQuestionForFollowingTags NotificationChannelConfigList `json:"all_new_question_for_following_tags"`
}

func (n *NotificationConfig) ToJsonString() string {
	data, _ := json.Marshal(n)
	return string(data)
}

func NewNotificationConfig(data string) *NotificationConfig {
	nc := &NotificationConfig{}
	nc.FromJsonString(data)
	return nc
}

func (n *NotificationConfig) FromJsonString(data string) {
	if len(data) > 0 {
		_ = json.Unmarshal([]byte(data), n)
		return
	}
	n.Inbox = make([]*NotificationChannelConfig, 0)
	n.AllNewQuestion = make([]*NotificationChannelConfig, 0)
	n.AllNewQuestionForFollowingTags = make([]*NotificationChannelConfig, 0)
	return
}

func (n *NotificationConfig) Format() {
	n.Inbox.Format([]constant.NotificationChannel{constant.EmailChannel})
	n.AllNewQuestion.Format([]constant.NotificationChannel{constant.EmailChannel})
	n.AllNewQuestionForFollowingTags.Format([]constant.NotificationChannel{constant.EmailChannel})
}

func (n *NotificationConfig) CheckEnable(
	source constant.NotificationSource, channel constant.NotificationChannel) bool {
	switch source {
	case constant.InboxChannel:
		return n.Inbox.CheckEnable(channel)
	case constant.AllNewQuestionChannel:
		return n.AllNewQuestion.CheckEnable(channel)
	case constant.AllNewQuestionForFollowingTagsChannel:
		return n.AllNewQuestionForFollowingTags.CheckEnable(channel)
	}
	return false
}

// UpdateUserNotificationConfigReq update user notification config request
type UpdateUserNotificationConfigReq struct {
	NotificationConfig
	UserID string `json:"-"`
}

// GetUserNotificationConfigResp get user notification config response
type GetUserNotificationConfigResp struct {
	NotificationConfig
}
