import React, { FC, FormEvent, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type { FormDataType } from '@/common/interface';
import { modifyPassword } from '@answer/api';
import { useToast } from '@answer/hooks';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.account',
  });
  const [showForm, setFormState] = useState(false);
  const toast = useToast();
  const [formData, setFormData] = useState<FormDataType>({
    old_pass: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    pass: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    pass2: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });

  const handleFormState = () => {
    setFormState((pre) => !pre);
  };

  const handleChange = (params: FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  const checkValidated = (): boolean => {
    let bol = true;
    const { old_pass, pass, pass2 } = formData;

    if (!old_pass.value) {
      bol = false;
      formData.old_pass = {
        value: '',
        isInvalid: true,
        errorMsg: t('current_pass.msg.empty'),
      };
    }

    if (!pass.value) {
      bol = false;
      formData.pass = {
        value: '',
        isInvalid: true,
        errorMsg: t('current_pass.msg.empty'),
      };
    }

    if (bol && pass.value && pass.value.length < 8) {
      bol = false;
      formData.pass = {
        value: pass.value,
        isInvalid: true,
        errorMsg: t('current_pass.msg.length'),
      };
    }

    if (!pass2.value) {
      bol = false;
      formData.pass2 = {
        value: '',
        isInvalid: true,
        errorMsg: t('current_pass.msg.empty'),
      };
    }

    if (bol && pass2.value && pass2.value.length < 8) {
      bol = false;
      formData.pass2 = {
        value: pass2.value,
        isInvalid: true,
        errorMsg: t('current_pass.msg.length'),
      };
    }
    if (bol && pass.value && pass2.value && pass.value !== pass2.value) {
      bol = false;
      formData.pass2 = {
        value: pass2.value,
        isInvalid: true,
        errorMsg: t('current_pass.msg.different'),
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
    modifyPassword({
      old_pass: formData.old_pass.value,
      pass: formData.pass.value,
    })
      .then(() => {
        toast.onShow({
          msg: t('update_password', { keyPrefix: 'toast' }),
          variant: 'success',
        });
        handleFormState();
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData[err.key].isInvalid = true;
          formData[err.key].errorMsg = err.value;
        }
        setFormData({ ...formData });
      });
  };

  return (
    <div className="mt-5">
      {showForm ? (
        <Form noValidate onSubmit={handleSubmit}>
          <Form.Group controlId="oldPass" className="mb-3">
            <Form.Label>{t('current_pass.label')}</Form.Label>
            <Form.Control
              autoComplete="off"
              required
              type="password"
              placeholder=""
              // value={formData.password.value}
              isInvalid={formData.old_pass.isInvalid}
              onChange={(e) =>
                handleChange({
                  old_pass: {
                    value: e.target.value,
                    isInvalid: false,
                    errorMsg: '',
                  },
                })
              }
            />
            <Form.Control.Feedback type="invalid">
              {formData.old_pass.errorMsg}
            </Form.Control.Feedback>
          </Form.Group>

          <Form.Group controlId="newPass" className="mb-3">
            <Form.Label>{t('new_pass.label')}</Form.Label>
            <Form.Control
              autoComplete="off"
              required
              type="password"
              maxLength={32}
              // value={formData.password.value}
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

          <Form.Group controlId="newPass2" className="mb-3">
            <Form.Label>{t('pass_confirm.label')}</Form.Label>
            <Form.Control
              autoComplete="off"
              required
              type="password"
              maxLength={32}
              // value={formData.password.value}
              isInvalid={formData.pass2.isInvalid}
              onChange={(e) =>
                handleChange({
                  pass2: {
                    value: e.target.value,
                    isInvalid: false,
                    errorMsg: '',
                  },
                })
              }
            />
            <Form.Control.Feedback type="invalid">
              {formData.pass2.errorMsg}
            </Form.Control.Feedback>
          </Form.Group>
          <div>
            <Button type="submit" variant="primary" className="me-2">
              {t('save', { keyPrefix: 'btns' })}
            </Button>

            <Button variant="link" onClick={() => handleFormState()}>
              {t('cancel', { keyPrefix: 'btns' })}
            </Button>
          </div>
        </Form>
      ) : (
        <>
          <Form.Label>{t('password_title')}</Form.Label>
          <br />
          <Button
            variant="outline-secondary"
            type="submit"
            onClick={handleFormState}>
            {t('change_pass_btn')}
          </Button>
        </>
      )}
    </div>
  );
};

export default React.memo(Index);
