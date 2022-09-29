import React, { FormEvent, useState, useEffect } from 'react';
import { Container, Form, Button, Col } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { Trans, useTranslation } from 'react-i18next';

import { login, checkImgCode } from '@answer/api';
import { PageTitle, Unactivate } from '@answer/components';
import { userInfoStore } from '@answer/stores';
import { isLogin, getQueryString } from '@answer/utils';

import { PicAuthCodeModal } from '@/components/Modal';
import type { LoginReqParams, ImgCodeRes } from '@/services/types';
import type { FormDataType } from '@/common/interface';
import Storage from '@/utils/storage';

const Index: React.FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'login' });
  const [refresh, setRefresh] = useState(0);
  const updateUser = userInfoStore((state) => state.update);
  const storeUser = userInfoStore((state) => state.user);
  const [formData, setFormData] = useState<FormDataType>({
    e_mail: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    pass: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    captcha_code: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });
  const [imgCode, setImgCode] = useState<ImgCodeRes>({
    captcha_id: '',
    captcha_img: '',
    verify: false,
  });
  const [showModal, setModalState] = useState(false);
  const [step, setStep] = useState(1);

  const handleChange = (params: FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  const getImgCode = () => {
    checkImgCode({
      action: 'login',
    }).then((res) => {
      setImgCode(res);
    });
  };

  const checkValidated = (): boolean => {
    let bol = true;
    const { e_mail, pass } = formData;

    if (!e_mail.value) {
      bol = false;
      formData.e_mail = {
        value: '',
        isInvalid: true,
        errorMsg: t('email.msg.empty'),
      };
    }

    if (!pass.value) {
      bol = false;
      formData.pass = {
        value: '',
        isInvalid: true,
        errorMsg: t('password.msg.empty'),
      };
    }

    setFormData({
      ...formData,
    });
    return bol;
  };

  const handleLogin = (event?: any) => {
    if (event) {
      event.preventDefault();
    }
    const params: LoginReqParams = {
      e_mail: formData.e_mail.value,
      pass: formData.pass.value,
    };
    if (imgCode.verify) {
      params.captcha_code = formData.captcha_code.value;
      params.captcha_id = imgCode.captcha_id;
    }

    login(params)
      .then((res) => {
        updateUser(res);
        if (res.mail_status === 2) {
          // inactive
          setStep(2);
          setRefresh((pre) => pre + 1);
        }
        if (res.mail_status === 1) {
          const path = Storage.get('ANSWER_PATH');
          Storage.remove('ANSWER_PATH');
          if (path) {
            window.location.href = path;
          } else {
            window.location.href = '/';
          }
        }

        setModalState(false);
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData[err.key].isInvalid = true;
          formData[err.key].errorMsg = err.value;
          if (err.key.indexOf('captcha') < 0) {
            setModalState(false);
          }
        }
        setFormData({ ...formData });
        setRefresh((pre) => pre + 1);
      });
  };

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();

    if (!checkValidated()) {
      return;
    }

    if (imgCode.verify) {
      setModalState(true);
      return;
    }

    handleLogin();
  };

  useEffect(() => {
    getImgCode();
  }, [refresh]);

  useEffect(() => {
    const isInactive = getQueryString('status');

    if ((storeUser.id && storeUser.mail_status === 2) || isInactive) {
      setStep(2);
    } else {
      isLogin();
    }
  }, []);

  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '5rem' }}>
      <h3 className="text-center mb-5">{t('page_title')}</h3>
      <PageTitle title={t('login', { keyPrefix: 'page_title' })} />
      {step === 1 && (
        <Col className="mx-auto" md={3}>
          <Form noValidate onSubmit={handleSubmit}>
            <Form.Group controlId="email" className="mb-3">
              <Form.Label>{t('email.label')}</Form.Label>
              <Form.Control
                required
                tabIndex={1}
                type="email"
                value={formData.e_mail.value}
                isInvalid={formData.e_mail.isInvalid}
                onChange={(e) =>
                  handleChange({
                    e_mail: {
                      value: e.target.value,
                      isInvalid: false,
                      errorMsg: '',
                    },
                  })
                }
              />
              <Form.Control.Feedback type="invalid">
                {formData.e_mail.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>

            <Form.Group controlId="password" className="mb-3">
              <div className="d-flex justify-content-between">
                <Form.Label>{t('password.label')}</Form.Label>
                <Link to="/users/account-recovery" tabIndex={2}>
                  <small>{t('forgot_pass')}</small>
                </Link>
              </div>

              <Form.Control
                required
                tabIndex={1}
                type="password"
                // value={formData.pass.value}
                maxLength={32}
                isInvalid={formData.pass.isInvalid}
                onChange={(e) =>
                  handleChange({
                    pass: {
                      value: e.target.value,
                      isInvalid: false,
                      errorMsg: '',
                    },
                  })
                }
              />
              <Form.Control.Feedback type="invalid">
                {formData.pass.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>

            <div className="d-grid mb-3">
              <Button variant="primary" type="submit" tabIndex={1}>
                {t('login', { keyPrefix: 'btns' })}
              </Button>
            </div>
          </Form>

          <div className="text-center">
            <Trans i18nKey="login.info_sign" ns="translation">
              Don’t have an account?
              <Link to="/users/register" tabIndex={2}>
                Sign up
              </Link>
            </Trans>
          </div>
        </Col>
      )}

      {step === 2 && <Unactivate visible={step === 2} />}

      <PicAuthCodeModal
        visible={showModal}
        data={{
          captcha: formData.captcha_code,
          imgCode,
        }}
        handleCaptcha={handleChange}
        clickSubmit={handleLogin}
        refreshImgCode={getImgCode}
        onClose={() => setModalState(false)}
      />
    </Container>
  );
};

export default React.memo(Index);
