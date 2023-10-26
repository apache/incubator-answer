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

// AcceptAnswerOperationInfo accept answer operation info
type AcceptAnswerOperationInfo struct {
	TriggerUserID    string
	QuestionObjectID string
	QuestionUserID   string
	AnswerObjectID   string
	AnswerUserID     string

	// vote activity info
	Activities []*AcceptAnswerActivity
}

// AcceptAnswerActivity accept answer activity
type AcceptAnswerActivity struct {
	ActivityType     int
	ActivityUserID   string
	TriggerUserID    string
	OriginalObjectID string
	Rank             int
}

func (v *AcceptAnswerActivity) HasRank() int {
	if v.Rank != 0 {
		return 1
	}
	return 0
}

func (a *AcceptAnswerOperationInfo) GetUserIDs() (userIDs []string) {
	for _, act := range a.Activities {
		userIDs = append(userIDs, act.ActivityUserID)
	}
	return userIDs
}
