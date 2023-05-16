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
        description: t('membership.text'),
        default: false,
      },
      allow_email_registrations: {
        type: 'boolean',
        title: t('email_registration.title'),
        description: t('email_registration.text'),
        default: true,
      },
      allow_email_domains: {
        type: 'string',
        title: t('allowed_email_domains.title'),
        description: t('allowed_email_domains.text'),
      },
      login_required: {
        type: 'boolean',
        title: t('private.title'),
        description: t('private.text'),
        default: false,
      },
    },
  };
  const uiSchema: UISchema = {
    allow_new_registrations: {
      'ui:widget': 'switch',
      'ui:options': {
        label: t('membership.label'),
      },
    },
    allow_email_registrations: {
      'ui:widget': 'switch',
      'ui:options': {
        label: t('email_registration.label'),
      },
    },
    allow_email_domains: {
      'ui:widget': 'textarea',
    },
    login_required: {
      'ui:widget': 'switch',
      'ui:options': {
        label: t('private.label'),
      },
    },
  };
  const [formData, setFormData] = useState(initFormData(schema));
  const { update: updateLoginSetting } = loginSettingStore((_) => _);

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();

    const allowedEmailDomains: string[] = [];
    if (formData.allow_email_domains.value) {
      const domainList = formData.allow_email_domains.value.split('\n');
      domainList.forEach((li) => {
        li = li.trim();
        if (li) {
          allowedEmailDomains.push(li);
        }
      });
    }
    const reqParams: Type.AdminSettingsLogin = {
      allow_new_registrations: formData.allow_new_registrations.value,
      allow_email_registrations: formData.allow_email_registrations.value,
      allow_email_domains: allowedEmailDomains,
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
        formMeta.allow_email_registrations.value =
          setting.allow_email_registrations;
        formMeta.allow_email_domains.value = '';
        if (Array.isArray(setting.allow_email_domains)) {
          formMeta.allow_email_domains.value =
            setting.allow_email_domains.join('\n');
        }
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
