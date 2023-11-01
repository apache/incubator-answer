/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package schema

import (
	"encoding/json"
	"github.com/apache/incubator-answer/internal/base/constant"
	"github.com/apache/incubator-answer/internal/entity"
)

type NotificationChannelConfig struct {
	Key    constant.NotificationChannelKey `json:"key"`
	Enable bool                            `json:"enable"`
}

type NotificationChannels []*NotificationChannelConfig

func NewNotificationChannelsFormJson(jsonStr string) NotificationChannels {
	var list NotificationChannels
	_ = json.Unmarshal([]byte(jsonStr), &list)
	return list
}

func (n *NotificationChannels) Format(sequences []constant.NotificationChannelKey) {
	if n == nil {
		*n = make([]*NotificationChannelConfig, 0)
		return
	}
	mapping := make(map[constant.NotificationChannelKey]*NotificationChannelConfig)
	for _, item := range *n {
		mapping[item.Key] = &NotificationChannelConfig{
			Key:    item.Key,
			Enable: item.Enable,
		}
	}
	newList := make([]*NotificationChannelConfig, 0)
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

func (n *NotificationChannels) CheckEnable(ch constant.NotificationChannelKey) bool {
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

func (n *NotificationChannels) ToJsonString() string {
	data, _ := json.Marshal(n)
	return string(data)
}

type NotificationConfig struct {
	Inbox                          NotificationChannels `json:"inbox"`
	AllNewQuestion                 NotificationChannels `json:"all_new_question"`
	AllNewQuestionForFollowingTags NotificationChannels `json:"all_new_question_for_following_tags"`
}

func (n *NotificationConfig) ToJsonString() string {
	data, _ := json.Marshal(n)
	return string(data)
}

func NewNotificationConfig(configs []*entity.UserNotificationConfig) NotificationConfig {
	nc := NotificationConfig{}
	nc.Inbox = make([]*NotificationChannelConfig, 0)
	nc.AllNewQuestion = make([]*NotificationChannelConfig, 0)
	nc.AllNewQuestionForFollowingTags = make([]*NotificationChannelConfig, 0)
	for _, item := range configs {
		switch item.Source {
		case string(constant.InboxSource):
			nc.Inbox = NewNotificationChannelsFormJson(item.Channels)
		case string(constant.AllNewQuestionSource):
			nc.AllNewQuestion = NewNotificationChannelsFormJson(item.Channels)
		case string(constant.AllNewQuestionForFollowingTagsSource):
			nc.AllNewQuestionForFollowingTags = NewNotificationChannelsFormJson(item.Channels)
		}
	}
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
	n.Inbox.Format([]constant.NotificationChannelKey{constant.EmailChannel})
	n.AllNewQuestion.Format([]constant.NotificationChannelKey{constant.EmailChannel})
	n.AllNewQuestionForFollowingTags.Format([]constant.NotificationChannelKey{constant.EmailChannel})
}

func (n *NotificationConfig) CheckEnable(
	source constant.NotificationSource, channel constant.NotificationChannelKey) bool {
	switch source {
	case constant.InboxSource:
		return n.Inbox.CheckEnable(channel)
	case constant.AllNewQuestionSource:
		return n.AllNewQuestion.CheckEnable(channel)
	case constant.AllNewQuestionForFollowingTagsSource:
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
