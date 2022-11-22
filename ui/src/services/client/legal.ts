import useSWR from 'swr';

import request from '@/utils/request';
import type * as Type from '@/common/interface';

export const useLegalTos = () => {
  const apiUrl = '/answer/api/v1/siteinfo/legal?info_type=tos';
  const { data, error, mutate } = useSWR<Type.AdminSettingsLegal, Error>(
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

export const useLegalPrivacy = () => {
  const apiUrl = '/answer/api/v1/siteinfo/legal?info_type=privacy';
  const { data, error, mutate } = useSWR<Type.AdminSettingsLegal, Error>(
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
