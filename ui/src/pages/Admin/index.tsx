import { FC } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Outlet } from 'react-router-dom';

import { AccordionNav, PageTitle } from '@answer/components';

import { ADMIN_NAV_MENUS } from '@answer/common/constants';

import './index.scss';

const Dashboard: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  return (
    <>
      <PageTitle title={t('admin')} />
      <Container className="admin-container">
        <Row>
          <Col lg={2}>
            <AccordionNav menus={ADMIN_NAV_MENUS} />
          </Col>
          <Col lg={10}>
            <Outlet />
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default Dashboard;
