import { memo } from 'react';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';

const Index = ({ httpCode = '' }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_error' });

  usePageTags({
    title: t(`http_${httpCode}`, { keyPrefix: 'page_title' }),
  });
  return (
    <>
      <div
        className="mb-4 text-secondary"
        style={{ fontSize: '120px', lineHeight: 1.2 }}>
        (=‘x‘=)
      </div>
      <h4 className="text-center">{t('http_error', { code: httpCode })}</h4>
      <div className="text-center mb-3 fs-5">{t(`desc_${httpCode}`)}</div>
      <div className="text-center">
        <Link to="/" className="btn btn-link">
          {t('back_home')}
        </Link>
      </div>
    </>
  );
};

export default memo(Index);
