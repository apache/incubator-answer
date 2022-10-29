import create from 'zustand';

import type { UserInfoRes } from '@answer/common/interface';
import Storage from '@answer/utils/storage';

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
};

const loggedUserInfoStore = create<UserInfoStore>((set) => ({
  user: initUser,
  update: (params) =>
    set(() => {
      Storage.set(LOGGED_TOKEN_STORAGE_KEY, params.access_token);
      Storage.set(LOGGED_USER_STORAGE_KEY, params);
      return { user: params };
    }),
  clear: () =>
    set(() => {
      Storage.remove(LOGGED_TOKEN_STORAGE_KEY);
      Storage.remove(LOGGED_USER_STORAGE_KEY);
      return { user: initUser };
    }),
}));

export default loggedUserInfoStore;
