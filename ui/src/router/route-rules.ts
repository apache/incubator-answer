import { isLogin } from '@answer/utils';

const RouteRules = {
  isLoginAndNormal: () => {
    return isLogin(true);
  },
};

export default RouteRules;
