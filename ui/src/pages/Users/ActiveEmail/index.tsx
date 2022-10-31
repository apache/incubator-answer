import { FC, memo, useEffect } from 'react';
import { useTranslation } from 'react-i18next';

import { loggedUserInfoStore } from '@answer/stores';
import { getQueryString } from '@answer/utils';

import { activateAccount } from '@/services';
import { PageTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const updateUser = loggedUserInfoStore((state) => state.update);
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
