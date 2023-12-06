import { useEffect } from 'react';
import { Spinner } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { logout } from '@/services';
import { loggedUserInfoStore } from '@/stores';
import Storage from '@/utils/storage';
import { RouteAlias } from '@/router/alias';
import { REDIRECT_PATH_STORAGE_KEY } from '@/common/constants';

const Index = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const { user: loggedUserInfo, clear: clearUserStore } = loggedUserInfoStore();

  usePageTags({
    title: t('logout'),
  });

  useEffect(() => {
    if (loggedUserInfo.username) {
      logout().then(() => {
        clearUserStore();
        const redirect =
          Storage.get(REDIRECT_PATH_STORAGE_KEY) || RouteAlias.home;
        Storage.remove(REDIRECT_PATH_STORAGE_KEY);
        window.location.replace(`${window.location.origin}${redirect}`);
      });
    }
    // auto height of container
    const pageWrap = document.querySelector('.page-wrap') as HTMLElement;
    if (pageWrap) {
      pageWrap.style.display = 'contents';
    }

    return () => {
      if (pageWrap) {
        pageWrap.style.display = 'block';
      }
    };
  }, []);
  return (
    <div className="d-flex flex-column flex-shrink-1 flex-grow-1 justify-content-center align-items-center">
      <Spinner variant="secondary" />
    </div>
  );
};

export default Index;
