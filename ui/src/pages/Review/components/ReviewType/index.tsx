import { FC } from 'react';
import { Card, Form } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import * as Type from '@/common/interface';

interface IProps {
  list: Type.ReviewTypeItem[] | undefined;
  checked: string;
  callback: (type: string) => void;
}

const Index: FC<IProps> = ({ list, checked, callback }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_review' });
  return (
    <Card>
      <Card.Header>{t('filter', { keyPrefix: 'btns' })}</Card.Header>
      <Card.Body>
        <Form.Group>
          <Form.Label>{t('filter_label')}</Form.Label>
          {list?.map((item) => {
            return (
              <Form.Check
                key={item.name}
                type="radio"
                id={item.name}
                disabled={item.todo_amount <= 0}
                label={`${item.label} (${item.todo_amount})`}
                checked={checked === item.name}
                onChange={() => callback(item.name)}
              />
            );
          })}
        </Form.Group>
      </Card.Body>
    </Card>
  );
};

export default Index;
