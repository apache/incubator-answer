import React from 'react';
import { useTranslation } from 'react-i18next';

import { ModifyEmail, ModifyPassword, MyLogins } from './components';

const Index = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.account',
  });
  return (
    <>
      <h3 className="mb-4">{t('heading')}</h3>
      <ModifyEmail />
      <ModifyPassword />
      <MyLogins />
    </>
  );
};

export default React.memo(Index);
