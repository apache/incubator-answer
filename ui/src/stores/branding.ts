import create from 'zustand';

import { AdminSettingBranding } from '@/common/interface';

interface IType {
  branding: AdminSettingBranding;
  update: (params: AdminSettingBranding) => void;
}

const brandingSetting = create<IType>((set) => ({
  branding: {
    logo: '',
    square_icon: '',
    mobile_logo: '',
    favicon: '',
  },
  update: (params) =>
    set(() => {
      return {
        branding: params,
      };
    }),
}));

export default brandingSetting;
