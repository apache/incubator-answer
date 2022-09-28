export const LOGIN_NEED_BACK = [
  '/users/login',
  '/users/register',
  '/users/account-recovery',
  '/users/password-reset',
];

export const ADMIN_LIST_STATUS = {
  // normal;
  1: {
    variant: 'success',
    name: 'normal',
  },
  // closed;
  2: {
    variant: 'warning',
    name: 'closed',
  },
  // deleted
  10: {
    variant: 'danger',
    name: 'deleted',
  },
  normal: {
    variant: 'success',
    name: 'normal',
  },
  closed: {
    variant: 'warning',
    name: 'closed',
  },
  deleted: {
    variant: 'danger',
    name: 'deleted',
  },
};

export const ADMIN_NAV_MENUS = [
  {
    name: 'dashboard',
    children: [],
  },
  {
    name: 'contents',
    child: [{ name: 'questions' }, { name: 'answers' }],
  },
  {
    name: 'users',
  },
  {
    name: 'flags',
    // badgeContent: 5,
  },
  {
    name: 'settings',
    child: [{ name: 'general' }, { name: 'interface' }],
  },
];
