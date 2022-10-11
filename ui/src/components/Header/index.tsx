import { FC, memo, useState, useEffect } from 'react';
import {
  Navbar,
  Container,
  Nav,
  Form,
  FormControl,
  Button,
  Col,
  Dropdown,
} from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useSearchParams, NavLink, Link, useNavigate } from 'react-router-dom';

import { Avatar, Icon } from '@answer/components';
import { userInfoStore, siteInfoStore, interfaceStore } from '@answer/stores';
import { logout, useQueryNotificationStatus } from '@answer/api';
import Storage from '@answer/utils/storage';

import './index.scss';

const Header: FC = () => {
  const navigate = useNavigate();
  const { user, clear } = userInfoStore();
  const { t } = useTranslation();
  const [urlSearch] = useSearchParams();
  const q = urlSearch.get('q');
  const [searchStr, setSearch] = useState('');
  const siteInfo = siteInfoStore((state) => state.siteInfo);
  const { interface: interfaceInfo } = interfaceStore();
  const { data: redDot } = useQueryNotificationStatus();
  const handleInput = (val) => {
    setSearch(val);
  };

  const handleLogout = async () => {
    await logout();
    Storage.remove('token');
    clear();
    navigate('/');
  };

  useEffect(() => {
    if (q) {
      handleInput(q);
    }
  }, [q]);
  return (
    <Navbar variant="dark" expand="lg" className="sticky-top" id="header">
      <Container className="d-flex align-items-center">
        <Navbar.Brand href="/">
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
        <Navbar.Toggle aria-controls="navBarContent" />
        <Navbar.Collapse id="navBarContent" className="me-auto">
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

          <Col md={4} className="d-none d-sm-flex justify-content-center">
            <Form action="/search" className="w-75 px-2">
              <FormControl
                placeholder={t('header.search.placeholder')}
                className="text-white placeholder-search"
                value={searchStr}
                name="q"
                onChange={(e) => handleInput(e.target.value)}
              />
            </Form>
          </Col>

          <Col
            md={4}
            className="d-flex justify-content-start justify-content-sm-end">
            {user?.username ? (
              <Nav className="d-flex align-items-center flex-lg-nowrap">
                <Nav.Item className="me-2">
                  <Link
                    to="/questions/ask"
                    className="text-capitalize text-nowrap btn btn-light">
                    {t('btns.add_question')}
                  </Link>
                </Nav.Item>
                <Nav.Link
                  as={NavLink}
                  to="/users/notifications/inbox"
                  className="icon-link d-flex align-items-center justify-content-center p-0 me-2 position-relative">
                  <div className="text-white text-opacity-75">
                    <Icon name="bell-fill" className="fs-5" />
                  </div>
                  {(redDot?.inbox || 0) > 0 && (
                    <div className="unread-dot bg-danger" />
                  )}
                </Nav.Link>

                <Nav.Link
                  as={Link}
                  to="/users/notifications/achievement"
                  className="icon-link d-flex align-items-center justify-content-center p-0 me-2 position-relative">
                  <div className="text-white text-opacity-75">
                    <Icon name="trophy-fill" className="fs-5" />
                  </div>
                  {(redDot?.achievement || 0) > 0 && (
                    <div className="unread-dot bg-danger" />
                  )}
                </Nav.Link>

                <Dropdown align="end">
                  <Dropdown.Toggle
                    variant="success"
                    id="dropdown-basic"
                    as="a"
                    className="no-toggle pointer">
                    <Avatar size="36px" avatar={user?.avatar} />
                  </Dropdown.Toggle>

                  <Dropdown.Menu>
                    <Dropdown.Item href={`/users/${user.username}`}>
                      {t('header.nav.profile')}
                    </Dropdown.Item>
                    <Dropdown.Item href="/users/settings/profile">
                      {t('header.nav.setting')}
                    </Dropdown.Item>
                    {user?.is_admin ? (
                      <Dropdown.Item href="/admin">
                        {t('header.nav.admin')}
                      </Dropdown.Item>
                    ) : null}
                    <Dropdown.Divider />
                    <Dropdown.Item onClick={handleLogout}>
                      {t('header.nav.logout')}
                    </Dropdown.Item>
                  </Dropdown.Menu>
                </Dropdown>
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
