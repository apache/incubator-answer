import { FC } from 'react';

import { useLegalTos } from '@/services';

const Index: FC = () => {
  const { data: tos } = useLegalTos();
  const contentText = tos?.terms_of_service_original_text;
  let matchUrl: URL | undefined;
  try {
    if (contentText) {
      matchUrl = new URL(contentText);
    }
    // eslint-disable-next-line no-empty
  } catch (ex) {}
  if (matchUrl) {
    window.location.replace(matchUrl.toString());
    return null;
  }
  return (
    <div
      className="fmt fs-14"
      dangerouslySetInnerHTML={{
        __html: tos?.terms_of_service_parsed_text || '',
      }}
    />
  );
};

export default Index;
