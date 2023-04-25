import create from 'zustand';

import type { UserInfoRes } from '@/common/interface';
import Storage from '@/utils/storage';
import { LOGGED_TOKEN_STORAGE_KEY } from '@/common/constants';

interface UserInfoStore {
  user: UserInfoRes;
  update: (params: UserInfoRes) => void;
  clear: (removeToken?: boolean) => void;
}

const initUser: UserInfoRes = {
  access_token: '',
  username: '',
  avatar: '',
  rank: 0,
  bio: '',
  bio_html: '',
  display_name: '',
  location: '',
  website: '',
  status: '',
  mail_status: 1,
  language: 'Default',
  is_admin: false,
  have_password: true,
  role_id: 1,
};

const loggedUserInfo = create<UserInfoStore>((set) => ({
  user: initUser,
  update: (params) => {
    if (typeof params !== 'object' || !params) {
      return;
    }
    if (!params?.language) {
      params.language = 'Default';
    }
    set(() => {
      Storage.set(LOGGED_TOKEN_STORAGE_KEY, params.access_token);
      return { user: params };
    });
  },
  clear: (removeToken = true) =>
    set(() => {
      if (removeToken) {
        Storage.remove(LOGGED_TOKEN_STORAGE_KEY);
      }
      return { user: initUser };
    }),
}));

export default loggedUserInfo;
