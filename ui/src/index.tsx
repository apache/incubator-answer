import React from 'react';

import ReactDOM from 'react-dom/client';

import App from './App';

import { pullLoggedUser } from '@/utils/guards';

import './i18n/init';
import './index.scss';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement,
);

async function bootstrapApp() {
  /**
   * NOTICE: must pre init logged user info for router
   */
  await pullLoggedUser();
  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  );
}

bootstrapApp();
