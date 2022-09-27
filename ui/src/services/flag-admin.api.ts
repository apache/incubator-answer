import useSWR from 'swr';

import request from '@answer/utils/request';

import type * as Type from './types';

export const putReport = (params) => {
  return request.instance.put('/answer/admin/api/report', params);
};

export const useFlagSearch = (params: Type.AdminFlagsReq) => {
  const apiUrl = `/answer/admin/api/reports/${params.status}/${params.object_type}?page=${params.page}&page_size=${params.page_size}`;
  const { data, error } = useSWR<Type.ListResult, Error>(
    [apiUrl],
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
  };
};
