import create from 'zustand';

import { AdminSettingsTheme } from '@/common/interface';

interface IType {
  theme: AdminSettingsTheme['theme'];
  theme_config: AdminSettingsTheme['theme_config'];
  theme_options: AdminSettingsTheme['theme_options'];
  update: (params: AdminSettingsTheme) => void;
}

const store = create<IType>((set) => ({
  theme: 'default',
  theme_options: [{ label: 'Default', value: 'default' }],
  theme_config: {
    default: {
      navbar_style: 'colored',
      primary_color: '#0033FF',
    },
  },
  update: (params) =>
    set((state) => {
      return {
        ...state,
        ...params,
      };
    }),
}));

export default store;
