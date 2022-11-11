import { initReactI18next } from 'react-i18next';

import i18next from 'i18next';
import Backend from 'i18next-http-backend';
import en_US from '@i18n/en_US.yaml';
import zh_CN from '@i18n/zh_CN.yaml';

import { DEFAULT_LANG } from '@/common/constants';

i18next
  // load translation using http
  .use(Backend)
  //  pass the i18n instance to react-i18next.
  .use(initReactI18next)
  .init({
    resources: {
      en_US: {
        translation: en_US.ui,
      },
      zh_CN: {
        translation: zh_CN.ui,
      },
      vi_VN: {
        translation: vi,
      },
    },
    // debug: process.env.NODE_ENV === 'development',
    fallbackLng: process.env.REACT_APP_LANG || DEFAULT_LANG,
    interpolation: {
      escapeValue: false,
    },
    react: {
      transSupportBasicHtmlNodes: true,
      // allow <br/> and simple html elements in translations
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
