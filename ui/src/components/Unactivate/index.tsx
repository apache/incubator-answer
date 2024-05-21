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

import React, { useState } from 'react';
import { Button, Col } from 'react-bootstrap';
import { Trans, useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import type { ImgCodeReq, FormDataType } from '@/common/interface';
import { loggedUserInfoStore } from '@/stores';
import { resendEmail } from '@/services';
import { handleFormError } from '@/utils';
import { useCaptchaPlugin } from '@/utils/pluginKit';

interface IProps {
  visible?: boolean;
}

const Index: React.FC<IProps> = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'inactive' });
  const [isSuccess, setSuccess] = useState(false);
  const { e_mail } = loggedUserInfoStore((state) => state.user);
  const [formData, setFormData] = useState<FormDataType>({
    captcha_code: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });

  const emailCaptcha = useCaptchaPlugin('email');

  const submit = () => {
    let req: ImgCodeReq = {};
    const imgCode = emailCaptcha?.getCaptcha();
    if (imgCode?.verify) {
      req = {
        captcha_code: imgCode.captcha_code,
        captcha_id: imgCode.captcha_id,
      };
    }
    resendEmail(req)
      .then(async () => {
        await emailCaptcha?.close();
        setSuccess(true);
      })
      .catch((err) => {
        if (err.isError) {
          emailCaptcha?.handleCaptchaError(err.list);
          const data = handleFormError(err, formData);
          setFormData({ ...data });
        }
      });
  };

  const onSentEmail = (evt) => {
    evt.preventDefault();
    if (!emailCaptcha) {
      submit();
      return;
    }
    emailCaptcha.check(() => {
      submit();
    });
  };

  return (
    <Col md={6} className="mx-auto text-center">
      {isSuccess ? (
        <p>
          <Trans
            i18nKey="inactive.another"
            values={{ mail: e_mail }}
            components={{ bold: <strong /> }}
          />
        </p>
      ) : (
        <>
          <p>
            <Trans
              i18nKey="inactive.first"
              values={{ mail: e_mail }}
              components={{ bold: <strong /> }}
            />
          </p>
          <p>{t('info')}</p>
          <Button variant="link" onClick={onSentEmail}>
            {t('btn_name')}
          </Button>
          <Link to="/users/change-email" replace className="btn btn-link ms-2">
            {t('change_btn_name')}
          </Link>
        </>
      )}
    </Col>
  );
};

export default React.memo(Index);
