import { FC, memo, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { useSearchParams, useNavigate } from 'react-router-dom';

import { usePageTags, useLoginRedirect } from '@/hooks';
import { loggedUserInfoStore } from '@/stores';
import { getLoggedUserInfo } from '@/services';
import Storage from '@/utils/storage';
import { LOGGED_TOKEN_STORAGE_KEY } from '@/common/constants';
import { guard } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const [searchParams] = useSearchParams();
  const { loginRedirect } = useLoginRedirect();
  const updateUser = loggedUserInfoStore((state) => state.update);
  const navigate = useNavigate();

  useEffect(() => {
    const token = searchParams.get('access_token');

    if (token) {
      Storage.set(LOGGED_TOKEN_STORAGE_KEY, token);
      getLoggedUserInfo().then((res) => {
        updateUser(res);
        const userStat = guard.deriveLoginState();
        if (userStat.isNotActivated) {
          // inactive
          navigate('/users/login?status=inactive', { replace: true });
        } else {
          setTimeout(() => {
            loginRedirect();
          }, 0);
        }
      });
    } else {
      navigate('/', { replace: true });
    }
  }, []);
  usePageTags({
    title: t('oauth_callback'),
  });
  return null;
};

export default memo(Index);
