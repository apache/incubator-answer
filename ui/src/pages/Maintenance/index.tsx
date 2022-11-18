import { Container } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { PageTitle } from '@/components';

const Index = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'page_maintenance',
  });
  return (
    <div className="bg-f5">
      <Container
        className="d-flex flex-column justify-content-center align-items-center"
        style={{ minHeight: '100vh' }}>
        <PageTitle title={t('maintenance', { keyPrefix: 'page_title' })} />
        <div
          className="mb-4 text-secondary"
          style={{ fontSize: '120px', lineHeight: 1.2 }}>
          (=‘_‘=)
        </div>
        <div className="text-center mb-4">{t('description')}</div>
      </Container>
    </div>
  );
};

export default Index;
