export const DEFAULT_LANG = 'en_US';
export const CURRENT_LANG_STORAGE_KEY = '_a_lang__';
export const LOGGED_USER_STORAGE_KEY = '_a_lui_';
export const LOGGED_TOKEN_STORAGE_KEY = '_a_ltk_';
export const REDIRECT_PATH_STORAGE_KEY = '_a_rp_';
export const CAPTCHA_CODE_STORAGE_KEY = '_a_captcha_';

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
    child: [{ name: 'general' }, { name: 'interface' }, { name: 'smtp' }],
  },
];
