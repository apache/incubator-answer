import create from 'zustand';

import { AdminSettingsGeneral } from '@/common/interface';
import { DEFAULT_SITE_NAME } from '@/common/constants';

interface SiteInfoType {
  siteInfo: AdminSettingsGeneral;
  update: (params: AdminSettingsGeneral) => void;
}

const siteInfo = create<SiteInfoType>((set) => ({
  siteInfo: {
    name: DEFAULT_SITE_NAME,
    description: '',
    short_description: '',
    site_url: '',
    contact_email: '',
    permalink: 1,
  },
  update: (params) =>
    set((_) => {
      const o = { ..._.siteInfo, ...params };
      if (!o.name) {
        o.name = DEFAULT_SITE_NAME;
      }
      return {
        siteInfo: o,
      };
    }),
}));

export default siteInfo;
