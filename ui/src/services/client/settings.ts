import useSWR from 'swr';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const getLanguageConfig = () => {
  return request.get('/answer/api/v1/language/config');
};

export const getLanguageOptions = () => {
  return request.get<Type.LangsType[]>('/answer/api/v1/language/options');
};

export const updateUserInterface = (lang: string) => {
  return request.put('/answer/api/v1/user/interface', {
    language: lang,
  });
};

export const useGetNotificationConfig = () => {
  return useSWR<Type.NotificationConfig>(
    '/answer/api/v1/user/notification/config',
    request.instance.get,
  );
};

export const putNotificationConfig = (data: Type.NotificationConfig) => {
  return request.put('/answer/api/v1/user/notification/config', data);
};
