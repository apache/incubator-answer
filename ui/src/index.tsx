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

document.addEventListener('error', handleImgLoad, true);
document.addEventListener('load', handleImgLoad, true);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
