import { FC, memo, useState, useEffect } from 'react';
import {
  Navbar,
  Container,
  Nav,
  Form,
  FormControl,
  Button,
  Col,
} from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import {
  useSearchParams,
  NavLink,
  Link,
  useNavigate,
  useLocation,
} from 'react-router-dom';

import { loggedUserInfoStore, siteInfoStore, interfaceStore } from '@/stores';
import { logout, useQueryNotificationStatus } from '@/services';
import { RouteAlias } from '@/router/alias';

import NavItems from './components/NavItems';

import './index.scss';

const Header: FC = () => {
  const navigate = useNavigate();
  const { user, clear } = loggedUserInfoStore();
  const { t } = useTranslation();
  const [urlSearch] = useSearchParams();
  const q = urlSearch.get('q');
  const [searchStr, setSearch] = useState('');
  const siteInfo = siteInfoStore((state) => state.siteInfo);
  const { interface: interfaceInfo } = interfaceStore();
  const { data: redDot } = useQueryNotificationStatus();
  const location = useLocation();
  const handleInput = (val) => {
    setSearch(val);
  };

  const handleLogout = async () => {
    await logout();
    clear();
    navigate(RouteAlias.home);
  };

  useEffect(() => {
    if (q) {
      handleInput(q);
    }
  }, [q]);

  useEffect(() => {
    const collapse = document.querySelector('#navBarContent');
    if (collapse && collapse.classList.contains('show')) {
      const toggle = document.querySelector('#navBarToggle') as HTMLElement;
      if (toggle) {
        toggle?.click();
      }
    }
  }, [location.pathname]);

  return (
    <Navbar variant="dark" expand="lg" className="sticky-top" id="header">
      <Container className="d-flex align-items-center">
        <Navbar.Toggle
          aria-controls="navBarContent"
          className="answer-navBar me-2"
          id="navBarToggle"
        />

        <div className="d-flex justify-content-between align-items-center nav-grow flex-nowrap">
          <Navbar.Brand to="/" as={Link} className="lh-1 me-0 me-sm-3">
            {interfaceInfo.logo ? (
              <img
                className="logo rounded-1 me-0"
                src={interfaceInfo.logo}
                alt=""
              />
            ) : (
              <span>{siteInfo.name || 'Answer'}</span>
            )}
          </Navbar.Brand>

          {/* mobile nav */}
          <div className="d-flex lg-none align-items-center flex-lg-nowrap">
            {user?.username ? (
              <NavItems redDot={redDot} userInfo={user} logOut={handleLogout} />
            ) : (
              <>
                <Button
                  variant="link"
                  className="me-2 text-white"
                  href="/users/login">
                  {t('btns.login')}
                </Button>
                <Button variant="light" href="/users/register">
                  {t('btns.signup')}
                </Button>
              </>
            )}
          </div>
        </div>

        <Navbar.Collapse id="navBarContent" className="me-auto">
          <hr className="hr lg-none mb-2" style={{ marginTop: '12px' }} />
          <Col md={4}>
            <Nav>
              <NavLink className="nav-link" to="/questions">
                {t('header.nav.question')}
              </NavLink>
              <NavLink className="nav-link" to="/tags">
                {t('header.nav.tag')}
              </NavLink>
              <NavLink className="nav-link d-none" to="/users">
                {t('header.nav.user')}
              </NavLink>
            </Nav>
          </Col>
          <hr className="hr lg-none mt-2" />

          <Col lg={4} className="d-flex justify-content-center">
            <Form action="/search" className="w-75 px-0 px-lg-2">
              <FormControl
                placeholder={t('header.search.placeholder')}
                className="text-white placeholder-search"
                value={searchStr}
                name="q"
                onChange={(e) => handleInput(e.target.value)}
              />
            </Form>
          </Col>

          <Nav.Item className="lg-none mt-3 pb-1">
            <Link
              to="/questions/ask"
              className="text-capitalize text-nowrap btn btn-light">
              {t('btns.add_question')}
            </Link>
          </Nav.Item>
          {/* pc nav */}
          <Col
            lg={4}
            className="d-none d-lg-flex justify-content-start justify-content-sm-end">
            {user?.username ? (
              <Nav className="d-flex align-items-center flex-lg-nowrap">
                <Nav.Item className="me-3">
                  <Link
                    to="/questions/ask"
                    className="text-capitalize text-nowrap btn btn-light">
                    {t('btns.add_question')}
                  </Link>
                </Nav.Item>

                <NavItems
                  redDot={redDot}
                  userInfo={user}
                  logOut={handleLogout}
                />
              </Nav>
            ) : (
              <>
                <Button
                  variant="link"
                  className="me-2 text-white"
                  href="/users/login">
                  {t('btns.login')}
                </Button>
                <Button variant="light" href="/users/register">
                  {t('btns.signup')}
                </Button>
              </>
            )}
          </Col>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
};

export default memo(Header);
