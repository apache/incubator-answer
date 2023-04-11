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
