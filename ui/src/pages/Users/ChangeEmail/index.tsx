import { FC, memo } from 'react';
import { Container, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { WelcomeTitle } from '@/components';

import SendEmail from './components/sendEmail';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'change_email' });
  usePageTags({
    title: t('change_email', { keyPrefix: 'page_title' }),
  });
  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '6rem' }}>
      <WelcomeTitle />
      <Col className="mx-auto" md={6} lg={4} xl={3}>
        <SendEmail />
      </Col>
    </Container>
  );
};

export default memo(Index);
