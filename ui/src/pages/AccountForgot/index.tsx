import React, { useState, useEffect } from 'react';
import { Container, Col } from 'react-bootstrap';
import { Trans, useTranslation } from 'react-i18next';

import { isLogin } from '@answer/utils';

import SendEmail from './components/sendEmail';

import { PageTitle } from '@/components';

const Index: React.FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'account_forgot' });
  const [step, setStep] = useState(1);
  const [email, setEmail] = useState('');

  const callback = (param: number, mail: string) => {
    setStep(param);
    setEmail(mail);
  };

  useEffect(() => {
    isLogin();
  }, []);

  return (
    <>
      <PageTitle title={t('account_recovery', { keyPrefix: 'page_title' })} />
      <Container style={{ paddingTop: '4rem', paddingBottom: '6rem' }}>
        <h3 className="text-center mb-5">{t('page_title')}</h3>
        {step === 1 && (
          <Col className="mx-auto" md={3}>
            <SendEmail visible={step === 1} callback={callback} />
          </Col>
        )}
        {step === 2 && (
          <Col className="mx-auto px-4" md={6}>
            <div className="text-center">
              <p>
                <Trans
                  i18nKey="account_forgot.send_success"
                  values={{ mail: email }}
                  components={{ bold: <strong /> }}
                />
              </p>
            </div>
          </Col>
        )}
      </Container>
    </>
  );
};

export default React.memo(Index);
