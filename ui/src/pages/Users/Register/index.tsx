import React, { useState } from 'react';
import { Container } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { PageTitle, Unactivate } from '@/components';

import SignUpForm from './components/SignUpForm';

const Index: React.FC = () => {
  const [showForm, setShowForm] = useState(true);
  const { t } = useTranslation('translation', { keyPrefix: 'login' });

  const onStep = () => {
    setShowForm((bol) => !bol);
  };

  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '5rem' }}>
      <h3 className="text-center mb-5">{t('page_title')}</h3>
      <PageTitle title={t('sign_up', { keyPrefix: 'page_title' })} />
      {showForm ? (
        <SignUpForm callback={onStep} />
      ) : (
        <Unactivate visible={!showForm} />
      )}
    </Container>
  );
};

export default React.memo(Index);
