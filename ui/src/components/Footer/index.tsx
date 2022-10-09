import React from 'react';
import { Container } from 'react-bootstrap';
import { Trans } from 'react-i18next';

const Index = () => {
  return (
    <footer className="bg-light py-3">
      <Container>
        <p className="text-center mb-0 fs-14 text-secondary">
          <Trans i18nKey="footer.build_on">
            Built on
            <a href="/"> Answer </a>
            - the open source source software that power knowledge communities.
            <br />
            Made with love. © 2022 Answer .
          </Trans>
        </p>
      </Container>
    </footer>
  );
};

export default React.memo(Index);
