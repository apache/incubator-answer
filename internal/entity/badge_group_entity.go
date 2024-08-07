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

// BadgeGroup badge_group
type BadgeGroup struct {
	ID        string    `json:"id" xorm:"not null pk autoincr BIGINT(20) id"`
	Name      string    `json:"name" xorm:"not null default '' VARCHAR(256) name"`
	CreatedAt time.Time `json:"created_at" xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated not null default CURRENT_TIMESTAMP TIMESTAMP updated_at"`
}

// TableName badge_group table name
func (BadgeGroup) TableName() string {
	return "badge_group"
}
