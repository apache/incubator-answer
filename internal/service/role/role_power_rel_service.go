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

package role

import (
	"context"
)

// RolePowerRelRepo rolePowerRel repository
type RolePowerRelRepo interface {
	GetRolePowerTypeList(ctx context.Context, roleID int) (powers []string, err error)
}

// RolePowerRelService user service
type RolePowerRelService struct {
	rolePowerRelRepo   RolePowerRelRepo
	userRoleRelService *UserRoleRelService
}

// NewRolePowerRelService new role power rel service
func NewRolePowerRelService(rolePowerRelRepo RolePowerRelRepo,
	userRoleRelService *UserRoleRelService) *RolePowerRelService {
	return &RolePowerRelService{
		rolePowerRelRepo:   rolePowerRelRepo,
		userRoleRelService: userRoleRelService,
	}
}

// GetRolePowerList get role power list
func (rs *RolePowerRelService) GetRolePowerList(ctx context.Context, roleID int) (powers []string, err error) {
	return rs.rolePowerRelRepo.GetRolePowerTypeList(ctx, roleID)
}

// GetUserPowerList get  list all
func (rs *RolePowerRelService) GetUserPowerList(ctx context.Context, userID string) (powers []string, err error) {
	roleID, err := rs.userRoleRelService.GetUserRole(ctx, userID)
	if err != nil {
		return nil, err
	}
	return rs.rolePowerRelRepo.GetRolePowerTypeList(ctx, roleID)
}
