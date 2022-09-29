import { FC, memo, useEffect, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { passRetrieve, checkImgCode } from '@answer/api';
import type { FormDataType } from '@/common/interface';
import type { ImgCodeRes, PssRetReq } from '@/services/types';
import { PicAuthCodeModal } from '@/components/Modal';

interface IProps {
  visible: boolean;
  callback: (param: number, email: string) => void;
}

const Index: FC<IProps> = ({ visible = false, callback }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'account_forgot' });
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

  const getImgCode = () => {
    checkImgCode({
      action: 'find_pass',
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
    const params: PssRetReq = {
      e_mail: formData.e_mail.value,
    };
    if (imgCode.verify) {
      params.captcha_code = formData.captcha_code.value;
      params.captcha_id = imgCode.captcha_id;
    }

    passRetrieve(params)
      .then(() => {
        callback?.(2, formData.e_mail.value);
        setModalState(false);
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData[err.key].isInvalid = true;
          formData[err.key].errorMsg = err.value;
          if (err.key.indexOf('captcha') < 0) {
            setModalState(false);
          }
        }
        setFormData({ ...formData });
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

  useEffect(() => {
    if (visible) {
      getImgCode();
    }
  }, [visible]);

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
            {t('btn_name')}
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
