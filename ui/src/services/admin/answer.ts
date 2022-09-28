import useSWR from 'swr';
import qs from 'qs';

import request from '@answer/utils/request';
import type * as Type from '../types';

export const useAnswerSearch = (params: Type.AdminContentsReq) => {
  const apiUrl = `/answer/admin/api/answer/page?${qs.stringify(params)}`;
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

export const changeAnswerStatus = (
  answer_id: string,
  status: Type.AdminAnswerStatus,
) => {
  return request.put('/answer/admin/api/answer/status', {
    answer_id,
    status,
  });
};
