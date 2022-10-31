import {
  pullLoggedUser,
  isLoggedAndNormal,
  isAdminLogged,
  isNotLogged,
  isNotLoggedOrNormal,
  isLoggedAndInactive,
  isLoggedAndSuspended,
  isNotLoggedOrInactive,
  isNotLoggedOrNotSuspend,
} from '@/utils/guards';

const RouteGuarder = {
  base: async () => {
    return isNotLoggedOrNotSuspend();
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
  notLogged: async () => {
    return isNotLogged();
  },
  notLoggedOrNormal: async () => {
    return isNotLoggedOrNormal();
  },
  notLoggedOrInactive: async () => {
    return isNotLoggedOrInactive();
  },
};

export default RouteGuarder;
