import dayjs from 'dayjs';
import i18next from 'i18next';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';

import { interfaceStore, loggedUserInfoStore } from '@/stores';
import {
  CURRENT_LANG_STORAGE_KEY,
  DEFAULT_LANG,
  LANG_RESOURCE_STORAGE_KEY,
} from '@/common/constants';
import { Storage } from '@/utils';
import {
  getAdminLanguageOptions,
  getLanguageConfig,
  getLanguageOptions,
} from '@/services';

export const loadLanguageOptions = async (forAdmin = false) => {
  const languageOptions = forAdmin
    ? await getAdminLanguageOptions()
    : await getLanguageOptions();
  if (process.env.NODE_ENV === 'development') {
    const { default: optConf } = await import('@/i18n/locales/i18n.yaml');
    optConf?.language_options.forEach((opt) => {
      if (!languageOptions.find((_) => opt.label === _.label)) {
        languageOptions.push(opt);
      }
    });
  }
  return languageOptions;
};

const addI18nResource = async (langName) => {
  const res = { lng: langName, resources: undefined };
  if (process.env.NODE_ENV === 'development') {
    try {
      const { default: resConf } = await import(
        `@/i18n/locales/${langName}.yaml`
      );
      res.resources = resConf.ui;
    } catch (ex) {
      console.log('ex: ', ex);
    }
  } else {
    const storageResource = Storage.get(LANG_RESOURCE_STORAGE_KEY);
    if (storageResource?.lng === res.lng) {
      res.resources = storageResource.resources;
    } else {
      const langConf = await getLanguageConfig();
      if (langConf) {
        res.resources = langConf;
      }
    }
  }
  if (res.resources) {
    i18next.addResourceBundle(
      res.lng,
      'translation',
      res.resources,
      true,
      true,
    );
    Storage.set(LANG_RESOURCE_STORAGE_KEY, res);
  }
};

dayjs.extend(utc);
dayjs.extend(timezone);
const localeDayjs = (langName) => {
  langName = langName.replace('_', '-').toLowerCase();
  dayjs.locale(langName);
};

export const getCurrentLang = () => {
  const loggedUser = loggedUserInfoStore.getState().user;
  const adminInterface = interfaceStore.getState().interface;
  const fallbackLang = Storage.get(CURRENT_LANG_STORAGE_KEY) || DEFAULT_LANG;
  let currentLang = loggedUser.language;
  // `default` mean use language value from admin interface
  if (/default/i.test(currentLang)) {
    currentLang = adminInterface.language;
  }
  currentLang ||= fallbackLang;
  return currentLang;
};

export const setupAppLanguage = async () => {
  const lang = getCurrentLang();
  if (!i18next.getDataByLanguage(lang)) {
    await addI18nResource(lang);
  }
  localeDayjs(lang);
  i18next.changeLanguage(lang);
};

export const setupAppTimeZone = () => {
  const adminInterface = interfaceStore.getState().interface;
  if (adminInterface.time_zone) {
    dayjs.tz.setDefault(adminInterface.time_zone);
  }
};
