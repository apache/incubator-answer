import { FC, ReactNode, useEffect, useState } from 'react';
import { useLocation, useNavigate, useLoaderData } from 'react-router-dom';

import { floppyNavigation } from '@/utils';
import { TGuardFunc } from '@/utils/guard';

import RouteErrorBoundary from './RouteErrorBoundary';

const RouteGuard: FC<{
  children: ReactNode;
  onEnter: TGuardFunc;
  path?: string;
  page?: string;
}> = ({ children, onEnter, path, page }) => {
  const navigate = useNavigate();
  const location = useLocation();
  const loaderData = useLoaderData();
  const [grOk, setGrOk] = useState(true);
  const [routeError, setRouteError] = useState<{
    code: string;
    msg: string;
  }>();
  const applyGuard = () => {
    if (typeof onEnter !== 'function') {
      return;
    }
    const gr = onEnter({
      loaderData,
      path,
      page,
    });

    setGrOk(gr.ok);
    if (gr.error?.code && /403|404|50X/i.test(gr.error.code.toString())) {
      setRouteError({
        code: `${gr.error.code}`,
        msg: gr.error.msg || '',
      });
      return;
    }
    if (gr.redirect) {
      floppyNavigation.navigate(gr.redirect, {
        handler: navigate,
        options: { replace: true },
      });
    }
  };
  useEffect(() => {
    /**
     * NOTICE:
     *  Must be put in `useEffect`,
     *  otherwise `guard` may not get `loggedUserInfo` correctly
     */
    applyGuard();
  }, [location]);
  return (
    <>
      {grOk ? children : null}
      {!grOk && routeError ? (
        <RouteErrorBoundary errCode={routeError.code} />
      ) : null}
    </>
  );
};

export default RouteGuard;
