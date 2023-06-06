import { FC } from 'react';
import { Row, Col, Nav } from 'react-bootstrap';
import { Outlet, NavLink } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'nav_menus' });
  return (
    <Row className="pt-4 mb-5">
      <Col xxl={12}>
        <Nav
          className="mb-4 flex-nowrap"
          variant="pills"
          style={{ overflow: 'auto' }}>
          <NavLink to="/tos" key="tos" className="nav-link">
            {t('tos')}
          </NavLink>
          <NavLink to="/privacy" key="privacy" className="nav-link">
            {t('privacy')}
          </NavLink>
        </Nav>
      </Col>
      <Col xxl={12}>
        <Outlet />
      </Col>
    </Row>
  );
};

export default Index;
