import { initReactI18next } from 'react-i18next';

import i18next from 'i18next';
import en_US from '@i18n/en_US.yaml';
import zh_CN from '@i18n/zh_CN.yaml';

import { DEFAULT_LANG, LANG_RESOURCE_STORAGE_KEY } from '@/common/constants';
import Storage from '@/utils/storage';

/**
 * Prevent i18n from re-initialising when the page is refreshed and switching to `fallbackLng`.
 */
const initLng = i18next.resolvedLanguage || DEFAULT_LANG;
const initResources = {
  en_US: {
    translation: en_US.ui,
  },
  zh_CN: {
    translation: zh_CN.ui,
  },
};

const storageLang = Storage.get(LANG_RESOURCE_STORAGE_KEY);
if (
  storageLang &&
  storageLang.resources &&
  storageLang.lng &&
  storageLang.lng !== 'en_US' &&
  storageLang.lng !== 'zh_CN'
) {
  initResources[storageLang.lng] = {
    translation: storageLang.resources,
  };
}

i18next
  //  pass the i18n instance to react-i18next.
  .use(initReactI18next)
  .init({
    resources: initResources,
    lng: initLng,
    fallbackLng: DEFAULT_LANG,
    interpolation: {
      escapeValue: false,
    },
    react: {
      transSupportBasicHtmlNodes: true,
      // allow <br/> and simple html elements in translations
      transKeepBasicHtmlNodesFor: ['br', 'strong', 'i'],
    },
  });

export default i18next;
