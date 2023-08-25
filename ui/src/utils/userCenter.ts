import { RouteAlias } from '@/router/alias';
import { userCenterStore } from '@/stores';
import { getUcAgent, UcAgent } from '@/services/user-center';

export const pullUcAgent = async () => {
  const uca = await getUcAgent();
  userCenterStore.getState().update(uca);
};

export const getLoginUrl = () => {
  return `${process.env.REACT_APP_LOGIN_URL}?redirect_url=${window.location.origin}/users/login`;
};

export const getSignUpUrl = (uca?: UcAgent) => {
  let ret = RouteAlias.signUp;
  uca ||= userCenterStore.getState().agent;
  if (uca?.enabled && uca?.agent_info?.sign_up_redirect_url) {
    ret = uca.agent_info.sign_up_redirect_url;
  }
  return ret;
};
