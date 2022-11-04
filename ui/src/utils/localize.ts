import dayjs from 'dayjs';
import i18next from 'i18next';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';

import { interfaceStore, loggedUserInfoStore } from '@/stores';
import { DEFAULT_LANG, CURRENT_LANG_STORAGE_KEY } from '@/common/constants';
import { Storage } from '@/utils';

dayjs.extend(utc);
dayjs.extend(timezone);
const localDayjs = (langName) => {
  langName = langName.replace('_', '-').toLowerCase();
  dayjs.locale(langName);
};

export const getCurrentLang = () => {
  const loggedUser = loggedUserInfoStore.getState().user;
  const adminInterface = interfaceStore.getState().interface;
  const storageLang = Storage.get(CURRENT_LANG_STORAGE_KEY);
  let currentLang = loggedUser.language;
  // `default` mean use language value from admin interface
  if (/default/i.test(currentLang) && adminInterface.language) {
    currentLang = adminInterface.language;
  }
  currentLang ||= storageLang || DEFAULT_LANG;
  return currentLang;
};

export const setupAppLanguage = () => {
  const lang = getCurrentLang();
  localDayjs(lang);
  i18next.changeLanguage(lang);
};

export const setupAppTimeZone = () => {
  const adminInterface = interfaceStore.getState().interface;
  if (adminInterface.time_zone) {
    dayjs.tz.setDefault(adminInterface.time_zone);
  }
};
