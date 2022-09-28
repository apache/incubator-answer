import React, { useEffect, useState, FormEvent } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import dayjs from 'dayjs';
import en from 'dayjs/locale/en';
import zh from 'dayjs/locale/zh-cn';

import { languages } from '@answer/api';
import type { FormDataType } from '@answer/common/interface';
import type { LangsType } from '@answer/services/types';
import { useToast } from '@answer/hooks';
import Storage from '@answer/utils/storage';

const Index = () => {
  const { t, i18n } = useTranslation('translation', {
    keyPrefix: 'settings.interface',
  });
  const toast = useToast();
  const [langs, setLangs] = useState<LangsType[]>();
  const [formData, setFormData] = useState<FormDataType>({
    lang: {
      value: true,
      isInvalid: false,
      errorMsg: '',
    },
  });

  const getLangs = async () => {
    const res: LangsType[] = await languages();
    setLangs(res);
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();

    Storage.set('LANG', formData.lang.value);
    dayjs.locale(formData.lang.value === 'en_US' ? en : zh);
    i18n.changeLanguage(formData.lang.value);
    toast.onShow({
      msg: t('update', { keyPrefix: 'toast' }),
      variant: 'success',
    });
  };

  useEffect(() => {
    getLangs();
    const lang = Storage.get('LANG');
    if (lang) {
      setFormData({
        lang: {
          value: lang,
          isInvalid: false,
          errorMsg: '',
        },
      });
    }
  }, []);
  return (
    <Form noValidate onSubmit={handleSubmit}>
      <Form.Group controlId="emailSend" className="mb-3">
        <Form.Label>{t('lang.label')}</Form.Label>

        <Form.Select
          value={formData.lang.value}
          isInvalid={formData.lang.isInvalid}
          onChange={(e) => {
            setFormData({
              lang: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            });
          }}>
          {langs?.map((item) => {
            return (
              <option value={item.value} key={item.value}>
                {item.label}
              </option>
            );
          })}
        </Form.Select>
        <Form.Text as="div">{t('lang.text')}</Form.Text>
        <Form.Control.Feedback type="invalid">
          {formData.lang.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <Button variant="primary" type="submit">
        {t('save', { keyPrefix: 'btns' })}
      </Button>
    </Form>
  );
};

export default React.memo(Index);
