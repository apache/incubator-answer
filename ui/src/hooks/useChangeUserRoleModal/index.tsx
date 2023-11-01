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

import { useLayoutEffect, useState } from 'react';
import { Modal, Form, Button, FormCheck } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

import { getUserRoles, changeUserRole } from '@/services';
import { UserRoleItem } from '@/common/interface';

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
  const [defaultId, setDefaultId] = useState(-1);
  const [isInvalid, setInvalidState] = useState(false);
  const [changedId, setChangeId] = useState(-1);
  const [show, setShow] = useState(false);
  const [list, setList] = useState<UserRoleItem[]>([]);

  const getRolesData = async () => {
    const res = await getUserRoles();
    setList(res);
  };

  const handleRadio = (val) => {
    setInvalidState(false);
    setChangeId(val.id);
  };

  const onClose = () => {
    setChangeId(-1);
    setDefaultId(-1);
    setShow(false);
  };

  const handleSubmit = () => {
    if (defaultId === changedId) {
      onClose();

      return;
    }

    changeUserRole({
      user_id: id,
      role_id: changedId,
    }).then(() => {
      callback?.();
      onClose();
    });
  };

  const onShow = (params) => {
    getRolesData();
    setId(params.id);
    setChangeId(params.role_id);
    setDefaultId(params.role_id);
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
                <div key={item?.id}>
                  <Form.Group controlId={item.name} className="mb-3">
                    <FormCheck>
                      <FormCheck.Input
                        id={item.name}
                        type="radio"
                        checked={changedId === item.id}
                        onChange={() => handleRadio(item)}
                        isInvalid={isInvalid}
                      />
                      <FormCheck.Label htmlFor={item.name}>
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
