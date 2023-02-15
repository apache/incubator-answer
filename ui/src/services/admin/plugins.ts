import qs from 'qs';
import useSWR from 'swr';

import type * as Types from '@/common/interface';
import request from '@/utils/request';

export const useQueryPlugins = (params) => {
  const apiUrl = `/answer/admin/api/plugins?${qs.stringify(params)}`;
  const { data, error, mutate } = useSWR<any[], Error>(
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

export const updatePluginStatus = (params) => {
  return request.put('/answer/admin/api/plugin/status', params);
};

export const useQueryPluginConfig = (params) => {
  const apiUrl = `/answer/admin/api/plugin/config?${qs.stringify(params)}`;
  const { data, error, mutate } = useSWR<Types.PluginConfig, Error>(
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

export const updatePluginConfig = (params) => {
  return request.put('/answer/admin/api/plugin/config', params);
};
