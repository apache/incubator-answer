import create from 'zustand';

interface updateParams {
  name: string;
  description: string;
  short_description: string;
}

interface SiteInfoType {
  siteInfo: updateParams;
  update: (params: updateParams) => void;
}

const siteInfo = create<SiteInfoType>((set) => ({
  siteInfo: {
    name: '',
    description: '',
    short_description: '',
  },
  update: (params) =>
    set(() => {
      return {
        siteInfo: params,
      };
    }),
}));

export default siteInfo;
