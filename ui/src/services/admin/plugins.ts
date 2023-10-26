/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import qs from 'qs';
import useSWR from 'swr';

import request from '@/utils/request';
import { UIOptions, UIWidget } from '@/components/SchemaForm';

export interface PluginOption {
  label: string;
  value: string;
}

export interface PluginItem {
  name: string;
  type: UIWidget;
  title: string;
  description: string;
  ui_options?: UIOptions;
  options?: PluginOption[];
  value?: string;
  required?: boolean;
}

export interface PluginConfig {
  name: string;
  slug_name: string;
  config_fields: PluginItem[];
}

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
  const { data, error, mutate } = useSWR<PluginConfig, Error>(
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
