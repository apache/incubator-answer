/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import React from 'react';
import { Modal } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import { loginToContinueStore, siteInfoStore } from '@/stores';
import { floppyNavigation } from '@/utils';
import { WelcomeTitle } from '@/components';

import './login.scss';

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
    floppyNavigation.storageLoginRedirect();
    closeModal();
  };
  return (
    <Modal
      show={visible}
      onHide={closeModal}
      centered
      className="loginToContinueModal"
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
