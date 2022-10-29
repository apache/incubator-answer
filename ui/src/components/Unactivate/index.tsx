import React, { useState, useEffect } from 'react';
import { Button, Col } from 'react-bootstrap';
import { Trans, useTranslation } from 'react-i18next';

import { resendEmail, checkImgCode } from '@/services';
import { PicAuthCodeModal } from '@answer/components/Modal';
import type {
  ImgCodeRes,
  ImgCodeReq,
  FormDataType,
} from '@answer/common/interface';
import { loggedUserInfoStore } from '@answer/stores';

import { CAPTCHA_CODE_STORAGE_KEY } from '@/common/constants';
import Storage from '@/utils/storage';

interface IProps {
  visible: boolean;
}

const Index: React.FC<IProps> = ({ visible = false }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'inactive' });
  const [isSuccess, setSuccess] = useState(false);
  const [showModal, setModalState] = useState(false);
  const { e_mail } = loggedUserInfoStore((state) => state.user);
  const [formData, setFormData] = useState<FormDataType>({
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

  const getImgCode = () => {
    checkImgCode({
      action: 'e_mail',
    }).then((res) => {
      setImgCode(res);
    });
  };

  const submit = (e?: any) => {
    if (e) {
      e.preventDefault();
    }
    let obj: ImgCodeReq = {};
    if (imgCode.verify) {
      const code = Storage.get(CAPTCHA_CODE_STORAGE_KEY) || '';
      obj = {
        captcha_code: code,
        captcha_id: imgCode.captcha_id,
      };
    }
    resendEmail(obj)
      .then(() => {
        setSuccess(true);
        setModalState(false);
      })
      .catch((err) => {
        if (err.isError && err.key) {
          formData[err.key].isInvalid = true;
          formData[err.key].errorMsg = err.value;
        }
        setFormData({ ...formData });
      })
      .finally(() => {
        getImgCode();
      });
  };

  const onSentEmail = () => {
    if (imgCode.verify) {
      setModalState(true);
      if (!formData.captcha_code.value) {
        setFormData({
          captcha_code: {
            value: '',
            isInvalid: false,
            errorMsg: t('msg.empty'),
          },
        });
      }
      return;
    }
    submit();
  };

  const handleChange = (params: FormDataType) => {
    setFormData({ ...formData, ...params });
  };

  useEffect(() => {
    if (visible) {
      getImgCode();
    }
  }, [visible]);

  return (
    <Col md={6} className="mx-auto text-center">
      {isSuccess ? (
        <p>
          <Trans
            i18nKey="inactive.another"
            values={{ mail: e_mail }}
            components={{ bold: <strong /> }}
          />
        </p>
      ) : (
        <>
          <p>
            <Trans
              i18nKey="inactive.first"
              values={{ mail: e_mail }}
              components={{ bold: <strong /> }}
            />
          </p>
          <p>{t('info')}</p>
          <Button variant="link" onClick={onSentEmail}>
            {t('btn_name')}
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
        clickSubmit={submit}
        refreshImgCode={getImgCode}
        onClose={() => setModalState(false)}
      />
    </Col>
  );
};

export default React.memo(Index);
