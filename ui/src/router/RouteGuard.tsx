import { FC, ReactNode, useEffect } from 'react';
import { useLocation, useNavigate, useLoaderData } from 'react-router-dom';

import { floppyNavigation } from '@/utils';
import { TGuardFunc } from '@/utils/guard';
import { errorCodeStore } from '@/stores';

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

  const { update: updateHttpError } = errorCodeStore();
  const handleGuardError = () => {
    const err = gr.error;
    let errCode = err?.code;
    if (errCode && typeof errCode !== 'string') {
      errCode = errCode.toString();
    }
    if (errCode === '403' || errCode === '404' || errCode === '50X') {
      updateHttpError(errCode, err?.msg);
    }
  };
  useEffect(() => {
    handleGuardError();
  }, [gr.error]);
  const handleGuardRedirect = () => {
    const redirectUrl = gr.redirect;
    if (redirectUrl) {
      floppyNavigation.navigate(redirectUrl, () => {
        navigate(redirectUrl, { replace: true });
      });
    }
  };
  useEffect(() => {
    handleGuardRedirect();
  }, [location]);
  return (
    <>
      {/* Route Guard */}
      {gr.ok ? children : null}
    </>
  );
};

export default RouteGuard;
