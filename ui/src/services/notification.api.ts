import useSWR from 'swr';
import qs from 'qs';

import request from '@answer/utils/request';
import { isLogin } from '@answer/utils';

import type * as Type from './types';

export const useQueryNotifications = (params) => {
  const apiUrl = `/answer/api/v1/notification/page?${qs.stringify(params, {
    skipNulls: true,
  })}`;

  const { data, error, mutate } = useSWR<Type.RecordResult>(
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

export const useQueryNotificationRedDot = () => {
  const apiUrl = '/answer/api/v1/notification/status';

  return useSWR<{ inbox: number; achievement: number }>(
    isLogin() ? apiUrl : null,
    request.instance.get,
    {
      refreshInterval: 3000,
    },
  );
};

export const clearNotificationRedDot = (type) => {
  return request.instance.put('/answer/api/v1/notification/status', {
    type,
  });
};

export const clearUnReadNotification = (type) => {
  return request.instance.put('/answer/api/v1/notification/read/state/all', {
    type,
  });
};
