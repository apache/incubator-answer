import React, { useState } from 'react';
import { Container, Col } from 'react-bootstrap';
import { Trans, useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import { usePageTags } from '@/hooks';
import { Unactivate, WelcomeTitle, PluginRender } from '@/components';
import { guard } from '@/utils';
import { loginSettingStore } from '@/stores';

import SignUpForm from './components/SignUpForm';

const Index: React.FC = () => {
  const [showForm, setShowForm] = useState(true);
  const { t } = useTranslation('translation', { keyPrefix: 'login' });
  const loginSetting = loginSettingStore((state) => state.login);
  const onStep = () => {
    setShowForm((bol) => !bol);
  };
  usePageTags({
    title: t('sign_up', { keyPrefix: 'page_title' }),
  });

  if (!guard.singUpAgent().ok) {
    return null;
  }

  const showSignupForm =
    loginSetting?.allow_new_registrations &&
    loginSetting.allow_email_registrations;

  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '5rem' }}>
      <WelcomeTitle />

      {showForm ? (
        <Col className="mx-auto" md={6} lg={4} xl={3}>
          <PluginRender type="Connector" className="mb-5" />
          {showSignupForm ? <SignUpForm callback={onStep} /> : null}
          <div className="text-center mt-5">
            <Trans i18nKey="login.info_login" ns="translation">
              Already have an account? <Link to="/users/login">Log in</Link>
            </Trans>
          </div>
        </Col>
      ) : (
        <Unactivate visible={!showForm} />
      )}
    </Container>
  );
};

export default React.memo(Index);
