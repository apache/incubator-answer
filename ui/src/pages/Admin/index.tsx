import { FC } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Outlet, useLocation } from 'react-router-dom';

import { usePageTags } from '@/hooks';
import { AccordionNav } from '@/components';
import { ADMIN_NAV_MENUS } from '@/common/constants';

import './index.scss';

const formPaths = [
  'general',
  'smtp',
  'interface',
  'branding',
  'legal',
  'write',
  'themes',
  'css-html',
];

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const { pathname } = useLocation();
  usePageTags({
    title: t('admin'),
  });
  return (
    <>
      <div className="bg-light py-2">
        <Container className="py-1">
          <h6 className="mb-0 fw-bold lh-base">
            {t('title', { keyPrefix: 'admin.admin_header' })}
          </h6>
        </Container>
      </div>
      <Container className="admin-container">
        <Row>
          <Col lg={2}>
            <AccordionNav menus={ADMIN_NAV_MENUS} path="/admin/" />
          </Col>
          <Col lg={formPaths.find((v) => pathname.includes(v)) ? 6 : 10}>
            <Outlet />
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default Index;
