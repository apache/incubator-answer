import create from 'zustand';

import type { UserInfoRes } from '@/common/interface';
import Storage from '@/utils/storage';
import {
  LOGGED_USER_STORAGE_KEY,
  LOGGED_TOKEN_STORAGE_KEY,
} from '@/common/constants';

interface UserInfoStore {
  user: UserInfoRes;
  update: (params: UserInfoRes) => void;
  clear: () => void;
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
  role_id: 1,
};

const loggedUserInfoStore = create<UserInfoStore>((set) => ({
  user: initUser,
  update: (params) => {
    if (!params?.language) {
      params.language = 'Default';
    }
    set(() => {
      Storage.set(LOGGED_TOKEN_STORAGE_KEY, params.access_token);
      Storage.set(LOGGED_USER_STORAGE_KEY, params);
      return { user: params };
    });
  },
  clear: () =>
    set(() => {
      Storage.remove(LOGGED_TOKEN_STORAGE_KEY);
      Storage.remove(LOGGED_USER_STORAGE_KEY);
      return { user: initUser };
    }),
}));

export default loggedUserInfoStore;
