import useSWR from 'swr';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const useQueryContributeUsers = () => {
  const apiUrl = '/answer/api/v1/user/ranking';
  return useSWR<{
    users_with_the_most_reputation: Type.User[];
    users_with_the_most_vote: Type.User[];
    staffs: Type.User[];
  }>(apiUrl, request.instance.get);
};
