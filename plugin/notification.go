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

package plugin

// NotificationType is the type of the notification
type NotificationType string

const (
	NotificationUpdateQuestion         NotificationType = "notification.action.update_question"
	NotificationAnswerTheQuestion      NotificationType = "notification.action.answer_the_question"
	NotificationUpVotedTheQuestion     NotificationType = "notification.action.up_voted_question"
	NotificationDownVotedTheQuestion   NotificationType = "notification.action.down_voted_question"
	NotificationUpdateAnswer           NotificationType = "notification.action.update_answer"
	NotificationAcceptAnswer           NotificationType = "notification.action.accept_answer"
	NotificationUpVotedTheAnswer       NotificationType = "notification.action.up_voted_answer"
	NotificationDownVotedTheAnswer     NotificationType = "notification.action.down_voted_answer"
	NotificationCommentQuestion        NotificationType = "notification.action.comment_question"
	NotificationCommentAnswer          NotificationType = "notification.action.comment_answer"
	NotificationUpVotedTheComment      NotificationType = "notification.action.up_voted_comment"
	NotificationReplyToYou             NotificationType = "notification.action.reply_to_you"
	NotificationMentionYou             NotificationType = "notification.action.mention_you"
	NotificationYourQuestionIsClosed   NotificationType = "notification.action.your_question_is_closed"
	NotificationYourQuestionWasDeleted NotificationType = "notification.action.your_question_was_deleted"
	NotificationYourAnswerWasDeleted   NotificationType = "notification.action.your_answer_was_deleted"
	NotificationYourCommentWasDeleted  NotificationType = "notification.action.your_comment_was_deleted"
	NotificationInvitedYouToAnswer     NotificationType = "notification.action.invited_you_to_answer"
	NotificationNewQuestion            NotificationType = "notification.action.new_question"
	NotificationNewQuestionFollowedTag NotificationType = "notification.action.new_question_followed_tag"
)

type Notification interface {
	Base

	// GetNewQuestionSubscribers returns the subscribers of the new question notification
	GetNewQuestionSubscribers() (userIDs []string)

	// Notify sends a notification to the user
	Notify(msg NotificationMessage)
}

type NotificationMessage struct {
	//  the type of the notification
	Type NotificationType `json:"notification_type"`
	// the receiver user id
	ReceiverUserID string `json:"receiver_user_id"`
	// the receiver user using language
	ReceiverLang string `json:"receiver_lang"`
	// the receiver user external id (optional)
	ReceiverExternalID string `json:"receiver_external_id"`

	// Who triggered the notification (optional, admin or system operation will not have this field)
	TriggerUserID string `json:"trigger_user_id"`
	// The trigger user's display name (optional, admin or system operation will not have this field)
	TriggerUserDisplayName string `json:"trigger_user_display_name"`
	// The trigger user's url (optional, admin or system operation will not have this field)
	TriggerUserUrl string `json:"trigger_user_url"`

	// the question title
	QuestionTitle string `json:"question_title"`
	// the question url
	QuestionUrl string `json:"question_url"`
	// the question tags (comma separated, optional, only for new question notification)
	QuestionTags string `json:"tags"`

	// the answer url (optional, only for new answer notification)
	AnswerUrl string `json:"answer_url"`
	// the comment url (optional, only for new comment notification)
	CommentUrl string `json:"comment_url"`
}

var (
	// CallNotification is a function that calls all registered notification plugins
	CallNotification,
	registerNotification = MakePlugin[Notification](false)
)
