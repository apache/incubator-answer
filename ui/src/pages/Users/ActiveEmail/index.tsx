import { FC, memo, useEffect } from 'react';
import { useTranslation } from 'react-i18next';

import { activateAccount } from '@answer/api';
import { userInfoStore } from '@answer/stores';
import { getQueryString } from '@answer/utils';

import { PageTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const updateUser = userInfoStore((state) => state.update);
  useEffect(() => {
    const code = getQueryString('code');

    if (code) {
      activateAccount(encodeURIComponent(code)).then((res) => {
        updateUser(res);
        setTimeout(() => {
          window.location.replace('/users/account-activation/success');
        }, 0);
      });
    }
  }, []);
  return <PageTitle title={t('account_activation')} />;
};

export default memo(Index);
