import request from '@/utils/request';

type loginConf = {
  key: string;
  redirect_url: string;
};

type loginResult = {
  is_login: boolean;
  token: string;
};

export const getLoginConf = () => {
  const apiUrl = `/answer/api/v1/wecom/login/url`;
  return request.get<loginConf>(apiUrl);
};

export const checkLoginResult = (key: loginConf['key']) => {
  const apiUrl = `/answer/api/v1/wecom/login/check?key=${key}`;
  return request.get<loginResult>(apiUrl);
};
