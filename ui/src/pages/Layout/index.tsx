import { FC, memo } from 'react';
import { Outlet } from 'react-router-dom';
import { HelmetProvider } from 'react-helmet-async';

import { SWRConfig } from 'swr';

import { toastStore } from '@/stores';
import {
  Header,
  Footer,
  Toast,
  Customize,
  CustomizeTheme,
  PageTags,
} from '@/components';

const Layout: FC = () => {
  const { msg: toastMsg, variant, clear: toastClear } = toastStore();
  const closeToast = () => {
    toastClear();
  };

  return (
    <HelmetProvider>
      <PageTags />
      <CustomizeTheme />
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
