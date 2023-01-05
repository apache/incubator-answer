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

import classnames from 'classnames';

import { floppyNavigation } from '@/utils';
import {
  loggedUserInfoStore,
  siteInfoStore,
  brandingStore,
  loginSettingStore,
  themeSettingStore,
} from '@/stores';
import { logout, useQueryNotificationStatus } from '@/services';
import { DEFAULT_SITE_NAME } from '@/common/constants';

import NavItems from './components/NavItems';

import './index.scss';

const Header: FC = () => {
  const navigate = useNavigate();
  const { user, clear: clearUserStore } = loggedUserInfoStore();
  const { t } = useTranslation();
  const [urlSearch] = useSearchParams();
  const q = urlSearch.get('q');
  const [searchStr, setSearch] = useState('');
  const siteInfo = siteInfoStore((state) => state.siteInfo);
  const brandingInfo = brandingStore((state) => state.branding);
  const loginSetting = loginSettingStore((state) => state.login);
  const { data: redDot } = useQueryNotificationStatus();
  const location = useLocation();
  const handleInput = (val) => {
    setSearch(val);
  };
  const handleSearch = (evt) => {
    evt.preventDefault();
    if (!searchStr) {
      return;
    }
    const searchUrl = `/search?q=${encodeURIComponent(searchStr)}`;
    navigate(searchUrl);
  };

  const handleLogout = async () => {
    await logout();
    clearUserStore();
  };
  const onLoginClick = (evt) => {
    evt.preventDefault();
    floppyNavigation.navigateToLogin((loginPath) => {
      navigate(loginPath, { replace: true });
    });
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

  let navbarStyle = 'theme-colored';
  const { theme, theme_config } = themeSettingStore((_) => _);
  if (theme_config?.[theme]?.navbar_style) {
    navbarStyle = `theme-${theme_config[theme].navbar_style}`;
  }

  return (
    <Navbar
      variant={navbarStyle === 'theme-colored' ? 'dark' : ''}
      expand="lg"
      className={classnames('sticky-top', navbarStyle)}
      id="header">
      <Container className="d-flex align-items-center">
        <Navbar.Toggle
          aria-controls="navBarContent"
          className="answer-navBar me-2"
          id="navBarToggle"
        />

        <div className="d-flex justify-content-between align-items-center nav-grow flex-nowrap">
          <Navbar.Brand to="/" as={Link} className="lh-1 me-0 me-sm-3">
            {brandingInfo.logo ? (
              <>
                <img
                  className="d-none d-lg-block logo rounded-1 me-0"
                  src={brandingInfo.logo}
                  alt=""
                />

                <img
                  className="lg-none logo rounded-1 me-0"
                  src={brandingInfo.mobile_logo || brandingInfo.logo}
                  alt=""
                />
              </>
            ) : (
              <span>{siteInfo.name || DEFAULT_SITE_NAME}</span>
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
                  className={classnames('me-2', {
                    'link-light': navbarStyle === 'theme-colored',
                    'link-primary': navbarStyle !== 'theme-colored',
                  })}
                  onClick={onLoginClick}
                  href="/users/login">
                  {t('btns.login')}
                </Button>
                {loginSetting.allow_new_registrations && (
                  <Button
                    variant={
                      navbarStyle === 'theme-colored' ? 'light' : 'primary'
                    }
                    href="/users/register">
                    {t('btns.signup')}
                  </Button>
                )}
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
              <NavLink className="nav-link" to="/users">
                {t('header.nav.user')}
              </NavLink>
            </Nav>
          </Col>
          <hr className="hr lg-none mt-2" />

          <Col lg={4} className="d-flex justify-content-center">
            <Form
              action="/search"
              className="w-75 px-0 px-lg-2"
              onSubmit={handleSearch}>
              <FormControl
                placeholder={t('header.search.placeholder')}
                className="placeholder-search"
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
                    className={classnames('text-capitalize text-nowrap btn', {
                      'btn-light': navbarStyle !== 'theme-light',
                      'btn-primary': navbarStyle === 'theme-light',
                    })}>
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
                  className={classnames('me-2', {
                    'link-light': navbarStyle === 'theme-colored',
                    'link-primary': navbarStyle !== 'theme-colored',
                  })}
                  onClick={onLoginClick}
                  href="/users/login">
                  {t('btns.login')}
                </Button>
                {loginSetting.allow_new_registrations && (
                  <Button
                    variant={
                      navbarStyle === 'theme-colored' ? 'light' : 'primary'
                    }
                    href="/users/register">
                    {t('btns.signup')}
                  </Button>
                )}
              </>
            )}
          </Col>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
};

export default memo(Header);
