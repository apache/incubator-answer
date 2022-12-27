import { FC, memo } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'unsubscribe' });
  usePageTags({
    title: t('page_title'),
  });
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col lg={6}>
          <h3 className="text-center mt-3 mb-5">{t('success_title')}</h3>
          <p className="text-center">{t('success_desc')}</p>
          <div className="text-center">
            <Link to="/users/settings/notify">{t('link')}</Link>
          </div>
        </Col>
      </Row>
    </Container>
  );
};

export default memo(Index);
