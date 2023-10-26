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
	"time"

	"github.com/answerdev/answer/internal/base/constant"
)

// AddRevisionDTO add revision request
type AddRevisionDTO struct {
	// user id
	UserID string
	// object id
	ObjectID string
	// title
	Title string
	// content
	Content string
	// log
	Log string
	// status
	Status int
}

// GetRevisionListReq get revision list all request
type GetRevisionListReq struct {
	// object id
	ObjectID string `validate:"required" comment:"object_id" form:"object_id"`
}

const RevisionAuditApprove = "approve"
const RevisionAuditReject = "reject"

type RevisionAuditReq struct {
	// object id
	ID                string `validate:"required" comment:"id" form:"id"`
	Operation         string `validate:"required" comment:"operation" form:"operation"` //approve or reject
	UserID            string `json:"-"`
	CanReviewQuestion bool   `json:"-"`
	CanReviewAnswer   bool   `json:"-"`
	CanReviewTag      bool   `json:"-"`
}

type RevisionSearch struct {
	Page              int    `json:"page" form:"page"` // Query number of pages
	CanReviewQuestion bool   `json:"-"`
	CanReviewAnswer   bool   `json:"-"`
	CanReviewTag      bool   `json:"-"`
	UserID            string `json:"-"`
}

func (r RevisionSearch) GetCanReviewObjectTypes() []int {
	objectType := make([]int, 0)
	if r.CanReviewAnswer {
		objectType = append(objectType, constant.ObjectTypeStrMapping[constant.AnswerObjectType])
	}
	if r.CanReviewQuestion {
		objectType = append(objectType, constant.ObjectTypeStrMapping[constant.QuestionObjectType])
	}
	if r.CanReviewTag {
		objectType = append(objectType, constant.ObjectTypeStrMapping[constant.TagObjectType])
	}
	return objectType
}

type GetUnreviewedRevisionResp struct {
	Type           string                      `json:"type"`
	Info           *UnreviewedRevisionInfoInfo `json:"info"`
	UnreviewedInfo *GetRevisionResp            `json:"unreviewed_info"`
}

// GetRevisionResp get revision response
type GetRevisionResp struct {
	// id
	ID string `json:"id"`
	// user id
	UserID string `json:"use_id"`
	// object id
	ObjectID string `json:"object_id"`
	// object type
	ObjectType int `json:"-"`
	// title
	Title string `json:"title"`
	// content
	Content string `json:"-"`
	// content parsed
	ContentParsed interface{} `json:"content"`
	// revision status(normal: 1; delete 2)
	Status int `json:"status"`
	// create time
	CreatedAt       time.Time     `json:"-"`
	CreatedAtParsed int64         `json:"create_at"`
	UserInfo        UserBasicInfo `json:"user_info"`
	Log             string        `json:"reason"`
}
