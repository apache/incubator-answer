import { useLayoutEffect, useState } from 'react';
import { Modal, Form, Button, FormCheck } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

const div = document.createElement('div');
const root = ReactDOM.createRoot(div);

const useEditStatusModal = ({
  editType = '',
  callback,
}: {
  editType: string;
  callback: (id, type) => void;
}) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.status_modal',
  });
  const [id, setId] = useState('');
  const [defaultType, setDefaultType] = useState('');
  const [isInvalid, setInvalidState] = useState(false);
  const [changeType, setChangeType] = useState('normal');

  const [show, setShow] = useState(false);
  const [list] = useState<any[]>([
    {
      type: 'normal',
      name: t('normal_name'),
      description: t('normal_desc'),
    },
    {
      type: 'closed',
      name: t('closed_name'),
      description: t('closed_desc'),
    },
    {
      type: 'deleted',
      name: t('deleted_name'),
      description: t('deleted_desc'),
    },
  ]);

  const handleRadio = (val) => {
    setInvalidState(false);
    setChangeType(val.type);
  };

  const onClose = () => {
    setChangeType('');
    setShow(false);
  };

  const handleSubmit = () => {
    if (changeType === '') {
      setInvalidState(true);
      return;
    }

    if (defaultType === changeType) {
      onClose();

      return;
    }

    onClose();
    callback?.(id, changeType);
  };

  const onShow = (params) => {
    setId(params.id);
    setChangeType(params.type);
    setDefaultType(params.type);
    setShow(true);
  };
  useLayoutEffect(() => {
    root.render(
      <Modal show={show} onHide={onClose}>
        <Modal.Header closeButton>
          <Modal.Title as="h5">{t('title', { type: editType })}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            {list.map((item) => {
              if (editType === 'answer' && item.type === 'closed') {
                return null;
              }
              return (
                <div key={item?.type}>
                  <Form.Group controlId={item.type} className="mb-3">
                    <FormCheck>
                      <FormCheck.Input
                        id={item.type}
                        type="radio"
                        checked={changeType === item.type}
                        onChange={() => handleRadio(item)}
                        isInvalid={isInvalid}
                      />
                      <FormCheck.Label htmlFor={item.type}>
                        <span className="fw-bold">{item.name}</span>
                        <br />
                        <span className="small text-secondary">
                          {item.description}
                        </span>
                      </FormCheck.Label>
                      <Form.Control.Feedback type="invalid">
                        {t('msg.empty')}
                      </Form.Control.Feedback>
                    </FormCheck>
                  </Form.Group>
                </div>
              );
            })}
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="link" onClick={() => onClose()}>
            {t('btn_cancel')}
          </Button>
          <Button variant="primary" onClick={handleSubmit}>
            {changeType !== 'normal' ? t('btn_next') : t('btn_submit')}
          </Button>
        </Modal.Footer>
      </Modal>,
    );
  });

  return {
    onClose,
    onShow,
  };
};

export default useEditStatusModal;
