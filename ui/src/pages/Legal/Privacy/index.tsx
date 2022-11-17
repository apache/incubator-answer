import { FC } from 'react';

import { useLegalPrivacy } from '@/services';

const Index: FC = () => {
  const { data: privacy } = useLegalPrivacy();
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
