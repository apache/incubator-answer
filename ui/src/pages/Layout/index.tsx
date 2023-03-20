import { FC, memo, useEffect } from 'react';
import { Outlet, useLocation } from 'react-router-dom';
import { HelmetProvider } from 'react-helmet-async';

import { SWRConfig } from 'swr';

import { toastStore, loginToContinueStore, errorCode } from '@/stores';
import {
  Header,
  Footer,
  Toast,
  Customize,
  CustomizeTheme,
  PageTags,
} from '@/components';
import { LoginToContinueModal } from '@/components/Modal';
import { useImgViewer } from '@/hooks';
import Component404 from '@/pages/404';
import Component50X from '@/pages/50X';

const Layout: FC = () => {
  const location = useLocation();
  const { msg: toastMsg, variant, clear: toastClear } = toastStore();
  const closeToast = () => {
    toastClear();
  };
  const { code: httpStatusCode, reset: httpStatusReset } = errorCode();

  const imgViewer = useImgViewer();
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
        <div
          className="position-relative page-wrap"
          onClick={imgViewer.checkClickForImgView}>
          {httpStatusCode === '404' ? (
            <Component404 />
          ) : httpStatusCode === '50X' ? (
            <Component50X />
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
