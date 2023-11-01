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

// AddTagListReq add tag list request
type AddTagListReq struct {
	// tag_id
	TagID int64 `validate:"required" comment:"tag_id" json:"tag_id"`
	// object_id
	ObjectID int64 `validate:"required" comment:"object_id" json:"object_id"`
	// tag_list_status(available: 1; deleted: 10)
	Status int `validate:"required" comment:"tag_list_status(available: 1; deleted: 10)" json:"status"`
}

// RemoveTagListReq delete tag list request
type RemoveTagListReq struct {
	// tag_list_id
	ID int64 `validate:"required" comment:"tag_list_id" json:"id"`
}

// UpdateTagListReq update tag list request
type UpdateTagListReq struct {
	// tag_list_id
	ID int64 `validate:"required" comment:"tag_list_id" json:"id"`
	// tag_id
	TagID int64 `validate:"omitempty" comment:"tag_id" json:"tag_id"`
	// object_id
	ObjectID int64 `validate:"omitempty" comment:"object_id" json:"object_id"`
	// tag_list_status(available: 1; deleted: 10)
	Status int `validate:"omitempty" comment:"tag_list_status(available: 1; deleted: 10)" json:"status"`
}

// GetTagListListReq get tag list list all request
type GetTagListListReq struct {
	// tag_id
	TagID int64 `validate:"omitempty" comment:"tag_id" form:"tag_id"`
	// object_id
	ObjectID int64 `validate:"omitempty" comment:"object_id" form:"object_id"`
	// tag_list_status(available: 1; deleted: 10)
	Status int `validate:"omitempty" comment:"tag_list_status(available: 1; deleted: 10)" form:"status"`
}

// GetTagListWithPageReq get tag list list page request
type GetTagListWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// tag_id
	TagID int64 `validate:"omitempty" comment:"tag_id" form:"tag_id"`
	// object_id
	ObjectID int64 `validate:"omitempty" comment:"object_id" form:"object_id"`
	// tag_list_status(available: 1; deleted: 10)
	Status int `validate:"omitempty" comment:"tag_list_status(available: 1; deleted: 10)" form:"status"`
}

// GetTagListResp get tag list response
type GetTagListResp struct {
	// tag_list_id
	ID int64 `json:"id"`
	// tag_id
	TagID int64 `json:"tag_id"`
	// object_id
	ObjectID int64 `json:"object_id"`
	// tag_list_status(available: 1; deleted: 10)
	Status int `json:"status"`
}
