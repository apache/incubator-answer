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

import { memo, FC } from 'react';
import { useTranslation } from 'react-i18next';

import {
  getTransNs,
  getTransKeyPrefix,
  PluginInfo,
} from '@/utils/pluginKit/utils';
import { SvgIcon } from '@/components';

import info from './info.yaml';
import { useGetSearchPLuginInfo } from './services';
import './i18n';

const pluginInfo: PluginInfo = {
  slug_name: info.slug_name,
  type: info.type,
};

const Index: FC = () => {
  const { t } = useTranslation(getTransNs(), {
    keyPrefix: getTransKeyPrefix(pluginInfo),
  });

  const { data } = useGetSearchPLuginInfo();
  if (!data?.icon) return null;

  return (
    <a
      className="d-flex align-items-center"
      href={data?.link}
      target="_blank"
      rel="noopener noreferrer">
      <span className="small text-secondary me-2">{t('search_by')}</span>
      <SvgIcon base64={data?.icon} svgClassName="max-width-200" />
    </a>
  );
};

export default {
  info: pluginInfo,
  component: memo(Index),
};
