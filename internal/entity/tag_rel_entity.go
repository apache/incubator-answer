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
	TagRelStatusAvailable = 1
	TagRelStatusHide      = 2
	TagRelStatusDeleted   = 10
)

// TagRel tag relation
type TagRel struct {
	ID        int64     `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	ObjectID  string    `xorm:"not null INDEX UNIQUE(s) BIGINT(20) object_id"`
	TagID     string    `xorm:"not null INDEX UNIQUE(s) BIGINT(20) tag_id"`
	Status    int       `xorm:"not null default 1 INT(11) status"`
}

// TableName tag list table name
func (TagRel) TableName() string {
	return "tag_rel"
}
