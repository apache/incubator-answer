import React, { FC, FormEvent, useEffect, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { useToast, useCaptchaModal } from '@/hooks';
import { getLoggedUserInfo, changeEmail } from '@/services';
import { handleFormError } from '@/utils';

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
  const emailCaptcha = useCaptchaModal('edit_userinfo');

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
        errorMsg: t('email.msg'),
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

    const imgCode = emailCaptcha.getCaptcha();
    if (imgCode.verify) {
      params.captcha_code = imgCode.captcha_code;
      params.captcha_id = imgCode.captcha_id;
    }
    changeEmail(params)
      .then(async () => {
        await emailCaptcha.close();
        setStep(1);
        toast.onShow({
          msg: t('change_email_info'),
          variant: 'warning',
        });
        initFormData();
      })
      .catch((err) => {
        if (err.isError) {
          emailCaptcha.handleCaptchaError(err.list);
          const data = handleFormError(err, formData);
          setFormData({ ...data });
        }
      });
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    if (!checkValidated()) {
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
          <Form.Group controlId="currentPass" className="mb-3">
            <Form.Label>{t('pass.label')}</Form.Label>
            <Form.Control
              autoComplete="new-password"
              required
              type="password"
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

          <Form.Group controlId="newEmail" className="mb-3">
            <Form.Label>{t('email.label')}</Form.Label>
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
