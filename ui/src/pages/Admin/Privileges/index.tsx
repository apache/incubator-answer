import { FC, FormEvent, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { useToast } from '@/hooks';
import { FormDataType } from '@/common/interface';
import { JSONSchema, SchemaForm, UISchema, initFormData } from '@/components';
import {
  getPrivilegeSetting,
  putPrivilegeSetting,
  AdminSettingsPrivilege,
} from '@/services';
import { handleFormError } from '@/utils';
import * as Type from '@/common/interface';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.privilege',
  });
  const Toast = useToast();
  const [privilege, setPrivilege] = useState<AdminSettingsPrivilege>();

  const schema: JSONSchema = {
    title: t('title'),
    properties: {
      level: {
        type: 'number',
        title: t('level.label'),
        description: t('level.text'),
        enum: privilege?.options.map((_) => _.level),
        enumNames: privilege?.options.map((_) => _.level_desc),
        default: 1,
      },
    },
  };

  const [formData, setFormData] = useState<FormDataType>(initFormData(schema));

  const uiSchema: UISchema = {
    level: {
      'ui:widget': 'select',
    },
  };

  const onSubmit = (evt: FormEvent) => {
    evt.preventDefault();
    evt.stopPropagation();
    const lv = Number(formData.level.value);
    putPrivilegeSetting(lv)
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
    getPrivilegeSetting().then((resp) => {
      setPrivilege(resp);
      const formMeta: Type.FormDataType = {};
      formMeta.level = {
        value: resp.selected_level,
        errorMsg: '',
        isInvalid: false,
      };
      setFormData({ ...formData, ...formMeta });
    });
  }, []);

  const handleOnChange = (data) => {
    setFormData(data);
  };

  return (
    <>
      <h3 className="mb-4">{t('title')}</h3>
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

export default Index;
