import React from 'react';
import { Container } from 'react-bootstrap';

const Index = () => {
  return (
    <footer className="bg-light py-3">
      <Container>
        <p className="text-center mb-2 fs-14">
          Built on
          <a href="/"> Answer </a>
          - the open source source software that power knowledge communities.
          <br />
          Made with love. © 2022 Answer .
        </p>
      </Container>
    </footer>
  );
};

export default React.memo(Index);
