import { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { getLoginSetting, putLoginSetting } from '@/services';
import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import { useToast } from '@/hooks';
import { handleFormError } from '@/utils';
import { loginSettingStore } from '@/stores';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.login',
  });
  const Toast = useToast();
  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      allow_new_registrations: {
        type: 'boolean',
        title: t('membership.title'),
        label: t('membership.label'),
        description: t('membership.text'),
        default: false,
      },
      login_required: {
        type: 'boolean',
        title: t('private.title'),
        label: t('private.label'),
        description: t('private.text'),
        default: false,
      },
    },
  };
  const uiSchema: UISchema = {
    allow_new_registrations: {
      'ui:widget': 'switch',
    },
    login_required: {
      'ui:widget': 'switch',
    },
  };
  const [formData, setFormData] = useState(initFormData(schema));
  const { update: updateLoginSetting } = loginSettingStore((_) => _);

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();

    const reqParams: Type.AdminSettingsLogin = {
      allow_new_registrations: formData.allow_new_registrations.value,
      login_required: formData.login_required.value,
    };

    putLoginSetting(reqParams)
      .then(() => {
        Toast.onShow({
          msg: t('update', { keyPrefix: 'toast' }),
          variant: 'success',
        });
        updateLoginSetting(reqParams);
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          setFormData({ ...data });
        }
      });
  };

  useEffect(() => {
    getLoginSetting().then((setting) => {
      if (setting) {
        const formMeta = { ...formData };
        formMeta.allow_new_registrations.value =
          setting.allow_new_registrations;
        formMeta.login_required.value = setting.login_required;
        setFormData({ ...formMeta });
      }
    });
  }, []);

  const handleOnChange = (data) => {
    setFormData(data);
  };

  return (
    <>
      <h3 className="mb-4">{t('page_title')}</h3>
      <SchemaForm
        schema={schema}
        formData={formData}
        onSubmit={onSubmit}
        uiSchema={uiSchema}
        onChange={handleOnChange}
      />
    </>
  );
};

export default Index;
