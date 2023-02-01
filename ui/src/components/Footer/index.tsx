import React from 'react';
import { Container } from 'react-bootstrap';
import { Trans } from 'react-i18next';

import dayjs from 'dayjs';

import { siteInfoStore } from '@/stores';

const Index = () => {
  const fullYear = dayjs().format('YYYY');
  const siteName = siteInfoStore((state) => state.siteInfo.name);
  const cc = `${fullYear} ${siteName}`;
  return (
    <footer className="bg-light py-3">
      <Container>
        <p className="text-center mb-0 fs-14 text-secondary">
          <Trans i18nKey="footer.build_on" values={{ cc }}>
            Built on
            {/* eslint-disable-next-line react/jsx-no-target-blank */}
            <a href="https://answer.dev/" target="_blank">
              Answer
            </a>
            - the open-source software that powers Q&A communities.
            <br />
            Made with love. © 2022 Answer.
          </Trans>
        </p>
      </Container>
    </footer>
  );
};

export default React.memo(Index);
