/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import type { NavigateFunction } from 'react-router-dom';

import { RouteAlias, REACT_BASE_PATH } from '@/router/alias';
import Storage from '@/utils/storage';
import { REDIRECT_PATH_STORAGE_KEY } from '@/common/constants';
import { getLoginUrl } from '@/utils/userCenter';

const equalToCurrentHref = (target: string, base?: string) => {
  base ||= window.location.origin;
  const targetUrl = new URL(
    target.startsWith(REACT_BASE_PATH) ? target : `${REACT_BASE_PATH}${target}`,
    base,
  );
  return targetUrl.toString() === window.location.href;
};
const matchToCurrentHref = (target: string) => {
  target = (target || '').trim();
  const hasBasePath = target.startsWith(REACT_BASE_PATH);
  // Empty string or `/` can match any path
  if (!target || target === '/') {
    return true;
  }
  const { pathname, search, hash } = window.location;
  let pathWithOutBase = pathname;
  if (!hasBasePath) {
    pathWithOutBase = pathWithOutBase.replace(REACT_BASE_PATH, '');
  }

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
  const lPath = pathWithOutBase.split('/').filter((_) => !!_);

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
  const filterPath = pathname.replace(REACT_BASE_PATH, '');
  if (filterPath !== RouteAlias.login && filterPath !== RouteAlias.signUp) {
    const loc = window.location;
    const redirectUrl = loc.href.replace(`${loc.origin}${REACT_BASE_PATH}`, '');
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
      if (
        to.startsWith('/') &&
        !to.startsWith('//') &&
        !to.startsWith(REACT_BASE_PATH)
      ) {
        to = `${REACT_BASE_PATH}${to}`;
      }
      window.location.href = to;
    } else if (handler === 'replace') {
      if (
        to.startsWith('/') &&
        !to.startsWith('//') &&
        !to.startsWith(REACT_BASE_PATH)
      ) {
        to = `${REACT_BASE_PATH}${to}`;
      }
      window.location.replace(to);
    } else if (typeof handler === 'function') {
      if (to === REACT_BASE_PATH) {
        to = '/';
      }

      if (to !== REACT_BASE_PATH && to.startsWith(REACT_BASE_PATH)) {
        to = to.replace(REACT_BASE_PATH, '');
      }
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
