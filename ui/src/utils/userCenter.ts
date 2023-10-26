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

import { RouteAlias } from '@/router/alias';
import { userCenterStore } from '@/stores';
import { getUcAgent, UcAgent } from '@/services/user-center';

export const pullUcAgent = async () => {
  const uca = await getUcAgent();
  userCenterStore.getState().update(uca);
};

export const getLoginUrl = (uca?: UcAgent) => {
  let ret = RouteAlias.login;
  uca ||= userCenterStore.getState().agent;
  if (
    uca?.enabled &&
    !uca.agent_info?.enabled_original_user_system &&
    uca.agent_info?.login_redirect_url
  ) {
    ret = uca.agent_info.login_redirect_url;
  }
  return ret;
};

export const getSignUpUrl = (uca?: UcAgent) => {
  let ret = RouteAlias.signUp;
  uca ||= userCenterStore.getState().agent;
  if (uca?.enabled && uca?.agent_info?.sign_up_redirect_url) {
    ret = uca.agent_info.sign_up_redirect_url;
  }
  return ret;
};
