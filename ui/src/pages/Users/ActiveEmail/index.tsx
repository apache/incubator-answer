import { FC, memo, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { useSearchParams, useNavigate } from 'react-router-dom';

import { usePageTags } from '@/hooks';
import { loggedUserInfoStore } from '@/stores';
import { activateAccount } from '@/services';

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
  usePageTags({
    title: t('account_activation'),
  });
  return null;
};

export default memo(Index);
