import create from 'zustand';

interface CommentReplyType {
  id: string | number;
  update: (id) => void;
}

const Index = create<CommentReplyType>((set) => ({
  id: '',
  update: (id) => {
    set(() => {
      return { id };
    });
  },
}));

export default Index;
