import { memo } from 'react';
import { Button } from 'react-bootstrap';

import { Icon, Modal } from '@/components';

const Index = () => {
  const deleteLogins = (type) => {
    Modal.confirm({
      title: 'Remove Login',
      content: 'Are you sure you want to delete this logins?',
      confirmBtnVariant: 'danger',
      confirmText: 'Remove',
      onConfirm: () => {
        console.log('delete login by: ', type);
      },
    });
  };
  return (
    <div className="mt-5">
      <div className="form-label">My Logins</div>
      <small className="form-text mt-0">
        Log in or sign up on this site using these accounts.
      </small>

      <div className="mt-3">
        <Button variant="outline-secondary" className="d-block mb-2">
          <Icon name="google" className="me-2" />
          <span>Connect with Google</span>
        </Button>

        <Button
          variant="outline-danger"
          className="mb-2"
          onClick={() => deleteLogins('github')}>
          <Icon name="github" className="me-2" />
          <span>Remove GitHub</span>
        </Button>
      </div>
    </div>
  );
};

export default memo(Index);
