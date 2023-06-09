import React, { useState } from 'react';
import { Container, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { Unactivate, WelcomeTitle, PluginRender } from '@/components';
import { guard } from '@/utils';

import SignUpForm from './components/SignUpForm';

const Index: React.FC = () => {
  const [showForm, setShowForm] = useState(true);
  const { t } = useTranslation('translation', { keyPrefix: 'login' });
  const onStep = () => {
    setShowForm((bol) => !bol);
  };
  usePageTags({
    title: t('sign_up', { keyPrefix: 'page_title' }),
  });
  if (!guard.singUpAgent().ok) {
    return null;
  }
  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '5rem' }}>
      <WelcomeTitle />

      {showForm ? (
        <Col className="mx-auto" md={6} lg={4} xl={3}>
          <PluginRender type="Connector" className="mb-5" />
          <SignUpForm callback={onStep} />
        </Col>
      ) : (
        <Unactivate visible={!showForm} />
      )}
    </Container>
  );
};

export default React.memo(Index);
