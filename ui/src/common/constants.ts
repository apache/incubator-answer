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
// timezones
export const TIMEZONES = [
  {
    label: 'UTC-12',
    value: 'UTC-12',
  },
  {
    label: 'UTC-11:30',
    value: 'UTC-11.5',
  },
  {
    label: 'UTC-11',
    value: 'UTC-11',
  },
  {
    label: 'UTC-10:30',
    value: 'UTC-10.5',
  },
  {
    label: 'UTC-10',
    value: 'UTC-10',
  },
  {
    label: 'UTC-9:30',
    value: 'UTC-9.5',
  },
  {
    label: 'UTC-9',
    value: 'UTC-9',
  },
  {
    label: 'UTC-8:30',
    value: 'UTC-8.5',
  },
  {
    label: 'UTC-8',
    value: 'UTC-8',
  },
  {
    label: 'UTC-7:30',
    value: 'UTC-7.5',
  },
  {
    label: 'UTC-7',
    value: 'UTC-7',
  },
  {
    label: 'UTC-6:30',
    value: 'UTC-6.5',
  },
  {
    label: 'UTC-6',
    value: 'UTC-6',
  },
  {
    label: 'UTC-5:30',
    value: 'UTC-5.5',
  },
  {
    label: 'UTC-5',
    value: 'UTC-5',
  },
  {
    label: 'UTC-4:30',
    value: 'UTC-4.5',
  },
  {
    label: 'UTC-4',
    value: 'UTC-4',
  },
  {
    label: 'UTC-3:30',
    value: 'UTC-3.5',
  },
  {
    label: 'UTC-3',
    value: 'UTC-3',
  },
  {
    label: 'UTC-2:30',
    value: 'UTC-2.5',
  },
  {
    label: 'UTC-2',
    value: 'UTC-2',
  },
  {
    label: 'UTC-1:30',
    value: 'UTC-1.5',
  },
  {
    label: 'UTC-1',
    value: 'UTC-1',
  },
  {
    label: 'UTC-0:30',
    value: 'UTC-0.5',
  },
  {
    label: 'UTC+0',
    value: 'UTC+0',
  },
  {
    label: 'UTC+0:30',
    value: 'UTC+0.5',
  },
  {
    label: 'UTC+1',
    value: 'UTC+1',
  },
  {
    label: 'UTC+1:30',
    value: 'UTC+1.5',
  },
  {
    label: 'UTC+2',
    value: 'UTC+2',
  },
  {
    label: 'UTC+2:30',
    value: 'UTC+2.5',
  },
  {
    label: 'UTC+3',
    value: 'UTC+3',
  },
  {
    label: 'UTC+3:30',

    value: 'UTC+3.5',
  },
  {
    label: 'UTC+4',
    value: 'UTC+4',
  },
  {
    label: 'UTC+4:30',
    value: 'UTC+4.5',
  },
  {
    label: 'UTC+5',
    value: 'UTC+5',
  },
  {
    label: 'UTC+5:30',
    value: 'UTC+5.5',
  },
  {
    label: 'UTC+5:45',
    value: 'UTC+5.75',
  },
  {
    label: 'UTC+6',
    value: 'UTC+6',
  },
  {
    label: 'UTC+6:30',

    value: 'UTC+6.5',
  },
  {
    label: 'UTC+7',
    value: 'UTC+7',
  },
  {
    label: 'UTC+7:30',
    value: 'UTC+7.5',
  },
  {
    label: 'UTC+8',
    value: 'UTC+8',
  },
  {
    label: 'UTC+8:30',
    value: 'UTC+8.5',
  },
  {
    label: 'UTC+8:45',
    value: 'UTC+8.75',
  },
  {
    label: 'UTC+9',
    value: 'UTC+9',
  },
  {
    label: 'UTC+9:30',
    value: 'UTC+9.5',
  },
  {
    label: 'UTC+10',
    value: 'UTC+10',
  },
  {
    label: 'UTC+10:30',
    value: 'UTC+10.5',
  },
  {
    label: 'UTC+11',
    value: 'UTC+11',
  },
  {
    label: 'UTC+11:30',
    value: 'UTC+11.5',
  },
  {
    label: 'UTC+12',
    value: 'UTC+12',
  },
  {
    label: 'UTC+12:45',
    value: 'UTC+12.75',
  },
  {
    label: 'UTC+13',
    value: 'UTC+13',
  },
  {
    label: 'UTC+13:45',
    value: 'UTC+13.75',
  },
  {
    label: 'UTC+14',
    value: 'UTC+14',
  },
];
export const DEFAULT_TIMEZONE = 'UTC+0';
