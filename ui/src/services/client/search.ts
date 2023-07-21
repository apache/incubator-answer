import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const getSearchResult = (params?: Type.SearchParams) => {
  const apiUrl = '/answer/api/v1/search';

  return request.get<Type.SearchRes>(apiUrl, {
    params,
  });
};
