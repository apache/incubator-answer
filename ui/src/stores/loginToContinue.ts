import create from 'zustand';

interface IProps {
  show: boolean;
  update: (params: { show: boolean }) => void;
}

const loginToContinueStore = create<IProps>((set) => ({
  show: false,
  update: (params) =>
    set({
      ...params,
    }),
}));

export default loginToContinueStore;
