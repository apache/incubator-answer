import { FC, memo } from 'react';
import { Outlet } from 'react-router-dom';
import { Helmet, HelmetProvider } from 'react-helmet-async';

import { SWRConfig } from 'swr';

import { toastStore, brandingStore, pageTagStore } from '@/stores';
import { Header, Footer, Toast, Customize } from '@/components';

const Layout: FC = () => {
  const { msg: toastMsg, variant, clear: toastClear } = toastStore();
  const closeToast = () => {
    toastClear();
  };
  const { favicon, square_icon } = brandingStore((state) => state.branding);
  const { pageTitle, keywords, description } = pageTagStore(
    (state) => state.items,
  );

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
        <title>{pageTitle}</title>
        {keywords && <meta name="keywords" content={keywords} />}
        {description && <meta name="description" content={description} />}
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
        <Customize />
      </SWRConfig>
    </HelmetProvider>
  );
};

export default memo(Layout);
