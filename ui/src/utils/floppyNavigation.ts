import type { NavigateFunction } from 'react-router-dom';

import { RouteAlias } from '@/router/alias';
import Storage from '@/utils/storage';
import { REDIRECT_PATH_STORAGE_KEY } from '@/common/constants';
import { getLoginUrl } from '@/utils/userCenter';

const equalToCurrentHref = (target: string, base?: string) => {
  base ||= window.location.origin;
  const targetUrl = new URL(target, base);
  return targetUrl.toString() === window.location.href;
};

const storageLoginRedirect = () => {
  const { pathname } = window.location;
  if (pathname !== RouteAlias.login && pathname !== RouteAlias.signUp) {
    const loc = window.location;
    const redirectUrl = loc.href.replace(loc.origin, '');
    Storage.set(REDIRECT_PATH_STORAGE_KEY, redirectUrl);
  }
};

/**
 * Determining if an url is a full link
 */
const isFullLink = (url = '') => {
  let ret = false;
  if (/^(http:|https:|\/\/)/i.test(url)) {
    ret = true;
  }
  return ret;
};

/**
 * Determining if a link is routable
 */
const isRoutableLink = (url = '') => {
  let ret = true;
  if (isFullLink(url)) {
    ret = false;
  }

  return ret;
};

/**
 * only navigate if not same as current url
 */
type NavigateHandler = 'href' | 'replace' | NavigateFunction;
export interface NavigateConfig {
  handler?: NavigateHandler;
  options?: any;
}
const navigate = (to: string | number, config: NavigateConfig = {}) => {
  let { handler = 'href' } = config;
  if (to && typeof to === 'string') {
    if (equalToCurrentHref(to)) {
      return;
    }
    /**
     * 1. Blocking redirection of two login pages
     * 2. Auto storage login redirect
     */
    if (to === RouteAlias.login || to === getLoginUrl()) {
      if (equalToCurrentHref(RouteAlias.login)) {
        return;
      }
      storageLoginRedirect();
    }

    if (!isRoutableLink(to) && handler !== 'href' && handler !== 'replace') {
      handler = 'href';
    }
    if (handler === 'href' && config.options?.replace) {
      handler = 'replace';
    }
    if (handler === 'href') {
      window.location.href = to;
    } else if (handler === 'replace') {
      window.location.replace(to);
    } else if (typeof handler === 'function') {
      handler(to, config.options);
    }
  }
  if (typeof to === 'number' && typeof handler === 'function') {
    handler(to);
  }
};

/**
 * auto navigate to login page with redirect info
 */
const navigateToLogin = (config?: NavigateConfig) => {
  const loginUrl = getLoginUrl();
  navigate(loginUrl, config);
};

/**
 * Determine if a Link click event should be handled
 */
const shouldProcessLinkClick = (evt) => {
  if (evt.defaultPrevented) {
    return false;
  }
  const nodeName = evt.currentTarget?.nodeName;
  if (nodeName?.toLowerCase() !== 'a') {
    return false;
  }
  const target = evt.currentTarget?.target;
  return (
    evt.button === 0 &&
    (!target || target === '_self') &&
    !(evt.metaKey || evt.ctrlKey || evt.shiftKey || evt.altKey)
  );
};

/**
 * Automatic handling of click events on route links
 */
const handleRouteLinkClick = (evt) => {
  if (!shouldProcessLinkClick(evt)) {
    return;
  }
  const curTarget = evt.currentTarget;
  const href = curTarget?.getAttribute('href');
  if (!isRoutableLink(href)) {
    evt.preventDefault();
    navigate(href);
  }
};

export const floppyNavigation = {
  navigate,
  navigateToLogin,
  shouldProcessLinkClick,
  isFullLink,
  isRoutableLink,
  handleRouteLinkClick,
  equalToCurrentHref,
};
