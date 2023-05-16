import { useNavigate } from 'react-router-dom';

import { floppyNavigation } from '@/utils';
import Storage from '@/utils/storage';
import { RouteAlias } from '@/router/alias';
import { REDIRECT_PATH_STORAGE_KEY } from '@/common/constants';

const Index = () => {
  const navigate = useNavigate();

  const loginRedirect = () => {
    const redirect = Storage.get(REDIRECT_PATH_STORAGE_KEY) || RouteAlias.home;
    Storage.remove(REDIRECT_PATH_STORAGE_KEY);
    floppyNavigation.navigate(redirect, {
      handler: navigate,
      options: {
        replace: true,
      },
    });
  };

  return { loginRedirect };
};

export default Index;
