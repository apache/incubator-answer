import create from 'zustand';

interface IType {
  custom_css: string;
  custom_head: string;
  custom_header: string;
  custom_footer: string;
  update: (params: {
    custom_css?: string;
    custom_head?: string;
    custom_header?: string;
    custom_footer?: string;
  }) => void;
}

const loginSetting = create<IType>((set) => ({
  custom_css: '',
  custom_head: '',
  custom_header: '',
  custom_footer: '',
  update: (params) =>
    set((state) => {
      return {
        ...state,
        ...params,
      };
    }),
}));

export default loginSetting;
