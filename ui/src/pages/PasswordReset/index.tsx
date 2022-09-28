import React, { FormEvent, useState, useEffect } from 'react';
import { Container, Col, Form, Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import type { FormDataType } from '@answer/common/interface';
import { passRetrieveSet } from '@answer/api';
import { userInfoStore } from '@answer/stores';
import Storage from '@answer/utils/storage';
import { getQueryString, isLogin } from '@answer/utils';

import { PageTitle } from '@/components';

const Index: React.FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'password_reset' });

  const [step, setStep] = useState(1);
  const clearUser = userInfoStore((state) => state.clear);
  const [formData, setFormData] = useState<FormDataType>({
    pass: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    passSecond: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });

  const handleChange = (params: FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  const checkValidated = (): boolean => {
    let bol = true;
    const { pass, passSecond } = formData;

    if (!pass.value) {
      bol = false;
      formData.pass = {
        value: '',
        isInvalid: true,
        errorMsg: t('password.msg.empty'),
      };
    }

    if (bol && pass.value && pass.value.length < 8) {
      bol = false;
      formData.pass = {
        value: pass.value,
        isInvalid: true,
        errorMsg: t('password.msg.length'),
      };
    }

    if (bol && !passSecond.value) {
      bol = false;
      formData.passSecond = {
        value: '',
        isInvalid: true,
        errorMsg: t('password.msg.empty'),
      };
    }

    if (bol && passSecond.value && passSecond.value.length < 8) {
      bol = false;
      formData.passSecond = {
        value: passSecond.value,
        isInvalid: true,
        errorMsg: t('password.msg.length'),
      };
    }

    if (bol && pass.value !== passSecond.value) {
      bol = false;
      formData.passSecond = {
        value: passSecond.value,
        isInvalid: true,
        errorMsg: t('password.msg.different'),
      };
    }
    setFormData({
      ...formData,
    });
    return bol;
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    if (checkValidated() === false) {
      return;
    }
    const code = getQueryString('code');
    if (!code) {
      console.error('code is required');
      return;
    }
    passRetrieveSet({
      code: encodeURIComponent(code),
      pass: formData.pass.value,
    })
      .then(() => {
        // clear login information then to login page
        clearUser();
        Storage.remove('token');
        setStep(2);
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData[err.key].isInvalid = true;
          formData[err.key].errorMsg = err.value;
        }
        setFormData({ ...formData });
      });
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
            <Form noValidate onSubmit={handleSubmit} autoComplete="off">
              <Form.Group controlId="email" className="mb-3">
                <Form.Label>{t('password.label')}</Form.Label>
                <Form.Control
                  autoComplete="off"
                  required
                  type="password"
                  maxLength={32}
                  isInvalid={formData.pass.isInvalid}
                  onChange={(e) => {
                    handleChange({
                      pass: {
                        value: e.target.value,
                        isInvalid: false,
                        errorMsg: '',
                      },
                    });
                  }}
                />
                <Form.Control.Feedback type="invalid">
                  {formData.pass.errorMsg}
                </Form.Control.Feedback>
              </Form.Group>

              <Form.Group controlId="password" className="mb-3">
                <Form.Label>{t('password_confirm.label')}</Form.Label>
                <Form.Control
                  autoComplete="off"
                  required
                  type="password"
                  maxLength={32}
                  isInvalid={formData.passSecond.isInvalid}
                  onChange={(e) => {
                    handleChange({
                      passSecond: {
                        value: e.target.value,
                        isInvalid: false,
                        errorMsg: '',
                      },
                    });
                  }}
                />
                <Form.Control.Feedback type="invalid">
                  {formData.passSecond.errorMsg}
                </Form.Control.Feedback>
              </Form.Group>

              <div className="d-grid mb-3">
                <Button variant="primary" type="submit">
                  {t('btn_name')}
                </Button>
              </div>
            </Form>
          </Col>
        )}

        {step === 2 && (
          <Col className="mx-auto px-4" md={6}>
            <div className="text-center">
              <p>{t('reset_success')}</p>
              <Link to="/users/login">{t('to_login')}</Link>
            </div>
          </Col>
        )}

        {step === 3 && (
          <Col className="mx-auto px-4" md={6}>
            <div className="text-center">
              <p>{t('link_invalid')}</p>
              <Link to="/users/login">{t('to_login')}</Link>
            </div>
          </Col>
        )}
      </Container>
    </>
  );
};

export default React.memo(Index);
