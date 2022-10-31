import { FC } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import Progress from '../Progress';

interface Props {
  visible: boolean;
}
const Index: FC<Props> = ({ visible }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'install' });

  if (!visible) return null;
  return (
    <Form>
      <h5>{t('site_information')}</h5>
      <Form.Group controlId="site_name" className="mb-3">
        <Form.Label>{t('site_name.label')}</Form.Label>
        <Form.Control required type="text" />
      </Form.Group>
      <Form.Group controlId="contact_email" className="mb-3">
        <Form.Label>{t('contact_email.label')}</Form.Label>
        <Form.Control required type="text" />
        <Form.Text>{t('contact_email.text')}</Form.Text>
      </Form.Group>

      <h5>{t('admin_account')}</h5>
      <Form.Group controlId="admin_name" className="mb-3">
        <Form.Label>{t('admin_name.label')}</Form.Label>
        <Form.Control required type="text" />
      </Form.Group>

      <Form.Group controlId="admin_password" className="mb-3">
        <Form.Label>{t('admin_password.label')}</Form.Label>
        <Form.Control required type="text" />
        <Form.Text>{t('admin_password.text')}</Form.Text>
      </Form.Group>

      <Form.Group controlId="admin_email" className="mb-3">
        <Form.Label>{t('admin_email.label')}</Form.Label>
        <Form.Control required type="text" />
        <Form.Text>{t('admin_email.text')}</Form.Text>
      </Form.Group>

      <div className="d-flex align-items-center justify-content-between">
        <Progress step={4} />
        <Button>{t('next')}</Button>
      </div>
    </Form>
  );
};

export default Index;
