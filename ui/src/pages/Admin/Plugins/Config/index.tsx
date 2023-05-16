import { useState, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { useParams } from 'react-router-dom';

import { useToast } from '@/hooks';
import type * as Types from '@/common/interface';
import { SchemaForm, JSONSchema, UISchema } from '@/components';
import { useQueryPluginConfig, updatePluginConfig } from '@/services';
import {
  InputOptions,
  FormKit,
  initFormData,
  mergeFormData,
} from '@/components/SchemaForm';

const Config = () => {
  const { t } = useTranslation('translation');
  const { slug_name } = useParams<{ slug_name: string }>();
  const { data, mutate: refreshPluginConfig } = useQueryPluginConfig({
    plugin_slug_name: slug_name,
  });
  const Toast = useToast();
  const [schema, setSchema] = useState<JSONSchema | null>(null);
  const [uiSchema, setUISchema] = useState<UISchema>();
  const required: string[] = [];

  const [formData, setFormData] = useState<Types.FormDataType | null>(null);

  useEffect(() => {
    if (!data) {
      return;
    }
    const properties: JSONSchema['properties'] = {};
    const uiConf: UISchema = {};
    data.config_fields?.forEach((item) => {
      properties[item.name] = {
        type: 'string',
        title: item.title,
        description: item.description,
        default: item.value,
      };

      if (item.options instanceof Array) {
        properties[item.name].enum = item.options.map((option) => option.value);
        properties[item.name].enumNames = item.options.map(
          (option) => option.label,
        );
      }
      uiConf[item.name] = {};
      uiConf[item.name]['ui:widget'] = item.type;
      if (item.ui_options) {
        if ((item.ui_options as InputOptions & { input_type })?.input_type) {
          (item.ui_options as InputOptions).inputType = (
            item.ui_options as InputOptions & { input_type }
          ).input_type;
        }
        uiConf[item.name]['ui:options'] = item.ui_options;
      }
      if (item.required) {
        required.push(item.name);
      }
    });
    const result = {
      title: data?.name || '',
      required,
      properties,
    };
    setSchema(result);
    setUISchema(uiConf);
    setFormData(mergeFormData(initFormData(result), formData));
  }, [data?.config_fields]);

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
  const refreshConfig: FormKit['refreshConfig'] = async () => {
    refreshPluginConfig();
  };
  const handleOnChange = (form) => {
    setFormData(form);
  };
  return (
    <>
      <h3 className="mb-4">{data?.name}</h3>
      <SchemaForm
        schema={schema}
        uiSchema={uiSchema}
        refreshConfig={refreshConfig}
        formData={formData}
        onSubmit={onSubmit}
        onChange={handleOnChange}
      />
    </>
  );
};

export default Config;
