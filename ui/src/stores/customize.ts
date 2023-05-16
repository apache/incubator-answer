import create from 'zustand';

interface IType {
  custom_css: string;
  custom_head: string;
  custom_header: string;
  custom_footer: string;
  custom_sidebar: string;
  update: (params: {
    custom_css?: string;
    custom_head?: string;
    custom_header?: string;
    custom_footer?: string;
    custom_sidebar?: string;
  }) => void;
}

const loginSetting = create<IType>((set) => ({
  custom_css: '',
  custom_head: '',
  custom_header: '',
  custom_footer: '',
  custom_sidebar: '',
  update: (params) =>
    set((state) => {
      return {
        ...state,
        ...params,
      };
    }),
}));

export default loginSetting;
