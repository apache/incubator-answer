import { FC, ReactNode, useEffect, useState } from 'react';
import { useLocation, useNavigate, useLoaderData } from 'react-router-dom';

import { floppyNavigation } from '@/utils';
import { TGuardFunc, TGuardResult } from '@/utils/guard';

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
  const [gk, setKeeper] = useState<TGuardResult>({
    ok: true,
  });
  const [gkError, setGkError] = useState<TGuardResult['error']>();
  const applyGuard = () => {
    if (typeof onEnter !== 'function') {
      return;
    }
    const gr = onEnter({
      loaderData,
      path,
      page,
    });

    setKeeper(gr);
    if (
      gk.ok === false &&
      gk.error?.code &&
      /403|404|50X/i.test(gk.error.code.toString())
    ) {
      setGkError(gk.error);
      return;
    }
    setGkError(undefined);
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

  let asOK = gk.ok;
  if (gk.ok === false && gk.redirect) {
    /**
     * It is possible that the route guard verification fails
     *    but the current page is already the target page for the route guard jump
     * This should render `children`!
     */

    asOK = floppyNavigation.equalToCurrentHref(gk.redirect);
  }

  return (
    <>
      {asOK ? children : null}
      {gkError ? <RouteErrorBoundary errCode={gkError.code as string} /> : null}
    </>
  );
};

export default RouteGuard;
