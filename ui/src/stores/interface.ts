import create from 'zustand';

import { AdminSettingsInterface } from '@/common/interface';
import { DEFAULT_LANG } from '@/common/constants';

interface InterfaceType {
  interface: AdminSettingsInterface;
  update: (params: AdminSettingsInterface) => void;
}

const interfaceSetting = create<InterfaceType>((set) => ({
  interface: {
    language: DEFAULT_LANG,
    time_zone: '',
    default_avatar: 'system',
  },
  update: (params) =>
    set(() => {
      return {
        interface: params,
      };
    }),
}));

export default interfaceSetting;
