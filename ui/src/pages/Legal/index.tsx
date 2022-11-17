import { FC } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Outlet } from 'react-router-dom';

import { AccordionNav } from '@/components';
import { ADMIN_LEGAL_MENUS } from '@/common/constants';

import './index.scss';

const Index: FC = () => {
  return (
    <Container className="sub-container">
      <Row>
        <Col lg={2}>
          <AccordionNav menus={ADMIN_LEGAL_MENUS} />
        </Col>
        <Col lg={6}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};

export default Index;
