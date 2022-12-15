import React, { useEffect, useState, FormEvent } from 'react';
import { useTranslation } from 'react-i18next';

import type { LangsType, FormDataType } from '@/common/interface';
import { useToast } from '@/hooks';
import { updateUserInterface } from '@/services';
import { localize } from '@/utils';
import { loggedUserInfoStore } from '@/stores';
import { SchemaForm, JSONSchema, UISchema, initFormData } from '@/components';

const Index = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.interface',
  });
  const loggedUserInfo = loggedUserInfoStore.getState().user;
  const toast = useToast();
  const [langs, setLangs] = useState<LangsType[]>();
  const schema: JSONSchema = {
    title: t('heading'),
    properties: {
      lang: {
        type: 'string',
        title: t('lang.label'),
        description: t('lang.text'),
        enum: langs?.map((_) => _.value),
        enumNames: langs?.map((_) => _.label),
        default: loggedUserInfo.language,
      },
    },
  };
  const uiSchema: UISchema = {
    lang: {
      'ui:widget': 'select',
    },
  };
  const [formData, setFormData] = useState<FormDataType>(initFormData(schema));

  const getLangs = async () => {
    const res: LangsType[] = await localize.loadLanguageOptions();
    setLangs(res);
  };

  const handleOnChange = (d) => {
    setFormData(d);
  };
  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    const lang = formData.lang.value;
    updateUserInterface(lang).then(() => {
      loggedUserInfoStore.getState().update({
        ...loggedUserInfo,
        language: lang,
      });
      localize.setupAppLanguage();
      toast.onShow({
        msg: t('update', { keyPrefix: 'toast' }),
        variant: 'success',
      });
    });
  };

  useEffect(() => {
    getLangs();
  }, []);
  return (
    <>
      <h3 className="mb-4">{t('heading')}</h3>
      <SchemaForm
        schema={schema}
        uiSchema={uiSchema}
        formData={formData}
        onChange={handleOnChange}
        onSubmit={handleSubmit}
      />
    </>
  );
};

export default React.memo(Index);
