import create from 'zustand';

type codeType = '403' | '404' | '50X' | '';

interface NotFoundType {
  code: codeType;
  update: (code: codeType) => void;
  reset: () => void;
}

const notFound = create<NotFoundType>((set) => ({
  code: '',
  update: (code: codeType) => {
    set(() => {
      return { code };
    });
  },
  reset: () => {
    set(() => {
      return { code: '' };
    });
  },
}));

export default notFound;
