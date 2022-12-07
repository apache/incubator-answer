import { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { getSeoSetting, putSeoSetting } from '@/services';
import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import { useToast } from '@/hooks';
import { handleFormError } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.login',
  });
  const Toast = useToast();
  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      membership: {
        type: 'boolean',
        title: t('membership.label'),
        label: t('membership.labelAlias'),
        description: t('membership.text'),
        default: true,
      },
      private: {
        type: 'string',
        title: t('private.label'),
        description: t('private.text'),
      },
    },
  };
  const uiSchema: UISchema = {
    membership: {
      'ui:widget': 'switch',
    },
    private: {
      'ui:widget': 'switch',
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
