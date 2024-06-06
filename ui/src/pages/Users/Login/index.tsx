/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import React, { FormEvent, useState, useEffect } from 'react';
import { Container, Form, Button, Col } from 'react-bootstrap';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';
import { Trans, useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import type { LoginReqParams, FormDataType } from '@/common/interface';
import { Unactivate, WelcomeTitle, PluginRender } from '@/components';
import {
  loggedUserInfoStore,
  loginSettingStore,
  userCenterStore,
} from '@/stores';
import {
  floppyNavigation,
  guard,
  handleFormError,
  userCenter,
  scrollToElementTop,
} from '@/utils';
import { PluginType, useCaptchaPlugin } from '@/utils/pluginKit';
import { login, UcAgent } from '@/services';
import { setupAppTheme } from '@/utils/localize';

const Index: React.FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'login' });
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const { user: storeUser, update: updateUser } = loggedUserInfoStore((_) => _);
  const loginSetting = loginSettingStore((state) => state.login);
  const ucAgent = userCenterStore().agent;
  let ucAgentInfo: UcAgent['agent_info'] | undefined;
  if (ucAgent?.enabled && ucAgent?.agent_info) {
    ucAgentInfo = ucAgent.agent_info;
  }
  const canOriginalLogin =
    (!ucAgentInfo || ucAgentInfo.enabled_original_user_system) &&
    loginSetting.allow_password_login;

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
  });

  const [step, setStep] = useState(1);

  const handleChange = (params: FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  const passwordCaptcha = useCaptchaPlugin('password');

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
    if (!bol) {
      const errObj = Object.keys(formData).filter(
        (key) => formData[key].isInvalid,
      );
      const ele = document.getElementById(errObj[0]);
      scrollToElementTop(ele);
    }

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

    const captcha = passwordCaptcha?.getCaptcha();
    if (captcha?.verify) {
      params.captcha_code = captcha.captcha_code;
      params.captcha_id = captcha.captcha_id;
    }

    login(params)
      .then(async (res) => {
        await passwordCaptcha?.close?.();
        updateUser(res);
        setupAppTheme();
        const userStat = guard.deriveLoginState();
        if (userStat.isNotActivated) {
          // inactive
          setStep(2);
        } else {
          guard.handleLoginRedirect(navigate);
        }
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          setFormData({ ...data });
          passwordCaptcha?.handleCaptchaError?.(err.list);
          const ele = document.getElementById(err.list[0].error_field);
          scrollToElementTop(ele);
        }
      });
  };

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();

    if (!checkValidated()) {
      return;
    }

    if (!passwordCaptcha) {
      handleLogin();
      return;
    }

    passwordCaptcha?.check?.(() => {
      handleLogin();
    });
  };

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
          <PluginRender
            type={PluginType.Captcha}
            slug_name="captcha_basic"
            className="mb-5"
          />

          <PluginRender
            type={PluginType.Captcha}
            slug_name="captcha_google_v2"
            className="mb-5"
          />
          {ucAgentInfo ? (
            <PluginRender
              type={PluginType.Connector}
              slug_name="hosting_connector"
              className="mb-5"
            />
          ) : (
            <PluginRender
              type={PluginType.Connector}
              slug_name="third_party_connector"
              className="mb-5"
            />
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

                <Form.Group controlId="pass" className="mb-3">
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
                    Don't have an account?
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
    </Container>
  );
};

export default React.memo(Index);
