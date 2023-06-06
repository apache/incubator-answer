import create from 'zustand';

type reviewData = {
  can_revision: boolean;
  revision: number;
};

interface ErrorCodeType {
  visible: boolean;
  can_revision: boolean;
  revision: number;
  updateVisible: () => void;
  updateReiview: (params: reviewData) => void;
}

const Index = create<ErrorCodeType>((set) => ({
  visible: false,
  can_revision: false,
  revision: 0,
  updateVisible: () => {
    set((state) => {
      return { visible: !state.visible };
    });
  },
  updateReiview: (params: reviewData) => {
    set(() => {
      return { ...params };
    });
  },
}));

export default Index;
