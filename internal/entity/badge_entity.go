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

type BadgeLevel int

const (
	BadgeStatusActive   = 1
	BadgeStatusDeleted  = 10
	BadgeStatusInactive = 11

	BadgeLevelBronze BadgeLevel = 1
	BadgeLevelSilver BadgeLevel = 2
	BadgeLevelGold   BadgeLevel = 3

	BadgeSingleAward = 1
	BadgeMultiAward  = 2
)

// Badge badge
type Badge struct {
	ID           string     `json:"id" xorm:"id"`
	CreatedAt    time.Time  `json:"created_at" xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt    time.Time  `json:"updated_at" xorm:"updated not null default CURRENT_TIMESTAMP TIMESTAMP updated_at"`
	Name         string     `json:"name" xorm:"not null default '' VARCHAR(256) name"`
	Icon         string     `json:"icon" xorm:"not null default '' VARCHAR(1024) icon"`
	AwardCount   int        `json:"award_count" xorm:"not null default 0 INT(11) award_count"`
	Description  string     `json:"description" xorm:"not null default '' MEDIUMTEXT description"`
	Status       int8       `json:"status" xorm:"not null default 1 INT(11) status"`
	BadgeGroupId int64      `json:"badge_group_id" xorm:"not null default 0 BIGINT(20) badge_group_id"`
	Level        BadgeLevel `json:"level" xorm:"not null default 1 TINYINT(4) level"`
	Single       int8       `json:"single" xorm:"not null default 1 TINYINT(4) single"`
	Collect      string     `json:"collect" xorm:"not null default '' VARCHAR(64) collect"`
	Handler      string     `json:"handler" xorm:"not null default '' VARCHAR(64) handler"`
	Param        string     `json:"param" xorm:"not null default '' VARCHAR(128) param"`
}

// TableName badge table name
func (Badge) TableName() string {
	return "badge"
}
