import { getLoggedUserInfo, getAppSettings } from '@/services';
import {
  loggedUserInfoStore,
  siteInfoStore,
  interfaceStore,
  brandingStore,
  loginSettingStore,
  customizeStore,
  themeSettingStore,
  seoSettingStore,
  loginToContinueStore,
} from '@/stores';
import { RouteAlias } from '@/router/alias';

import { setupAppLanguage, setupAppTimeZone } from './localize';
import { floppyNavigation } from './floppyNavigation';
import { pullUcAgent, getLoginUrl, getSignUpUrl } from './userCenter';

type TLoginState = {
  isLogged: boolean;
  isNotActivated: boolean;
  isActivated: boolean;
  isForbidden: boolean;
  isNormal: boolean;
  isAdmin: boolean;
  isModerator: boolean;
};

export type TGuardResult = {
  ok: boolean;
  redirect?: string;
  error?: {
    code?: number | string;
    msg?: string;
  };
};
export type TGuardFunc = (args: {
  loaderData?: any;
  path?: string;
  page?: string;
}) => TGuardResult;

export const deriveLoginState = (): TLoginState => {
  const ls: TLoginState = {
    isLogged: false,
    isNotActivated: false,
    isActivated: false,
    isForbidden: false,
    isNormal: false,
    isAdmin: false,
    isModerator: false,
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
  if (ls.isNormal && user.role_id === 2) {
    ls.isAdmin = true;
  }
  if (ls.isNormal && user.role_id === 3) {
    ls.isModerator = true;
  }
  return ls;
};

export const isIgnoredPath = (ignoredPath: string | string[]) => {
  if (!Array.isArray(ignoredPath)) {
    ignoredPath = [ignoredPath];
  }
  const { pathname } = window.location;
  const matchingPath = ignoredPath.find((_) => {
    return pathname.indexOf(_) !== -1;
  });
  return !!matchingPath;
};

let pluLock = false;
let pluTimestamp = 0;
export const pullLoggedUser = async (forceRePull = false) => {
  // only pull once if not force re-pull
  if (pluLock && !forceRePull) {
    return;
  }
  // dedupe pull requests in this time span in 10 seconds
  if (Date.now() - pluTimestamp < 1000 * 10) {
    return;
  }
  pluTimestamp = Date.now();
  const loggedUserInfo = await getLoggedUserInfo().catch((ex) => {
    pluTimestamp = 0;
    loggedUserInfoStore.getState().clear(false);
    console.error(ex);
  });
  if (loggedUserInfo) {
    pluLock = true;
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
  const gr: TGuardResult = { ok: true };
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
  const gr: TGuardResult = { ok: true };
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
    gr.error = {
      code: '403',
      msg: '',
    };
    gr.redirect = '';
  }
  return gr;
};

export const isAdminOrModerator = () => {
  const gr = logged();
  const us = deriveLoginState();
  if (gr.ok && !us.isAdmin && !us.isModerator) {
    gr.ok = false;
    gr.error = {
      code: '403',
      msg: '',
    };
    gr.redirect = '';
  }
  return gr;
};

export const isEditable = (args) => {
  const loaderData = args?.loaderData || {};
  const gr: TGuardResult = { ok: true };
  if (loaderData.code === 400) {
    gr.ok = false;
    gr.error = {
      code: '403',
      msg: loaderData.msg,
    };
  }
  return gr;
};

export const allowNewRegistration = () => {
  const gr: TGuardResult = { ok: true };
  const loginSetting = loginSettingStore.getState().login;
  if (!loginSetting.allow_new_registrations) {
    gr.ok = false;
    gr.redirect = RouteAlias.home;
  }
  return gr;
};

export const loginAgent = () => {
  const gr: TGuardResult = { ok: true };
  const loginUrl = getLoginUrl();
  if (loginUrl !== RouteAlias.login) {
    gr.ok = false;
    gr.redirect = loginUrl;
  }
  return gr;
};

export const singUpAgent = () => {
  const gr: TGuardResult = { ok: true };
  const signUpUrl = getSignUpUrl();
  if (signUpUrl !== RouteAlias.signUp) {
    gr.ok = false;
    gr.redirect = signUpUrl;
  }
  return gr;
};

export const shouldLoginRequired = () => {
  const gr: TGuardResult = { ok: true };
  const loginSetting = loginSettingStore.getState().login;
  if (!loginSetting.login_required) {
    return gr;
  }
  const us = deriveLoginState();
  if (us.isLogged) {
    return gr;
  }
  if (
    isIgnoredPath([
      RouteAlias.login,
      RouteAlias.signUp,
      '/users/account-recovery',
      'users/change-email',
      'users/password-reset',
      'users/account-activation',
      'users/account-activation/success',
      '/users/account-activation/failed',
      '/users/confirm-new-email',
    ])
  ) {
    return gr;
  }
  gr.ok = false;
  gr.redirect = RouteAlias.login;
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
      loginToContinueStore.getState().update({ show: true });
    }
    return false;
  }
  if (us.isNotActivated) {
    floppyNavigation.navigate(RouteAlias.activation);
  } else if (us.isForbidden) {
    floppyNavigation.navigate(RouteAlias.suspended, {
      handler: 'replace',
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

/**
 * Initialize app configuration
 */
let appInitialized = false;
export const initAppSettingsStore = async () => {
  const appSettings = await getAppSettings();
  if (appSettings) {
    siteInfoStore.getState().update(appSettings.general);
    siteInfoStore.getState().updateVersion(appSettings.version);
    interfaceStore.getState().update(appSettings.interface);
    brandingStore.getState().update(appSettings.branding);
    loginSettingStore.getState().update(appSettings.login);
    customizeStore.getState().update(appSettings.custom_css_html);
    themeSettingStore.getState().update(appSettings.theme);
    seoSettingStore.getState().update(appSettings.site_seo);
  }
};

export const setupApp = async () => {
  if (appInitialized) {
    return;
  }
  /**
   * WARN:
   * 1. must pre init logged user info for router guard
   * 2. must pre init app settings for app render
   */
  // TODO: optimize `initAppSettingsStore` by server render
  await Promise.allSettled([
    pullLoggedUser(),
    pullUcAgent(),
    initAppSettingsStore(),
  ]);
  setupAppLanguage();
  setupAppTimeZone();
  appInitialized = true;
};
