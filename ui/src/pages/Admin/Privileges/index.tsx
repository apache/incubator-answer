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

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.privilege',
  });
  const Toast = useToast();
  const [privilege, setPrivilege] = useState<AdminSettingsPrivilege>();
  const [schema, setSchema] = useState<JSONSchema>({
    title: t('title'),
    properties: {},
  });
  const [uiSchema, setUiSchema] = useState<UISchema>({
    level: {
      'ui:widget': 'select',
    },
  });
  const [formData, setFormData] = useState<FormDataType>(initFormData(schema));

  const setFormConfig = (selectedLevel: number = 1) => {
    selectedLevel = Number(selectedLevel);
    const levelOptions = privilege?.options;
    const curLevel = levelOptions?.find((li) => {
      return li.level === selectedLevel;
    });
    if (!levelOptions || !curLevel) {
      return;
    }
    const uiState = {
      level: uiSchema.level,
    };
    const props: JSONSchema['properties'] = {
      level: {
        type: 'number',
        title: t('level.label'),
        description: t('level.text'),
        enum: levelOptions.map((_) => _.level),
        enumNames: levelOptions.map((_) => _.level_desc),
        default: selectedLevel,
      },
    };
    curLevel.privileges.forEach((li) => {
      props[li.key] = {
        type: 'number',
        title: li.label,
        default: li.value,
      };
      uiState[li.key] = {
        'ui:options': {
          readOnly: true,
        },
      };
    });
    const schemaState = {
      ...schema,
      properties: props,
    };
    const formState = initFormData(schemaState);
    curLevel.privileges.forEach((li) => {
      formState[li.key] = {
        value: li.value,
        isInvalid: false,
        errorMsg: '',
      };
    });
    setSchema(schemaState);
    setUiSchema(uiState);
    setFormData(formState);
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
    if (!privilege) {
      return;
    }
    setFormConfig(privilege.selected_level);
  }, [privilege]);
  useEffect(() => {
    getPrivilegeSetting().then((resp) => {
      setPrivilege(resp);
    });
  }, []);
  const handleOnChange = (state) => {
    setFormConfig(state.level.value);
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
