import { FC, memo, useEffect } from 'react';
import { Outlet, useLocation } from 'react-router-dom';
import { HelmetProvider } from 'react-helmet-async';

import { SWRConfig } from 'swr';

import { toastStore, loginToContinueStore, notFoundStore } from '@/stores';
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

const Layout: FC = () => {
  const location = useLocation();
  const { msg: toastMsg, variant, clear: toastClear } = toastStore();
  const closeToast = () => {
    toastClear();
  };
  const { visible: show404, hide: notFoundHide } = notFoundStore();

  const imgViewer = useImgViewer();
  const { show: showLoginToContinueModal } = loginToContinueStore();

  useEffect(() => {
    notFoundHide();
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
          {show404 ? <Component404 /> : <Outlet />}
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
