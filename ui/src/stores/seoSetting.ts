import create from 'zustand';

import { AdminSettingsSeo } from '@/common/interface';

interface IProps {
  seo: AdminSettingsSeo;
  update: (params: AdminSettingsSeo) => void;
}

const Index = create<IProps>((set) => ({
  seo: {
    robots: '',
    permalink: 1,
  },
  update: (params) =>
    set((state) => {
      const o = { ...state.seo, ...params };
      // @ts-ignore
      if (!/[1234]/.test(o.permalink)) {
        o.permalink = 1;
      }
      return {
        seo: o,
      };
    }),
}));

export default Index;
