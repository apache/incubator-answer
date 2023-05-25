import { useState, memo } from 'react';
import { Button, Form } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { TextArea, Mentions } from '@/components';
import { usePageUsers, usePromptWithUnload } from '@/hooks';

const Index = ({ userName, onSendReply, onCancel, mode }) => {
  const [value, setValue] = useState('');
  const pageUsers = usePageUsers();
  const { t } = useTranslation('translation', { keyPrefix: 'comment' });
  const [validationErrorMsg, setValidationErrorMsg] = useState('');

  usePromptWithUnload({
    when: Boolean(value),
  });

  const handleChange = (e) => {
    setValue(e.target.value);
  };
  const handleSelected = (val) => {
    setValue(val);
  };
  const handleSendReply = () => {
    onSendReply(value).catch((ex) => {
      if (ex.isError) {
        setValidationErrorMsg(ex.msg);
      }
    });
  };

  return (
    <div className="mb-2">
      <div className="small mb-2">
        {t('reply_to')} {userName}
      </div>
      <div className="d-flex mb-1 align-items-start flex-column flex-md-row">
        <div className="w-100">
          <div
            className={classNames('custom-form-control', {
              'is-invalid': validationErrorMsg,
            })}>
            <Mentions
              pageUsers={pageUsers.getUsers()}
              onSelected={handleSelected}>
              <TextArea size="sm" value={value} onChange={handleChange} />
            </Mentions>
            <div className="form-text">{t(`tip_${mode}`)}</div>
          </div>
          <Form.Control.Feedback type="invalid">
            {validationErrorMsg}
          </Form.Control.Feedback>
        </div>
        <div className="d-flex flex-row flex-md-column ms-0 ms-md-2 mt-2 mt-md-0">
          <Button
            size="sm"
            className="text-nowrap"
            onClick={() => handleSendReply()}>
            {t('btn_add_comment')}
          </Button>
          <Button
            variant="link"
            size="sm"
            className="text-nowrap btn-no-border ms-2 ms-md-0"
            onClick={onCancel}>
            {t('btn_cancel')}
          </Button>
        </div>
      </div>
    </div>
  );
};

export default memo(Index);
