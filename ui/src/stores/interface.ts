import create from 'zustand';

import { AdminSettingsInterface } from '@/common/interface';
import { DEFAULT_LANG } from '@/common/constants';

interface InterfaceType {
  interface: AdminSettingsInterface;
  update: (params: AdminSettingsInterface) => void;
}

const interfaceSetting = create<InterfaceType>((set) => ({
  interface: {
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
}));

export default interfaceSetting;
