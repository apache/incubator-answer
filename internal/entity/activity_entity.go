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
	ActivityAvailable = 0
	ActivityCancelled = 1
)

// Activity activity
type Activity struct {
	ID               string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt        time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt        time.Time `xorm:"updated TIMESTAMP updated_at"`
	CancelledAt      time.Time `xorm:"TIMESTAMP cancelled_at"`
	UserID           string    `xorm:"not null index BIGINT(20) user_id"`
	TriggerUserID    int64     `xorm:"not null default 0 index BIGINT(20) trigger_user_id"`
	ObjectID         string    `xorm:"not null default 0 index BIGINT(20) object_id"`
	OriginalObjectID string    `xorm:"not null default 0 BIGINT(20) original_object_id"`
	ActivityType     int       `xorm:"not null INT(11) activity_type"`
	Cancelled        int       `xorm:"not null default 0 TINYINT(4) cancelled"`
	Rank             int       `xorm:"not null default 0 INT(11) rank"`
	HasRank          int       `xorm:"not null default 0 TINYINT(4) has_rank"`
	RevisionID       int64     `xorm:"not null default 0 BIGINT(20) revision_id"`
}

type ActivityRankSum struct {
	Rank int `xorm:"not null default 0 INT(11) rank"`
}

type ActivityUserRankStat struct {
	UserID string `xorm:"user_id"`
	Rank   int    `xorm:"rank_amount"`
}

type ActivityUserVoteStat struct {
	UserID    string `xorm:"user_id"`
	VoteCount int    `xorm:"vote_count"`
}

// TableName activity table name
func (Activity) TableName() string {
	return "activity"
}
