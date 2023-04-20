import { FC, memo, useEffect, useState } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Link, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { loggedUserInfoStore } from '@/stores';
import { changeEmailVerify } from '@/services';
import { WelcomeTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'account_result' });
  const [searchParams] = useSearchParams();
  const [step, setStep] = useState('loading');

  const updateUser = loggedUserInfoStore((state) => state.update);

  useEffect(() => {
    const code = searchParams.get('code');
    if (code) {
      // do
      changeEmailVerify({ code })
        .then((res) => {
          setStep('success');
          if (res?.access_token) {
            // update user info
            updateUser(res);
          }
        })
        .catch(() => {
          setStep('invalid');
        });
    }
  }, []);
  usePageTags({
    title: t('confirm_email', { keyPrefix: 'page_title' }),
  });
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col lg={6}>
          <WelcomeTitle className="mt-3 mb-5" />
          {step === 'success' && (
            <>
              <p className="text-center">{t('confirm_new_email')}</p>
              <div className="text-center">
                <Link to="/">{t('link')}</Link>
              </div>
            </>
          )}

          {step === 'invalid' && (
            <p className="text-center">{t('confirm_new_email_invalid')}</p>
          )}
        </Col>
      </Row>
    </Container>
  );
};

export default memo(Index);
