import useSWR from 'swr';
import qs from 'qs';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const useTimelineData = (params: Type.TimelineReq) => {
  const apiUrl = '/answer/api/v1/activity/timeline';
  const { data, error, mutate } = useSWR<Type.TimelineRes, Error>(
    `${apiUrl}?${qs.stringify(params, { skipNulls: true })}`,
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
    mutate,
  };
};
