import { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { getSeoSetting, putSeoSetting } from '@/services';
import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import { useToast } from '@/hooks';
import { handleFormError } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.themes',
  });
  const Toast = useToast();
  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      themes: {
        type: 'string',
        title: t('themes.label'),
        description: t('themes.text'),
        enum: ['default'],
        enumNames: ['Default'],
      },
      navbar_style: {
        type: 'string',
        title: t('navbar_style.label'),
        description: t('navbar_style.text'),
        enum: ['colored', 'light'],
        enumNames: ['Colored', 'Light'],
      },
      primary_color: {
        type: 'string',
        title: t('primary_color.label'),
        description: t('primary_color.text'),
      },
    },
  };
  const uiSchema: UISchema = {
    themes: {
      'ui:widget': 'select',
    },
    navbar_style: {
      'ui:widget': 'select',
    },
    primary_color: {
      'ui:options': {
        type: 'color',
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
        // const formMeta = { ...formData };
        // formMeta.robots.value = setting.robots;
        // setFormData(formMeta);
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
