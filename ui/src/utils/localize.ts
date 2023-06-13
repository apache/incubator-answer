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
import {
  getAdminLanguageOptions,
  getLanguageConfig,
  getLanguageOptions,
} from '@/services';

import Storage from './storage';

/**
 * localize kit for i18n
 */
export const loadLanguageOptions = async (forAdmin = false) => {
  const languageOptions = forAdmin
    ? await getAdminLanguageOptions()
    : await getLanguageOptions();
  if (process.env.NODE_ENV === 'development') {
    const { default: optConf } = await import('@i18n/i18n.yaml');
    optConf?.language_options.forEach((opt) => {
      if (!languageOptions.find((_) => opt.value === _.value)) {
        languageOptions.push(opt);
      }
    });
  }
  return languageOptions;
};

const pullLanguageConf = (res) => {
  return getLanguageConfig().then((langConf) => {
    if (langConf && langConf.ui) {
      res.resources = langConf.ui;
      Storage.set(LANG_RESOURCE_STORAGE_KEY, res);
    }
  });
};
const addI18nResource = async (langName) => {
  const res = { lng: langName, resources: undefined };
  const storageResource = Storage.get(LANG_RESOURCE_STORAGE_KEY);
  if (process.env.NODE_ENV === 'development') {
    try {
      const { default: resConf } = await import(`@i18n/${langName}.yaml`);
      res.resources = resConf.ui;
    } catch (ex) {
      console.log('ex: ', ex);
    }
  } else if (storageResource && storageResource.lng === res.lng) {
    res.resources = storageResource.resources;
    pullLanguageConf(res);
  } else {
    await pullLanguageConf(res);
  }

  if (res.resources) {
    i18next.addResourceBundle(
      res.lng,
      'translation',
      res.resources,
      true,
      true,
    );
  }
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

/**
 * localize for Day.js
 */
dayjs.extend(utc);
dayjs.extend(timezone);
const localeDayjs = (langName) => {
  langName = langName.replace('_', '-').toLowerCase();
  dayjs.locale(langName);
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
