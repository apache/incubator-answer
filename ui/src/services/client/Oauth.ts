import useSWR from 'swr';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const oAuthBindEmail = (data: Type.OauthBindEmailReq) => {
  return request.post('/answer/api/v1/connector/binding/email', data);
};

export const useOauthConnectorInfoByUser = () => {
  const { data, error, mutate } = useSWR<Type.UserOauthConnectorItem[]>(
    '/answer/api/v1/connector/user/info',
    request.instance.get,
  );
  return {
    data,
    mutate,
    isLoading: !data && !error,
    error,
  };
};

export const userOauthUnbind = (data: { external_id: string }) => {
  return request.delete('/answer/api/v1/connector/user/unbinding', data);
};
