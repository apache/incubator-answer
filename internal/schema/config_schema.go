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

// AddConfigReq add config request
type AddConfigReq struct {
	// the config key
	Key string `validate:"omitempty,gt=0,lte=32" comment:"the config key" json:"key"`
	// the config value, custom data structures and types
	Value string `validate:"omitempty,gt=0,lte=128" comment:"the config value, custom data structures and types" json:"value"`
}

// RemoveConfigReq delete config request
type RemoveConfigReq struct {
	// config id
	ID int `validate:"required" comment:"config id" json:"id"`
}

// UpdateConfigReq update config request
type UpdateConfigReq struct {
	// config id
	ID int `validate:"required" comment:"config id" json:"id"`
	// the config key
	Key string `validate:"omitempty,gt=0,lte=32" comment:"the config key" json:"key"`
	// the config value, custom data structures and types
	Value string `validate:"omitempty,gt=0,lte=128" comment:"the config value, custom data structures and types" json:"value"`
}

// GetConfigListReq get config list all request
type GetConfigListReq struct {
	// the config key
	Key string `validate:"omitempty,gt=0,lte=32" comment:"the config key" form:"key"`
	// the config value, custom data structures and types
	Value string `validate:"omitempty,gt=0,lte=128" comment:"the config value, custom data structures and types" form:"value"`
}

// GetConfigWithPageReq get config list page request
type GetConfigWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// the config key
	Key string `validate:"omitempty,gt=0,lte=32" comment:"the config key" form:"key"`
	// the config value, custom data structures and types
	Value string `validate:"omitempty,gt=0,lte=128" comment:"the config value, custom data structures and types" form:"value"`
}

// GetConfigResp get config response
type GetConfigResp struct {
	// config id
	ID int `json:"id"`
	// the config key
	Key string `json:"key"`
	// the config value, custom data structures and types
	Value string `json:"value"`
}
