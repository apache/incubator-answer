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
