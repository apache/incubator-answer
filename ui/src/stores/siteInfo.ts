import create from 'zustand';

import { AdminSettingsGeneral, AdminSettingsUsers } from '@/common/interface';
import { DEFAULT_SITE_NAME } from '@/common/constants';

interface SiteInfoType {
  siteInfo: AdminSettingsGeneral;
  version: string;
  revision: string;
  update: (params: AdminSettingsGeneral) => void;
  updateVersion: (ver: string, revision: string) => void;
  users: AdminSettingsUsers;
  updateUsers: (users: SiteInfoType['users']) => void;
}

const defaultUsersConf: AdminSettingsUsers = {
  allow_update_avatar: false,
  allow_update_bio: false,
  allow_update_display_name: false,
  allow_update_location: false,
  allow_update_username: false,
  allow_update_website: false,
  default_avatar: 'system',
  gravatar_base_url: 'https://www.gravatar.com/avatar/',
};

const siteInfo = create<SiteInfoType>((set) => ({
  siteInfo: {
    name: DEFAULT_SITE_NAME,
    description: '',
    short_description: '',
    site_url: '',
    contact_email: '',
    permalink: 1,
  },
  users: defaultUsersConf,
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
  updateUsers: (users) => {
    set(() => {
      users ||= defaultUsersConf;
      return { users };
    });
  },
}));

export default siteInfo;
