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

export interface UcSettingAgent {
  enabled: boolean;
  redirect_url: string;
}
export interface UcSettings {
  profile_setting_agent: UcSettingAgent;
  account_setting_agent: UcSettingAgent;
}

export const getUcAgent = () => {
  const apiUrl = `/answer/api/v1/user-center/agent`;
  return request.get<UcAgent>(apiUrl);
};

export const getUcSettings = () => {
  const apiUrl = `/answer/api/v1/user-center/user/settings`;
  return request.get<UcSettings>(apiUrl);
};
