import create from 'zustand';

import { AdminSettingsGeneral } from '@/common/interface';
import { DEFAULT_SITE_NAME } from '@/common/constants';

interface SiteInfoType {
  siteInfo: AdminSettingsGeneral;
  version: string;
  revision: string;
  update: (params: AdminSettingsGeneral) => void;
  updateVersion: (ver: string, revision: string) => void;
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
  version: '',
  revision: '',
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
  updateVersion: (ver, revision) => {
    set(() => {
      return { version: ver, revision };
    });
  },
}));

export default siteInfo;
