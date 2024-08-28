/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import { initReactI18next } from 'react-i18next';

import i18next from 'i18next';
import en_US from '@i18n/en_US.yaml';
import es_ES from '@i18n/es_ES.yaml';
import pt_BR from '@i18n/pt_BR.yaml';
import pt_PT from '@i18n/pt_PT.yaml';
import de_DE from '@i18n/de_DE.yaml';
import fr_FR from '@i18n/fr_FR.yaml';
import ja_JP from '@i18n/ja_JP.yaml';
import it_IT from '@i18n/it_IT.yaml';
import ru_RU from '@i18n/ru_RU.yaml';
import zh_CN from '@i18n/zh_CN.yaml';
import zh_TW from '@i18n/zh_TW.yaml';
import ko_KR from '@i18n/ko_KR.yaml';
import vi_VN from '@i18n/vi_VN.yaml';
import sk_SK from '@i18n/sk_SK.yaml';
import fa_IR from '@i18n/fa_IR.yaml';

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
  es_ES: {
    translation: es_ES.ui,
  },
  pt_BR: {
    translation: pt_BR.ui,
  },
  pt_PT: {
    translation: pt_PT.ui,
  },
  de_DE: {
    translation: de_DE.ui,
  },
  fr_FR: {
    translation: fr_FR.ui,
  },
  ja_JP: {
    translation: ja_JP.ui,
  },
  it_IT: {
    translation: it_IT.ui,
  },
  ru_RU: {
    translation: ru_RU.ui,
  },
  zh_CN: {
    translation: zh_CN.ui,
  },
  zh_TW: {
    translation: zh_TW.ui,
  },
  ko_KR: {
    translation: ko_KR.ui,
  },
  vi_VN: {
    translation: vi_VN.ui,
  },
  sk_SK: {
    translation: sk_SK.ui,
  },
  fa_IR: {
    translation: fa_IR.ui,
  },
};

const storageLang = Storage.get(LANG_RESOURCE_STORAGE_KEY);
if (
  storageLang &&
  storageLang.resources &&
  storageLang.lng &&
  storageLang.lng !== 'en_US'
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
