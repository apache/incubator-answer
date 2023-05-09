import { useTranslation } from 'react-i18next';
import { Button } from 'react-bootstrap';

import { siteInfoStore } from '@/stores';
import { usePageTags } from '@/hooks';

const Suspended = () => {
  const { contact_email = '' } = siteInfoStore((state) => state.siteInfo);
  const { t } = useTranslation('translation', { keyPrefix: 'suspended' });
  usePageTags({
    title: t('account_suspended', { keyPrefix: 'page_title' }),
  });

  return (
    <div className="d-flex flex-column align-items-center mt-5 pt-3">
      <h3 className="mb-5">{t('title')}</h3>
      <p className="text-center">
        {t('forever')}
        <br />
        {t('end')}
      </p>
      <Button href={`mailto:${contact_email}`} variant="link">
        {t('contact_us')}
      </Button>
    </div>
  );
};

export default Suspended;
