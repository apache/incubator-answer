import create from 'zustand';

import { AdminSettingsInterface } from '@/common/interface';
import { DEFAULT_LANG } from '@/common/constants';

interface InterfaceType {
  interface: AdminSettingsInterface;
  update: (params: AdminSettingsInterface) => void;
  updateLogo: (logo: string) => void;
}

const interfaceSetting = create<InterfaceType>((set) => ({
  interface: {
    logo: '',
    theme: '',
    language: DEFAULT_LANG,
    time_zone: '',
  },
  update: (params) =>
    set(() => {
      return {
        interface: params,
      };
    }),
  updateLogo: (logo) =>
    set((state) => {
      return {
        interface: {
          ...state.interface,
          logo,
        },
      };
    }),
}));

export default interfaceSetting;
