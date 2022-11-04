import React from 'react';

import ReactDOM from 'react-dom/client';

import { guard } from '@/utils';

import App from './App';

import './index.scss';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement,
);

async function bootstrapApp() {
  await guard.setupApp();
  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  );
}

bootstrapApp();
