import create from 'zustand';

interface NotFoundType {
  visible: boolean;
  show: () => void;
  hide: () => void;
}

const notFound = create<NotFoundType>((set) => ({
  visible: false,
  show: () => {
    set(() => {
      return { visible: true };
    });
  },
  hide: () => {
    set(() => {
      return { visible: false };
    });
  },
}));

export default notFound;
