import { FC } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Outlet, useLocation } from 'react-router-dom';

import { AccordionNav, AdminHeader, PageTitle } from '@/components';
import { ADMIN_NAV_MENUS } from '@/common/constants';

import './index.scss';

const formPaths = ['general', 'smtp', 'interface', 'branding'];

const Dashboard: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const { pathname } = useLocation();
  return (
    <>
      <PageTitle title={t('admin')} />
      <AdminHeader />
      <Container className="admin-container">
        <Row>
          <Col lg={2}>
            <AccordionNav menus={ADMIN_NAV_MENUS} />
          </Col>
          <Col lg={formPaths.find((v) => pathname.includes(v)) ? 6 : 10}>
            <Outlet />
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default Dashboard;
