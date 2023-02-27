import React, { useState } from 'react';
import { Container } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { Unactivate, WelcomeTitle } from '@/components';
import { PluginOauth } from '@/plugins';

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
  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '5rem' }}>
      <WelcomeTitle />
      <PluginOauth className="mb-5" />
      {showForm ? (
        <SignUpForm callback={onStep} />
      ) : (
        <Unactivate visible={!showForm} />
      )}
    </Container>
  );
};

export default React.memo(Index);
