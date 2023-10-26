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

// Code generated by MockGen. DO NOT EDIT.
// Source: ./siteinfo_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/answerdev/answer/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockSiteInfoRepo is a mock of SiteInfoRepo interface.
type MockSiteInfoRepo struct {
	ctrl     *gomock.Controller
	recorder *MockSiteInfoRepoMockRecorder
}

// MockSiteInfoRepoMockRecorder is the mock recorder for MockSiteInfoRepo.
type MockSiteInfoRepoMockRecorder struct {
	mock *MockSiteInfoRepo
}

// NewMockSiteInfoRepo creates a new mock instance.
func NewMockSiteInfoRepo(ctrl *gomock.Controller) *MockSiteInfoRepo {
	mock := &MockSiteInfoRepo{ctrl: ctrl}
	mock.recorder = &MockSiteInfoRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSiteInfoRepo) EXPECT() *MockSiteInfoRepoMockRecorder {
	return m.recorder
}

// GetByType mocks base method.
func (m *MockSiteInfoRepo) GetByType(ctx context.Context, siteType string) (*entity.SiteInfo, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByType", ctx, siteType)
	ret0, _ := ret[0].(*entity.SiteInfo)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByType indicates an expected call of GetByType.
func (mr *MockSiteInfoRepoMockRecorder) GetByType(ctx, siteType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByType", reflect.TypeOf((*MockSiteInfoRepo)(nil).GetByType), ctx, siteType)
}

// SaveByType mocks base method.
func (m *MockSiteInfoRepo) SaveByType(ctx context.Context, siteType string, data *entity.SiteInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveByType", ctx, siteType, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveByType indicates an expected call of SaveByType.
func (mr *MockSiteInfoRepoMockRecorder) SaveByType(ctx, siteType, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveByType", reflect.TypeOf((*MockSiteInfoRepo)(nil).SaveByType), ctx, siteType, data)
}
