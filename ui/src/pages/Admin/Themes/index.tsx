import { FC, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { getThemeSetting, putThemeSetting } from '@/services';
import { SchemaForm, JSONSchema, initFormData, UISchema } from '@/components';
import { useToast } from '@/hooks';
import { handleFormError } from '@/utils';
import { themeSettingStore } from '@/stores';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.themes',
  });
  const Toast = useToast();
  const [themeSetting, setThemeSetting] = useState<Type.AdminSettingsTheme>();
  const schema: JSONSchema = {
    title: t('page_title'),
    properties: {
      themes: {
        type: 'string',
        title: t('themes.label'),
        description: t('themes.text'),
        enum: themeSetting?.theme_options?.map((_) => _.value),
        enumNames: themeSetting?.theme_options?.map((_) => _.label),
        default: themeSetting?.theme_options?.[0]?.value,
      },
      navbar_style: {
        type: 'string',
        title: t('navbar_style.label'),
        description: t('navbar_style.text'),
        enum: ['colored', 'light'],
        enumNames: ['Colored', 'Light'],
        default: 'colored',
      },
      primary_color: {
        type: 'string',
        title: t('primary_color.label'),
        description: t('primary_color.text'),
        default: '#0033FF',
      },
    },
  };
  const uiSchema: UISchema = {
    themes: {
      'ui:widget': 'select',
    },
    navbar_style: {
      'ui:widget': 'select',
    },
    primary_color: {
      'ui:options': {
        inputType: 'color',
      },
    },
  };
  const [formData, setFormData] = useState(initFormData(schema));
  const { update: updateThemeSetting } = themeSettingStore((_) => _);
  const onSubmit = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();
    const themeName = formData.themes.value;
    const reqParams: Type.AdminSettingsTheme = {
      theme: themeName,
      theme_config: {
        [themeName]: {
          navbar_style: formData.navbar_style.value,
          primary_color: formData.primary_color.value,
        },
      },
    };

    putThemeSetting(reqParams)
      .then(() => {
        Toast.onShow({
          msg: t('update', { keyPrefix: 'toast' }),
          variant: 'success',
        });
        updateThemeSetting(reqParams);
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          setFormData({ ...data });
        }
      });
  };

  useEffect(() => {
    getThemeSetting().then((setting) => {
      if (setting) {
        setThemeSetting(setting);
        const themeName = setting.theme;
        const themeConfig = setting.theme_config[themeName];
        const formMeta = { ...formData };
        formMeta.themes.value = themeName;
        formMeta.navbar_style.value = themeConfig?.navbar_style;
        formMeta.primary_color.value = themeConfig?.primary_color;
        setFormData({ ...formMeta });
      }
    });
  }, []);

  const handleOnChange = (cd) => {
    setFormData(cd);
    const themeConfig = themeSetting?.theme_config[cd.themes.value];
    if (themeConfig) {
      themeConfig.navbar_style = cd.navbar_style.value;
      themeConfig.primary_color = cd.primary_color.value;
      setThemeSetting({
        ...themeSetting,
        theme: themeSetting?.theme,
        theme_config: themeSetting?.theme_config,
      });
    }
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
