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

import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { userCenterStore } from '@/stores';
import { getUcSettings, UcSettingAgent } from '@/services';

import { ModifyEmail, ModifyPassword, MyLogins } from './components';

const Index = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.account',
  });
  const { agent: ucAgent } = userCenterStore();
  const [accountAgent, setAccountAgent] = useState<UcSettingAgent>();

  const initData = () => {
    if (ucAgent?.enabled) {
      getUcSettings().then((resp) => {
        setAccountAgent(resp.account_setting_agent);
      });
    }
  };
  useEffect(() => {
    initData();
  }, []);
  return (
    <>
      <h3 className="mb-4">{t('heading')}</h3>
      {accountAgent?.enabled && accountAgent?.redirect_url ? (
        <a href={accountAgent.redirect_url}>
          {t('goto_modify', { keyPrefix: 'settings' })}
        </a>
      ) : null}
      {!ucAgent?.enabled || accountAgent?.enabled === false ? (
        <>
          <ModifyEmail />
          <ModifyPassword />
          <MyLogins />
        </>
      ) : null}
    </>
  );
};

export default React.memo(Index);
