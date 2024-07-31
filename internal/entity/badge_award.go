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

// BadgeAward badge_award
type BadgeAward struct {
	ID             string    `json:"id" xorm:"id"`
	CreatedAt      time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"updated_at"`
	UserId         int64     `json:"user_id" xorm:"user_id"`
	BadgeId        int64     `json:"badge_id" xorm:"badge_id"`
	ObjectId       int64     `json:"object_id" xorm:"object_id"`
	BadgeGroupId   int8      `json:"badge_group_id" xorm:"badge_group_id"`
	IsBadgeDeleted int8      `json:"is_badge_deleted" xorm:"is_badge_deleted"`
}

// TableName badge_award table name
func (*BadgeAward) TableName() string {
	return "badge_award"
}
