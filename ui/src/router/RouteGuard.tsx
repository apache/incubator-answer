import { FC, ReactNode, useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import { floppyNavigation } from '@/utils';
import { TGuardFunc } from '@/utils/guard';
import { loggedUserInfoStore } from '@/stores';

const Index: FC<{
  children: ReactNode;
  onEnter?: TGuardFunc;
  path?: string;
}> = ({
  children,
  onEnter,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  path,
}) => {
  const navigate = useNavigate();
  const location = useLocation();
  const { user } = loggedUserInfoStore();
  const runGuards = () => {
    if (onEnter) {
      const gr = onEnter();
      const redirectUrl = gr.redirect;
      if (redirectUrl) {
        floppyNavigation.navigate(redirectUrl, () => {
          navigate(redirectUrl, { replace: true });
        });
      }
    }
  };
  useEffect(() => {
    runGuards();
  }, [location]);
  useEffect(() => {
    if (!user.access_token) {
      runGuards();
    }
  }, [user]);
  return (
    <>
      {/* Route Guard */}
      {children}
    </>
  );
};

export default Index;
