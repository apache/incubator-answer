import { useState } from 'react';
import { Card, Form } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

const Index = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_review' });
  const [checked, setValue] = useState(false);
  return (
    <Card>
      <Card.Header>{t('filter', { keyPrefix: 'btns' })}</Card.Header>
      <Card.Body>
        <Form.Group>
          <Form.Label>{t('filter_label')}</Form.Label>
          <Form.Check
            type="radio"
            label="Queued post (99+)"
            checked={checked}
            onChange={(e) => setValue(e.target.checked)}
          />

          <Form.Check
            type="radio"
            label="Queued post (199+)"
            checked={checked}
            onChange={(e) => setValue(e.target.checked)}
          />
        </Form.Group>
      </Card.Body>
    </Card>
  );
};

export default Index;
