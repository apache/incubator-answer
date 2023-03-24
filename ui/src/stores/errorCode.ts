import create from 'zustand';

type codeType = '403' | '404' | '50X' | '';

interface ErrorCodeType {
  code: codeType;
  msg: string;
  update: (code: codeType, msg?: string) => void;
  reset: () => void;
}

const Index = create<ErrorCodeType>((set) => ({
  code: '',
  msg: '',
  update: (code: codeType, msg: string = '') => {
    set(() => {
      return { code, msg };
    });
  },
  reset: () => {
    set(() => {
      return { code: '', msg: '' };
    });
  },
}));

export default Index;
