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
	"github.com/apache/incubator-answer/internal/base/constant"
	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/entity"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/internal/service/badge"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
	"strconv"
)

// eventRuleRepo event rule repo
type eventRuleRepo struct {
	data             *data.Data
	EventRuleMapping map[constant.EventType][]badge.EventRuleHandler
}

// NewEventRuleRepo creates a new badge repository
func NewEventRuleRepo(data *data.Data) badge.EventRuleRepo {
	b := &eventRuleRepo{
		data: data,
	}
	b.EventRuleMapping = map[constant.EventType][]badge.EventRuleHandler{
		constant.EventUserUpdate:     {b.FirstUpdateUserProfile},
		constant.EventUserShare:      {b.FirstSharedPost},
		constant.EventQuestionCreate: nil,
		constant.EventQuestionUpdate: {b.FirstPostEdit},
		constant.EventQuestionDelete: nil,
		constant.EventQuestionVote:   {b.FirstVotedPost, b.ReachQuestionVote},
		constant.EventQuestionAccept: {b.FirstAcceptAnswer, b.ReachAnswerAcceptedAmount},
		constant.EventQuestionFlag:   {b.FirstFlaggedPost},
		constant.EventQuestionReact:  {b.FirstReactedPost},
		constant.EventAnswerCreate:   nil,
		constant.EventAnswerUpdate:   {b.FirstPostEdit},
		constant.EventAnswerDelete:   nil,
		constant.EventAnswerVote:     {b.FirstVotedPost, b.ReachAnswerVote},
		constant.EventAnswerFlag:     {b.FirstFlaggedPost},
		constant.EventAnswerReact:    {b.FirstReactedPost},
		constant.EventCommentCreate:  nil,
		constant.EventCommentUpdate:  nil,
		constant.EventCommentDelete:  nil,
		constant.EventCommentVote:    {b.FirstVotedPost},
		constant.EventCommentFlag:    {b.FirstFlaggedPost},
	}
	return b
}

// HandleEventWithRule handle event with rule
func (br *eventRuleRepo) HandleEventWithRule(ctx context.Context, msg *schema.EventMsg) (
	awards []*entity.BadgeAward) {
	handlers := br.EventRuleMapping[msg.EventType]
	for _, h := range handlers {
		t, err := h(ctx, msg)
		if err != nil {
			log.Errorf("error handling badge event %+v: %v", msg, err)
		} else {
			awards = append(awards, t...)
		}
	}
	return awards
}

// FirstUpdateUserProfile first update user profile
func (br *eventRuleRepo) FirstUpdateUserProfile(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "FirstUpdateUserProfile")
	if b == nil {
		return nil, nil
	}
	bean := &entity.User{ID: event.UserID}
	exist, err := br.data.DB.Context(ctx).Get(bean)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	if !exist {
		return nil, nil
	}
	if len(bean.Bio) > 0 {
		return append(awards, br.createBadgeAward(event.UserID, b.ID, entity.BadgeOnceAwardKey)), nil
	}
	return nil, nil
}

// FirstPostEdit first post edit
func (br *eventRuleRepo) FirstPostEdit(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "FirstPostEdit")
	if b == nil {
		return nil, nil
	}
	return append(awards, br.createBadgeAward(event.UserID, b.ID, event.GetObjectID())), nil
}

// FirstFlaggedPost first flagged post.
func (br *eventRuleRepo) FirstFlaggedPost(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "FirstFlaggedPost")
	if b == nil {
		return nil, nil
	}
	return append(awards, br.createBadgeAward(event.UserID, b.ID, event.GetObjectID())), nil
}

// FirstVotedPost first voted post
func (br *eventRuleRepo) FirstVotedPost(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "FirstVotedPost")
	if b == nil {
		return nil, nil
	}
	return append(awards, br.createBadgeAward(event.UserID, b.ID, event.GetObjectID())), nil
}

// FirstReactedPost first reacted post
func (br *eventRuleRepo) FirstReactedPost(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "FirstReactedPost")
	if b == nil {
		return nil, nil
	}
	return append(awards, br.createBadgeAward(event.UserID, b.ID, event.GetObjectID())), nil
}

// FirstSharedPost first shared post
func (br *eventRuleRepo) FirstSharedPost(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "FirstSharedPost")
	if b == nil {
		return nil, nil
	}
	return append(awards, br.createBadgeAward(event.UserID, b.ID, event.GetObjectID())), nil
}

// FirstAcceptAnswer user first accept answer
func (br *eventRuleRepo) FirstAcceptAnswer(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "FirstAcceptAnswer")
	if b == nil {
		return nil, nil
	}
	return append(awards, br.createBadgeAward(event.UserID, b.ID, event.GetObjectID())), nil
}

// ReachAnswerAcceptedAmount reach answer accepted amount
func (br *eventRuleRepo) ReachAnswerAcceptedAmount(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "ReachAnswerAcceptedAmount")
	if b == nil {
		return nil, nil
	}
	if len(event.AnswerUserID) == 0 {
		return nil, nil
	}

	// count user's accepted answer amount
	amount, err := br.data.DB.Context(ctx).Count(&entity.Answer{
		UserID:   event.AnswerUserID,
		Accepted: schema.AnswerAcceptedEnable,
		Status:   entity.AnswerStatusAvailable,
	})
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	// get badge requirement
	requirement := b.GetIntParam("amount")
	if requirement == 0 || amount < requirement {
		return nil, nil
	}

	return append(awards, br.createBadgeAward(event.UserID, b.ID, event.GetObjectID())), nil
}

// ReachAnswerVote reach answer vote
func (br *eventRuleRepo) ReachAnswerVote(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "ReachAnswerVote")
	if b == nil {
		return nil, nil
	}

	// get vote amount
	amount, _ := strconv.Atoi(event.GetExtra("vote_up_amount"))
	if amount == 0 {
		return nil, nil
	}

	// get badge requirement
	requirement := b.GetIntParam("amount")
	if requirement == 0 || int64(amount) < requirement {
		return nil, nil
	}

	return append(awards, br.createBadgeAward(event.AnswerUserID, b.ID, event.AnswerID)), nil
}

// ReachQuestionVote reach question vote
func (br *eventRuleRepo) ReachQuestionVote(ctx context.Context,
	event *schema.EventMsg) (awards []*entity.BadgeAward, err error) {
	b := br.getBadgeByHandler(ctx, "ReachQuestionVote")
	if b == nil {
		return nil, nil
	}

	// get vote amount
	amount, _ := strconv.Atoi(event.GetExtra("vote_up_amount"))
	if amount == 0 {
		return nil, nil
	}

	// get badge requirement
	requirement := b.GetIntParam("amount")
	if requirement == 0 || int64(amount) < requirement {
		return nil, nil
	}

	return append(awards, br.createBadgeAward(event.QuestionUserID, b.ID, event.QuestionID)), nil
}

func (br *eventRuleRepo) getBadgeByHandler(ctx context.Context, handler string) (b *entity.Badge) {
	b = &entity.Badge{Handler: handler}
	exist, err := br.data.DB.Context(ctx).Get(b)
	if err != nil {
		log.Errorf("error getting badge by handler %s: %v", handler, err)
		return nil
	}
	if !exist {
		log.Errorf("badge not found by handler %s", handler)
		return nil
	}
	return b
}

func (br *eventRuleRepo) createBadgeAward(userID, badgeID, awardKey string) (awards *entity.BadgeAward) {
	return &entity.BadgeAward{
		UserID:   userID,
		BadgeID:  badgeID,
		AwardKey: awardKey,
	}
}
