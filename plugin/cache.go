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

package plugin

import (
	"context"
	"time"
)

type Cache interface {
	Base

	GetString(ctx context.Context, key string) (data string, exist bool, err error)
	SetString(ctx context.Context, key, value string, ttl time.Duration) (err error)
	GetInt64(ctx context.Context, key string) (data int64, exist bool, err error)
	SetInt64(ctx context.Context, key string, value int64, ttl time.Duration) (err error)
	Increase(ctx context.Context, key string, value int64) (data int64, err error)
	Decrease(ctx context.Context, key string, value int64) (data int64, err error)
	Del(ctx context.Context, key string) (err error)
	Flush(ctx context.Context) (err error)
}

var (
	// CallCache is a function that calls all registered cache
	CallCache,
	registerCache = MakePlugin[Cache](false)
)
