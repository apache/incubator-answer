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
	"github.com/apache/incubator-answer/internal/service/event_queue"
	"github.com/segmentfault/pacman/log"

	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/schema"
)

type BadgeEventService struct {
	data              *data.Data
	eventQueueService event_queue.EventQueueService
}

func NewBadgeEventService(
	data *data.Data,
	eventQueueService event_queue.EventQueueService,
) *BadgeEventService {
	n := &BadgeEventService{
		data:              data,
		eventQueueService: eventQueueService,
	}
	eventQueueService.RegisterHandler(n.Handler)
	return n
}

func (ns *BadgeEventService) Handler(ctx context.Context, msg *schema.EventMsg) error {
	log.Debugf("received badge event %+v", msg)
	// TODO: Check if badge already exists

	// TODO: Check rule

	// TODO: Distribute badge

	return nil
}
