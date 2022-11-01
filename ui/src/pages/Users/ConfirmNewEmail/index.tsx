import { FC, memo, useEffect, useState } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Link, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { changeEmailVerify, getUserInfo } from '@answer/api';
import { userInfoStore } from '@answer/stores';

import { PageTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'account_result' });
  const [searchParams] = useSearchParams();
  const [step, setStep] = useState('loading');

  const updateUser = userInfoStore((state) => state.update);

  useEffect(() => {
    const code = searchParams.get('code');
    if (code) {
      // do
      changeEmailVerify({ code })
        .then(() => {
          setStep('success');
          getUserInfo().then((res) => {
            // update user info
            updateUser(res);
          });
        })
        .catch(() => {
          setStep('invalid');
        });
    }
  }, []);

  return (
    <>
      <PageTitle title={t('confirm_email', { keyPrefix: 'page_title' })} />
      <Container className="pt-4 mt-2 mb-5">
        <Row className="justify-content-center">
          <Col lg={6}>
            <h3 className="text-center mt-3 mb-5">{t('page_title')}</h3>
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
    </>
  );
};

export default memo(Index);
