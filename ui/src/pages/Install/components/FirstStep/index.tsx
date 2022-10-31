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
      <Form.Group controlId="langSelect" className="mb-3">
        <Form.Label>{t('choose_lang.label')}</Form.Label>
        <Form.Select>
          <option>English</option>
        </Form.Select>
      </Form.Group>

      <div className="d-flex align-items-center justify-content-between">
        <Progress step={1} />
        <Button>{t('next')}</Button>
      </div>
    </Form>
  );
};

export default Index;
