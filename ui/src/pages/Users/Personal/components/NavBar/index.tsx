import { FC, memo } from 'react';
import { Nav } from 'react-bootstrap';
import { NavLink } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

interface Props {
  slug: string;
  isSelf: boolean;
  tabName: string;
}

const list = [
  {
    path: '',
    name: 'overview',
  },
  {
    path: '/answers',
    name: 'answers',
  },
  {
    path: '/questions',
    name: 'questions',
  },
  {
    role: 'self', // Only visible to author
    path: '/bookmarks',
    name: 'bookmarks',
  },
  {
    path: '/reputation',
    name: 'reputation',
  },
  {
    path: '/comments',
    name: 'comments',
  },
  {
    role: 'self', // Only visible to author
    path: '/votes',
    name: 'votes',
  },
];
const Index: FC<Props> = ({ slug, tabName = 'overview', isSelf }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });
  return (
    <Nav
      className="pt-2 mb-4 flex-nowrap"
      variant="pills"
      style={{ overflow: 'auto' }}>
      {list.map((item) => {
        if (item.role && !isSelf) {
          return null;
        }
        if (item.path) {
          return (
            <NavLink
              to={`/users/${slug}${item.path}`}
              key={item.name}
              className="nav-link">
              {t(item.name)}
            </NavLink>
          );
        }
        return (
          <NavLink
            key={item.name}
            to={`/users/${slug}`}
            className={({ isActive }) =>
              isActive && tabName === 'overview'
                ? 'nav-link active'
                : 'nav-link'
            }>
            {t(item.name)}
          </NavLink>
        );
      })}
    </Nav>
  );
};

export default memo(Index);
