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

import { useLayoutEffect, useState, useRef } from 'react';
import { Modal, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

import type * as Type from '@/common/interface';
import { SchemaForm, JSONSchema, UISchema, initFormData } from '@/components';
import { handleFormError } from '@/utils';
import { getUserActivation, postUserActivation } from '@/services';
import { useToast } from '@/hooks';

const div = document.createElement('div');
const root = ReactDOM.createRoot(div);

interface IProps {
  title?: string;
  onConfirm?: (formData: any) => Promise<any>;
}
const useChangePasswordModal = (props: IProps = {}) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'inactive',
  });

  const { title = t('btn_name') } = props;
  const [visible, setVisibleState] = useState(false);
  const userId = useRef('');
  const isLoading = useRef(false);
  const Toast = useToast();

  const schema: JSONSchema = {
    title: t('btn_name'),
    properties: {
      activationUrl: {
        type: 'string',
        title: t('resend_email.url_label'),
        description: t('resend_email.url_text'),
      },
    },
  };
  const uiSchema: UISchema = {
    activationUrl: {
      'ui:options': {
        readOnly: true,
      },
    },
  };
  const [formData, setFormData] = useState<Type.FormDataType>(
    initFormData(schema),
  );

  const formRef = useRef<{
    validator: () => Promise<boolean>;
  }>(null);

  const getActivationUrl = () => {
    return getUserActivation(userId.current).then((resp) => {
      if (resp?.activation_url) {
        setFormData({
          ...formData,
          activationUrl: {
            value: resp.activation_url,
            isInvalid: false,
            errorMsg: '',
          },
        });
      }
    });
  };

  const onClose = () => {
    setVisibleState(false);
    userId.current = '';
    setFormData(initFormData(schema));
  };

  const onShow = async (user_id: string) => {
    if (!user_id) {
      return;
    }
    userId.current = user_id;
    await getActivationUrl();
    setVisibleState(true);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    event.stopPropagation();
    isLoading.current = true;
    postUserActivation(userId.current)
      .then(() => {
        Toast.onShow({
          msg: t('sent_success', { keyPrefix: 'toast' }),
          variant: 'success',
        });
        onClose();
      })
      .catch((err) => {
        if (err.isError) {
          const data = handleFormError(err, formData);
          setFormData({ ...data });
        }
      })
      .finally(() => {
        isLoading.current = false;
      });
  };

  const handleOnChange = (data) => {
    setFormData(data);
  };

  useLayoutEffect(() => {
    root.render(
      <Modal show={visible} title={title} onHide={onClose}>
        <Modal.Header closeButton>
          <Modal.Title as="h5">{title}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <SchemaForm
            ref={formRef}
            schema={schema}
            uiSchema={uiSchema}
            formData={formData}
            onChange={handleOnChange}
            hiddenSubmit
          />
        </Modal.Body>
        <Modal.Footer>
          <Button variant="link" onClick={() => onClose()}>
            {t('cancel', { keyPrefix: 'btns' })}
          </Button>
          <Button
            disabled={isLoading.current}
            variant="primary"
            onClick={handleSubmit}>
            {t('resend', { keyPrefix: 'btns' })}
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

export default useChangePasswordModal;
