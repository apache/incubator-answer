import { MouseEvent, useCallback } from 'react';

import { useLegalPrivacy, useLegalTos } from '@/services/client/legal';

export const useLegalClick = () => {
  const { data: tos } = useLegalTos();
  const { data: privacy } = useLegalPrivacy();

  const legalClick = useCallback(
    (evt: MouseEvent, type: 'tos' | 'privacy') => {
      evt.stopPropagation();
      const contentText =
        type === 'tos'
          ? tos?.terms_of_service_original_text
          : privacy?.privacy_policy_original_text;
      let matchUrl: URL | undefined;
      try {
        if (contentText) {
          matchUrl = new URL(contentText);
        }
        // eslint-disable-next-line no-empty
      } catch (ex) {}
      if (matchUrl) {
        evt.preventDefault();
        window.open(matchUrl.toString());
      }
    },
    [tos, privacy],
  );

  return legalClick;
};
