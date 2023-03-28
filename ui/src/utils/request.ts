import axios, { AxiosResponse } from 'axios';
import type { AxiosInstance, AxiosRequestConfig, AxiosError } from 'axios';

import { Modal } from '@/components';
import { loggedUserInfoStore, toastStore, errorCode } from '@/stores';
import { LOGGED_TOKEN_STORAGE_KEY, IGNORE_PATH_LIST } from '@/common/constants';
import { RouteAlias } from '@/router/alias';
import { getCurrentLang } from '@/utils/localize';

import Storage from './storage';
import { floppyNavigation } from './floppyNavigation';
import { isIgnoredPath } from './guard';

const baseConfig = {
  timeout: 10000,
  withCredentials: true,
};

interface APIconfig extends AxiosRequestConfig {
  allow404: boolean;
}

class Request {
  instance: AxiosInstance;

  constructor(config: AxiosRequestConfig) {
    this.instance = axios.create(config);
    this.instance.interceptors.request.use(
      (requestConfig: AxiosRequestConfig) => {
        const token = Storage.get(LOGGED_TOKEN_STORAGE_KEY) || '';
        const lang = getCurrentLang();
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
        const { status, data: respData } = error.response || {};
        const { data = {}, msg = '', reason = '' } = respData || {};

        console.log('response error:', error);

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

            if (data.err_type === 'alert') {
              return Promise.reject({
                msg,
                ...data,
              });
            }

            if (data.err_type === 'modal') {
              // modal error message
              Modal.confirm({
                content: msg,
              });
            }

            return Promise.reject(false);
          }

          if (data instanceof Array && data.length > 0) {
            // handle form error
            return Promise.reject({
              code: status,
              msg,
              reason,
              isError: true,
              list: data,
            });
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
          errorCode.getState().reset();
          loggedUserInfoStore.getState().clear();
          floppyNavigation.navigateToLogin();
          return Promise.reject(false);
        }
        if (status === 403) {
          // Permission interception
          errorCode.getState().reset();
          if (data?.type === 'url_expired') {
            // url expired
            floppyNavigation.navigate(RouteAlias.activationFailed, {
              handler: 'replace',
            });
            return Promise.reject(false);
          }
          if (data?.type === 'inactive') {
            // inactivated
            floppyNavigation.navigate(RouteAlias.activation);
            return Promise.reject(false);
          }

          if (data?.type === 'suspended') {
            floppyNavigation.navigate(RouteAlias.suspended, {
              handler: 'replace',
            });
            return Promise.reject(false);
          }

          if (msg) {
            toastStore.getState().show({
              msg,
              variant: 'danger',
            });
          }
          return Promise.reject(false);
        }

        if (status === 404 && error.config?.allow404) {
          if (isIgnoredPath(IGNORE_PATH_LIST)) {
            return Promise.reject(false);
          }
          errorCode.getState().update('404');
          return Promise.reject(false);
        }
        if (status >= 500) {
          if (isIgnoredPath(IGNORE_PATH_LIST)) {
            return Promise.reject(false);
          }
          errorCode.getState().update('50X');
          console.error(
            `Request failed with status code ${status}, ${msg || ''}`,
          );
        }
        return Promise.reject(false);
      },
    );
  }

  public request(config: AxiosRequestConfig): Promise<AxiosResponse> {
    return this.instance.request(config);
  }

  public get<T = any>(url: string, config?: APIconfig): Promise<T> {
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
