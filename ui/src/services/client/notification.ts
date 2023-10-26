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

import useSWR from 'swr';
import qs from 'qs';

import request from '@/utils/request';
import type * as Type from '@/common/interface';
import { tryLoggedAndActivated } from '@/utils/guard';

export const useQueryNotifications = (params) => {
  const apiUrl = `/answer/api/v1/notification/page?${qs.stringify(params, {
    skipNulls: true,
  })}`;

  const { data, error, mutate } = useSWR<Type.ListResult>(
    apiUrl,
    request.instance.get,
  );

  return {
    data,
    isLoading: !data && !error,
    error,
    mutate,
  };
};

export const readNotification = (id) => {
  return request.instance.put('/answer/api/v1/notification/read/state', {
    id,
  });
};

export const useQueryNotificationStatus = () => {
  const apiUrl = '/answer/api/v1/notification/status';

  return useSWR<Type.NotificationStatus>(
    tryLoggedAndActivated().ok ? apiUrl : null,
    (url) => request.get(url, { ignoreError: '50X' }),
    {
      refreshInterval: 3000,
    },
  );
};

export const clearNotificationStatus = (type) => {
  return request.instance.put('/answer/api/v1/notification/status', {
    type,
  });
};

export const clearUnreadNotification = (type) => {
  return request.instance.put('/answer/api/v1/notification/read/state/all', {
    type,
  });
};
