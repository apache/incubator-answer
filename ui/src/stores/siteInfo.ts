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
    permalink: 1,
  },
  update: (params) =>
    set((_) => {
      const o = { ..._.siteInfo, ...params };
      if (o.permalink !== 1 && o.permalink !== 2) {
        o.permalink = 1;
      }
      return {
        siteInfo: o,
      };
    }),
}));

export default siteInfo;
