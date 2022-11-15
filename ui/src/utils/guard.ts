import { getLoggedUserInfo, getAppSettings } from '@/services';
import { loggedUserInfoStore, siteInfoStore, interfaceStore } from '@/stores';
import { RouteAlias } from '@/router/alias';
import Storage from '@/utils/storage';
import { LOGGED_USER_STORAGE_KEY } from '@/common/constants';
import { setupAppLanguage, setupAppTimeZone } from '@/utils/localize';

import { floppyNavigation } from './floppyNavigation';

type TLoginState = {
  isLogged: boolean;
  isNotActivated: boolean;
  isActivated: boolean;
  isForbidden: boolean;
  isNormal: boolean;
  isAdmin: boolean;
};

export type TGuardResult = {
  ok: boolean;
  redirect?: string;
};

export const deriveLoginState = (): TLoginState => {
  const ls: TLoginState = {
    isLogged: false,
    isNotActivated: false,
    isActivated: false,
    isForbidden: false,
    isNormal: false,
    isAdmin: false,
  };
  const { user } = loggedUserInfoStore.getState();
  if (user.access_token) {
    ls.isLogged = true;
  }
  if (ls.isLogged && user.mail_status === 1) {
    ls.isActivated = true;
  }
  if (ls.isLogged && user.mail_status === 2) {
    ls.isNotActivated = true;
  }
  if (ls.isLogged && user.status === 'forbidden') {
    ls.isForbidden = true;
  }
  if (ls.isActivated && !ls.isForbidden) {
    ls.isNormal = true;
  }
  if (ls.isNormal && user.is_admin === true) {
    ls.isAdmin = true;
  }

  return ls;
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
    if (!deriveLoginState().isLogged) {
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

export const logged = () => {
  const gr: TGuardResult = { ok: true };
  const us = deriveLoginState();
  if (!us.isLogged) {
    gr.ok = false;
    gr.redirect = RouteAlias.login;
  }
  return gr;
};

export const notLogged = () => {
  const gr: TGuardResult = { ok: true };
  const us = deriveLoginState();
  if (us.isLogged) {
    gr.ok = false;
    gr.redirect = RouteAlias.home;
  }
  return gr;
};

export const notActivated = () => {
  const gr = logged();
  const us = deriveLoginState();
  if (us.isActivated) {
    gr.ok = false;
    gr.redirect = RouteAlias.home;
  }
  return gr;
};

export const activated = () => {
  const gr = logged();
  const us = deriveLoginState();
  if (us.isNotActivated) {
    gr.ok = false;
    gr.redirect = RouteAlias.activation;
  }
  return gr;
};

export const forbidden = () => {
  const gr = logged();
  const us = deriveLoginState();
  if (gr.ok && !us.isForbidden) {
    gr.ok = false;
    gr.redirect = RouteAlias.home;
  }
  return gr;
};

export const notForbidden = () => {
  const gr: TGuardResult = { ok: true };
  const us = deriveLoginState();
  if (us.isForbidden) {
    gr.ok = false;
    gr.redirect = RouteAlias.suspended;
  }
  return gr;
};

export const admin = () => {
  const gr = logged();
  const us = deriveLoginState();
  if (gr.ok && !us.isAdmin) {
    gr.ok = false;
    gr.redirect = RouteAlias.home;
  }
  return gr;
};

/**
 * try user was logged and all state ok
 * @param canNavigate // if true, will navigate to login page if not logged
 */
export const tryNormalLogged = (canNavigate: boolean = false) => {
  const us = deriveLoginState();

  if (us.isNormal) {
    return true;
  }
  // must assert logged state first and return
  if (!us.isLogged) {
    if (canNavigate) {
      floppyNavigation.navigateToLogin();
    }
    return false;
  }
  if (us.isNotActivated) {
    floppyNavigation.navigate(RouteAlias.activation, () => {
      window.location.href = RouteAlias.activation;
    });
  } else if (us.isForbidden) {
    floppyNavigation.navigate(RouteAlias.suspended, () => {
      window.location.replace(RouteAlias.suspended);
    });
  }

  return false;
};

export const tryLoggedAndActivated = () => {
  const gr: TGuardResult = { ok: true };
  const us = deriveLoginState();

  if (!us.isLogged || !us.isActivated) {
    gr.ok = false;
  }
  return gr;
};

export const initAppSettingsStore = async () => {
  const appSettings = await getAppSettings();
  if (appSettings) {
    siteInfoStore.getState().update(appSettings.general);
    interfaceStore.getState().update(appSettings.interface);
  }
};

export const shouldInitAppFetchData = () => {
  const { pathname } = window.location;
  if (pathname === '/install') {
    return false;
  }

  return true;
};

export const setupApp = async () => {
  /**
   * WARN:
   * 1. must pre init logged user info for router guard
   * 2. must pre init app settings for app render
   */
  // TODO: optimize `initAppSettingsStore` by server render

  if (shouldInitAppFetchData()) {
    await Promise.allSettled([pullLoggedUser(), initAppSettingsStore()]);
    setupAppLanguage();
    setupAppTimeZone();
  }
};
