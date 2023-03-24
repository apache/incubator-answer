import { useEffect } from 'react';
import { Container } from 'react-bootstrap';

import { HttpErrorContent } from '@/components';

const Index = () => {
  useEffect(() => {
    // auto height of container
    const pageWrap = document.querySelector('.page-wrap') as HTMLElement;
    pageWrap.style.display = 'contents';

    return () => {
      pageWrap.style.display = 'block';
    };
  }, []);

  return (
    <Container
      className="d-flex flex-column justify-content-center align-items-center"
      style={{ flex: 1 }}>
      <HttpErrorContent httpCode="404" />
    </Container>
  );
};

export default Index;
