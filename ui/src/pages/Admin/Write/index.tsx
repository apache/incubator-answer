import React, { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import type * as Type from '@/common/interface';
// import { useToast } from '@/hooks';
// import { siteInfoStore } from '@/stores';
import { useGeneralSetting } from '@/services';

import '../index.scss';

const Legal: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.write',
  });
  // const Toast = useToast();
  // const updateSiteInfo = siteInfoStore((state) => state.update);

  const { data: setting } = useGeneralSetting();
  const schema: JSONSchema = {
    title: t('page_title'),
    required: ['terms_of_service', 'privacy_policy'],
    properties: {
      recommend_tags: {
        type: 'string',
        title: t('recommend_tags.label'),
        description: t('recommend_tags.text'),
      },
      required_tag: {
        type: 'boolean',
        title: t('required_tag.label'),
        description: t('required_tag.text'),
      },
      reserved_tags: {
        type: 'string',
        title: t('reserved_tags.label'),
        description: t('reserved_tags.text'),
      },
    },
  };
  const uiSchema: UISchema = {
    recommend_tags: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 5,
      },
    },
    required_tag: {
      'ui:widget': 'switch',
    },
    reserved_tags: {
      'ui:widget': 'textarea',
      'ui:options': {
        rows: 5,
      },
    },
  };
  const [formData, setFormData] = useState(initFormData(schema));

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();

    const reqParams: Type.AdminSettingsWrite = {
      recommend_tags: formData.recommend_tags.value,
      required_tag: formData.required_tag.value,
    };

    console.log(reqParams);
    // updateGeneralSetting(reqParams)
    //   .then(() => {
    //     Toast.onShow({
    //       msg: t('update', { keyPrefix: 'toast' }),
    //       variant: 'success',
    //     });
    //     updateSiteInfo(reqParams);
    //   })
    //   .catch((err) => {
    //     if (err.isError && err.key) {
    //       formData[err.key].isInvalid = true;
    //       formData[err.key].errorMsg = err.value;
    //     }
    //     setFormData({ ...formData });
    //   });
  };

  useEffect(() => {
    if (!setting) {
      return;
    }
    const formMeta = {};
    Object.keys(setting).forEach((k) => {
      formMeta[k] = { ...formData[k], value: setting[k] };
    });
    setFormData({ ...formData, ...formMeta });
  }, [setting]);

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
