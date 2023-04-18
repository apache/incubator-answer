import { useSearchParams } from 'react-router-dom';

import { HttpErrorContent } from '@/components';

const Index = () => {
  const [searchParams] = useSearchParams();
  const msg = searchParams.get('msg') || '';
  return <HttpErrorContent httpCode="50X" errMsg={msg} showErrorCode={!msg} />;
};

export default Index;
