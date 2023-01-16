import { useState, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { useParams } from 'react-router-dom';

import { isEmpty } from 'lodash';

import { useToast } from '@/hooks';
import type * as Types from '@/common/interface';
import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import { useQueryPluginConfig, updatePluginConfig } from '@/services';

const Config = () => {
  const { slug_name } = useParams<{ slug_name: string }>();
  const { data } = useQueryPluginConfig({ plugin_slug_name: slug_name });
  const Toast = useToast();
  const { t } = useTranslation('translation');
  const [schema, setSchema] = useState<JSONSchema | null>(null);

  const uiSchema: UISchema = {};
  const required: string[] = [];

  const [formData, setFormData] = useState<Types.FormDataType | null>(null);

  useEffect(() => {
    if (!data) {
      return;
    }
    const properties: JSONSchema['properties'] = {};

    data.config_fields?.forEach((item) => {
      properties[item.name] = {
        type: 'string',
        title: item.title,
        description: item.description,
        default: item.value,
      };

      if (item.ui_options) {
        uiSchema[item.name] = {
          'ui:options': item.ui_options,
        };
      }
      if (item.required) {
        required.push(item.name);
      }
    });

    setSchema({
      title: data?.name || '',
      required,
      properties,
    });
  }, [data?.config_fields]);

  useEffect(() => {
    if (!schema || formData) {
      return;
    }
    setFormData(initFormData(schema));
  }, [schema, formData]);

  const onSubmit = (evt) => {
    if (!formData) {
      return;
    }
    evt.preventDefault();
    evt.stopPropagation();
    const config_fields = {};
    Object.keys(formData).forEach((key) => {
      config_fields[key] = formData[key].value;
    });
    const params = {
      plugin_slug_name: slug_name,
      config_fields,
    };
    updatePluginConfig(params).then(() => {
      Toast.onShow({
        msg: t('update', { keyPrefix: 'toast' }),
        variant: 'success',
      });
    });
  };

  const handleOnChange = (form) => {
    setFormData(form);
  };

  if (!data || !schema || !formData) {
    return null;
  }

  if (isEmpty(schema.properties)) {
    return <h3 className="mb-4">{data?.name}</h3>;
  }
  return (
    <>
      <h3 className="mb-4">{data?.name}</h3>
      <SchemaForm
        schema={schema}
        uiSchema={uiSchema}
        formData={formData}
        onSubmit={onSubmit}
        onChange={handleOnChange}
      />
    </>
  );
};

export default Config;
