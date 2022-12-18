import { Suspense, lazy } from 'react';
import { RouteObject, createBrowserRouter } from 'react-router-dom';

import Layout from '@/pages/Layout';
import ErrorBoundary from '@/pages/50X';
import baseRoutes, { RouteNode } from '@/router/routes';
import RouteGuard from '@/router/RouteGuard';

const routes: RouteObject[] = [];

const routeWrapper = (routeNodes: RouteNode[], root: RouteObject[]) => {
  routeNodes.forEach((rn) => {
    if (rn.page === 'pages/Layout') {
      rn.element = rn.guard ? (
        <RouteGuard onEnter={rn.guard} path={rn.path}>
          <Layout />
        </RouteGuard>
      ) : (
        <Layout />
      );
      rn.errorElement = <ErrorBoundary />;
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
            <RouteGuard onEnter={rn.guard} path={rn.path}>
              <Ctrl />
            </RouteGuard>
          ) : (
            <Ctrl />
          )}
        </Suspense>
      );
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

export { routes, createBrowserRouter };
