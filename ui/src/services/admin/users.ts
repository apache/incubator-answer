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
