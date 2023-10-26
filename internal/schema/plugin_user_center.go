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

type UserCenterAgentResp struct {
	Enabled   bool       `json:"enabled"`
	AgentInfo *AgentInfo `json:"agent_info"`
}

type AgentInfo struct {
	Name                      string           `json:"name"`
	DisplayName               string           `json:"display_name"`
	Icon                      string           `json:"icon"`
	Url                       string           `json:"url"`
	LoginRedirectURL          string           `json:"login_redirect_url"`
	SignUpRedirectURL         string           `json:"sign_up_redirect_url"`
	ControlCenterItems        []*ControlCenter `json:"control_center"`
	EnabledOriginalUserSystem bool             `json:"enabled_original_user_system"`
}

type ControlCenter struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Url   string `json:"url"`
}

type UserCenterPersonalBranding struct {
	Enabled          bool                `json:"enabled"`
	PersonalBranding []*PersonalBranding `json:"personal_branding"`
}

type PersonalBranding struct {
	Icon  string `json:"icon"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Url   string `json:"url"`
}
