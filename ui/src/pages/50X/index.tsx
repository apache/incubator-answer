import { useSearchParams } from 'react-router-dom';

import { HttpErrorContent } from '@/components';

const Index = () => {
  const [searchParams] = useSearchParams();
  const errMsg = searchParams.get('msg') || '';
  const title = searchParams.get('title') || '';
  return (
    <HttpErrorContent
      httpCode="50X"
      title={title}
      errMsg={errMsg}
      showErrorCode={!errMsg}
    />
  );
};

export default Index;
