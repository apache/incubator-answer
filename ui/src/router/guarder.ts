import {
  pullLoggedUser,
  isLoggedAndNormal,
  isAdminLogged,
  isNotLogged,
  isNotLoggedOrNormal,
  isLoggedAndInactive,
  isLoggedAndSuspended,
  isNotLoggedOrInactive,
} from '@/utils/guards';

const RouteGuarder = {
  base: async () => {
    return isNotLoggedOrNormal();
  },
  notLogged: async () => {
    return isNotLogged();
  },
  notLoggedOrInactive: async () => {
    return isNotLoggedOrInactive();
  },
  loggedAndNormal: async () => {
    await pullLoggedUser(true);
    return isLoggedAndNormal();
  },
  loggedAndInactive: async () => {
    return isLoggedAndInactive();
  },
  loggedAndSuspended: async () => {
    return isLoggedAndSuspended();
  },
  adminLogged: async () => {
    await pullLoggedUser(true);
    return isAdminLogged();
  },
};

export default RouteGuarder;
