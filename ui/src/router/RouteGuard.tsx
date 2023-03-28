import { FC, ReactNode, useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import { floppyNavigation } from '@/utils';
import { TGuardFunc } from '@/utils/guard';

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
  const callGuards = () => {
    if (onEnter) {
      const gr = onEnter();
      const redirectUrl = gr.redirect;
      if (redirectUrl) {
        floppyNavigation.navigate(redirectUrl, {
          handler: navigate,
          options: { replace: true },
        });
      }
    }
  };
  useEffect(() => {
    callGuards();
  }, [location]);
  return (
    <>
      {/* Route Guard */}
      {children}
    </>
  );
};

export default Index;
