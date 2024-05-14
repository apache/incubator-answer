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
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import {
  getTransNs,
  getTransKeyPrefix,
  PluginInfo,
} from '@/utils/pluginKit/utils';
import { SvgIcon } from '@/components';
import { userCenterStore } from '@/stores';
import './i18n';

import info from './info.yaml';

interface Props {
  className?: classnames.Argument;
}

const pluginInfo: PluginInfo = {
  slug_name: info.slug_name,
  type: info.type,
};

const Index: FC<Props> = ({ className }) => {
  const { t } = useTranslation(getTransNs(), {
    keyPrefix: getTransKeyPrefix(pluginInfo),
  });
  const ucAgent = userCenterStore().agent;
  const ucLoginRedirect =
    ucAgent?.enabled && ucAgent?.agent_info?.login_redirect_url;

  if (ucLoginRedirect) {
    return (
      <Button
        className={classnames('w-100', className)}
        variant="outline-secondary"
        href={ucAgent?.agent_info.login_redirect_url}>
        <SvgIcon base64={ucAgent?.agent_info.icon} svgClassName="btnSvg me-2" />
        <span>
          {t('connect', { auth_name: ucAgent?.agent_info.display_name })}
        </span>
      </Button>
    );
  }
  return null;
};
export default {
  info: pluginInfo,
  component: memo(Index),
};
