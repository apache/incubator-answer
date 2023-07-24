import { useEffect, useRef, useState, useLayoutEffect } from 'react';
import { Modal, Form, Button, InputGroup } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

import { Icon } from '@/components';
import type {
  FormValue,
  ImgCodeRes,
  CaptchaKey,
  FieldError,
  ImgCodeReq,
} from '@/common/interface';
import { checkImgCode } from '@/services';

type SubmitCallback = {
  (): void;
};

const Index = (captchaKey: CaptchaKey) => {
  const refRoot = useRef(null);
  if (refRoot.current === null) {
    // @ts-ignore
    refRoot.current = ReactDOM.createRoot(document.createElement('div'));
  }

  const { t } = useTranslation('translation', { keyPrefix: 'pic_auth_code' });
  const refKey = useRef<CaptchaKey>(captchaKey);
  const refCallback = useRef<SubmitCallback>();
  const pending = useRef(false);
  const autoInitCaptchaData = /email/i.test(refKey.current);

  const [stateShow, setStateShow] = useState(false);
  const [captcha, setCaptcha] = useState<ImgCodeRes>({
    captcha_id: '',
    captcha_img: '',
    verify: false,
  });
  const [imgCode, setImgCode] = useState<FormValue>({
    value: '',
    isInvalid: false,
    errorMsg: '',
  });
  const refCaptcha = useRef(captcha);
  const refImgCode = useRef(imgCode);

  const fetchCaptchaData = () => {
    pending.current = true;
    checkImgCode(refKey.current)
      .then((resp) => {
        setCaptcha(resp);
      })
      .finally(() => {
        pending.current = false;
      });
  };

  const resetCapture = () => {
    setCaptcha({
      captcha_id: '',
      captcha_img: '',
      verify: false,
    });
  };

  const show = () => {
    if (!stateShow) {
      setStateShow(true);
    }
  };
  /**
   * There are some cases where the React scheduler cancels the execution of some functions,
   * which prevents them from closing properly:
   *  for example, if the parent component uninstalls the child component directly,
   *  and the `captchaModal.close()` call is inside the child component.
   * In this case, call `await captchaModal.close()` and wait for the close action to complete.
   */
  const close = (reset = true) => {
    setStateShow(false);
    if (reset) {
      resetCapture();
    }
    const p = new Promise<void>((resolve) => {
      setTimeout(resolve);
    });
    return p;
  };

  const handleCaptchaError = (fel: FieldError[] = []) => {
    const captchaErr = fel.find((o) => {
      return o.error_field === 'captcha_code';
    });

    const ri = refImgCode.current;
    if (captchaErr) {
      /**
       * `imgCode.value` No value but a validation error is received,
       * indicating that it is the first time the interface has returned a CAPTCHA error,
       * triggering the CAPTCHA logic. There is no need to display the error message at this point.
       */
      if (ri.value) {
        setImgCode({
          ...ri,
          isInvalid: true,
          errorMsg: captchaErr.error_msg,
        });
      }
      fetchCaptchaData();
      show();
    } else {
      setImgCode({
        ...ri,
        isInvalid: false,
        errorMsg: '',
      });
      close();
    }
  };

  const handleChange = (evt) => {
    evt.preventDefault();
    setImgCode({
      value: evt.target.value || '',
      isInvalid: false,
      errorMsg: '',
    });
  };

  const getCaptcha = () => {
    const rc = refCaptcha.current;
    const ri = refImgCode.current;
    const r = {
      verify: !!rc?.verify,
      captcha_id: rc?.captcha_id,
      captcha_code: ri.value,
    };

    return r;
  };

  const resolveCaptchaReq = (req: ImgCodeReq) => {
    const r = getCaptcha();
    if (r.verify) {
      req.captcha_code = r.captcha_code;
      req.captcha_id = r.captcha_id;
    }
  };

  const handleSubmit = (evt) => {
    evt.preventDefault();
    if (!imgCode.value) {
      return;
    }

    if (refCallback.current) {
      refCallback.current();
    }
  };

  useEffect(() => {
    if (autoInitCaptchaData) {
      fetchCaptchaData();
    }
  }, []);

  useLayoutEffect(() => {
    refImgCode.current = imgCode;
    refCaptcha.current = captcha;
  }, [captcha, imgCode]);

  useEffect(() => {
    // @ts-ignore
    refRoot.current.render(
      <Modal
        size="sm"
        title="Captcha"
        show={stateShow}
        onHide={() => close(false)}
        centered>
        <Modal.Header closeButton>
          <Modal.Title as="h5">{t('title')}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form noValidate onSubmit={handleSubmit}>
            <Form.Group controlId="code" className="mb-3">
              <div className="mb-3 p-2 d-flex align-items-center justify-content-center bg-light rounded-2">
                <img
                  src={captcha?.captcha_img}
                  alt="captcha img"
                  width="auto"
                  height="40px"
                />
              </div>
              <InputGroup>
                <Form.Control
                  type="text"
                  autoComplete="off"
                  placeholder={t('placeholder')}
                  isInvalid={imgCode?.isInvalid}
                  onChange={handleChange}
                  value={imgCode.value}
                />
                <Button
                  onClick={fetchCaptchaData}
                  variant="outline-secondary"
                  title={t('refresh', { keyPrefix: 'btns' })}
                  style={{
                    borderTopRightRadius: '0.375rem',
                    borderBottomRightRadius: '0.375rem',
                  }}>
                  <Icon name="arrow-repeat" />
                </Button>

                <Form.Control.Feedback type="invalid">
                  {imgCode?.errorMsg}
                </Form.Control.Feedback>
              </InputGroup>
            </Form.Group>

            <div className="d-grid">
              <Button type="submit" disabled={!imgCode.value}>
                {t('verify', { keyPrefix: 'btns' })}
              </Button>
            </div>
          </Form>
        </Modal.Body>
      </Modal>,
    );
  });

  const r = {
    close,
    show,
    check: (submitFunc: SubmitCallback) => {
      if (pending.current) {
        return false;
      }
      refCallback.current = submitFunc;
      if (captcha?.verify) {
        show();
        return false;
      }
      return submitFunc();
    },
    getCaptcha,
    resolveCaptchaReq,
    fetchCaptchaData,
    handleCaptchaError,
  };

  return r;
};

export default Index;
