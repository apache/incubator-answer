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
    keyPrefix: 'admin.legal',
  });
  // const Toast = useToast();
  // const updateSiteInfo = siteInfoStore((state) => state.update);

  const { data: setting } = useGeneralSetting();
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
      terms_of_service: formData.terms_of_service.value,
      privacy_policy: formData.privacy_policy.value,
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
