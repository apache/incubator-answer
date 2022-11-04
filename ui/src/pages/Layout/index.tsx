import { FC, memo } from 'react';
import { Outlet } from 'react-router-dom';
import { Helmet, HelmetProvider } from 'react-helmet-async';

import { SWRConfig } from 'swr';

import { siteInfoStore, toastStore } from '@/stores';
import { Header, Footer, Toast } from '@/components';

const Layout: FC = () => {
  const { msg: toastMsg, variant, clear: toastClear } = toastStore();
  const { siteInfo } = siteInfoStore.getState();
  const closeToast = () => {
    toastClear();
  };

  return (
    <HelmetProvider>
      <Helmet>
        {siteInfo && <meta name="description" content={siteInfo.description} />}
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
