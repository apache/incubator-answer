import { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { getSeoSetting, putSeoSetting } from '@/services';
import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import { useToast } from '@/hooks';
import { handleFormError } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.css_and_html',
  });
  const Toast = useToast();
  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      custom_css: {
        type: 'string',
        title: t('custom_css.label'),
        description: t('custom_css.text'),
      },
      head: {
        type: 'string',
        title: t('head.label'),
        description: t('head.text'),
      },
      header: {
        type: 'string',
        title: t('header.label'),
        description: t('header.text'),
      },
      footer: {
        type: 'string',
        title: t('footer.label'),
        description: t('footer.text'),
      },
    },
  };
  const uiSchema: UISchema = {
    custom_css: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 10,
      },
    },
    head: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 10,
      },
    },
    header: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 10,
      },
    },
    footer: {
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

    const reqParams: Type.AdminSettingsSeo = {
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
