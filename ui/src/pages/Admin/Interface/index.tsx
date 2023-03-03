import { FC, FormEvent, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import { useToast } from '@/hooks';
import {
  LangsType,
  FormDataType,
  AdminSettingsInterface,
} from '@/common/interface';
import { interfaceStore } from '@/stores';
import { JSONSchema, SchemaForm, UISchema } from '@/components';
import { DEFAULT_TIMEZONE, SYSTEM_AVATAR_OPTIONS } from '@/common/constants';
import { updateInterfaceSetting, useInterfaceSetting } from '@/services';
import {
  setupAppLanguage,
  loadLanguageOptions,
  setupAppTimeZone,
} from '@/utils/localize';
import { handleFormError } from '@/utils';

const Interface: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.interface',
  });
  const storeInterface = interfaceStore.getState().interface;
  const Toast = useToast();
  const [langs, setLangs] = useState<LangsType[]>();
  const { data: setting } = useInterfaceSetting();

  console.log('setting', langs);

  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      language: {
        type: 'string',
        title: t('language.label'),
        description: t('language.text'),
        enum: langs?.map((lang) => lang.value),
        enumNames: langs?.map((lang) => lang.label),
      },
      time_zone: {
        type: 'string',
        title: t('time_zone.label'),
        description: t('time_zone.text'),
      },
      default_avatar: {
        type: 'string',
        title: t('avatar.label'),
        description: t('avatar.text'),
        enum: SYSTEM_AVATAR_OPTIONS?.map((v) => v.value),
        enumNames: SYSTEM_AVATAR_OPTIONS?.map((v) => v.label),
      },
    },
  };

  const [formData, setFormData] = useState<FormDataType>({
    language: {
      value: setting?.language || storeInterface.language,
      isInvalid: false,
      errorMsg: '',
    },
    time_zone: {
      value: setting?.time_zone || DEFAULT_TIMEZONE,
      isInvalid: false,
      errorMsg: '',
    },
    default_avatar: {
      value: setting?.default_avatar || 'System',
      isInvalid: false,
      errorMsg: '',
    },
  });

  const uiSchema: UISchema = {
    language: {
      'ui:widget': 'select',
    },
    time_zone: {
      'ui:widget': 'timezone',
    },
    default_avatar: {
      'ui:widget': 'select',
    },
  };
  const getLangs = async () => {
    const res: LangsType[] = await loadLanguageOptions(true);
    setLangs(res);
  };

  const checkValidated = (): boolean => {
    let ret = true;
    const { language } = formData;
    const formCheckData = { ...formData };
    if (!language.value) {
      ret = false;
      formCheckData.language = {
        value: '',
        isInvalid: true,
        errorMsg: t('language.msg'),
      };
    }
    setFormData({
      ...formCheckData,
    });
    return ret;
  };
  const onSubmit = (evt: FormEvent) => {
    evt.preventDefault();
    evt.stopPropagation();
    if (checkValidated() === false) {
      return;
    }
    const reqParams: AdminSettingsInterface = {
      language: formData.language.value,
      time_zone: formData.time_zone.value,
      default_avatar: formData.default_avatar.value,
    };

    updateInterfaceSetting(reqParams)
      .then(() => {
        interfaceStore.getState().update(reqParams);
        setupAppLanguage();
        setupAppTimeZone();
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
    if (setting) {
      const formMeta = {};
      Object.keys(setting).forEach((k) => {
        formMeta[k] = { ...formData[k], value: setting[k] };
      });
      setFormData({ ...formData, ...formMeta });
    }
  }, [setting]);
  useEffect(() => {
    getLangs();
  }, []);

  const handleOnChange = (data) => {
    setFormData(data);
  };
  return (
    <>
      <h3 className="mb-4">{t('page_title')}</h3>
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
