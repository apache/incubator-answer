import create from 'zustand';

import type { UcAgent } from '@/services/user-center';

interface UserCenterStore {
  agent?: UcAgent;
  update: (uca: UcAgent) => void;
}

const store = create<UserCenterStore>((set) => ({
  agent: undefined,
  update: (uca: UcAgent) => {
    if (uca) {
      set({
        agent: uca,
      });
    }
  },
}));

export default store;
