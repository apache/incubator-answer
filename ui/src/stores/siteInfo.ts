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
    site_url: '',
    contact_email: '',
    permalink: true,
  },
  update: (params) =>
    set((_) => {
      const o = { ..._.siteInfo, ...params };
      return {
        siteInfo: o,
      };
    }),
}));

export default siteInfo;
