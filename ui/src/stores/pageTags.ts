import create from 'zustand';

import { HelmetBase, HelmetUpdate } from '@/common/interface';

import siteInfoStore from './siteInfo';

interface HelmetStore {
  items: HelmetBase;
  update: (params: HelmetUpdate) => void;
}

const makePageTitle = (title = '', subtitle = '') => {
  const { siteInfo } = siteInfoStore.getState();
  if (!subtitle) {
    subtitle = `${siteInfo.name}`;
  }
  let pageTitle = subtitle;
  if (title && title !== subtitle) {
    pageTitle = `${title}${subtitle ? ` - ${subtitle}` : ''}`;
  }
  return pageTitle;
};

const pageTags = create<HelmetStore>((set) => ({
  items: {
    pageTitle: '',
    description: '',
    keywords: '',
  },
  update: (params) => {
    const o: HelmetBase = {};
    if (params.title || params.subtitle) {
      o.pageTitle = makePageTitle(params.title, params.subtitle);
    }
    o.description =
      params.description ||
      siteInfoStore.getState().siteInfo?.description ||
      '';
    o.keywords = params.keywords || '';

    set({
      items: o,
    });
  },
}));

export default pageTags;
