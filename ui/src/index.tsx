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

import React from 'react';

import ReactDOM from 'react-dom/client';

import App from './App';

import './index.scss';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement,
);

/**
 * Uniformly hide broken images
 *  - The `onload` event for elements such as `img` can only be `capture` on `document` (window cannot).
 *  - For images with an empty `src` attribute, sometimes the browser will simply display the broken image without reporting an 'error' event.
 */
const handleImgLoad = (evt: Event | UIEvent) => {
  const { target } = evt;

  if (target === null || !(target instanceof Element)) {
    return;
  }
  if (!/IMG/i.test(target.nodeName)) {
    return;
  }

  if (/error/i.test(evt.type)) {
    target.classList.add('broken');
    const attrSrc = target.getAttribute('src');
    const attrAlt = target.getAttribute('alt')?.trim();
    // Images without the `src` attribute are hidden directly by `css`.
    // Images with `alt` content are not hidden - the display of the `alt` content is also hidden.
    if (attrSrc && !attrAlt) {
      target.classList.add('invisible');
    }
  }

  if (/load/i.test(evt.type)) {
    target.classList.remove('broken', 'invisible');
  }
};

/**
 *Automatically jump when the href of a Link component within a matching project is not a front-end route.
 *
 */
const handleClickLink = (evt: Event) => {
  const { target } = evt;

  if (target === null || !(target instanceof Element)) {
    return;
  }
  if (!/A/i.test(target.nodeName)) {
    return;
  }

  if (target.getAttribute('href')?.includes('/answer/api/')) {
    evt.preventDefault();
    window.location.href = target.getAttribute('href') || '';
  }
};

document.addEventListener('error', handleImgLoad, true);
document.addEventListener('load', handleImgLoad, true);
document.addEventListener('click', handleClickLink, true);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
