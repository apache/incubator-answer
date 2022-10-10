import { Container, Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

const Index = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_50X' });
  return (
    <Container className="d-flex flex-column justify-content-center align-items-center page-wrap">
      <div
        className="mb-4 text-secondary"
        style={{ fontSize: '120px', lineHeight: 1.2 }}>
        (=T^T=)
      </div>
      <div className="text-center mb-3">{t('description')}</div>
      <div className="text-center">
        <Button as={Link} to="/" variant="link">
          {t('back_home')}
        </Button>
      </div>
    </Container>
  );
};

export default Index;
