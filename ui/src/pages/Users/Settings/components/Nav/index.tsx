import React, { FC } from 'react';
import { Nav } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { NavLink, useMatch } from 'react-router-dom';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'settings.nav' });
  const settingMatch = useMatch('/users/settings/:setting');
  return (
    <Nav variant="pills" className="flex-column">
      <NavLink
        className={({ isActive }) =>
          isActive || !settingMatch ? 'nav-link active' : 'nav-link'
        }
        to="/users/settings/profile">
        {t('profile')}
      </NavLink>
      <NavLink className="nav-link" to="/users/settings/notify">
        {t('notification')}
      </NavLink>
      <NavLink className="nav-link" to="/users/settings/account">
        {t('account')}
      </NavLink>
      <NavLink className="nav-link" to="/users/settings/interface">
        {t('interface')}
      </NavLink>
    </Nav>
  );
};

export default React.memo(Index);
