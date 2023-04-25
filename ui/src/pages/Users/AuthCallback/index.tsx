import { FC, memo, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { useSearchParams, useNavigate } from 'react-router-dom';

import { usePageTags } from '@/hooks';
import { guard } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  useEffect(() => {
    const token = searchParams.get('access_token');
    guard.handleLoginWithToken(token, navigate);
  }, []);
  usePageTags({
    title: t('oauth_callback'),
  });
  return null;
};

export default memo(Index);
