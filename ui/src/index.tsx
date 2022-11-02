import React from 'react';

import ReactDOM from 'react-dom/client';

import { Guard } from '@/utils';

import App from './App';

import './i18n/init';
import './index.scss';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement,
);

async function bootstrapApp() {
  /**
   * NOTICE: must pre init logged user info for router
   */
  await Guard.pullLoggedUser();
  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  );
}

bootstrapApp();
