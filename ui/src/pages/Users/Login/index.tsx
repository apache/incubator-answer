import React, { FormEvent, useState, useEffect } from 'react';
import { Container, Form, Button, Col } from 'react-bootstrap';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';
import { Trans, useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import type {
  LoginReqParams,
  ImgCodeRes,
  FormDataType,
} from '@/common/interface';
import { Unactivate, WelcomeTitle, PluginRender } from '@/components';
import {
  loggedUserInfoStore,
  loginSettingStore,
  userCenterStore,
} from '@/stores';
import { floppyNavigation, guard, handleFormError, userCenter } from '@/utils';
import { login, checkImgCode, UcAgent } from '@/services';
import { PicAuthCodeModal } from '@/components/Modal';

const Index: React.FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'login' });
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const [refresh, setRefresh] = useState(0);
  const { user: storeUser, update: updateUser } = loggedUserInfoStore((_) => _);
  const loginSetting = loginSettingStore((state) => state.login);
  const ucAgent = userCenterStore().agent;
  let ucAgentInfo: UcAgent['agent_info'] | undefined;
  if (ucAgent?.enabled && ucAgent?.agent_info) {
    ucAgentInfo = ucAgent.agent_info;
  }
  const canOriginalLogin =
    !ucAgentInfo || ucAgentInfo.enabled_original_user_system;

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
    if (!canOriginalLogin) {
      return;
    }
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
        const userStat = guard.deriveLoginState();
        if (userStat.isNotActivated) {
          // inactive
          setStep(2);
          setRefresh((pre) => pre + 1);
        } else {
          guard.handleLoginRedirect(navigate);
        }

        setModalState(false);
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          if (!err.list.find((v) => v.error_field.indexOf('captcha') >= 0)) {
            setModalState(false);
          }
          setFormData({ ...data });
        }
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
    const isInactive = searchParams.get('status');

    if (storeUser.id && (storeUser.mail_status === 2 || isInactive)) {
      setStep(2);
    }
  }, []);
  usePageTags({
    title: t('login', { keyPrefix: 'page_title' }),
  });

  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '5rem' }}>
      <WelcomeTitle />
      {step === 1 ? (
        <Col className="mx-auto" md={6} lg={4} xl={3}>
          {ucAgentInfo ? (
            <PluginRender slug_name="uc_login" className="mb-5" />
          ) : (
            <PluginRender type="Connector" className="mb-5" />
          )}
          {canOriginalLogin ? (
            <>
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

                <div className="d-grid">
                  <Button variant="primary" type="submit" tabIndex={1}>
                    {t('login', { keyPrefix: 'btns' })}
                  </Button>
                </div>
              </Form>
              {loginSetting.allow_new_registrations && (
                <div className="text-center mt-5">
                  <Trans i18nKey="login.info_sign" ns="translation">
                    Don’t have an account?
                    <Link
                      to={userCenter.getSignUpUrl()}
                      tabIndex={2}
                      onClick={floppyNavigation.handleRouteLinkClick}>
                      Sign up
                    </Link>
                  </Trans>
                </div>
              )}
            </>
          ) : null}
        </Col>
      ) : null}

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
