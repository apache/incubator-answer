import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { loggedUserInfoStore } from '@/stores';

const Suspended = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'suspended' });
  const userInfo = loggedUserInfoStore((state) => state.user);
  usePageTags({
    title: t('account_suspended', { keyPrefix: 'page_title' }),
  });
  if (userInfo.status !== 'forbidden') {
    window.location.replace('/');
    return null;
  }
  return (
    <div className="d-flex flex-column align-items-center mt-5 pt-3">
      <h3 className="mb-5">{t('title')}</h3>
      <p className="text-center">
        {t('forever')}
        <br />
        {t('end')}
      </p>
    </div>
  );
};

export default Suspended;
