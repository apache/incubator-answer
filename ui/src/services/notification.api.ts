import useSWR from 'swr';
import qs from 'qs';

import request from '@answer/utils/request';
import { isLogin } from '@answer/utils';

import type * as Type from './types';

export const useQueryNotifications = (params) => {
  const apiUrl = `/answer/api/v1/notification/list?${qs.stringify(params, {
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
  return request.instance.post('/answer/api/v1/notification/read', {
    id,
  });
};

export const useQueryNotificationRedDot = () => {
  const apiUrl = '/answer/api/v1/notification/reddot';

  return useSWR<{ inbox: number; achievement: number }>(
    isLogin() ? apiUrl : null,
    request.instance.get,
  );
};

export const clearNotificationRedDot = (type) => {
  return request.instance.post('/answer/api/v1/notification/reddot/clear', {
    type,
  });
};

export const clearUnReadNotification = (type) => {
  return request.instance.post('/answer/api/v1/notification/unread/clear', {
    type,
  });
};
