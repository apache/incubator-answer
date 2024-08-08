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

import (
	"context"
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
)

// BadgeRuleRepo collection repository
type BadgeRuleRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

// FilledPersonalProfile filled personal profile
func (br *BadgeRuleRepo) FilledPersonalProfile(ctx context.Context, userID string) (reach bool, err error) {
	bean := &entity.User{ID: userID}
	exist, err := br.data.DB.Context(ctx).Get(bean)
	if err != nil {
		return false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	if !exist {
		return false, nil
	}
	if len(bean.Bio) > 0 {
		return true, nil
	}
	return false, nil
}

// FirstPostEdit first post edit
func (br *BadgeRuleRepo) FirstPostEdit(ctx context.Context, userID string, objectID string) {

}

// FirstFlaggedPost first flagged post.
func (br *BadgeRuleRepo) FirstFlaggedPost(ctx context.Context, userID string, reportID string) {
}

// FirstVotedPost first voted post
func (br *BadgeRuleRepo) FirstVotedPost(ctx context.Context) {

}

// FirstReactedPost first reacted post
func (br *BadgeRuleRepo) FirstReactedPost(ctx context.Context) {

}

// FirstSharedPost first shared post
func (br *BadgeRuleRepo) FirstSharedPost(ctx context.Context) {

}

// AskQuestionAcceptAnswer ask question accept answer
func (br *BadgeRuleRepo) AskQuestionAcceptAnswer(ctx context.Context) {

}

// AnswerAccepted answer accepted
func (br *BadgeRuleRepo) AnswerAccepted(ctx context.Context) {

}

// ReachAnswerScore reach answer score
func (br *BadgeRuleRepo) ReachAnswerScore(ctx context.Context) {

}

// ReachQuestionScore reach question score
func (br *BadgeRuleRepo) ReachQuestionScore(ctx context.Context) {

}
