import create from 'zustand';

interface updateParams {
  logo: string;
  theme: string;
  language: string;
}

interface InterfaceType {
  interface: updateParams;
  update: (params: updateParams) => void;
}

const interfaceSetting = create<InterfaceType>((set) => ({
  interface: {
    logo: '',
    theme: '',
    language: '',
  },
  update: (params) =>
    set(() => {
      return {
        interface: params,
      };
    }),
}));

export default interfaceSetting;
