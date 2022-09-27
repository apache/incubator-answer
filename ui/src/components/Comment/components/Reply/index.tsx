import { useState, memo } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { TextArea, Mentions } from '@answer/components';
import { usePageUsers } from '@answer/hooks';

const Form = ({ userName, onSendReply, onCancel, mode }) => {
  const [value, setValue] = useState('');
  const pageUsers = usePageUsers();
  const { t } = useTranslation('translation', { keyPrefix: 'comment' });

  const handleChange = (e) => {
    setValue(e.target.value);
  };

  return (
    <div className="mb-2">
      <div className="fs-14 mb-2">Reply to {userName}</div>
      <div className="d-flex mb-1 align-items-start">
        <div>
          <Mentions pageUsers={pageUsers.getUsers()}>
            <TextArea size="sm" value={value} onChange={handleChange} />
          </Mentions>
          <div className="text-muted fs-14">{t(`tip_${mode}`)}</div>
        </div>
        <div className="d-flex flex-column">
          <Button
            size="sm"
            className="text-nowrap ms-2"
            onClick={() => onSendReply(value)}>
            {t('btn_add_comment')}
          </Button>
          <Button
            variant="link"
            size="sm"
            className="text-nowrap ms-2 btn-no-border"
            onClick={onCancel}>
            {t('btn_cancel')}
          </Button>
        </div>
      </div>
    </div>
  );
};

export default memo(Form);
