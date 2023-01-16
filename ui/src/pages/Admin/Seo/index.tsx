import { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { getSeoSetting, putSeoSetting } from '@/services';
import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import { useToast } from '@/hooks';
import { handleFormError } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.seo',
  });
  const Toast = useToast();
  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      permalink: {
        type: 'number',
        title: t('permalink.label'),
        description: t('permalink.text'),
        enum: [1, 2],
        enumNames: ['/questions/123/post-title', '/questions/123'],
        default: 1,
      },
      robots: {
        type: 'string',
        title: t('robots.label'),
        description: t('robots.text'),
      },
    },
  };
  const uiSchema: UISchema = {
    permalink: {
      'ui:widget': 'select',
    },
    robots: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 10,
        className: 'font-monospace',
      },
    },
  };
  const [formData, setFormData] = useState(initFormData(schema));

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();

    const reqParams: Type.AdminSettingsSeo = {
      permalink: Number(formData.permalink.value),
      robots: formData.robots.value,
    };

    putSeoSetting(reqParams)
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
    getSeoSetting().then((setting) => {
      if (setting) {
        const formMeta = { ...formData };
        formMeta.robots.value = setting.robots;
        formMeta.permalink.value = setting.permalink;
        if (formMeta.permalink.value !== 1 && formMeta.permalink.value !== 2) {
          formMeta.permalink.value = 1;
        }
        setFormData(formMeta);
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
