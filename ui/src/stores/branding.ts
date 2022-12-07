import create from 'zustand';

import { AdminSettingBranding } from '@/common/interface';
import { DEFAULT_LANG } from '@/common/constants';

interface InterfaceType {
  branding: AdminSettingBranding;
  update: (params: AdminSettingBranding) => void;
}

const interfaceSetting = create<InterfaceType>((set) => ({
  branding: {
    logo: '',
    square_icon: '',
    mobile_logo: '',
    favicon: '',
  },
  interface: {
    theme: '',
    language: DEFAULT_LANG,
    time_zone: '',
  },
  update: (params) =>
    set(() => {
      return {
        branding: params,
      };
    }),
}));

export default interfaceSetting;
