import create from 'zustand';

import { AdminSettingsGeneral } from '@/common/interface';

interface SiteInfoType {
  siteInfo: AdminSettingsGeneral;
  update: (params: AdminSettingsGeneral) => void;
}

const siteInfo = create<SiteInfoType>((set) => ({
  siteInfo: {
    name: '',
    description: '',
    short_description: '',
  },
  update: (params) =>
    set(() => {
      return {
        siteInfo: params,
      };
    }),
}));

export default siteInfo;
