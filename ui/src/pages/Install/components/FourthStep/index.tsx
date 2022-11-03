import { FC, FormEvent } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type { FormDataType } from '@/common/interface';
import Progress from '../Progress';

interface Props {
  data: FormDataType;
  changeCallback: (value: FormDataType) => void;
  nextCallback: () => void;
  visible: boolean;
}
const Index: FC<Props> = ({ visible, data, changeCallback, nextCallback }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'install' });

  const checkValidated = (): boolean => {
    let bol = true;
    const {
      site_name,
      contact_email,
      admin_name,
      admin_password,
      admin_email,
    } = data;

    if (!site_name.value) {
      bol = false;
      data.site_name = {
        value: '',
        isInvalid: true,
        errorMsg: t('site_name.msg'),
      };
    }

    if (!contact_email.value) {
      bol = false;
      data.contact_email = {
        value: '',
        isInvalid: true,
        errorMsg: t('contact_email.msg'),
      };
    }

    if (!admin_name.value) {
      bol = false;
      data.admin_name = {
        value: '',
        isInvalid: true,
        errorMsg: t('admin_name.msg'),
      };
    }

    if (!admin_password.value) {
      bol = false;
      data.admin_password = {
        value: '',
        isInvalid: true,
        errorMsg: t('admin_password.msg'),
      };
    }

    if (!admin_email.value) {
      bol = false;
      data.admin_email = {
        value: '',
        isInvalid: true,
        errorMsg: t('admin_email.msg'),
      };
    }

    changeCallback({
      ...data,
    });
    return bol;
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    event.stopPropagation();
    if (!checkValidated()) {
      return;
    }
    nextCallback();
  };

  if (!visible) return null;
  return (
    <Form noValidate onSubmit={handleSubmit}>
      <h5>{t('site_information')}</h5>
      <Form.Group controlId="site_name" className="mb-3">
        <Form.Label>{t('site_name.label')}</Form.Label>
        <Form.Control
          required
          value={data.site_name.value}
          isInvalid={data.site_name.isInvalid}
          onChange={(e) => {
            changeCallback({
              site_name: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            });
          }}
        />
        <Form.Control.Feedback type="invalid">
          {data.site_name.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>
      <Form.Group controlId="contact_email" className="mb-3">
        <Form.Label>{t('contact_email.label')}</Form.Label>
        <Form.Control
          required
          value={data.contact_email.value}
          isInvalid={data.contact_email.isInvalid}
          onChange={(e) => {
            changeCallback({
              contact_email: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            });
          }}
        />
        <Form.Text>{t('contact_email.text')}</Form.Text>
        <Form.Control.Feedback type="invalid">
          {data.contact_email.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <h5>{t('admin_account')}</h5>
      <Form.Group controlId="admin_name" className="mb-3">
        <Form.Label>{t('admin_name.label')}</Form.Label>
        <Form.Control
          required
          value={data.admin_name.value}
          isInvalid={data.admin_name.isInvalid}
          onChange={(e) => {
            changeCallback({
              admin_name: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            });
          }}
        />
        <Form.Control.Feedback type="invalid">
          {data.admin_name.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <Form.Group controlId="admin_password" className="mb-3">
        <Form.Label>{t('admin_password.label')}</Form.Label>
        <Form.Control
          required
          value={data.admin_password.value}
          isInvalid={data.admin_password.isInvalid}
          onChange={(e) => {
            changeCallback({
              admin_password: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            });
          }}
        />
        <Form.Text>{t('admin_password.text')}</Form.Text>
        <Form.Control.Feedback type="invalid">
          {data.admin_password.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <Form.Group controlId="admin_email" className="mb-3">
        <Form.Label>{t('admin_email.label')}</Form.Label>
        <Form.Control
          required
          value={data.admin_email.value}
          isInvalid={data.admin_email.isInvalid}
          onChange={(e) => {
            changeCallback({
              admin_email: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            });
          }}
        />
        <Form.Text>{t('admin_email.text')}</Form.Text>
        <Form.Control.Feedback type="invalid">
          {data.admin_email.errorMsg}
        </Form.Control.Feedback>
      </Form.Group>

      <div className="d-flex align-items-center justify-content-between">
        <Progress step={4} />
        <Button type="submit">{t('next')}</Button>
      </div>
    </Form>
  );
};

export default Index;
