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
import { Modal, Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

import { TAG_SLUG_NAME_MAX_LENGTH } from '@/common/constants';

const div = document.createElement('div');
const root = ReactDOM.createRoot(div);

interface IProps {
  title?: string;
  onConfirm?: (formData: any) => void;
}
const useTagModal = (props: IProps = {}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'tag_modal' });

  const { title = t('title'), onConfirm } = props;
  const [visible, setVisibleState] = useState(false);
  const [formData, setFormData] = useState({
    displayName: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    slugName: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    description: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });

  const onClose = () => {
    setVisibleState(false);
  };

  const onShow = (searchStr = '') => {
    setVisibleState(true);
    setFormData({
      ...formData,
      displayName: {
        value: searchStr,
        isInvalid: false,
        errorMsg: '',
      },
      slugName: {
        value: searchStr,
        isInvalid: false,
        errorMsg: '',
      },
    });
  };

  const checkValidated = (): boolean => {
    let bol = true;
    const { displayName, slugName } = formData;
    if (!displayName.value) {
      bol = false;
      formData.displayName = {
        value: '',
        isInvalid: true,
        errorMsg: t('form.fields.display_name.msg.empty'),
      };
    } else if (displayName.value.length > TAG_SLUG_NAME_MAX_LENGTH) {
      bol = false;
      formData.displayName = {
        value: displayName.value,
        isInvalid: true,
        errorMsg: t('form.fields.display_name.msg.range'),
      };
    } else {
      formData.displayName = {
        value: displayName.value,
        isInvalid: false,
        errorMsg: '',
      };
    }

    if (!slugName.value) {
      bol = false;
      formData.slugName = {
        value: '',
        isInvalid: true,
        errorMsg: t('form.fields.slug_name.msg.empty'),
      };
    } else if (slugName.value.length > TAG_SLUG_NAME_MAX_LENGTH) {
      bol = false;
      formData.slugName = {
        value: slugName.value,
        isInvalid: true,
        errorMsg: t('form.fields.slug_name.msg.range'),
      };
      // } else if (/[^a-z0-9+#\-.]/.test(slugName.value)) {
      //   bol = false;
      //   formData.slugName = {
      //     value: slugName.value,
      //     isInvalid: true,
      //     errorMsg: t('form.fields.slug_name.msg.character'),
      //   };
    } else {
      formData.slugName = {
        value: slugName.value,
        isInvalid: false,
        errorMsg: '',
      };
    }

    setFormData({
      ...formData,
    });
    return bol;
  };

  const handleSubmit = (event: React.MouseEvent<HTMLElement>) => {
    event.preventDefault();
    event.stopPropagation();

    if (!checkValidated()) {
      return;
    }

    if (onConfirm instanceof Function) {
      onConfirm({
        slug_name: formData.slugName.value,
        display_name: formData.displayName.value,
        original_text: formData.description.value,
      });
      setFormData({
        displayName: {
          value: '',
          isInvalid: false,
          errorMsg: '',
        },
        slugName: {
          value: '',
          isInvalid: false,
          errorMsg: '',
        },
        description: {
          value: '',
          isInvalid: false,
          errorMsg: '',
        },
      });
    }
    onClose();
  };

  const handleDisplayNameChange = (
    event: React.ChangeEvent<HTMLInputElement>,
  ) => {
    const { value } = event.target;
    setFormData({
      ...formData,
      displayName: {
        value,
        isInvalid: false,
        errorMsg: '',
      },
    });
  };

  const handleSlugNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { value } = event.target;
    setFormData({
      ...formData,
      slugName: {
        value,
        isInvalid: false,
        errorMsg: '',
      },
    });
  };

  const handleDescriptionChange = (
    event: React.ChangeEvent<HTMLInputElement>,
  ) => {
    const { value } = event.target;
    setFormData({
      ...formData,
      description: {
        value,
        isInvalid: false,
        errorMsg: '',
      },
    });
  };
  useLayoutEffect(() => {
    root.render(
      <Modal show={visible} title={title} onHide={onClose}>
        <Modal.Header closeButton>
          <Modal.Title as="h5">{title}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group controlId="displayName" className="mb-3">
              <Form.Label>{t('form.fields.display_name.label')}</Form.Label>
              <Form.Control
                type="text"
                value={formData.displayName.value}
                onChange={handleDisplayNameChange}
                isInvalid={formData.displayName.isInvalid}
              />
              <Form.Control.Feedback type="invalid">
                {formData.displayName.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>
            <Form.Group controlId="slugName" className="mb-3">
              <Form.Label>{t('form.fields.slug_name.label')}</Form.Label>
              <Form.Control
                type="text"
                value={formData.slugName.value}
                onChange={handleSlugNameChange}
                isInvalid={formData.slugName.isInvalid}
              />

              <Form.Text as="div">
                {t('form.fields.slug_name.msg.range')}
              </Form.Text>
              <Form.Control.Feedback type="invalid">
                {formData.slugName.errorMsg}
              </Form.Control.Feedback>
            </Form.Group>
            <Form.Group controlId="description">
              <Form.Label>{`${t('form.fields.desc.label')} ${t('optional', {
                keyPrefix: 'form',
              })}`}</Form.Label>
              <Form.Control
                className="font-monospace"
                value={formData.description.value}
                onChange={handleDescriptionChange}
                as="textarea"
                rows={2}
              />
            </Form.Group>
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

export default useTagModal;
