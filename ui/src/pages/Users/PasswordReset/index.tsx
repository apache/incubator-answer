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

import React, { FormEvent, useState } from 'react';
import { Container, Col, Form, Button } from 'react-bootstrap';
import { Link, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { loggedUserInfoStore } from '@/stores';
import type { FormDataType } from '@/common/interface';
import { replacementPassword } from '@/services';
import { handleFormError, scrollToElementTop } from '@/utils';

const Index: React.FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'password_reset' });
  const [searchParams] = useSearchParams();
  const [step, setStep] = useState(1);
  const clearUser = loggedUserInfoStore((state) => state.clear);
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
    if (!bol) {
      const errObj = Object.keys(formData).filter(
        (key) => formData[key].isInvalid,
      );
      const ele = document.getElementById(errObj[0]);
      scrollToElementTop(ele);
    }
    return bol;
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    if (checkValidated() === false) {
      return;
    }
    const code = searchParams.get('code');
    if (!code) {
      console.error('code is required');
      return;
    }
    replacementPassword({
      code: encodeURIComponent(code),
      pass: formData.pass.value,
    })
      .then(() => {
        // clear login information then to login page
        clearUser();
        setStep(2);
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          setFormData({ ...data });
          const ele = document.getElementById(err.list[0].error_field);
          scrollToElementTop(ele);
        }
      });
  };
  usePageTags({
    title: t('account_recovery', { keyPrefix: 'page_title' }),
  });
  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '6rem' }}>
      <h3 className="text-center mb-5">{t('page_title')}</h3>
      {step === 1 && (
        <Col className="mx-auto" md={6} lg={4} xl={3}>
          <Form noValidate onSubmit={handleSubmit} autoComplete="off">
            <Form.Group controlId="pass" className="mb-3">
              <Form.Label>{t('password.label')}</Form.Label>
              <Form.Control
                autoComplete="off"
                required
                type="password"
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

            <Form.Group controlId="passSecond" className="mb-3">
              <Form.Label>{t('password_confirm.label')}</Form.Label>
              <Form.Control
                autoComplete="off"
                required
                type="password"
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
  );
};

export default React.memo(Index);
