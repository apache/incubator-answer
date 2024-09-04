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

const (
	IsBadgeNotDeleted = 0
	IsBadgeDeleted    = 1

	BadgeEmptyAwardKey = "0"
)

// BadgeAward badge_award
type BadgeAward struct {
	ID             string    `xorm:"not null pk BIGINT(20) id"`
	CreatedAt      time.Time `xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt      time.Time `xorm:"updated not null default CURRENT_TIMESTAMP TIMESTAMP updated_at"`
	UserID         string    `xorm:"not null index BIGINT(20) user_id"`
	BadgeID        string    `xorm:"not null index BIGINT(20) badge_id"`
	AwardKey       string    `xorm:"not null index VARCHAR(64) award_key"`
	BadgeGroupID   int64     `xorm:"not null index BIGINT(20) badge_group_id"`
	IsBadgeDeleted int8      `xorm:"not null TINYINT(1) is_badge_deleted"`
}

// TableName badge_award table name
func (BadgeAward) TableName() string {
	return "badge_award"
}

type BadgeEarnedCount struct {
	BadgeID     string `xorm:"badge_id"`
	EarnedCount int64  `xorm:"earned_count"`
}

// TableName badge_award table name
func (BadgeEarnedCount) TableName() string {
	return "badge_award"
}

type BadgeAwardRecent struct {
	Created        time.Time `xorm:"created"`
	BadgeID        string    `xorm:"badge_id"`
	AwardKey       string    `xorm:"award_key"`
	EarnedCount    int64     `xorm:"earned_count"`
	IsBadgeDeleted int8      `xorm:"is_badge_deleted"`
}

// TableName badge_award table name
func (BadgeAwardRecent) TableName() string {
	return "badge_award"
}
