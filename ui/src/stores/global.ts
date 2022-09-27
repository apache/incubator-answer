import create from 'zustand';

interface GlobalStoreType {
  title: string;
  description: string;
  keywords: string;
  updateTitle: (title: string) => void;
  updateSeo: (params: { description: string; keywords: string }) => void;
}

const globalStore = create<GlobalStoreType>((set) => ({
  title: '',
  description: '',
  keywords: '',
  updateTitle: (params) =>
    set((state) => {
      return {
        ...state,
        title: params,
      };
    }),
  updateSeo: (params) =>
    set((state) => {
      return {
        ...state,
        ...params,
      };
    }),
}));

export default globalStore;
