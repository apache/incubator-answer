import useSWR from 'swr';
import qs from 'qs';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const useSearch = (params?: Type.SearchParams) => {
  const apiUrl = '/answer/api/v1/search';
  const queryParams = qs.stringify(params, { skipNulls: true });
  const { data, error, mutate } = useSWR<Type.SearchRes, Error>(
    params?.q ? `${apiUrl}?${queryParams}` : null,
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
    mutate,
  };
};
