import React from 'react';

import ReactDOM from 'react-dom/client';

import App from './App';

import './index.scss';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement,
);

/**
 * Uniformly hide broken images
 */
document.addEventListener(
  'error',
  (err) => {
    const { target } = err;
    if (target === null || !(target instanceof Element)) {
      return;
    }

    if (/IMG/.test(target.nodeName)) {
      if (!target.getAttribute('alt')) {
        target.classList.add('invisible');
      }
    }
  },
  true,
);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
