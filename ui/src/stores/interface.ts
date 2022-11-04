import create from 'zustand';

import { AdminSettingsInterface } from '@/common/interface';

interface InterfaceType {
  interface: AdminSettingsInterface;
  update: (params: AdminSettingsInterface) => void;
}

const interfaceSetting = create<InterfaceType>((set) => ({
  interface: {
    logo: '',
    theme: '',
    language: '',
    time_zone: '',
  },
  update: (params) =>
    set(() => {
      return {
        interface: params,
      };
    }),
}));

export default interfaceSetting;
