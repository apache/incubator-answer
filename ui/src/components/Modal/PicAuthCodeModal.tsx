import { useEffect } from 'react';
import { Modal, Form, Button, InputGroup } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { Icon } from '@/components';
import type { FormValue, FormDataType, ImgCodeRes } from '@/common/interface';
import { CAPTCHA_CODE_STORAGE_KEY } from '@/common/constants';
import Storage from '@/utils/storage';

interface IProps {
  /** control visible */
  visible: boolean;
  data: {
    captcha: FormValue;
    imgCode: ImgCodeRes;
  };
  handleCaptcha: (parma: FormDataType) => void;
  clickSubmit: (e: any) => void;
  refreshImgCode: () => void;
  onClose: () => void;
}

const Index: React.FC<IProps> = ({
  visible,
  data,
  handleCaptcha,
  clickSubmit,
  refreshImgCode,
  onClose,
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'pic_auth_code' });
  const { captcha, imgCode } = data;

  useEffect(() => {
    if (visible) {
      refreshImgCode();
    }
  }, [visible]);

  return (
    <Modal size="sm" title="Captcha" show={visible} onHide={onClose} centered>
      <Modal.Header closeButton>
        <Modal.Title as="h5">{t('title')}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form noValidate onSubmit={clickSubmit}>
          <Form.Group controlId="code" className="mb-3">
            <div className="mb-3">
              <img
                src={imgCode?.captcha_img}
                alt="code"
                width="auto"
                height="40px"
              />
            </div>
            <InputGroup>
              <Form.Control
                type="text"
                autoComplete="off"
                placeholder={t('placeholder')}
                isInvalid={captcha?.isInvalid}
                onChange={(e) => {
                  Storage.set(CAPTCHA_CODE_STORAGE_KEY, e.target.value);
                  handleCaptcha({
                    captcha_code: {
                      value: e.target.value,
                      isInvalid: false,
                      errorMsg: '',
                    },
                  });
                }}
              />
              <Button
                onClick={refreshImgCode}
                variant="outline-secondary"
                style={{
                  borderTopRightRadius: '0.375rem',
                  borderBottomRightRadius: '0.375rem',
                }}>
                <Icon name="arrow-repeat" />
              </Button>

              <Form.Control.Feedback type="invalid">
                {captcha?.errorMsg}
              </Form.Control.Feedback>
            </InputGroup>
          </Form.Group>

          <div className="d-grid">
            <Button type="submit">{t('verify', { keyPrefix: 'btns' })}</Button>
          </div>
        </Form>
      </Modal.Body>
    </Modal>
  );
};
export default Index;
