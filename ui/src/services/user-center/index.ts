import request from '@/utils/request';

export interface UcAgentControl {
  name: string;
  label: string;
  url: string;
}
export interface UcAgent {
  enabled: boolean;
  agent_info: {
    name: string;
    icon: string;
    url: string;
    login_redirect_url: string;
    sign_up_redirect_url: string;
    control_center: UcAgentControl[];
  };
}

export const getUcAgent = () => {
  const apiUrl = `/answer/api/v1/user-center/agent`;
  return request.get<UcAgent>(apiUrl);
};
