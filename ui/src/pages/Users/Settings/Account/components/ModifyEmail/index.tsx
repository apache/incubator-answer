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

import React, { FC, FormEvent, useEffect, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { useToast } from '@/hooks';
import { useCaptchaPlugin } from '@/utils/pluginKit';
import { getLoggedUserInfo, changeEmail } from '@/services';
import { handleFormError, scrollToElementTop } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.account',
  });
  const [step, setStep] = useState(1);
  const [formData, setFormData] = useState<Type.FormDataType>({
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
  const [userInfo, setUserInfo] = useState<Type.UserInfoRes>();
  const toast = useToast();
  const emailCaptcha = useCaptchaPlugin('edit_userinfo');

  useEffect(() => {
    getLoggedUserInfo().then((resp) => {
      setUserInfo(resp);
    });
  }, []);

  const handleChange = (params: Type.FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  const checkValidated = (): boolean => {
    let bol = true;
    const { e_mail, pass } = formData;

    if (!e_mail.value) {
      bol = false;
      formData.e_mail = {
        value: '',
        isInvalid: true,
        errorMsg: t('new_email.msg'),
      };
    }

    if (!pass.value) {
      bol = false;
      formData.pass = {
        value: '',
        isInvalid: true,
        errorMsg: t('pass.msg'),
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

  const initFormData = () => {
    setFormData({
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
  };

  const postEmail = (event?: any) => {
    if (event) {
      event.preventDefault();
    }
    const params: any = {
      e_mail: formData.e_mail.value,
      pass: formData.pass.value,
    };

    const imgCode = emailCaptcha?.getCaptcha();
    if (imgCode?.verify) {
      params.captcha_code = imgCode.captcha_code;
      params.captcha_id = imgCode.captcha_id;
    }
    changeEmail(params)
      .then(async () => {
        await emailCaptcha?.close();
        setStep(1);
        toast.onShow({
          msg: t('change_email_info'),
          variant: 'warning',
        });
        initFormData();
      })
      .catch((err) => {
        if (err.isError) {
          emailCaptcha?.handleCaptchaError(err.list);
          const data = handleFormError(err, formData);
          setFormData({ ...data });
          const ele = document.getElementById(err.list[0].error_field);
          scrollToElementTop(ele);
        }
      });
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    if (!checkValidated()) {
      return;
    }
    if (!emailCaptcha) {
      postEmail();
      return;
    }
    emailCaptcha.check(() => {
      postEmail();
    });
  };

  return (
    <div>
      {step === 1 && (
        <Form>
          <Form.Group controlId="oldEmail" className="mb-3">
            <Form.Label>{t('email.label')}</Form.Label>
            <Form.Control
              type="text"
              disabled
              defaultValue={userInfo?.e_mail?.replace(
                /(.{2})(.+)(@.+)/i,
                '$1****$3',
              )}
            />
          </Form.Group>

          <Button
            variant="outline-secondary"
            onClick={() => {
              setStep(2);
            }}>
            {t('change_email_btn')}
          </Button>
        </Form>
      )}
      {step === 2 && (
        <Form noValidate onSubmit={handleSubmit}>
          <Form.Group controlId="pass" className="mb-3">
            <Form.Label>{t('pass.label')}</Form.Label>
            <Form.Control
              autoComplete="new-password"
              required
              type="password"
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

          <Form.Group controlId="e_mail" className="mb-3">
            <Form.Label>{t('new_email.label')}</Form.Label>
            <Form.Control
              autoComplete="off"
              required
              type="email"
              placeholder=""
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

          <div>
            <Button type="submit" variant="primary" className="me-2">
              {t('save', { keyPrefix: 'btns' })}
            </Button>

            <Button variant="link" onClick={() => setStep(1)}>
              {t('cancel', { keyPrefix: 'btns' })}
            </Button>
          </div>
        </Form>
      )}
    </div>
  );
};

export default React.memo(Index);
