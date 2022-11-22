import { FC, memo, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { useSearchParams, useNavigate } from 'react-router-dom';

import { loggedUserInfoStore } from '@/stores';
import { activateAccount } from '@/services';
import { PageTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const [searchParams] = useSearchParams();
  const updateUser = loggedUserInfoStore((state) => state.update);
  const navigate = useNavigate();
  useEffect(() => {
    const code = searchParams.get('code');

    if (code) {
      activateAccount(encodeURIComponent(code)).then((res) => {
        updateUser(res);
        setTimeout(() => {
          navigate('/users/account-activation/success', { replace: true });
        }, 0);
      });
    } else {
      navigate('/', { replace: true });
    }
  }, []);
  return <PageTitle title={t('account_activation')} />;
};

export default memo(Index);
