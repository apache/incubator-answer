import React, { useState, FormEvent, useEffect } from 'react';
import { useTranslation } from 'react-i18next';

import type { FormDataType } from '@/common/interface';
import { useToast } from '@/hooks';
import { setNotice, getLoggedUserInfo } from '@/services';
import { SchemaForm, JSONSchema, UISchema, initFormData } from '@/components';

const Index = () => {
  const toast = useToast();
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.notification',
  });
  const schema: JSONSchema = {
    title: t('heading'),
    properties: {
      notice_switch: {
        type: 'boolean',
        title: t('email.label'),
        default: false,
      },
    },
  };
  const uiSchema: UISchema = {
    notice_switch: {
      'ui:widget': 'switch',
      'ui:options': {
        label: t('email.radio'),
      },
    },
  };
  const [formData, setFormData] = useState<FormDataType>(initFormData(schema));

  const getProfile = () => {
    getLoggedUserInfo().then((res) => {
      if (res) {
        setFormData({
          notice_switch: {
            value: res.notice_status === 1,
            isInvalid: false,
            errorMsg: '',
          },
        });
      }
    });
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    setNotice({
      notice_switch: formData.notice_switch.value,
    }).then(() => {
      toast.onShow({
        msg: t('update', { keyPrefix: 'toast' }),
        variant: 'success',
      });
    });
  };

  useEffect(() => {
    getProfile();
  }, []);
  const handleChange = (ud) => {
    setFormData(ud);
  };
  return (
    <>
      <h3 className="mb-4">{t('heading')}</h3>
      <SchemaForm
        schema={schema}
        uiSchema={uiSchema}
        formData={formData}
        onChange={handleChange}
        onSubmit={handleSubmit}
      />
    </>
  );
};

export default React.memo(Index);
