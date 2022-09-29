import { useState, useEffect, memo } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { TextArea, Mentions } from '@answer/components';
import { usePageUsers } from '@answer/hooks';

const Form = ({
  className = '',
  value: initialValue = '',
  onSendReply,
  type = '',
  onCancel,
  mode,
}) => {
  const [value, setValue] = useState('');
  const pageUsers = usePageUsers();
  const { t } = useTranslation('translation', { keyPrefix: 'comment' });

  useEffect(() => {
    if (!initialValue) {
      return;
    }
    setValue(initialValue);
  }, [initialValue]);

  const handleChange = (e) => {
    setValue(e.target.value);
  };

  return (
    <div className={classNames('d-flex align-items-start', className)}>
      <div>
        <Mentions pageUsers={pageUsers.getUsers()}>
          <TextArea size="sm" value={value} onChange={handleChange} />
        </Mentions>
        <div className="form-text">{t(`tip_${mode}`)}</div>
      </div>
      {type === 'edit' ? (
        <div className="d-flex flex-column">
          <Button
            size="sm"
            className="text-nowrap ms-2"
            onClick={() => onSendReply(value)}>
            {t('btn_save_edits')}
          </Button>
          <Button
            variant="link"
            size="sm"
            className="text-nowrap ms-2 btn-no-border"
            onClick={onCancel}>
            {t('btn_cancel')}
          </Button>
        </div>
      ) : (
        <Button
          size="sm"
          className="text-nowrap ms-2"
          onClick={() => onSendReply(value)}>
          {t('btn_add_comment')}
        </Button>
      )}
    </div>
  );
};

export default memo(Form);
