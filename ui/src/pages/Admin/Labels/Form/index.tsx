import React, { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Button } from 'react-bootstrap';

import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import type * as Type from '@/common/interface';
import { useToast } from '@/hooks';
import { siteInfoStore } from '@/stores';
import { useGeneralSetting, updateGeneralSetting } from '@/services';

interface IProps {
  onClose: () => void;
}
const LabelForm: FC<IProps> = ({ onClose }) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.labels.form',
  });
  const Toast = useToast();
  const updateSiteInfo = siteInfoStore((state) => state.update);

  const { data: setting } = useGeneralSetting();
  const schema: JSONSchema = {
    title: t('title'),
    required: ['name', 'site_url', 'contact_email'],
    properties: {
      display_name: {
        type: 'string',
        title: t('display_name.label'),
      },
      url_slug: {
        type: 'string',
        title: t('url_slug.label'),
        description: t('url_slug.text'),
      },
      description: {
        type: 'string',
        title: t('description.label'),
        description: t('description.text'),
      },
      color: {
        type: 'string',
        title: t('color.label'),
      },
    },
  };
  const uiSchema: UISchema = {
    color: {
      'ui:options': {
        type: 'color',
      },
    },
  };
  const [formData, setFormData] = useState(initFormData(schema));

  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();

    const reqParams: Type.AdminSettingsGeneral = {
      name: formData.name.value,
      description: formData.description.value,
      short_description: formData.short_description.value,
      site_url: formData.site_url.value,
      contact_email: formData.contact_email.value,
    };

    updateGeneralSetting(reqParams)
      .then(() => {
        Toast.onShow({
          msg: t('update', { keyPrefix: 'toast' }),
          variant: 'success',
        });
        onClose();
        updateSiteInfo(reqParams);
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData[err.key].isInvalid = true;
          formData[err.key].errorMsg = err.value;
        }
        setFormData({ ...formData });
      });
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
      <Button
        size="sm"
        className="mb-4"
        variant="outline-secondary"
        onClick={onClose}>
        ← {t('back')}
      </Button>
      <h3 className="mb-4">{t('title')}</h3>
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

export default LabelForm;
