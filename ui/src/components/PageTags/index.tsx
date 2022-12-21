import { FC } from 'react';
import { Helmet } from 'react-helmet-async';

import { brandingStore, pageTagStore } from '@/stores';

const doInsertCustomCSS = !document.querySelector('link[href*="custom.css"]');

const Index: FC = () => {
  const { favicon, square_icon } = brandingStore((state) => state.branding);
  const { pageTitle, keywords, description } = pageTagStore(
    (state) => state.items,
  );
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
