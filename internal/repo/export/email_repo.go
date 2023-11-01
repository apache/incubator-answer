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

package export

import (
	"context"
	"time"

	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/reason"
	"github.com/apache/incubator-answer/internal/service/export"
	"github.com/segmentfault/pacman/errors"
)

// emailRepo email repository
type emailRepo struct {
	data *data.Data
}

// NewEmailRepo new repository
func NewEmailRepo(data *data.Data) export.EmailRepo {
	return &emailRepo{
		data: data,
	}
}

// SetCode The email code is used to verify that the link in the message is out of date
func (e *emailRepo) SetCode(ctx context.Context, code, content string, duration time.Duration) error {
	err := e.data.Cache.SetString(ctx, code, content, duration)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// VerifyCode verify the code if out of date
func (e *emailRepo) VerifyCode(ctx context.Context, code string) (content string, err error) {
	content, exist, err := e.data.Cache.GetString(ctx, code)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", nil
	}
	return content, nil
}
