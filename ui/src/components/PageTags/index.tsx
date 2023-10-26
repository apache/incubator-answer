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

import { FC, useEffect, useLayoutEffect } from 'react';
import { Helmet } from 'react-helmet-async';

import { brandingStore, pageTagStore, siteInfoStore } from '@/stores';
import { getCurrentLang } from '@/utils/localize';

const doInsertCustomCSS = !document.querySelector('link[href*="custom.css"]');

const Index: FC = () => {
  const { favicon, square_icon } = brandingStore((state) => state.branding);
  const { pageTitle, keywords, description } = pageTagStore(
    (state) => state.items,
  );
  const appVersion = siteInfoStore((_) => _.version);
  const hashVersion = siteInfoStore((_) => _.revision);
  const setAppGenerator = () => {
    if (!appVersion) {
      return;
    }
    const generatorMetaNode = document.querySelector('meta[name="generator"]');
    if (generatorMetaNode) {
      generatorMetaNode.setAttribute(
        'content',
        `Answer ${appVersion} - https://github.com/answerdev/answer version ${hashVersion}`,
      );
    }
  };
  const setDocTitle = () => {
    try {
      if (pageTitle) {
        document.title = pageTitle;
      }
      // eslint-disable-next-line no-empty
    } catch (ex) {}
  };
  const currentLang = getCurrentLang();
  const setDocLang = () => {
    if (currentLang) {
      document.documentElement.setAttribute(
        'lang',
        currentLang.replace('_', '-'),
      );
    }
  };

  useEffect(() => {
    setDocLang();
  }, [currentLang]);
  useEffect(() => {
    setAppGenerator();
  }, [appVersion]);
  useLayoutEffect(() => {
    setDocTitle();
  }, [pageTitle]);
  return (
    <Helmet>
      <link
        rel="icon"
        type="image/png"
        href={favicon || square_icon || '/favicon.ico'}
      />
      <link rel="icon" type="image/png" sizes="192x192" href={square_icon} />
      <link rel="apple-touch-icon" type="image/png" href={square_icon} />
      <title>{pageTitle}</title>
      {keywords && <meta name="keywords" content={keywords} />}
      {description && <meta name="description" content={description} />}
      {doInsertCustomCSS && (
        <link rel="stylesheet" href={`${process.env.PUBLIC_URL}/custom.css`} />
      )}
    </Helmet>
  );
};

export default Index;
