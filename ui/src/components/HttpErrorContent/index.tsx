import { memo, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';

const Index = ({
  httpCode = '',
  title = '',
  errMsg = '',
  showErrorCode = true,
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_error' });
  useEffect(() => {
    // auto height of container
    const pageWrap = document.querySelector('.page-wrap') as HTMLElement;
    if (pageWrap) {
      pageWrap.style.display = 'contents';
    }

    return () => {
      if (pageWrap) {
        pageWrap.style.display = 'block';
      }
    };
  }, []);

  usePageTags({
    title: t(`http_${httpCode}`, { keyPrefix: 'page_title' }),
  });

  return (
    <div className="d-flex flex-column flex-shrink-1 flex-grow-1 justify-content-center align-items-center">
      <div
        className="mb-4 text-secondary"
        style={{ fontSize: '120px', lineHeight: 1.2 }}>
        (=‘x‘=)
      </div>
      {showErrorCode && (
        <h4 className="text-center">{t('http_error', { code: httpCode })}</h4>
      )}
      {title && <h4 className="text-center">{title}</h4>}
      <div className="text-center mb-3 fs-5">
        {errMsg || t(`desc_${httpCode}`)}
      </div>
      <div className="text-center">
        <Link to="/" className="btn btn-link">
          {t('back_home')}
        </Link>
      </div>
    </div>
  );
};

export default memo(Index);
