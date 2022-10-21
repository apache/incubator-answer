import React, { FC, useEffect, useState } from 'react';
import { Form, Button, Stack } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type * as Type from '@answer/common/interface';
import { useToast } from '@answer/hooks';
import { useSmtpSetting, updateSmtpSetting } from '@answer/api';

import pattern from '@/common/pattern';

const Smtp: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.smtp',
  });
  const Toast = useToast();
  const { data: setting } = useSmtpSetting();
  const [formData, setFormData] = useState<Type.FormDataType>({
    from_email: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    from_name: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    smtp_host: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    encryption: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    smtp_port: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    smtp_authentication: {
      value: 'yes',
      isInvalid: false,
      errorMsg: '',
    },
    smtp_username: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    smtp_password: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    test_email_recipient: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });
  const checkValidated = (): boolean => {
    let ret = true;
    const { smtp_port, test_email_recipient } = formData;
    if (
      !/^[1-9][0-9]*$/.test(smtp_port.value) ||
      Number(smtp_port.value) > 65535
    ) {
      ret = false;
      formData.smtp_port = {
        value: smtp_port.value,
        isInvalid: true,
        errorMsg: t('smtp_port.msg'),
      };
    }
    if (
      test_email_recipient.value &&
      !pattern.email.test(test_email_recipient.value)
    ) {
      ret = false;
      formData.test_email_recipient = {
        value: test_email_recipient.value,
        isInvalid: true,
        errorMsg: t('test_email_recipient.msg'),
      };
    }
    setFormData({
      ...formData,
    });
    return ret;
  };

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();
    if (!checkValidated()) {
      return;
    }
    const reqParams: Type.AdminSettingsSmtp = {
      from_email: formData.from_email.value,
      from_name: formData.from_name.value,
      smtp_host: formData.smtp_host.value,
      encryption: formData.encryption.value,
      smtp_port: Number(formData.smtp_port.value),
      smtp_authentication: formData.smtp_authentication.value,
      smtp_username: formData.smtp_username.value,
      smtp_password: formData.smtp_password.value,
      test_email_recipient: formData.test_email_recipient.value,
    };

    updateSmtpSetting(reqParams)
      .then(() => {
        Toast.onShow({
          msg: t('update', { keyPrefix: 'toast' }),
          variant: 'success',
        });
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData[err.key].isInvalid = true;
          formData[err.key].errorMsg = err.value;
        }
        setFormData({ ...formData });
      });
  };
  const onFieldChange = (fieldName, fieldValue) => {
    if (!formData[fieldName]) {
      return;
    }
    const fieldData: Type.FormDataType = {
      [fieldName]: {
        value: fieldValue,
        isInvalid: false,
        errorMsg: '',
      },
    };
    setFormData({ ...formData, ...fieldData });
  };
  useEffect(() => {
    if (!setting) {
      return;
    }
    const formState = {};
    Object.keys(formData).forEach((k) => {
      let v = setting[k];
      if (v === null || v === undefined) {
        v = '';
      }
      formState[k] = { ...formData[k], value: v };
    });
    setFormData(formState);
  }, [setting]);

  return (
    <>
      <h3 className="mb-4">{t('page_title')}</h3>
      <Form noValidate onSubmit={onSubmit}>
        <Form.Group controlId="fromEmail" className="mb-3">
          <Form.Label>{t('from_email.label')}</Form.Label>
          <Form.Control
            required
            type="text"
            value={formData.from_email.value}
            isInvalid={formData.from_email.isInvalid}
            onChange={(evt) => onFieldChange('from_email', evt.target.value)}
          />
          <Form.Text as="div">{t('from_email.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.from_email.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="fromName" className="mb-3">
          <Form.Label>{t('from_name.label')}</Form.Label>
          <Form.Control
            required
            type="text"
            value={formData.from_name.value}
            isInvalid={formData.from_name.isInvalid}
            onChange={(evt) => onFieldChange('from_name', evt.target.value)}
          />
          <Form.Text as="div">{t('from_name.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.from_name.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="smtpHost" className="mb-3">
          <Form.Label>{t('smtp_host.label')}</Form.Label>
          <Form.Control
            required
            type="text"
            value={formData.smtp_host.value}
            isInvalid={formData.smtp_host.isInvalid}
            onChange={(evt) => onFieldChange('smtp_host', evt.target.value)}
          />
          <Form.Text as="div">{t('smtp_host.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.smtp_host.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="encryption" className="mb-3">
          <Form.Label>{t('encryption.label')}</Form.Label>
          <Stack direction="horizontal">
            <Form.Check
              inline
              label={t('encryption.ssl')}
              name="smtp_encryption"
              id="smtp_encryption_ssl"
              checked={formData.encryption.value === 'SSL'}
              onChange={() => onFieldChange('encryption', 'SSL')}
              type="radio"
            />
            <Form.Check
              inline
              label={t('encryption.none')}
              name="smtp_encryption"
              id="smtp_encryption_none"
              checked={!formData.encryption.value}
              onChange={() => onFieldChange('encryption', '')}
              type="radio"
            />
          </Stack>
          <Form.Text as="div">{t('encryption.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.encryption.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="smtpPort" className="mb-3">
          <Form.Label>{t('smtp_port.label')}</Form.Label>
          <Form.Control
            required
            type="text"
            value={formData.smtp_port.value}
            isInvalid={formData.smtp_port.isInvalid}
            onChange={(evt) => onFieldChange('smtp_port', evt.target.value)}
          />
          <Form.Text as="div">{t('smtp_port.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.smtp_port.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="smtpAuthentication" className="mb-3">
          <Form.Label>{t('smtp_authentication.label')}</Form.Label>
          <Stack direction="horizontal">
            <Form.Check
              inline
              label={t('smtp_authentication.yes')}
              name="smtp_authentication"
              id="smtp_authentication_yes"
              checked={!!formData.smtp_authentication.value}
              onChange={() => onFieldChange('smtp_authentication', true)}
              type="radio"
            />
            <Form.Check
              inline
              label={t('smtp_authentication.no')}
              name="smtp_authentication"
              id="smtp_authentication_no"
              checked={!formData.smtp_authentication.value}
              onChange={() => onFieldChange('smtp_authentication', false)}
              type="radio"
            />
          </Stack>
          <Form.Control.Feedback type="invalid">
            {formData.smtp_authentication.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="smtpUsername" className="mb-3">
          <Form.Label>{t('smtp_username.label')}</Form.Label>
          <Form.Control
            required
            type="text"
            value={formData.smtp_username.value}
            isInvalid={formData.smtp_username.isInvalid}
            onChange={(evt) => onFieldChange('smtp_username', evt.target.value)}
          />
          <Form.Text as="div">{t('smtp_username.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.smtp_username.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="smtpPassword" className="mb-3">
          <Form.Label>{t('smtp_password.label')}</Form.Label>
          <Form.Control
            required
            type="password"
            value={formData.smtp_password.value}
            isInvalid={formData.smtp_password.isInvalid}
            onChange={(evt) => onFieldChange('smtp_password', evt.target.value)}
          />
          <Form.Text as="div">{t('smtp_password.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.smtp_password.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="testEmailRecipient" className="mb-3">
          <Form.Label>{t('test_email_recipient.label')}</Form.Label>
          <Form.Control
            required
            type="text"
            value={formData.test_email_recipient.value}
            isInvalid={formData.test_email_recipient.isInvalid}
            onChange={(evt) =>
              onFieldChange('test_email_recipient', evt.target.value)
            }
          />
          <Form.Text as="div">{t('test_email_recipient.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.test_email_recipient.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>

        <Button variant="primary" type="submit">
          {t('save', { keyPrefix: 'btns' })}
        </Button>
      </Form>
    </>
  );
};

export default Smtp;
