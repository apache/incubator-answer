import qs from 'qs';
import useSWR from 'swr';

import request from '@answer/utils/request';

import type * as Type from './types';

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

export const useQuestionSearch = (params: Type.AdminContentsReq) => {
  const apiUrl = `/answer/admin/api/question/page?${qs.stringify(params)}`;
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

export const changeQuestionStatus = (
  question_id: string,
  status: Type.AdminQuestionStatus,
) => {
  return request.put('/answer/admin/api/question/status', {
    question_id,
    status,
  });
};
