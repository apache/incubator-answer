import React from 'react';
import { Modal } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import { loginToContinueStore, siteInfoStore } from '@/stores';
import { WelcomeTitle } from '@/components';

interface IProps {
  visible: boolean;
}

const Index: React.FC<IProps> = ({ visible = false }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'login' });
  const { update: updateStore } = loginToContinueStore();
  const { siteInfo } = siteInfoStore((_) => _);
  const closeModal = () => {
    updateStore({ show: false });
  };
  const linkClick = (evt) => {
    evt.stopPropagation();
    closeModal();
  };
  return (
    <Modal
      title="LoginToContinue"
      show={visible}
      onHide={closeModal}
      centered
      fullscreen="sm-down">
      <Modal.Header closeButton>
        <Modal.Title as="h5">{t('login_to_continue')}</Modal.Title>
      </Modal.Header>
      <Modal.Body className="p-5">
        <div className="d-flex flex-column align-items-center text-center text-body">
          <WelcomeTitle className="mb-2" />
          <p>{siteInfo.description}</p>
        </div>
        <div className="d-grid gap-2">
          <Link
            to="/users/login"
            className="btn btn-primary"
            onClick={linkClick}>
            {t('login', { keyPrefix: 'btns' })}
          </Link>
          <Link
            to="/users/register"
            className="btn btn-link"
            onClick={linkClick}>
            {t('signup', { keyPrefix: 'btns' })}
          </Link>
        </div>
      </Modal.Body>
    </Modal>
  );
};
export default Index;
