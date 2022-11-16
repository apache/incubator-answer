import useSWR from 'swr';

import * as Type from '@/common/interface';
import request from '@/utils/request';

export const useDashBoard = () => {
  const apiUrl = `/answer/admin/api/dashboard`;
  const { data, error } = useSWR<Type.AdminDashboard, Error>(
    [apiUrl],
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
  };
};
