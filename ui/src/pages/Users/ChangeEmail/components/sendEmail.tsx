import { FC, memo, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';

import type { PasswordResetReq, FormDataType } from '@/common/interface';
import { loggedUserInfoStore } from '@/stores';
import { changeEmail } from '@/services';
import { handleFormError } from '@/utils';
import { useCaptchaModal } from '@/hooks';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'change_email' });
  const [formData, setFormData] = useState<FormDataType>({
    e_mail: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });

  const navigate = useNavigate();
  const { user: userInfo, update: updateUser } = loggedUserInfoStore();

  const emailCaptcha = useCaptchaModal('email');

  const handleChange = (params: FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  const checkValidated = (): boolean => {
    let bol = true;

    if (!formData.e_mail.value) {
      bol = false;
      formData.e_mail = {
        value: '',
        isInvalid: true,
        errorMsg: t('email.msg.empty'),
      };
    }
    setFormData({
      ...formData,
    });
    return bol;
  };

  const sendEmail = (e?: any) => {
    if (e) {
      e.preventDefault();
    }
    const params: PasswordResetReq = {
      e_mail: formData.e_mail.value,
    };
    const imgCode = emailCaptcha.getCaptcha();
    if (imgCode.verify) {
      params.captcha_code = imgCode.captcha_code;
      params.captcha_id = imgCode.captcha_id;
    }

    changeEmail(params)
      .then(async () => {
        await emailCaptcha.close();
        userInfo.e_mail = formData.e_mail.value;
        updateUser(userInfo);
        navigate('/users/login', { replace: true });
      })
      .catch((err) => {
        if (err.isError) {
          emailCaptcha.handleCaptchaError(err.list);
          const data = handleFormError(err, formData);
          setFormData({ ...data });
        }
      });
  };

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    event.stopPropagation();
    if (!checkValidated()) {
      return;
    }

    emailCaptcha.check(() => {
      sendEmail();
    });
  };

  const goBack = () => {
    navigate('/users/login?status=inactive', { replace: true });
  };

  return (
    <Form noValidate onSubmit={handleSubmit} autoComplete="off">
      <Form.Group controlId="email" className="mb-3">
        <Form.Label>{t('email.label')}</Form.Label>
        <Form.Control
          required
          type="email"
          value={formData.e_mail.value}
          isInvalid={formData.e_mail.isInvalid}
          onChange={(e) => {
            handleChange({
              e_mail: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            });
          }}
        />
        <Form.Control.Feedback type="invalid">
          {formData.e_mail.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <div className="d-grid mb-3">
        <Button variant="primary" type="submit">
          {t('btn_update')}
        </Button>
        <Button variant="link" className="mt-2 d-block" onClick={goBack}>
          {t('btn_cancel')}
        </Button>
      </div>
    </Form>
  );
};

export default memo(Index);
