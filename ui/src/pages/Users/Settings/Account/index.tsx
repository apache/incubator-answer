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
