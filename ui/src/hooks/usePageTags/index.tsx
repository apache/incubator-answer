import { useEffect } from 'react';

import { HelmetUpdate } from '@/common/interface';
import { pageTagStore } from '@/stores';

export default function usePageTags(info: HelmetUpdate) {
  const { update } = pageTagStore.getState();
  useEffect(() => {
    update(info);
  }, [info.title, info.subtitle, info.description, info.keywords]);
}
