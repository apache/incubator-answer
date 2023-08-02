import qs from 'qs';
import useSWR from 'swr';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const changeUserStatus = (params) => {
  return request.put('/answer/admin/api/user/status', params);
};

export const useQueryUsers = (params) => {
  const apiUrl = `/answer/admin/api/users/page?${qs.stringify(params)}`;
  const { data, error, mutate } = useSWR<Type.ListResult, Error>(
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

export const getUserRoles = () => {
  return request.get('/answer/admin/api/roles');
};

export const changeUserRole = (params) => {
  return request.put('/answer/admin/api/user/role', params);
};

export const addUser = (params: {
  display_name: string;
  email: string;
  password: string;
}) => {
  return request.post('/answer/admin/api/user', params);
};

export const updateUserPassword = (params: {
  password: string;
  user_id: string;
}) => {
  return request.put('/answer/admin/api/user/password', params);
};

export const getUserActivation = (userId: string) => {
  const apiUrl = `/answer/admin/api/user/activation`;
  return request.get<{
    activation_url: string;
  }>(apiUrl, {
    params: {
      user_id: userId,
    },
  });
};

export const postUserActivation = (userId: string) => {
  const apiUrl = `/answer/admin/api/user/activation`;
  return request.post(apiUrl, {
    user_id: userId,
  });
};
