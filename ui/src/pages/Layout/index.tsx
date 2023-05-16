import { FC, memo, useEffect } from 'react';
import { Outlet, useLocation } from 'react-router-dom';
import { HelmetProvider } from 'react-helmet-async';

import { SWRConfig } from 'swr';

import { toastStore, loginToContinueStore, errorCodeStore } from '@/stores';
import {
  Header,
  Footer,
  Toast,
  Customize,
  CustomizeTheme,
  PageTags,
  HttpErrorContent,
} from '@/components';
import { LoginToContinueModal } from '@/components/Modal';

const Layout: FC = () => {
  const location = useLocation();
  const { msg: toastMsg, variant, clear: toastClear } = toastStore();
  const closeToast = () => {
    toastClear();
  };
  const { code: httpStatusCode, reset: httpStatusReset } = errorCodeStore();
  const { show: showLoginToContinueModal } = loginToContinueStore();

  useEffect(() => {
    httpStatusReset();
  }, [location]);
  return (
    <HelmetProvider>
      <PageTags />
      <CustomizeTheme />
      <SWRConfig
        value={{
          revalidateOnFocus: false,
        }}>
        <Header />
        {/* eslint-disable-next-line jsx-a11y/click-events-have-key-events */}
        <div className="position-relative page-wrap">
          {httpStatusCode ? (
            <HttpErrorContent httpCode={httpStatusCode} />
          ) : (
            <Outlet />
          )}
        </div>
        <Toast msg={toastMsg} variant={variant} onClose={closeToast} />
        <Footer />
        <Customize />
        <LoginToContinueModal visible={showLoginToContinueModal} />
      </SWRConfig>
    </HelmetProvider>
  );
};

export default memo(Layout);
