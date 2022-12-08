import React from 'react';
import { useTranslation } from 'react-i18next';

import ModifyEmail from './components/ModifyEmail';
import ModifyPassword from './components/ModifyPass';

const Index = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.account',
  });
  return (
    <>
      <h3 className="mb-4">{t('heading')}</h3>
      <ModifyEmail />
      <ModifyPassword />
    </>
  );
};

export default React.memo(Index);
