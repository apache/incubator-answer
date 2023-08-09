import useSWR from 'swr';

import request from '@/utils/request';

export interface AlgoliaRes {
  name: string;
  icon: string;
}

export const useGetAlgoliaInfo = () => {
  const { data, error } = useSWR<AlgoliaRes>(
    '/answer/api/v1/search/desc',
    request.instance.get,
  );

  return {
    data,
    error,
  };
};
