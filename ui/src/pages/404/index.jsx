import { useEffect } from 'react';
import { Container, Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import './index.scss';

const Index = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_404' });
  useEffect(() => {
    // auto height of container
    const pageWrap = document.querySelector('.page-wrap');
    pageWrap.style.display = 'contents';

    return () => {
      pageWrap.style.display = 'block';
    };
  }, []);
  return (
    <Container
      className="d-flex flex-column justify-content-center align-items-center"
      style={{ flex: 1 }}>
      <div
        className="mb-4 text-secondary"
        style={{ fontSize: '120px', lineHeight: 1.2 }}>
        (=‘x‘=)
      </div>
      <h4 className="text-center">{t('http_error')}</h4>
      <div className="text-center mb-3 fs-5">{t('desc')}</div>
      <div className="text-center">
        <Button as={Link} to="/" variant="link">
          {t('back_home')}
        </Button>
      </div>
    </Container>
  );
};

export default Index;
