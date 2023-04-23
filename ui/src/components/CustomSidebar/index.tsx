import { memo } from 'react';

import { customizeStore } from '@/stores';

const Index = () => {
  const { custom_sidebar } = customizeStore((state) => state);
  if (!custom_sidebar) return null;
  return <div dangerouslySetInnerHTML={{ __html: custom_sidebar }} />;
};

export default memo(Index);
