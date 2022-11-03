import { initReactI18next } from 'react-i18next';

import i18next from 'i18next';
import Backend from 'i18next-http-backend';

import { DEFAULT_LANG } from '@/common/constants';

import en from './locales/en.json';
import zh from './locales/zh_CN.json';

i18next
  // load translation using http
  .use(Backend)
  //  pass the i18n instance to react-i18next.
  .use(initReactI18next)
  .init({
    resources: {
      en_US: {
        translation: en,
      },
      zh_CN: {
        translation: zh,
      },
    },
    // debug: process.env.NODE_ENV === 'development',
    fallbackLng: process.env.REACT_APP_LANG || DEFAULT_LANG,
    interpolation: {
      escapeValue: false,
    },
    react: {
      transSupportBasicHtmlNodes: true, // allow <br/> and simple html elements in translations
      transKeepBasicHtmlNodesFor: ['br', 'strong', 'i'],
    },
    // backend: {
    //   loadPath: (lngs, namespace) => {
    //     console.log(lngs, namespace);
    //     return 'https://cdn.jsdelivr.net/npm/echarts@4.8.0/map/js/china.js';
    //   },
    // },
  });

export default i18next;
