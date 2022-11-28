import { useLayoutEffect, useState } from 'react';
import { Modal, Form, Button, FormCheck } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

import { Modal as AnswerModal } from '@/components';
import { changeUserStatus } from '@/services';

const div = document.createElement('div');
const root = ReactDOM.createRoot(div);

interface Props {
  callback?: () => void;
}

const useChangeUserRoleModal = ({ callback }: Props) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.user_role_modal',
  });
  const [id, setId] = useState('');
  const [defaultType, setDefaultType] = useState('');
  const [isInvalid, setInvalidState] = useState(false);
  const [changeType, setChangeType] = useState({
    type: '',
    haveContent: false,
  });
  const [content, setContent] = useState({
    value: '',
    isInvalid: false,
    errorMsg: '',
  });
  const [show, setShow] = useState(false);
  const [list] = useState<any[]>([
    {
      type: 'user',
      name: t('user.name'),
      description: t('user.description'),
    },
    {
      type: 'admin',
      name: t('admin.name'),
      description: t('admin.description'),
    },
    {
      type: 'moderator',
      name: t('moderator.name'),
      description: t('moderator.description'),
    },
  ]);

  const handleRadio = (val) => {
    setInvalidState(false);
    setContent({
      value: '',
      isInvalid: false,
      errorMsg: '',
    });
    setChangeType({
      type: val.type,
      haveContent: val.have_content,
    });
  };

  const onClose = () => {
    setChangeType({
      type: '',
      haveContent: false,
    });
    setContent({
      value: '',
      isInvalid: false,
      errorMsg: '',
    });
    setContent({
      value: '',
      isInvalid: false,
      errorMsg: '',
    });
    setShow(false);
  };

  const handleSubmit = () => {
    if (changeType.type === '') {
      setInvalidState(true);
      return;
    }
    if (changeType.haveContent && !content.value) {
      setContent({
        value: content.value,
        isInvalid: true,
        errorMsg: t('remark.empty'),
      });
      return;
    }
    if (defaultType === changeType.type) {
      onClose();

      return;
    }
    if (changeType.type === 'deleted') {
      onClose();

      AnswerModal.confirm({
        title: t('confirm_title'),
        content: t('confirm_content'),
        confirmText: t('confirm_btn'),
        confirmBtnVariant: 'danger',
        onConfirm: () => {
          changeUserStatus({
            user_id: id,
            status: changeType.type,
          }).then(() => {
            callback?.();
            onClose();
          });
        },
      });
      return;
    }
    changeUserStatus({
      user_id: id,
      status: changeType.type,
    }).then(() => {
      callback?.();
      onClose();
    });
  };

  const onShow = (params) => {
    setId(params.id);
    setChangeType({
      ...changeType,
      type: params.type,
    });
    setDefaultType(params.type);
    setShow(true);
  };
  useLayoutEffect(() => {
    root.render(
      <Modal show={show} onHide={onClose}>
        <Modal.Header closeButton>
          <Modal.Title as="h5">{t('title')}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            {list.map((item) => {
              return (
                <div key={item?.type}>
                  <Form.Group controlId={item.type} className="mb-3">
                    <FormCheck>
                      <FormCheck.Input
                        id={item.type}
                        type="radio"
                        checked={changeType.type === item.type}
                        onChange={() => handleRadio(item)}
                        isInvalid={isInvalid}
                      />
                      <FormCheck.Label htmlFor={item.type}>
                        <span className="fw-bold">{item.name}</span>
                        <br />
                        <span className="text-secondary">
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
            {t('btn_submit')}
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

export default useChangeUserRoleModal;
