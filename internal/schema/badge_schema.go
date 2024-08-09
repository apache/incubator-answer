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

import "github.com/apache/incubator-answer/internal/entity"

// BadgeListInfo get badge list response
type BadgeListInfo struct {
	ID         string            `json:"id" `
	Name       string            `json:"name" `
	Icon       string            `json:"icon" `
	AwardCount int               `json:"award_count" `
	Earned     bool              `json:"earned" `
	Level      entity.BadgeLevel `json:"level" `
}

type GetBadgeListResp struct {
	Badges    []*BadgeListInfo `json:"badges" `
	GroupName string           `json:"group_name" `
}

type GetBadgeInfoResp struct {
	ID          string            `json:"id" `
	Name        string            `json:"name" `
	Description string            `json:"description" `
	Icon        string            `json:"icon" `
	AwardCount  int               `json:"award_count" `
	EarnedCount int64             `json:"earned_count" `
	IsSingle    bool              `json:"is_single" `
	Level       entity.BadgeLevel `json:"level" `
}

type GetBadgeAwardWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// badge id
	BadgeID string `validate:"required" form:"badge_id"`
	// user id
	UserID string `json:"-"`
}

type GetBadgeAwardWithPageResp struct {
	CreatedAt      int64         `json:"created_at"`
	ObjectID       string        `json:"object_id"`
	QuestionID     string        `json:"question_id"`
	AnswerID       string        `json:"answer_id"`
	CommentID      string        `json:"comment_id"`
	ObjectType     string        `json:"object_type" enums:"question,answer,comment"`
	UrlTitle       string        `json:"url_title"`
	AuthorUserInfo UserBasicInfo `json:"author_user_info"`
}

type GetUserBadgeAwardListReq struct {
	Username string `validate:"omitempty,gt=0,lte=100" form:"username"`
	UserID   string `json:"-"`
}
type GetUserBadgeAwardListResp struct {
	ID          string            `json:"id" `
	Name        string            `json:"name" `
	Icon        string            `json:"icon" `
	EarnedCount int64             `json:"earned_count" `
	Level       entity.BadgeLevel `json:"level" `
}
