import axios, { AxiosResponse } from 'axios';
import type { AxiosInstance, AxiosRequestConfig, AxiosError } from 'axios';

import { Modal } from '@answer/components';
import { userInfoStore, toastStore } from '@answer/stores';

import Storage from './storage';

const API = {
  development: '',
  production: '',
  test: '',
};

const baseApiUrl = process.env.REACT_APP_API_URL || API[process.env.NODE_ENV];

const baseConfig = {
  baseUrl: baseApiUrl,
  timeout: 10000,
  withCredentials: true,
};

class Request {
  instance: AxiosInstance;

  constructor(config: AxiosRequestConfig) {
    this.instance = axios.create(config);

    this.instance.interceptors.request.use(
      (requestConfig: AxiosRequestConfig) => {
        const token = Storage.get('token') || '';
        // default lang en_US
        const lang = Storage.get('LANG') || 'en_US';
        requestConfig.headers = {
          Authorization: token,
          'Accept-Language': lang,
        };
        return requestConfig;
      },
      (err: AxiosError) => {
        console.error('request interceptors error:', err);
      },
    );

    this.instance.interceptors.response.use(
      (res: AxiosResponse) => {
        const { status, data } = res.data;

        if (status === 204) {
          // no content
          return true;
        }

        return data;
      },
      (error) => {
        const { status, data, msg } = error.response;
        const { data: realData, msg: realMsg = '' } = data;
        if (status === 400) {
          // show error message
          if (realData instanceof Object && realData.err_type) {
            if (realData.err_type === 'toast') {
              // toast error message
              toastStore.getState().show({
                msg: realMsg,
                variant: 'danger',
              });
            }

            if (realData.type === 'modal') {
              // modal error message
              Modal.confirm({
                content: realMsg,
              });
            }

            return Promise.reject(false);
          }

          if (
            realData instanceof Object &&
            Object.keys(realData).length > 0 &&
            realData.key
          ) {
            // handle form error
            return Promise.reject({ ...realData, isError: true });
          }

          if (!realData || Object.keys(realData).length <= 0) {
            // default error msg will show modal
            Modal.confirm({
              content: realMsg,
            });
            return Promise.reject(false);
          }
        }

        if (status === 401) {
          // clear userinfo;
          Storage.remove('token');
          userInfoStore.getState().clear();
          // need login
          const { pathname } = window.location;
          if (pathname !== '/users/login' && pathname !== '/users/register') {
            Storage.set('ANSWER_PATH', window.location.pathname);
          }
          window.location.href = '/users/login';

          return Promise.reject(false);
        }

        if (status === 403) {
          // Permission interception

          if (realData?.type === 'inactive') {
            // inactivated
            window.location.href = '/users/login?status=inactive';
            return Promise.reject(false);
          }

          if (realData?.type === 'url_expired') {
            // url expired
            window.location.href = '/users/account-activation/failed';
            return Promise.reject(false);
          }

          if (realData?.type === 'suspended') {
            if (window.location.pathname !== '/users/account-suspended') {
              window.location.href = '/users/account-suspended';
            }

            return Promise.reject(false);
          }
        }

        toastStore.getState().show({
          msg: `statusCode: ${status}; ${msg || ''}`,
          variant: 'danger',
        });
        return Promise.reject(false);
      },
    );
  }

  public request(config: AxiosRequestConfig): Promise<AxiosResponse> {
    return this.instance.request(config);
  }

  public get<T = any>(url: string, config?: AxiosRequestConfig): Promise<T> {
    return this.instance.get(url, config);
  }

  public post<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig,
  ): Promise<T> {
    return this.instance.post(url, data, config);
  }

  public put<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig,
  ): Promise<T> {
    return this.instance.put(url, data, config);
  }

  public delete<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig,
  ): Promise<T> {
    return this.instance.delete(url, {
      data,
      ...config,
    });
  }
}

// export const Request;

export default new Request(baseConfig);
