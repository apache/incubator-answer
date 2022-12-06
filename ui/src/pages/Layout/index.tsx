import { FC, memo } from 'react';
import { Outlet } from 'react-router-dom';
import { Helmet, HelmetProvider } from 'react-helmet-async';

import { SWRConfig } from 'swr';

import {
  siteInfoStore,
  toastStore,
  brandingStore,
  headInfoStore,
} from '@/stores';
import { Header, Footer, Toast } from '@/components';

const Layout: FC = () => {
  const { msg: toastMsg, variant, clear: toastClear } = toastStore();
  const siteInfo = siteInfoStore((state) => state.siteInfo);
  const headInfo = headInfoStore((state) => state.headInfo);
  const { favicon, square_icon } = brandingStore((state) => state.branding);

  const closeToast = () => {
    toastClear();
  };
  const {
    title,
    keywords,
    description = siteInfo.description,
    ldJSON,
  } = headInfo;

  return (
    <HelmetProvider>
      <Helmet>
        <link
          rel="icon"
          type="image/png"
          href={favicon || square_icon || '/favicon.ico'}
        />
        <link rel="icon" type="image/png" sizes="192x192" href={square_icon} />
        <link rel="apple-touch-icon" type="image/png" href={square_icon} />
        {title && <title>{title}</title>}
        {keywords && <meta name="keywords" content={keywords} />}
        {description && <meta name="description" content={description} />}
        {ldJSON && <script type="application/ld+json">{ldJSON}</script>}
      </Helmet>
      <SWRConfig
        value={{
          revalidateOnFocus: false,
        }}>
        <Header />
        <div className="position-relative page-wrap">
          <Outlet />
        </div>
        <Toast msg={toastMsg} variant={variant} onClose={closeToast} />
        <Footer />
      </SWRConfig>
    </HelmetProvider>
  );
};

export default memo(Layout);
