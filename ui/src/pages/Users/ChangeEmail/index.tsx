import { FC, memo } from 'react';
import { Container, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';

import SendEmail from './components/sendEmail';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'change_email' });
  usePageTags({
    title: t('change_email', { keyPrefix: 'page_title' }),
  });
  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '6rem' }}>
      <h3 className="text-center mb-5">{t('page_title')}</h3>
      <Col className="mx-auto" md={3}>
        <SendEmail />
      </Col>
    </Container>
  );
};

export default memo(Index);
