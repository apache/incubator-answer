import { FC } from 'react';

import { useLegalTos } from '@/services';

const Index: FC = () => {
  const { data: tos } = useLegalTos();
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
