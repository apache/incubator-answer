import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { userCenterStore } from '@/stores';
import { getUcSettings } from '@/services';

import { ModifyEmail, ModifyPassword, MyLogins } from './components';

const Index = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.account',
  });
  const { agent: ucAgent } = userCenterStore();
  const [accountAgent, setAccountAgent] = useState('');

  const initData = () => {
    if (ucAgent?.enabled) {
      getUcSettings().then((resp) => {
        if (
          resp.account_setting_agent?.enabled &&
          resp.account_setting_agent?.redirect_url
        ) {
          setAccountAgent(resp.account_setting_agent.redirect_url);
        }
      });
    }
  };
  useEffect(() => {
    initData();
  }, []);
  return (
    <>
      <h3 className="mb-4">{t('heading')}</h3>
      {accountAgent ? (
        <a href={accountAgent}>{t('goto_modify', { keyPrefix: 'settings' })}</a>
      ) : null}
      {!ucAgent?.enabled ? (
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
