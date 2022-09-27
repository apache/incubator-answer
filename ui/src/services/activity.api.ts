import useSWR from 'swr';

import request from '@answer/utils/request';

import type * as Type from './types';

export const useFollow = (params?: Type.FollowParams) => {
  const apiUrl = '/answer/api/v1/follow';
  const { data, error, mutate } = useSWR<any, Error>(
    params ? [apiUrl, params] : null,
    request.instance.post,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
    mutate,
  };
};
