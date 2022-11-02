import { FC, useEffect, memo } from 'react';
import { useTranslation } from 'react-i18next';
import { Outlet } from 'react-router-dom';
import { Helmet, HelmetProvider } from 'react-helmet-async';

import { SWRConfig } from 'swr';

import { siteInfoStore, interfaceStore, toastStore } from '@/stores';
import { Header, AdminHeader, Footer, Toast } from '@/components';
import { useSiteSettings } from '@/services';
import Storage from '@/utils/storage';
import { CURRENT_LANG_STORAGE_KEY } from '@/common/constants';

let isMounted = false;
const Layout: FC = () => {
  const { siteInfo, update: siteStoreUpdate } = siteInfoStore();
  const { update: interfaceStoreUpdate } = interfaceStore();
  const { data: siteSettings } = useSiteSettings();
  const { msg: toastMsg, variant, clear: toastClear } = toastStore();
  const { i18n } = useTranslation();

  const closeToast = () => {
    toastClear();
  };

  useEffect(() => {
    if (siteSettings) {
      siteStoreUpdate(siteSettings.general);
      interfaceStoreUpdate(siteSettings.interface);
    }
  }, [siteSettings]);
  if (!isMounted) {
    isMounted = true;
    const lang = Storage.get(CURRENT_LANG_STORAGE_KEY);
    if (lang) {
      i18n.changeLanguage(lang);
    }
  }

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
        <AdminHeader />
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
