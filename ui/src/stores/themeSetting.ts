import create from 'zustand';

import { AdminSettingsTheme } from '@/common/interface';

interface IType {
  theme: AdminSettingsTheme['theme'];
  theme_config: AdminSettingsTheme['theme_config'];
  update: (params: AdminSettingsTheme) => void;
}

const store = create<IType>((set) => ({
  theme: '',
  theme_config: {},
  update: (params) =>
    set((state) => {
      return {
        ...state,
        ...params,
      };
    }),
}));

export default store;
