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

import request from '@/utils/request';

export interface UcAgentControl {
  name: string;
  label: string;
  url: string;
}
export interface UcAgent {
  enabled: boolean;
  agent_info: {
    name: string;
    icon: string;
    url: string;
    display_name: string;
    login_redirect_url: string;
    sign_up_redirect_url: string;
    control_center: UcAgentControl[];
    enabled_original_user_system: boolean;
  };
}

export interface UcSettingAgent {
  enabled: boolean;
  redirect_url: string;
}
export interface UcSettings {
  profile_setting_agent: UcSettingAgent;
  account_setting_agent: UcSettingAgent;
}

export interface UcBrandingEntry {
  icon: string;
  name: string;
  label: string;
  url: string;
}
export interface UcBranding {
  enabled: boolean;
  personal_branding: UcBrandingEntry[];
}

export interface AdminUcAgent {
  allow_create_user: boolean;
  allow_update_user_status: boolean;
  allow_update_user_password: boolean;
  allow_update_user_role: boolean;
}

export const getUcAgent = () => {
  const apiUrl = `/answer/api/v1/user-center/agent`;
  return request.get<UcAgent>(apiUrl);
};
export const getAdminUcAgent = () => {
  const apiUrl = `/answer/admin/api/user-center/agent`;
  return request.get<AdminUcAgent>(apiUrl);
};
export const getUcSettings = () => {
  const apiUrl = `/answer/api/v1/user-center/user/settings`;
  return request.get<UcSettings>(apiUrl);
};

export const getUcBranding = (username: string) => {
  const apiUrl = `/answer/api/v1/user-center/personal/branding?username=${username}`;
  return request.get<UcBranding>(apiUrl);
};
