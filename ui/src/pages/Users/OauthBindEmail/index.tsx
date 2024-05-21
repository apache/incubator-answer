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

import { FC, memo, useState, useEffect } from 'react';
import { Container, Col, Form, Button } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';
import { useSearchParams, useNavigate } from 'react-router-dom';

import { Modal, WelcomeTitle } from '@/components';
import type { FormDataType } from '@/common/interface';
import { usePageTags } from '@/hooks';
import { loggedUserInfoStore } from '@/stores';
import { oAuthBindEmail, getLoggedUserInfo } from '@/services';
import Storage from '@/utils/storage';
import { LOGGED_TOKEN_STORAGE_KEY } from '@/common/constants';
import { handleFormError, scrollToElementTop } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'oauth_bind_email',
  });
  const navigate = useNavigate();
  const [searchParams, setUrlSearchParams] = useSearchParams();
  const updateUser = loggedUserInfoStore((state) => state.update);
  const binding_key = searchParams.get('binding_key') || '';
  const [showResult, setShowResult] = useState(false);

  usePageTags({
    title: t('confirm_email', { keyPrefix: 'page_title' }),
  });
  const [formData, setFormData] = useState<FormDataType>({
    email: {
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

    if (!formData.email.value) {
      bol = false;
      formData.email = {
        value: '',
        isInvalid: true,
        errorMsg: t('email.msg.empty'),
      };
    }
    setFormData({
      ...formData,
    });
    if (!bol) {
      const ele = document.getElementById('email');
      scrollToElementTop(ele);
    }

    return bol;
  };

  const getUserInfo = (token) => {
    Storage.set(LOGGED_TOKEN_STORAGE_KEY, token);
    getLoggedUserInfo().then((user) => {
      updateUser(user);
      setTimeout(() => {
        navigate('/users/login?status=inactive', { replace: true });
      }, 0);
    });
  };

  const connectConfirm = () => {
    Modal.confirm({
      title: t('modal_title'),
      content: t('modal_content'),
      cancelText: t('modal_cancel'),
      confirmText: t('modal_confirm'),
      onConfirm: () => {
        // send activation email
        oAuthBindEmail({
          binding_key,
          email: formData.email.value,
          must: true,
        }).then((result) => {
          if (result.access_token) {
            getUserInfo(result.access_token);
          } else {
            searchParams.delete('binding_key');
            setUrlSearchParams('');
            setShowResult(true);
          }
        });
      },
      onCancel: () => {
        setFormData({
          email: {
            value: '',
            isInvalid: false,
            errorMsg: '',
          },
        });
      },
    });
  };

  const handleSubmit = (event: any) => {
    event.preventDefault();
    event.stopPropagation();
    if (!checkValidated()) {
      return;
    }

    if (binding_key) {
      oAuthBindEmail({
        binding_key,
        email: formData.email.value,
        must: false,
      })
        .then((res) => {
          if (res.email_exist_and_must_be_confirmed) {
            connectConfirm();
          }
          if (res.access_token) {
            getUserInfo(res.access_token);
          }
        })
        .catch((err) => {
          if (err.isError) {
            const data = handleFormError(err, formData);
            setFormData({ ...data });
            const ele = document.getElementById(err.list[0].error_field);
            scrollToElementTop(ele);
          }
        });
    }
  };

  useEffect(() => {
    if (!binding_key) {
      navigate('/', { replace: true });
    }
  }, []);
  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '6rem' }}>
      <WelcomeTitle />
      {showResult ? (
        <Col md={6} className="mx-auto text-center">
          <p>
            <Trans
              i18nKey="inactive.first"
              values={{ mail: formData.email.value }}
              components={{ bold: <strong /> }}
            />
          </p>
          <p>{t('info', { keyPrefix: 'inactive' })}</p>
        </Col>
      ) : (
        <Col className="mx-auto" md={6} lg={4} xl={3}>
          <div className="text-center mb-5">{t('subtitle')}</div>
          <Form noValidate onSubmit={handleSubmit} autoComplete="off">
            <Form.Group controlId="email" className="mb-3">
              <Form.Label>{t('email.label')}</Form.Label>
              <Form.Control
                required
                type="email"
                value={formData.email.value}
                isInvalid={formData.email.isInvalid}
                onChange={(e) => {
                  handleChange({
                    email: {
                      value: e.target.value,
                      isInvalid: false,
                      errorMsg: '',
                    },
                  });
                }}
              />
              <Form.Control.Feedback type="invalid">
                {formData.email.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>

            <div className="d-grid mb-3">
              <Button variant="primary" type="submit">
                {t('btn_update')}
              </Button>
            </div>
          </Form>
        </Col>
      )}
    </Container>
  );
};

export default memo(Index);
