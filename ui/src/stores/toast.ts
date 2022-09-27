import create from 'zustand';

type Variant = 'warning' | 'success' | 'danger';
interface ToastStore {
  msg: string;
  variant: Variant;
  show: (params: { msg: string; variant?: Variant }) => void;
  clear: () => void;
}

const toastStore = create<ToastStore>((set) => ({
  msg: '',
  variant: 'warning',
  show: (params) =>
    set((state) => {
      return {
        ...state,
        ...params,
      };
    }),
  clear: () => set({ msg: '' }),
}));

export default toastStore;
