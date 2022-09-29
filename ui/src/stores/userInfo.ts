import create from 'zustand';

import type { UserInfoRes } from '@answer/common/interface';
import Storage from '@answer/utils/storage';

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
  mail_status: 0,
};

const userInfoStore = create<UserInfoStore>((set) => ({
  user: initUser,
  update: (params) =>
    set(() => {
      Storage.set('token', params.access_token);
      Storage.set('userInfo', params);
      return { user: params };
    }),
  clear: () =>
    set(() => {
      // Storage.remove('token');
      Storage.remove('userInfo');
      return { user: initUser };
    }),
}));

export default userInfoStore;
