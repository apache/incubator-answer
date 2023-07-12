import React, { FC, FormEvent, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classname from 'classnames';

import { useToast } from '@/hooks';
import type { FormDataType, ImgCodeRes } from '@/common/interface';
import { modifyPassword, checkImgCode } from '@/services';
import { handleFormError } from '@/utils';
import { loggedUserInfoStore } from '@/stores';
import { PicAuthCodeModal } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.account',
  });
  const { user } = loggedUserInfoStore();
  const [showForm, setFormState] = useState(false);
  const toast = useToast();
  const [formData, setFormData] = useState<FormDataType>({
    old_pass: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    pass: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    pass2: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });
  const [showModal, setModalState] = useState(false);
  const [imgCode, setImgCode] = useState<ImgCodeRes>({
    captcha_id: '',
    captcha_img: '',
    verify: false,
  });

  const getImgCode = () => {
    checkImgCode({
      action: 'modify_pass',
    }).then((res) => {
      setImgCode(res);
    });
  };

  const handleFormState = () => {
    setFormState((pre) => !pre);
  };

  const handleChange = (params: FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  const checkValidated = (): boolean => {
    let bol = true;
    const { old_pass, pass, pass2 } = formData;
    if (!old_pass.value && user.have_password) {
      bol = false;
      formData.old_pass = {
        value: '',
        isInvalid: true,
        errorMsg: t('current_pass.msg.empty'),
      };
    }

    if (!pass.value) {
      bol = false;
      formData.pass = {
        value: '',
        isInvalid: true,
        errorMsg: t('current_pass.msg.empty'),
      };
    }

    if (bol && pass.value && pass.value.length < 8) {
      bol = false;
      formData.pass = {
        value: pass.value,
        isInvalid: true,
        errorMsg: t('current_pass.msg.length'),
      };
    }

    if (!pass2.value) {
      bol = false;
      formData.pass2 = {
        value: '',
        isInvalid: true,
        errorMsg: t('current_pass.msg.empty'),
      };
    }

    if (bol && pass2.value && pass2.value.length < 8) {
      bol = false;
      formData.pass2 = {
        value: pass2.value,
        isInvalid: true,
        errorMsg: t('current_pass.msg.length'),
      };
    }
    if (bol && pass.value && pass2.value && pass.value !== pass2.value) {
      bol = false;
      formData.pass2 = {
        value: pass2.value,
        isInvalid: true,
        errorMsg: t('current_pass.msg.different'),
      };
    }
    setFormData({
      ...formData,
    });
    return bol;
  };

  const postModifyPass = (event?: any) => {
    if (event) {
      event.preventDefault();
    }
    const params: any = {
      old_pass: formData.old_pass.value,
      pass: formData.pass.value,
    };

    if (imgCode.verify) {
      params.captcha_code = formData.captcha_code.value;
      params.captcha_id = imgCode.captcha_id;
    }
    modifyPassword(params)
      .then(() => {
        setModalState(false);
        toast.onShow({
          msg: t('update_password', { keyPrefix: 'toast' }),
          variant: 'success',
        });
        handleFormState();
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          if (!err.list.find((v) => v.error_field.indexOf('captcha') >= 0)) {
            setModalState(false);
          }
          setFormData({ ...data });
        }
      })
      .finally(() => {
        getImgCode();
      });
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    if (!checkValidated()) {
      return;
    }

    if (imgCode.verify) {
      setModalState(true);
      return;
    }
    postModifyPass();
  };

  return (
    <div className="mt-5">
      {showForm ? (
        <Form noValidate onSubmit={handleSubmit}>
          <Form.Group
            controlId="oldPass"
            className={classname('mb-3', user.have_password ? '' : 'd-none')}>
            <Form.Label>{t('current_pass.label')}</Form.Label>
            <Form.Control
              autoComplete="off"
              required
              type="password"
              placeholder=""
              isInvalid={formData.old_pass.isInvalid}
              onChange={(e) =>
                handleChange({
                  old_pass: {
                    value: e.target.value,
                    isInvalid: false,
                    errorMsg: '',
                  },
                })
              }
            />
            <Form.Control.Feedback type="invalid">
              {formData.old_pass.errorMsg}
            </Form.Control.Feedback>
          </Form.Group>

          <Form.Group controlId="newPass" className="mb-3">
            <Form.Label>{t('new_pass.label')}</Form.Label>
            <Form.Control
              autoComplete="off"
              required
              type="password"
              maxLength={32}
              isInvalid={formData.pass.isInvalid}
              onChange={(e) =>
                handleChange({
                  pass: {
                    value: e.target.value,
                    isInvalid: false,
                    errorMsg: '',
                  },
                })
              }
            />
            <Form.Control.Feedback type="invalid">
              {formData.pass.errorMsg}
            </Form.Control.Feedback>
          </Form.Group>

          <Form.Group controlId="newPass2" className="mb-3">
            <Form.Label>{t('pass_confirm.label')}</Form.Label>
            <Form.Control
              autoComplete="off"
              required
              type="password"
              maxLength={32}
              isInvalid={formData.pass2.isInvalid}
              onChange={(e) =>
                handleChange({
                  pass2: {
                    value: e.target.value,
                    isInvalid: false,
                    errorMsg: '',
                  },
                })
              }
            />
            <Form.Control.Feedback type="invalid">
              {formData.pass2.errorMsg}
            </Form.Control.Feedback>
          </Form.Group>
          <div>
            <Button type="submit" variant="primary" className="me-2">
              {t('save', { keyPrefix: 'btns' })}
            </Button>

            <Button variant="link" onClick={() => handleFormState()}>
              {t('cancel', { keyPrefix: 'btns' })}
            </Button>
          </div>
        </Form>
      ) : (
        <>
          <Form.Label>{t('password_title')}</Form.Label>
          <br />
          <Button
            variant="outline-secondary"
            type="submit"
            onClick={() => {
              handleFormState();
              getImgCode();
            }}>
            {t('change_pass_btn')}
          </Button>
        </>
      )}

      <PicAuthCodeModal
        visible={showModal}
        data={{
          captcha: formData.captcha_code,
          imgCode,
        }}
        handleCaptcha={handleChange}
        clickSubmit={postModifyPass}
        refreshImgCode={getImgCode}
        onClose={() => setModalState(false)}
      />
    </div>
  );
};

export default React.memo(Index);
