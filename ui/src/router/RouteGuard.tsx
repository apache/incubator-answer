import { FC, ReactNode, useEffect } from 'react';
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
  const gr = onEnter({
    loaderData,
    path,
    page,
  });

  let guardError;
  const errCode = gr.error?.code;
  if (errCode === '403' || errCode === '404' || errCode === '50X') {
    guardError = {
      code: errCode,
      msg: gr.error?.msg,
    };
  }
  const handleGuardRedirect = () => {
    const redirectUrl = gr.redirect;
    if (redirectUrl) {
      floppyNavigation.navigate(redirectUrl, {
        handler: navigate,
        options: { replace: true },
      });
    }
  };
  useEffect(() => {
    handleGuardRedirect();
  }, [location]);
  return (
    <>
      {gr.ok ? children : null}
      {!gr.ok && guardError ? (
        <RouteErrorBoundary errCode={guardError.code} />
      ) : null}
    </>
  );
};

export default RouteGuard;
