import useSWR from 'swr';
import qs from 'qs';

import request from '@answer/utils/request';
import type * as Type from '@answer/common/interface';

export const putReport = (params) => {
  return request.instance.put('/answer/admin/api/report', params);
};

export const useFlagSearch = (params: Type.AdminFlagsReq) => {
  const apiUrl = `/answer/admin/api/reports/page?${qs.stringify(params)}`;
  const { data, error, mutate } = useSWR<Type.ListResult, Error>(
    [apiUrl],
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
    mutate,
  };
};
