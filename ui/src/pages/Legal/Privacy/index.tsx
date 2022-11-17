import { FC } from 'react';

import { useLegalPrivacy } from '@/services';

const Index: FC = () => {
  const { data: privacy } = useLegalPrivacy();
  const contentText = privacy?.privacy_policy_original_text;
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
        __html: privacy?.privacy_policy_parsed_text || '',
      }}
    />
  );
};

export default Index;
