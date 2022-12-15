import create from 'zustand';

import { AdminSettingsSeo } from '@/common/interface';

interface IProps {
  seo: AdminSettingsSeo;
  update: (params: AdminSettingsSeo) => void;
}

const siteInfo = create<IProps>((set) => ({
  seo: {
    robots: '',
    permalink: 1,
  },
  update: (params) =>
    set((state) => {
      const o = { ...state.seo, ...params };
      if (o.permalink !== 1 && o.permalink !== 2) {
        o.permalink = 1;
      }
      return {
        seo: o,
      };
    }),
}));

export default siteInfo;
