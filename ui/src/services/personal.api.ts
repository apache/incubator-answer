import useSWR from 'swr';
import qs from 'qs';

import request from '@answer/utils/request';

import type * as Type from './types';

export const usePersonalInfoByName = (username: string) => {
  const apiUrl = '/answer/api/v1/user/info/username';
  const { data, error, mutate } = useSWR<Type.UserInfoRes, Error>(
    username ? `${apiUrl}?username=${username}` : null,
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
    mutate,
  };
};

interface ListReq {
  username?: string;
  page: number;
  page_size: number;
  order?: string;
}

interface ListRes {
  count: number;
  list: any[];
}

export const usePersonalTop = (username: string, tabName: string) => {
  const apiUrl = '/answer/api/v1/question/user/top?username=';
  const { data, error } = useSWR<{ answer: any[]; question: any[] }, Error>(
    tabName === 'overview' ? `${apiUrl}${username}` : null,
    request.instance.get,
  );
  return {
    data,
    isLoading: !data && !error,
    error,
  };
};

export const usePersonalListByTabName = (params: ListReq, tabName: string) => {
  let apiUrl = '';
  if (tabName === 'answers') {
    apiUrl = '/answer/api/v1/answer/user';
  }
  if (tabName === 'questions') {
    apiUrl = '/answer/api/v1/question/user';
  }
  if (tabName === 'bookmarks') {
    delete params.order;
    apiUrl = '/answer/api/v1/collection/user';
  }
  if (tabName === 'comments') {
    delete params.order;
    apiUrl = '/answer/api/v1/comment/personal/page';
  }
  if (tabName === 'reputation') {
    delete params.order;
    apiUrl = '/answer/api/v1/rank/personal/page';
  }
  if (tabName === 'votes') {
    delete params.username;
    apiUrl = '/answer/api/v1/votes/personal';
  }

  const queryParams = qs.stringify(params, { skipNulls: true });
  const { data, error, mutate } = useSWR<ListRes, Error>(
    tabName !== 'overview' ? `${apiUrl}?${queryParams}` : null,
    request.instance.get,
  );
  return {
    data: {
      [tabName]: data,
    },
    isLoading: !data && !error,
    error,
    mutate,
  };
};
