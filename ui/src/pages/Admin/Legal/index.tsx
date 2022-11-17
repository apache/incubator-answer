import { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { marked } from 'marked';

import type * as Type from '@/common/interface';
import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import { useToast } from '@/hooks';
import { getLegalSetting, putLegalSetting } from '@/services';
import { handleFormError } from '@/utils';
import '../index.scss';

const Legal: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.legal',
  });
  const Toast = useToast();

  const schema: JSONSchema = {
    title: t('page_title'),
    required: ['terms_of_service', 'privacy_policy'],
    properties: {
      terms_of_service: {
        type: 'string',
        title: t('terms_of_service.label'),
        description: t('terms_of_service.text'),
      },
      privacy_policy: {
        type: 'string',
        title: t('privacy_policy.label'),
        description: t('privacy_policy.text'),
      },
    },
  };
  const uiSchema: UISchema = {
    terms_of_service: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 10,
      },
    },
    privacy_policy: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 10,
      },
    },
  };
  const [formData, setFormData] = useState(initFormData(schema));

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();

    const reqParams: Type.AdminSettingsLegal = {
      terms_of_service_original_text: formData.terms_of_service.value,
      terms_of_service_parsed_text: marked.parse(
        formData.terms_of_service.value,
      ),
      privacy_policy_original_text: formData.privacy_policy.value,
      privacy_policy_parsed_text: marked.parse(formData.privacy_policy.value),
    };

    putLegalSetting(reqParams)
      .then(() => {
        Toast.onShow({
          msg: t('update', { keyPrefix: 'toast' }),
          variant: 'success',
        });
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          setFormData({ ...data });
        }
      });
  };

  useEffect(() => {
    getLegalSetting().then((setting) => {
      const formMeta = { ...formData };
      formMeta.terms_of_service.value = setting.terms_of_service_original_text;
      formMeta.privacy_policy.value = setting.privacy_policy_original_text;

      setFormData(formMeta);
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

export default Legal;
