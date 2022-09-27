import React from 'react';
import { useTranslation } from 'react-i18next';

import ModifyEmail from './components/ModifyEmail';
import ModifyPassword from './components/ModifyPass';

const Index = () => {
  const { t } = useTranslation();
  return (
    <>
      <h4 className="mb-3">{t('settings.nav.account')}</h4>
      <ModifyEmail />
      <ModifyPassword />
    </>
  );
};

export default React.memo(Index);
