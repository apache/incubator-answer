import React, { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { useToast } from '@/hooks';
import { useSmtpSetting, updateSmtpSetting } from '@/services';
import pattern from '@/common/pattern';
import { SchemaForm, JSONSchema, UISchema } from '@/components';
import { initFormData } from '../../../components/SchemaForm/index';

const Smtp: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.smtp',
  });
  const Toast = useToast();
  const { data: setting } = useSmtpSetting();
  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      from_email: {
        type: 'string',
        title: t('from_email.label'),
        description: t('from_email.text'),
      },
      from_name: {
        type: 'string',
        title: t('from_name.label'),
        description: t('from_name.text'),
      },
      smtp_host: {
        type: 'string',
        title: t('smtp_host.label'),
        description: t('smtp_host.text'),
      },
      encryption: {
        type: 'boolean',
        title: t('encryption.label'),
        description: t('encryption.text'),
        enum: [true, false],
        enumNames: ['SSL', ''],
      },
      smtp_port: {
        type: 'string',
        title: t('smtp_port.label'),
        description: t('smtp_port.text'),
      },
      smtp_authentication: {
        type: 'boolean',
        title: t('smtp_authentication.label'),
        enum: [true, false],
        enumNames: [t('smtp_authentication.yes'), t('smtp_authentication.no')],
      },
      smtp_username: {
        type: 'string',
        title: t('smtp_username.label'),
        description: t('smtp_username.text'),
      },
      smtp_password: {
        type: 'string',
        title: t('smtp_password.label'),
        description: t('smtp_password.text'),
      },
      test_email_recipient: {
        type: 'string',
        title: t('test_email_recipient.label'),
        description: t('test_email_recipient.text'),
      },
    },
  };
  const uiSchema: UISchema = {
    encryption: {
      'ui:widget': 'radio',
    },
    smtp_password: {
      'ui:options': {
        type: 'password',
      },
    },
    smtp_authentication: {
      'ui:widget': 'radio',
    },
    smtp_port: {
      'ui:options': {
        invalid: t('smtp_port.msg'),
        validator: (value) => {
          if (!/^[1-9][0-9]*$/.test(value) || Number(value) > 65535) {
            return false;
          }
          return true;
        },
      },
    },
    test_email_recipient: {
      'ui:options': {
        invalid: t('test_email_recipient.msg'),
        validator: (value) => {
          if (value && !pattern.email.test(value)) {
            return false;
          }
          return true;
        },
      },
    },
  };
  const [formData, setFormData] = useState<Type.FormDataType>(
    initFormData(schema),
  );

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();

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

  const handleOnChange = (data) => {
    setFormData(data);
  };
  return (
    <>
      <h3 className="mb-4">{t('page_title')}</h3>
      <SchemaForm
        schema={schema}
        uiSchema={uiSchema}
        formData={formData}
        onChange={handleOnChange}
        onSubmit={onSubmit}
      />
    </>
  );
};

export default Smtp;
