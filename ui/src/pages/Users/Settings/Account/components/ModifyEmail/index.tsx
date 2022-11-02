import React, { FC, FormEvent, useEffect, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { useToast } from '@/hooks';
import { getLoggedUserInfo, changeEmail } from '@/services';

const reg = /(?<=.{2}).+(?=@)/gi;

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
  });
  const [userInfo, setUserInfo] = useState<Type.UserInfoRes>();
  const toast = useToast();
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
    const { e_mail } = formData;

    if (!e_mail.value) {
      bol = false;
      formData.e_mail = {
        value: '',
        isInvalid: true,
        errorMsg: t('email.msg'),
      };
    }
    setFormData({
      ...formData,
    });
    return bol;
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    if (!checkValidated()) {
      return;
    }
    changeEmail({
      e_mail: formData.e_mail.value,
    })
      .then(() => {
        setStep(1);
        toast.onShow({
          msg: t('change_email_info'),
          variant: 'warning',
        });
        setFormData({
          e_mail: {
            value: '',
            isInvalid: false,
            errorMsg: '',
          },
        });
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData.e_mail.isInvalid = true;
          formData.e_mail.errorMsg = err.value;
        }
        setFormData({ ...formData });
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
              defaultValue={userInfo?.e_mail?.replace(reg, () => '*'.repeat(4))}
            />
          </Form.Group>

          <Button variant="outline-secondary" onClick={() => setStep(2)}>
            {t('change_email_btn')}
          </Button>
        </Form>
      )}
      {step === 2 && (
        <Form noValidate onSubmit={handleSubmit}>
          <Form.Group controlId="newEmail" className="mb-3">
            <Form.Label>{t('email.label')}</Form.Label>
            <Form.Control
              autoComplete="off"
              required
              type="text"
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
