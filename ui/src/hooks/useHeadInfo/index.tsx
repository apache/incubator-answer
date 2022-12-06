import { useEffect } from 'react';

import { HeadInfo } from '@/common/interface';
import { headInfoStore } from '@/stores';

export default function useHeadInfo(info: HeadInfo) {
  const { update } = headInfoStore.getState();

  useEffect(() => {
    update(info);
  }, [info]);
}
