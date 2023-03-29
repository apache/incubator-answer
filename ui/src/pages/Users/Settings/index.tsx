import { FC, memo } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Outlet } from 'react-router-dom';

import { usePageTags } from '@/hooks';

import Nav from './components/Nav';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.profile',
  });

  usePageTags({
    title: t('settings', { keyPrefix: 'page_title' }),
  });
  return (
    <Container className="mt-4 mb-5 pb-5">
      <Row className="justify-content-center">
        <Col xxl={10} md={12}>
          <h3 className="mb-4">{t('page_title', { keyPrefix: 'settings' })}</h3>
        </Col>
      </Row>

      <Row>
        <Col xxl={1} />
        <Col md={3} lg={2} className="mb-3">
          <Nav />
        </Col>
        <Col md={9} lg={6}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};

export default memo(Index);
