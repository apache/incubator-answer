import { useState } from 'react';
import { Modal, Button, Form } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

const DeleteUserModal = ({ show, onClose, onDelete }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.users' });
  const [checkVal, setCheckVal] = useState(false);

  const handleClose = () => {
    onClose();
    setCheckVal(false);
  };

  return (
    <Modal show={show} onHide={handleClose}>
      <Modal.Header closeButton>
        <Modal.Title>{t('delete_user.title')}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <p>{t('delete_user.content')}</p>
        <div className="text-danger mb-2">
          {t('delete_user.remove')} {t('optional', { keyPrefix: 'form' })}
        </div>
        <Form>
          <Form.Group controlId="delete_user">
            <Form.Check type="checkbox" id="delete_user">
              <Form.Check.Input
                type="checkbox"
                checked={checkVal}
                onChange={(e) => {
                  setCheckVal(e.target.checked);
                }}
              />
              <Form.Check.Label htmlFor="delete_user">
                <span>{t('delete_user.label')}</span>
                <br />
                <span className="small text-secondary">
                  {t('delete_user.text')}
                </span>
              </Form.Check.Label>
            </Form.Check>
          </Form.Group>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="link" onClick={handleClose}>
          {t('cancel', { keyPrefix: 'btns' })}
        </Button>
        <Button variant="danger" onClick={() => onDelete(checkVal)}>
          {t('delete', { keyPrefix: 'btns' })}
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default DeleteUserModal;
