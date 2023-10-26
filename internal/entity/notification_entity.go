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

// Notification notification
type Notification struct {
	ID        string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"TIMESTAMP updated_at"`
	UserID    string    `xorm:"not null default 0 BIGINT(20) INDEX user_id"`
	ObjectID  string    `xorm:"not null default 0 INDEX BIGINT(20) object_id"`
	Content   string    `xorm:"not null TEXT content"`
	Type      int       `xorm:"not null default 0 INT(11) type"`
	MsgType   int       `xorm:"not null default 0 INT(11) msg_type"`
	IsRead    int       `xorm:"not null default 1 INT(11) is_read"`
	Status    int       `xorm:"not null default 1 INT(11) status"`
}

// TableName notification table name
func (Notification) TableName() string {
	return "notification"
}
