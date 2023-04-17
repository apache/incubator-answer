import { FC, FormEvent, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { useToast } from '@/hooks';
import { FormDataType } from '@/common/interface';
import { JSONSchema, SchemaForm, UISchema, initFormData } from '@/components';
import { SYSTEM_AVATAR_OPTIONS } from '@/common/constants';
import {
  getUsersSetting,
  putUsersSetting,
  AdminSettingsUsers,
} from '@/services';
import { handleFormError } from '@/utils';
import * as Type from '@/common/interface';

const Interface: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.settings_users',
  });
  const Toast = useToast();

  const schema: JSONSchema = {
    title: t('title'),
    properties: {
      default_avatar: {
        type: 'string',
        title: t('avatar.label'),
        description: t('avatar.text'),
        enum: SYSTEM_AVATAR_OPTIONS?.map((v) => v.value),
        enumNames: SYSTEM_AVATAR_OPTIONS?.map((v) => v.label),
        default: 'system',
      },
      profile_editable: {
        type: 'string',
        title: t('profile_editable.title'),
      },
    },
  };

  const [formData, setFormData] = useState<FormDataType>(initFormData(schema));

  const uiSchema: UISchema = {
    default_avatar: {
      'ui:widget': 'select',
    },
    profile_editable: {
      'ui:widget': 'legend',
    },
    profile_displayname: {
      'ui:widget': 'legend',
    },
  };

  const onSubmit = (evt: FormEvent) => {
    evt.preventDefault();
    evt.stopPropagation();
    // @ts-ignore
    const reqParams: AdminSettingsUsers = {
      default_avatar: '',
    };

    // @ts-ignore
    putUsersSetting(reqParams)
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
    getUsersSetting().then((resp) => {
      const formMeta: Type.FormDataType = {};
      Object.keys(formData).forEach((k) => {
        let v = resp[k];
        if (k === 'default_avatar' && !v) {
          v = 'system';
        }
        formMeta[k] = { ...formData[k], value: v };
      });
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

export default Interface;
