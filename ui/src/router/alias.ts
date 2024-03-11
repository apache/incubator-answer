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
export const BASE_URL_PATH = process.env.REACT_APP_BASE_URL_PATH;
export const RouteAlias = {
  home: `${BASE_URL_PATH}/`,
  login: `${BASE_URL_PATH}/users/login`,
  signUp: `${BASE_URL_PATH}/users/register`,
  inactive: `${BASE_URL_PATH}/users/login?status=inactive`,
  accountRecovery: `${BASE_URL_PATH}/users/account-recovery`,
  changeEmail: `${BASE_URL_PATH}/users/change-email`,
  passwordReset: `${BASE_URL_PATH}/users/password-reset`,
  accountActivation: `${BASE_URL_PATH}/users/account-activation`,
  activationSuccess: `${BASE_URL_PATH}/users/account-activation/success`,
  activationFailed: `${BASE_URL_PATH}/users/account-activation/failed`,
  suspended: `${BASE_URL_PATH}/users/account-suspended`,
  confirmNewEmail: `${BASE_URL_PATH}/users/confirm-new-email`,
  confirmEmail: `${BASE_URL_PATH}/users/confirm-email`,
  authLanding: `${BASE_URL_PATH}/users/auth-landing`,
};
