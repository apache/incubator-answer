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
