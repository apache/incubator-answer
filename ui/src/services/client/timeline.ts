import qs from 'qs';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

// export const useTimelineData = (params: Type.TimelineReq) => {
//   const apiUrl = '/answer/api/v1/activity/timeline';
//   const { data, error, mutate } = useSWR<Type.TimelineRes, Error>(
//     `${apiUrl}?${qs.stringify(params, { skipNulls: true })}`,
//     request.instance.get,
//   );
//   return {
//     data,
//     isLoading: !data && !error,
//     error,
//     mutate,
//   };
// };

export const getTimelineData = (params: Type.TimelineReq) => {
  return request.get<Type.TimelineRes>(
    `/answer/api/v1/activity/timeline?${qs.stringify(params, {
      skipNulls: true,
    })}`,
  );
};

export const getTimelineDetail = (params: {
  new_revision_id: string;
  old_revision_id: string;
}) => {
  return request.get(
    `/answer/api/v1/activity/timeline/detail?${qs.stringify(params, {
      skipNulls: true,
    })}`,
  );
};
