import useSWR from 'swr';

import request from '@/utils/request';

export interface OauthConnectorItem {
  icon: string;
  name: string;
  link: string;
}

export const useGetStartUseOauthConnector = () => {
  const { data, error } = useSWR<OauthConnectorItem[]>(
    '/answer/api/v1/connector/info',
    request.instance.get,
  );

  return {
    data,
    error,
  };
};
