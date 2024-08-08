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

package badge

import "github.com/apache/incubator-answer/internal/base/constant"

var (
	EventRuleMapping = map[constant.EventType][]string{
		constant.EventUserUpdate:     {FilledPersonalProfile},
		constant.EventUserShare:      {FirstSharedPost},
		constant.EventQuestionCreate: {""},
		constant.EventQuestionUpdate: {FirstPostEdit},
		constant.EventQuestionDelete: {""},
		constant.EventQuestionVote:   {FirstVotedPost, ReachQuestionScore},
		constant.EventQuestionAccept: {AskQuestionAcceptAnswer, AnswerAccepted},
		constant.EventQuestionFlag:   {FirstFlaggedPost},
		constant.EventQuestionReact:  {FirstReactedPost},
		constant.EventAnswerCreate:   {""},
		constant.EventAnswerUpdate:   {FirstPostEdit},
		constant.EventAnswerDelete:   {""},
		constant.EventAnswerVote:     {FirstVotedPost, ReachAnswerScore},
		constant.EventAnswerFlag:     {FirstFlaggedPost},
		constant.EventAnswerReact:    {FirstReactedPost},
		constant.EventCommentCreate:  {""},
		constant.EventCommentUpdate:  {""},
		constant.EventCommentDelete:  {""},
		constant.EventCommentVote:    {FirstVotedPost},
		constant.EventCommentFlag:    {FirstFlaggedPost},
	}
)
