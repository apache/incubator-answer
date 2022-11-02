import React, { FC, FormEvent, useEffect, useState } from 'react';
import { Form, Button, Image, Stack } from 'react-bootstrap';
import { Trans, useTranslation } from 'react-i18next';

import { useToast } from '@answer/hooks';
import {
  LangsType,
  FormDataType,
  AdminSettingsInterface,
} from '@answer/common/interface';
import { interfaceStore } from '@answer/stores';
import { UploadImg } from '@answer/components';
import { TIMEZONES, DEFAULT_TIMEZONE } from '@answer/common/constants';

import {
  languages,
  uploadAvatar,
  updateInterfaceSetting,
  useInterfaceSetting,
  useThemeOptions,
} from '@/services';

const Interface: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.interface',
  });
  const { update: interfaceStoreUpdate } = interfaceStore();
  const { data: themes } = useThemeOptions();
  const Toast = useToast();
  const [langs, setLangs] = useState<LangsType[]>();
  const { data: setting } = useInterfaceSetting();

  const [formData, setFormData] = useState<FormDataType>({
    logo: {
      value: setting?.logo || '',
      isInvalid: false,
      errorMsg: '',
    },
    theme: {
      value: setting?.theme || '',
      isInvalid: false,
      errorMsg: '',
    },
    language: {
      value: setting?.language || '',
      isInvalid: false,
      errorMsg: '',
    },
    time_zone: {
      value: setting?.time_zone || DEFAULT_TIMEZONE,
      isInvalid: false,
      errorMsg: '',
    },
  });
  const getLangs = async () => {
    const res: LangsType[] = await languages();
    setLangs(res);
    if (!formData.language.value) {
      // set default theme value
      setFormData({
        ...formData,
        language: {
          value: res[0].value,
          isInvalid: false,
          errorMsg: '',
        },
      });
    }
  };
  // set default theme value
  if (!formData.theme.value && Array.isArray(themes) && themes.length) {
    setFormData({
      ...formData,
      theme: {
        value: themes[0].value,
        isInvalid: false,
        errorMsg: '',
      },
    });
  }

  const checkValidated = (): boolean => {
    let ret = true;
    const { theme, language } = formData;
    const formCheckData = { ...formData };
    if (!theme.value) {
      ret = false;
      formCheckData.theme = {
        value: '',
        isInvalid: true,
        errorMsg: t('theme.msg'),
      };
    }
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
      logo: formData.logo.value,
      theme: formData.theme.value,
      language: formData.language.value,
      time_zone: formData.time_zone.value,
    };

    updateInterfaceSetting(reqParams)
      .then(() => {
        Toast.onShow({
          msg: t('update', { keyPrefix: 'toast' }),
          variant: 'success',
        });
        interfaceStoreUpdate(reqParams);
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData[err.key].isInvalid = true;
          formData[err.key].errorMsg = err.value;
        }
        setFormData({ ...formData });
      });
  };
  const imgUpload = (file: any) => {
    return new Promise((resolve) => {
      uploadAvatar(file).then((res) => {
        setFormData({
          ...formData,
          logo: {
            value: res,
            isInvalid: false,
            errorMsg: '',
          },
        });
        resolve(true);
      });
    });
  };
  const onChange = (fieldName, fieldValue) => {
    if (!formData[fieldName]) {
      return;
    }
    const fieldData: FormDataType = {
      [fieldName]: {
        value: fieldValue,
        isInvalid: false,
        errorMsg: '',
      },
    };
    setFormData({ ...formData, ...fieldData });
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

  console.log('formData', formData);
  return (
    <>
      <h3 className="mb-4">{t('page_title')}</h3>
      <Form noValidate onSubmit={onSubmit}>
        <Form.Group controlId="logo" className="mb-3">
          <Form.Label>{t('logo.label')}</Form.Label>
          <Stack gap={2}>
            <div
              className="bg-light overflow-hidden"
              style={{ width: '288px', height: '96px' }}>
              {formData.logo.value ? (
                <Image
                  width="288"
                  height="96"
                  className="object-fit-contain"
                  src={formData.logo.value}
                />
              ) : null}
            </div>
            <div className="d-inline-flex">
              <UploadImg type="logo" upload={imgUpload} />
            </div>
          </Stack>
          <Form.Text as="div" className="text-muted">
            <Trans i18nKey="admin.interface.logo.text">
              You can upload your image or
              <Button
                variant="link"
                size="sm"
                className="p-0 mx-1"
                onClick={(evt) => {
                  evt.preventDefault();
                  onChange('logo', '');
                }}>
                reset it
              </Button>
              to the site title text.
            </Trans>
          </Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.logo.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="theme" className="mb-3">
          <Form.Label>{t('theme.label')}</Form.Label>
          <Form.Select
            value={formData.theme.value}
            isInvalid={formData.theme.isInvalid}
            onChange={(evt) => {
              onChange('theme', evt.target.value);
            }}>
            {themes?.map((item) => {
              return (
                <option value={item.value} key={item.value}>
                  {item.label}
                </option>
              );
            })}
          </Form.Select>
          <Form.Text as="div">{t('theme.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.theme.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="language" className="mb-3">
          <Form.Label>{t('language.label')}</Form.Label>
          <Form.Select
            value={formData.language.value}
            isInvalid={formData.language.isInvalid}
            onChange={(evt) => {
              onChange('language', evt.target.value);
            }}>
            {langs?.map((item) => {
              return (
                <option value={item.value} key={item.value}>
                  {item.label}
                </option>
              );
            })}
          </Form.Select>
          <Form.Text as="div">{t('language.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.language.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Form.Group controlId="time-zone" className="mb-3">
          <Form.Label>{t('time_zone.label')}</Form.Label>
          <Form.Select
            value={formData.time_zone.value}
            isInvalid={formData.time_zone.isInvalid}
            onChange={(evt) => {
              onChange('time_zone', evt.target.value);
            }}>
            {TIMEZONES?.map((item) => {
              return (
                <option value={item.value} key={item.value}>
                  {item.label}
                </option>
              );
            })}
          </Form.Select>
          <Form.Text as="div">{t('time_zone.text')}</Form.Text>
          <Form.Control.Feedback type="invalid">
            {formData.time_zone.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
        <Button variant="primary" type="submit">
          {t('save', { keyPrefix: 'btns' })}
        </Button>
      </Form>
    </>
  );
};

export default Interface;
