import useSWR from 'swr';

import request from '@answer/utils/request';

export const useCheckUserStatus = () => {
  const apiUrl = '/answer/api/v1/user/status';
  const hasToken = localStorage.getItem('token');
  const { data, error } = useSWR<{ status: string }, Error>(
    hasToken ? apiUrl : null,
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
  };
};
