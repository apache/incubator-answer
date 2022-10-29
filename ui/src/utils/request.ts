import axios, { AxiosResponse } from 'axios';
import type { AxiosInstance, AxiosRequestConfig, AxiosError } from 'axios';

import { Modal } from '@answer/components';
import { loggedUserInfoStore, toastStore } from '@answer/stores';

import Storage from './storage';
import { floppyNavigation } from './floppyNavigation';

import {
  LOGGED_TOKEN_STORAGE_KEY,
  CURRENT_LANG_STORAGE_KEY,
  DEFAULT_LANG,
} from '@/common/constants';
import { RouteAlias } from '@/router/alias';

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
        const token = Storage.get(LOGGED_TOKEN_STORAGE_KEY) || '';
        // default lang en_US
        const lang = Storage.get(CURRENT_LANG_STORAGE_KEY) || DEFAULT_LANG;
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
        const { status, data: respData, msg: respMsg } = error.response;
        const { data, msg = '' } = respData;
        if (status === 400) {
          // show error message
          if (data instanceof Object && data.err_type) {
            if (data.err_type === 'toast') {
              // toast error message
              toastStore.getState().show({
                msg,
                variant: 'danger',
              });
            }

            if (data.type === 'modal') {
              // modal error message
              Modal.confirm({
                content: msg,
              });
            }

            return Promise.reject(false);
          }

          if (
            data instanceof Object &&
            Object.keys(data).length > 0 &&
            data.key
          ) {
            // handle form error
            return Promise.reject({ ...data, isError: true });
          }

          if (!data || Object.keys(data).length <= 0) {
            // default error msg will show modal
            Modal.confirm({
              content: msg,
            });
            return Promise.reject(false);
          }
        }
        // 401: Re-login required
        if (status === 401) {
          // clear userinfo
          loggedUserInfoStore.getState().clear();
          floppyNavigation.navigateToLogin();
          return Promise.reject(false);
        }
        if (status === 403) {
          // Permission interception
          if (data?.type === 'url_expired') {
            // url expired
            floppyNavigation.navigate(RouteAlias.activationFailed, () => {
              window.location.replace(RouteAlias.activationFailed);
            });
            return Promise.reject(false);
          }
          if (data?.type === 'inactive') {
            // inactivated
            floppyNavigation.navigate(RouteAlias.activation, () => {
              window.location.href = RouteAlias.activation;
            });
            return Promise.reject(false);
          }

          if (data?.type === 'suspended') {
            floppyNavigation.navigate(RouteAlias.suspended, () => {
              window.location.replace(RouteAlias.suspended);
            });
            return Promise.reject(false);
          }
        }

        toastStore.getState().show({
          msg: `statusCode: ${status}; ${respMsg || ''}`,
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

export default new Request(baseConfig);
