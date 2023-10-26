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

package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/auth"
	"github.com/stretchr/testify/assert"
)

var (
	token  = "token"
	userID = "1"
)

func Test_authRepo_SetUserCacheInfo(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetUserCacheInfo(context.TODO(), token, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	cacheInfo, err := authRepo.GetUserCacheInfo(context.TODO(), token)
	assert.NoError(t, err)
	assert.Equal(t, userID, cacheInfo.UserID)
}

func Test_authRepo_RemoveUserCacheInfo(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetUserCacheInfo(context.TODO(), token, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	err = authRepo.RemoveUserCacheInfo(context.TODO(), token)
	assert.NoError(t, err)

	_, err = authRepo.GetUserCacheInfo(context.TODO(), token)
	assert.Error(t, err)
}

func Test_authRepo_SetUserStatus(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetUserStatus(context.TODO(), userID, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	cacheInfo, err := authRepo.GetUserStatus(context.TODO(), userID)
	assert.NoError(t, err)
	assert.Equal(t, userID, cacheInfo.UserID)
}
func Test_authRepo_RemoveUserStatus(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetUserStatus(context.TODO(), userID, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	err = authRepo.RemoveUserStatus(context.TODO(), userID)
	assert.NoError(t, err)

	_, err = authRepo.GetUserStatus(context.TODO(), userID)
	assert.Error(t, err)
}

func Test_authRepo_SetAdminUserCacheInfo(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetAdminUserCacheInfo(context.TODO(), token, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	cacheInfo, err := authRepo.GetAdminUserCacheInfo(context.TODO(), token)
	assert.NoError(t, err)
	assert.Equal(t, userID, cacheInfo.UserID)
}

func Test_authRepo_RemoveAdminUserCacheInfo(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetAdminUserCacheInfo(context.TODO(), token, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	err = authRepo.RemoveAdminUserCacheInfo(context.TODO(), token)
	assert.NoError(t, err)

	_, err = authRepo.GetAdminUserCacheInfo(context.TODO(), token)
	assert.Error(t, err)
}
