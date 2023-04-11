import { Suspense, lazy } from 'react';
import { RouteObject } from 'react-router-dom';

import Layout from '@/pages/Layout';

import baseRoutes, { RouteNode } from './routes';
import RouteGuard from './RouteGuard';
import RouteErrorBoundary from './RouteErrorBoundary';

const routes: RouteNode[] = [];

const routeWrapper = (routeNodes: RouteNode[], root: RouteNode[]) => {
  routeNodes.forEach((rn) => {
    if (rn.page === 'pages/Layout') {
      rn.element = rn.guard ? (
        <RouteGuard onEnter={rn.guard} path={rn.path} page={rn.page}>
          <Layout />
        </RouteGuard>
      ) : (
        <Layout />
      );
      rn.errorElement = <RouteErrorBoundary />;
    } else {
      /**
       * cannot use a fully dynamic import statement
       * ref: https://webpack.js.org/api/module-methods/#import-1
       */
      const pagePath = rn.page.replace('pages/', '');
      const Ctrl = lazy(() => import(`@/pages/${pagePath}`));
      rn.element = (
        <Suspense>
          {rn.guard ? (
            <RouteGuard onEnter={rn.guard} path={rn.path} page={rn.page}>
              <Ctrl />
            </RouteGuard>
          ) : (
            <Ctrl />
          )}
        </Suspense>
      );
      rn.errorElement = <RouteErrorBoundary />;
    }
    root.push(rn);
    const children = Array.isArray(rn.children) ? rn.children : null;
    if (children) {
      rn.children = [];
      routeWrapper(children, rn.children);
    }
  });
};

routeWrapper(baseRoutes, routes);

export default routes as RouteObject[];
