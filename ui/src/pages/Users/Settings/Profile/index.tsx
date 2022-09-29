import React, { FormEvent, useState, useEffect } from 'react';
import { Form, Button } from 'react-bootstrap';
import { Trans, useTranslation } from 'react-i18next';

import { marked } from 'marked';

import { modifyUserInfo, uploadAvatar, getUserInfo } from '@answer/api';
import type { FormDataType } from '@/common/interface';
import { UploadImg, Avatar } from '@answer/components';
import { userInfoStore } from '@answer/stores';
import { useToast } from '@answer/hooks';

const Index: React.FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.profile',
  });
  const toast = useToast();
  const { user, update } = userInfoStore();
  const [formData, setFormData] = useState<FormDataType>({
    display_name: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    avatar: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    bio: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    website: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    location: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });

  const handleChange = (params: FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  const avatarUpload = (file: any) => {
    return new Promise((resolve) => {
      uploadAvatar(file).then((res) => {
        setFormData({
          ...formData,
          avatar: {
            value: res,
            isInvalid: false,
            errorMsg: '',
          },
        });
        resolve(true);
      });
    });
  };

  const checkValidated = (): boolean => {
    let bol = true;
    const { display_name, website } = formData;
    if (!display_name.value) {
      bol = false;
      formData.display_name = {
        value: '',
        isInvalid: true,
        errorMsg: t('display_name.msg'),
      };
    }

    const reg = /^(http|https):\/\//g;
    if (website.value && !website.value.match(reg)) {
      bol = false;
      formData.website = {
        value: formData.website.value,
        isInvalid: true,
        errorMsg: t('website.msg'),
      };
    }
    setFormData({
      ...formData,
    });
    return bol;
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    if (checkValidated() === false) {
      return;
    }

    const params = {
      display_name: formData.display_name.value,
      avatar: formData.avatar.value,
      bio: formData.bio.value,
      website: formData.website.value,
      location: formData.location.value,
      bio_html: marked.parse(formData.bio.value),
    };

    modifyUserInfo(params)
      .then(() => {
        update({
          ...user,
          ...params,
        });
        toast.onShow({
          msg: t('update', { keyPrefix: 'toast' }),
          variant: 'success',
        });
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData[err.key].isInvalid = true;
          formData[err.key].errorMsg = err.value;
        }
        setFormData({ ...formData });
      });
  };

  const getProfile = () => {
    getUserInfo().then((res) => {
      formData.display_name.value = res.display_name;
      formData.bio.value = res.bio;
      formData.avatar.value = res.avatar;
      formData.location.value = res.location;
      formData.website.value = res.website;
      setFormData({ ...formData });
    });
  };

  useEffect(() => {
    getProfile();
  }, []);
  return (
    <Form noValidate onSubmit={handleSubmit}>
      <Form.Group controlId="displayName" className="mb-3">
        <Form.Label>{t('display_name.label')}</Form.Label>
        <Form.Control
          required
          type="text"
          value={formData.display_name.value}
          isInvalid={formData.display_name.isInvalid}
          onChange={(e) =>
            handleChange({
              display_name: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            })
          }
        />
        <Form.Control.Feedback type="invalid">
          {formData.display_name.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <Form.Group className="mb-3">
        <Form.Label>{t('avatar.label')}</Form.Label>
        <div className="d-flex align-items-center">
          <Avatar
            size="128px"
            avatar={formData.avatar.value}
            className="me-3 rounded"
          />

          <div>
            <UploadImg type="avatar" upload={avatarUpload} />
            <div>
              <Form.Text className="text-muted mt-0">
                <Trans i18nKey="settings.profile.avatar.text">
                  You can upload your image or
                  <a
                    href="@/pages/Users/Settings/Profile/index##"
                    onClick={(e) => {
                      e.preventDefault();
                      handleChange({
                        avatar: {
                          value: '',
                          isInvalid: false,
                          errorMsg: '',
                        },
                      });
                    }}>
                    reset
                  </a>
                  it to
                </Trans>
                <a href="https://gravatar.com"> gravatar.com</a>
              </Form.Text>
            </div>
          </div>
        </div>
      </Form.Group>

      <Form.Group controlId="bio" className="mb-3">
        <Form.Label>{t('bio.label')}</Form.Label>
        <Form.Control
          className="font-monospace"
          required
          as="textarea"
          rows={5}
          value={formData.bio.value}
          isInvalid={formData.bio.isInvalid}
          onChange={(e) =>
            handleChange({
              bio: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            })
          }
        />
        <Form.Control.Feedback type="invalid">
          {formData.bio.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <Form.Group controlId="website" className="mb-3">
        <Form.Label>{t('website.label')}</Form.Label>
        <Form.Control
          required
          type="text"
          placeholder={t('website.placeholder')}
          value={formData.website.value}
          isInvalid={formData.website.isInvalid}
          onChange={(e) =>
            handleChange({
              website: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            })
          }
        />
        <Form.Control.Feedback type="invalid">
          {formData.website.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <Form.Group controlId="email" className="mb-3">
        <Form.Label>{t('location.label')}</Form.Label>
        <Form.Control
          required
          type="text"
          placeholder={t('location.placeholder')}
          value={formData.location.value}
          isInvalid={formData.location.isInvalid}
          onChange={(e) =>
            handleChange({
              location: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            })
          }
        />
        <Form.Control.Feedback type="invalid">
          {formData.location.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <Button variant="primary" type="submit">
        {t('btn_name')}
      </Button>
    </Form>
  );
};

export default React.memo(Index);
