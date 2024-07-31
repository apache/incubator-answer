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

package entity

import "time"

// Badge badge
type Badge struct {
	ID           string    `json:"id" xorm:"id"`
	CreatedAt    time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"updated_at"`
	Name         string    `json:"name" xorm:"name"`
	AwardTotal   int64     `json:"award_total" xorm:"award_total"`
	Description  string    `json:"description" xorm:"description"`
	Status       int8      `json:"status" xorm:"status"`
	BadgeGroupId int64     `json:"badge_group_id" xorm:"badge_group_id"`
	Single       int8      `json:"single" xorm:"single"`
	Collect      string    `json:"collect" xorm:"collect"`
	Handler      string    `json:"handler" xorm:"handler"`
	Param        string    `json:"param" xorm:"param"`
}

// TableName badge table name
func (*Badge) TableName() string {
	return "badge"
}
