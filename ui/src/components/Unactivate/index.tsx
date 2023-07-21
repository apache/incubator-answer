import React, { useState } from 'react';
import { Button, Col } from 'react-bootstrap';
import { Trans, useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import type { ImgCodeReq, FormDataType } from '@/common/interface';
import { loggedUserInfoStore } from '@/stores';
import { resendEmail } from '@/services';
import { handleFormError } from '@/utils';
import { useCaptchaModal } from '@/hooks';

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

  const emailCaptcha = useCaptchaModal('email');

  const submit = () => {
    let req: ImgCodeReq = {};
    const imgCode = emailCaptcha.getCaptcha();
    if (imgCode.verify) {
      req = {
        captcha_code: imgCode.captcha_code,
        captcha_id: imgCode.captcha_id,
      };
    }
    resendEmail(req)
      .then(() => {
        emailCaptcha.close();
        setSuccess(true);
      })
      .catch((err) => {
        if (err.isError) {
          emailCaptcha.handleCaptchaError(err.list);
          const data = handleFormError(err, formData);
          setFormData({ ...data });
        }
      });
  };

  const onSentEmail = (evt) => {
    evt.preventDefault();
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
