import request from '@/utils/request';

export const checkConfigFileExists = () => {
  return request.post('/installation/config-file/check');
};

export const dbCheck = (params) => {
  return request.post('/installation/db/check', params);
};

export const installInit = (params) => {
  return request.post('/installation/init', params);
};

export const installBaseInfo = (params) => {
  return request.post('/installation/base-info', params);
};

export const getInstallLangOptions = () => {
  return request.get('/installation/language/options');
};
