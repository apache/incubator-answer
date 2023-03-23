/* eslint-disable import/no-unresolved */
import { useEffect } from 'react';
import { Container } from 'react-bootstrap';

import { HttpErrorContent } from '@/components';

const Index = () => {
  useEffect(() => {
    // auto height of container
    const pageWrap = document.querySelector('.page-wrap');
    pageWrap.style.display = 'contents';

    return () => {
      pageWrap.style.display = 'block';
    };
  }, []);

  return (
    <Container
      className="d-flex flex-column justify-content-center align-items-center"
      style={{ flex: 1 }}>
      <HttpErrorContent httpCode="50X" />
    </Container>
  );
};

export default Index;
