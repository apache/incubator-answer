import type { NavigateFunction } from 'react-router-dom';

import { RouteAlias } from '@/router/alias';
import Storage from '@/utils/storage';
import { REDIRECT_PATH_STORAGE_KEY } from '@/common/constants';
import { getLoginUrl } from '@/utils/userCenter';

const differentCurrent = (target: string, base?: string) => {
  base ||= window.location.origin;
  const targetUrl = new URL(target, base);
  return targetUrl.toString() !== window.location.href;
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
 * Determining whether an url is an external link
 */
const isExternalLink = (url = '') => {
  let ret = false;
  try {
    const urlObject = new URL(url, document.baseURI);
    if (urlObject && urlObject.origin !== window.location.origin) {
      ret = true;
    }
    // eslint-disable-next-line no-empty
  } catch (ex) {}
  return ret;
};

/**
 * only navigate if not same as current url
 */
type NavigateHandler = 'href' | 'replace' | NavigateFunction;
interface NavigateConfig {
  handler: NavigateHandler;
  options?: any;
}
const navigate = (
  to: string | number,
  config: NavigateConfig = { handler: 'href' },
) => {
  let { handler } = config;
  if (to && typeof to === 'string') {
    if (!differentCurrent(to)) {
      return;
    }
    if (to === RouteAlias.login || to === getLoginUrl()) {
      storageLoginRedirect();
    }
    if (isExternalLink(to)) {
      handler = 'href';
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
  const { target, nodeName } = evt.currentTarget;
  if (nodeName.toLowerCase() !== 'a') {
    return false;
  }
  return (
    evt.button === 0 &&
    (!target || target === '_self') &&
    !(evt.metaKey || evt.ctrlKey || evt.shiftKey || evt.altKey)
  );
};

export const floppyNavigation = {
  differentCurrent,
  navigate,
  navigateToLogin,
  shouldProcessLinkClick,
};
