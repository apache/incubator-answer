import { FC, memo } from 'react';
import { Nav, Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { NavLink, useNavigate } from 'react-router-dom';

import type * as Type from '@/common/interface';
import { Avatar, Icon } from '@/components';
import { floppyNavigation } from '@/utils';
import { userCenterStore } from '@/stores';

interface Props {
  redDot: Type.NotificationStatus | undefined;
  userInfo;
  logOut: () => void;
}

const Index: FC<Props> = ({ redDot, userInfo, logOut }) => {
  const { t } = useTranslation();
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const navigate = useNavigate();
  const { agent: ucAgent } = userCenterStore();
  const handleLinkClick = (evt) => {
    if (floppyNavigation.shouldProcessLinkClick(evt)) {
      evt.preventDefault();
      const href = evt.currentTarget.getAttribute('href');
      navigate(href);
    }
  };
  return (
    <>
      <Nav className="flex-row">
        <Nav.Link
          as={NavLink}
          to="/users/notifications/inbox"
          className="icon-link d-flex align-items-center justify-content-center p-0 me-3 position-relative">
          <Icon name="bell-fill" className="fs-4" />
          {(redDot?.inbox || 0) > 0 && <div className="unread-dot bg-danger" />}
        </Nav.Link>

        <Nav.Link
          as={NavLink}
          to="/users/notifications/achievement"
          className="icon-link d-flex align-items-center justify-content-center p-0 me-3 position-relative">
          <Icon name="trophy-fill" className="fs-4" />
          {(redDot?.achievement || 0) > 0 && (
            <div className="unread-dot bg-danger" />
          )}
        </Nav.Link>
      </Nav>

      <Dropdown align="end">
        <Dropdown.Toggle
          variant="success"
          id="dropdown-basic"
          as="a"
          className="no-toggle pointer">
          <Avatar size="36px" avatar={userInfo?.avatar} searchStr="s=96" />
        </Dropdown.Toggle>

        <Dropdown.Menu>
          <Dropdown.Item
            href={`/users/${userInfo.username}`}
            onClick={handleLinkClick}>
            {t('header.nav.profile')}
          </Dropdown.Item>
          <Dropdown.Item
            href={`/users/${userInfo.username}/bookmarks`}
            onClick={handleLinkClick}>
            {t('header.nav.bookmark')}
          </Dropdown.Item>
          <Dropdown.Item
            href="/users/settings/profile"
            onClick={handleLinkClick}>
            {t('header.nav.setting')}
          </Dropdown.Item>
          <Dropdown.Divider />
          <Dropdown.Item onClick={logOut}>
            {t('header.nav.logout')}
          </Dropdown.Item>
        </Dropdown.Menu>
      </Dropdown>
      {/* Dropdown for user center agent info */}
      {ucAgent?.enabled &&
      (ucAgent?.agent_info?.url ||
        ucAgent?.agent_info?.control_center?.length) ? (
        <Dropdown align="end">
          <Dropdown.Toggle
            variant="success"
            id="dropdown-uca"
            as="span"
            className="no-toggle">
            <Nav>
              <Icon
                name="grid-3x3-gap-fill"
                className="nav-link pointer p-0 fs-4 ms-3"
              />
            </Nav>
          </Dropdown.Toggle>

          <Dropdown.Menu>
            {ucAgent.agent_info.url ? (
              <Dropdown.Item href={ucAgent.agent_info.url}>
                {ucAgent.agent_info.name}
              </Dropdown.Item>
            ) : null}
            {ucAgent.agent_info.url &&
            ucAgent.agent_info.control_center?.length ? (
              <Dropdown.Divider />
            ) : null}
            {ucAgent.agent_info.control_center?.map((ctrl) => {
              return (
                <Dropdown.Item key={ctrl.name} href={ctrl.url}>
                  {ctrl.label}
                </Dropdown.Item>
              );
            })}
          </Dropdown.Menu>
        </Dropdown>
      ) : null}
    </>
  );
};

export default memo(Index);
