import { FC, memo, useEffect, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';

import type {
  ImgCodeRes,
  PasswordResetReq,
  FormDataType,
} from '@/common/interface';
import { loggedUserInfoStore } from '@/stores';
import { changeEmail, checkImgCode } from '@/services';
import { PicAuthCodeModal } from '@/components/Modal';
import { handleFormError } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'change_email' });
  const [formData, setFormData] = useState<FormDataType>({
    e_mail: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    captcha_code: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });
  const [imgCode, setImgCode] = useState<ImgCodeRes>({
    captcha_id: '',
    captcha_img: '',
    verify: false,
  });
  const [showModal, setModalState] = useState(false);
  const navigate = useNavigate();
  const { user: userInfo, update: updateUser } = loggedUserInfoStore();

  const getImgCode = () => {
    checkImgCode({
      action: 'e_mail',
    }).then((res) => {
      setImgCode(res);
    });
  };

  const handleChange = (params: FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  const checkValidated = (): boolean => {
    let bol = true;

    if (!formData.e_mail.value) {
      bol = false;
      formData.e_mail = {
        value: '',
        isInvalid: true,
        errorMsg: t('email.msg.empty'),
      };
    }
    setFormData({
      ...formData,
    });
    return bol;
  };

  const sendEmail = (e?: any) => {
    if (e) {
      e.preventDefault();
    }
    const params: PasswordResetReq = {
      e_mail: formData.e_mail.value,
    };
    if (imgCode.verify) {
      params.captcha_code = formData.captcha_code.value;
      params.captcha_id = imgCode.captcha_id;
    }

    changeEmail(params)
      .then(() => {
        userInfo.e_mail = formData.e_mail.value;
        updateUser(userInfo);
        navigate('/users/login', { replace: true });
        setModalState(false);
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

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    event.stopPropagation();

    if (!checkValidated()) {
      return;
    }

    if (imgCode.verify) {
      setModalState(true);
      return;
    }

    sendEmail();
  };

  const goBack = () => {
    navigate('/users/login?status=inactive', { replace: true });
  };

  useEffect(() => {
    getImgCode();
  }, []);

  return (
    <>
      <Form noValidate onSubmit={handleSubmit} autoComplete="off">
        <Form.Group controlId="email" className="mb-3">
          <Form.Label>{t('email.label')}</Form.Label>
          <Form.Control
            required
            type="email"
            value={formData.e_mail.value}
            isInvalid={formData.e_mail.isInvalid}
            onChange={(e) => {
              handleChange({
                e_mail: {
                  value: e.target.value,
                  isInvalid: false,
                  errorMsg: '',
                },
              });
            }}
          />
          <Form.Control.Feedback type="invalid">
            {formData.e_mail.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>

        <div className="d-grid mb-3">
          <Button variant="primary" type="submit">
            {t('btn_update')}
          </Button>
          <Button variant="link" className="mt-2 d-block" onClick={goBack}>
            {t('btn_cancel')}
          </Button>
        </div>
      </Form>

      <PicAuthCodeModal
        visible={showModal}
        data={{
          captcha: formData.captcha_code,
          imgCode,
        }}
        handleCaptcha={handleChange}
        clickSubmit={sendEmail}
        refreshImgCode={getImgCode}
        onClose={() => setModalState(false)}
      />
    </>
  );
};

export default memo(Index);
