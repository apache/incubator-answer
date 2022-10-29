import { getLoggedUserInfo } from '@/services';
import { loggedUserInfoStore } from '@/stores';
import { RouteAlias } from '@/router/alias';
import Storage from '@/utils/storage';
import { LOGGED_USER_STORAGE_KEY } from '@/common/constants';
import { floppyNavigation } from '@/utils/floppyNavigation';

type UserStat = {
  isLogged: boolean;
  isActivated: boolean;
  isSuspended: boolean;
  isNormal: boolean;
  isAdmin: boolean;
};
export const deriveUserStat = (): UserStat => {
  const stat: UserStat = {
    isLogged: false,
    isActivated: false,
    isSuspended: false,
    isNormal: false,
    isAdmin: false,
  };
  const { user } = loggedUserInfoStore.getState();
  if (user.id && user.username) {
    stat.isLogged = true;
  }
  if (stat.isLogged && user.mail_status === 1) {
    stat.isActivated = true;
  }
  if (stat.isLogged && user.status === 'forbidden') {
    stat.isSuspended = true;
  }
  if (stat.isLogged && stat.isActivated && !stat.isSuspended) {
    stat.isNormal = true;
  }
  if (stat.isNormal && user.is_admin === true) {
    stat.isAdmin = true;
  }

  return stat;
};

type GuardResult = {
  ok: boolean;
  redirect?: string;
};
let pullLock = false;
let dedupeTimestamp = 0;
export const pullLoggedUser = async (forceRePull = false) => {
  // only pull once if not force re-pull
  if (pullLock && !forceRePull) {
    return;
  }
  // dedupe pull requests in this time span in 10 seconds
  if (Date.now() - dedupeTimestamp < 1000 * 10) {
    return;
  }
  dedupeTimestamp = Date.now();
  const loggedUserInfo = await getLoggedUserInfo().catch((ex) => {
    dedupeTimestamp = 0;
    if (!deriveUserStat().isLogged) {
      // load fallback userInfo from local storage
      const storageLoggedUserInfo = Storage.get(LOGGED_USER_STORAGE_KEY);
      if (storageLoggedUserInfo) {
        loggedUserInfoStore.getState().update(storageLoggedUserInfo);
      }
    }
    console.error(ex);
  });
  if (loggedUserInfo) {
    pullLock = true;
    loggedUserInfoStore.getState().update(loggedUserInfo);
  }
};

export const isLogged = () => {
  const ret: GuardResult = { ok: true, redirect: undefined };
  const userStat = deriveUserStat();
  if (!userStat.isLogged) {
    ret.ok = false;
    ret.redirect = RouteAlias.login;
  }
  return ret;
};

export const isNotLogged = () => {
  const ret: GuardResult = { ok: true, redirect: undefined };
  const userStat = deriveUserStat();
  if (userStat.isLogged) {
    ret.ok = false;
    ret.redirect = RouteAlias.home;
  }
  return ret;
};

export const isLoggedAndInactive = () => {
  const ret: GuardResult = { ok: false, redirect: undefined };
  const userStat = deriveUserStat();
  if (!userStat.isActivated) {
    ret.ok = true;
    ret.redirect = RouteAlias.activation;
  }
  return ret;
};

export const isLoggedAndSuspended = () => {
  const ret: GuardResult = { ok: false, redirect: undefined };
  const userStat = deriveUserStat();
  if (userStat.isSuspended) {
    ret.redirect = RouteAlias.suspended;
    ret.ok = true;
  }
  return ret;
};

export const isLoggedAndNormal = () => {
  const ret: GuardResult = { ok: false, redirect: undefined };
  const userStat = deriveUserStat();
  if (userStat.isNormal) {
    ret.ok = true;
  } else if (!userStat.isActivated) {
    ret.redirect = RouteAlias.activation;
  } else if (!userStat.isSuspended) {
    ret.redirect = RouteAlias.suspended;
  } else if (!userStat.isLogged) {
    ret.redirect = RouteAlias.login;
  }
  return ret;
};

export const isNotLoggedOrNormal = () => {
  const ret: GuardResult = { ok: true, redirect: undefined };
  const userStat = deriveUserStat();
  const gr = isLoggedAndNormal();
  if (!gr.ok && userStat.isLogged) {
    ret.ok = false;
    ret.redirect = gr.redirect;
  }
  return ret;
};

export const isNotLoggedOrInactive = () => {
  const ret: GuardResult = { ok: true, redirect: undefined };
  const userStat = deriveUserStat();
  if (userStat.isLogged || userStat.isActivated) {
    ret.ok = false;
    ret.redirect = RouteAlias.home;
  }
  return ret;
};

export const isAdminLogged = () => {
  const ret: GuardResult = { ok: true, redirect: undefined };
  const userStat = deriveUserStat();
  if (!userStat.isAdmin) {
    ret.redirect = RouteAlias.home;
    ret.ok = false;
  }
  return ret;
};

/**
 * try user was logged and all state ok
 * @param autoLogin
 */
export const tryNormalLogged = (autoLogin: boolean = false) => {
  const gr = isLoggedAndNormal();
  if (gr.ok) {
    return true;
  }

  if (gr.redirect === RouteAlias.login && autoLogin) {
    floppyNavigation.navigateToLogin();
  } else if (gr.redirect) {
    floppyNavigation.navigate(gr.redirect, () => {
      // @ts-ignore
      window.location.replace(gr.redirect);
    });
  }

  return false;
};
