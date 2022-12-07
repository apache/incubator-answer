import { Container, Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

const Index = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_404' });
  return (
    <Container className="d-flex flex-column justify-content-center align-items-center page-wrap">
      <div
        className="mb-4 text-secondary"
        style={{ fontSize: '120px', lineHeight: 1.2 }}>
        (=‘x‘=)
      </div>
      <div className="text-center mb-4">{t('desc')}</div>
      <div className="text-center">
        <Button as={Link} to="/" variant="link">
          {t('back_home')}
        </Button>
      </div>
    </Container>
  );
};

export default Index;
