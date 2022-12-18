import create from 'zustand';

import { AdminSettingsLogin } from '@/common/interface';

interface IType {
  login: AdminSettingsLogin;
  update: (params: AdminSettingsLogin) => void;
}

const loginSetting = create<IType>((set) => ({
  login: {
    allow_new_registrations: true,
    login_required: false,
  },
  update: (params) =>
    set(() => {
      return {
        login: params,
      };
    }),
}));

export default loginSetting;
