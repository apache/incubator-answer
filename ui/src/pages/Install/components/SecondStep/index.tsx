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
      <Form.Group controlId="database_engine" className="mb-3">
        <Form.Label>{t('database_engine.label')}</Form.Label>
        <Form.Select>
          <option>English</option>
        </Form.Select>
      </Form.Group>

      <Form.Group controlId="username" className="mb-3">
        <Form.Label>{t('username.label')}</Form.Label>
        <Form.Control placeholder={t('username.placeholder')} />
      </Form.Group>

      <Form.Group controlId="password" className="mb-3">
        <Form.Label>{t('password.label')}</Form.Label>
        <Form.Control placeholder={t('password.placeholder')} />
      </Form.Group>

      <Form.Group controlId="database_host" className="mb-3">
        <Form.Label>{t('database_host.label')}</Form.Label>
        <Form.Control placeholder={t('database_host.placeholder')} />
      </Form.Group>

      <Form.Group controlId="database_name" className="mb-3">
        <Form.Label>{t('database_name.label')}</Form.Label>
        <Form.Control placeholder={t('database_name.placeholder')} />
      </Form.Group>

      <Form.Group controlId="table_prefix" className="mb-3">
        <Form.Label>{t('table_prefix.label')}</Form.Label>
        <Form.Control placeholder={t('table_prefix.placeholder')} />
      </Form.Group>

      <div className="d-flex align-items-center justify-content-between">
        <Progress step={2} />
        <Button>{t('next')}</Button>
      </div>
    </Form>
  );
};

export default Index;
