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

type UpdateReactionReq struct {
	ObjectID string `validate:"required" json:"object_id"`
	Emoji    string `validate:"required,oneof=heart smile frown" json:"emoji"`
	Reaction string `validate:"required,oneof=activate deactivate" json:"reaction"`
	UserID   string `json:"-"`
}

type GetReactionReq struct {
	ObjectID string `validate:"required" form:"object_id"`
}

type ReactSummaryMeta map[string][]string

type ReactionResp struct {
	// reaction summary is a map, key is emoji, value is username list
	// such as {"heart": ["jack", "tom"], "smile": ["andy"], "frown": ["bob"]}
	ReactionSummary ReactSummaryMeta `json:"reaction_summary"`
}
