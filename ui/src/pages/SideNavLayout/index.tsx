import { FC, memo } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Outlet } from 'react-router-dom';

import { SideNav } from '@/components';

import '@/common/sideNavLayout.scss';

const Index: FC = () => {
  return (
    <Container>
      <Row>
        <SideNav />
        <Col xl={10} lg={9} md={12}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};

export default memo(Index);
