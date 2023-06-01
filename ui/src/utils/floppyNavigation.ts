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
const matchToCurrentHref = (target: string) => {
  target = (target || '').trim();
  // Empty string or `/` can match any path
  if (!target || target === '/') {
    return true;
  }
  const { pathname, search, hash } = window.location;
  const tPart = target.split('?');

  /**
   * With the current requirements, `hash` and `search` can simply be matched
   * Later extended to field-by-field matching if necessary
   */
  if (tPart[1]) {
    const tChip = tPart[1].split('#');
    const tSearch = tChip[0] || '';
    const tHash = tChip[1] || '';
    if (tHash && hash.indexOf(tHash) === -1) {
      return false;
    }
    if (tSearch && search.indexOf(tSearch) === -1) {
      return false;
    }
  }

  /**
   * As determination above, `tPart[0]` must be a valid string
   */
  let pathMatch = true;
  const tPath = tPart[0].split('/').filter((_) => !!_);
  const lPath = pathname.split('/').filter((_) => !!_);

  tPath.forEach((p, i) => {
    const lp = lPath[i];
    if (p !== lp) {
      pathMatch = false;
    }
  });

  return pathMatch;
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
     * Note: The or judgement cannot be missing here, both jumps will be used
     */
    if (to === RouteAlias.login || to === getLoginUrl()) {
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
 * auto navigate to login page
 * Note: Only the internal login page is jumped here, `userAgent` login is handled on the internal login page.
 */
const navigateToLogin = (config?: NavigateConfig) => {
  const loginUrl = RouteAlias.login;
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
  matchToCurrentHref,
  storageLoginRedirect,
};
