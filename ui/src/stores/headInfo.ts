import create from 'zustand';

import { HeadInfo } from '@/common/interface';

interface HeadInfoType {
  headInfo: HeadInfo;
  update: (params: HeadInfo) => void;
}

const headInfo = create<HeadInfoType>((set) => ({
  headInfo: {
    title: '',
    description: '',
    keywords: '',
  },
  update: (params) =>
    set((state) => {
      return {
        headInfo: {
          ...state.headInfo,
          ...params,
        },
      };
    }),
}));

export default headInfo;
