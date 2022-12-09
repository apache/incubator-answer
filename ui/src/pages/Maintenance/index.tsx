import { Container } from 'react-bootstrap';
import { Helmet, HelmetProvider } from 'react-helmet-async';
import { useTranslation } from 'react-i18next';

const Index = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'page_maintenance',
  });
  return (
    <HelmetProvider>
      <Helmet>
        <title>{t('maintenance', { keyPrefix: 'page_title' })}</title>
      </Helmet>
      <div className="bg-f5">
        <Container
          className="d-flex flex-column justify-content-center align-items-center"
          style={{ minHeight: '100vh' }}>
          <div
            className="mb-4 text-secondary"
            style={{ fontSize: '120px', lineHeight: 1.2 }}>
            (=‘_‘=)
          </div>
          <div className="text-center mb-4">{t('desc')}</div>
        </Container>
      </div>
    </HelmetProvider>
  );
};

export default Index;
