import { useLayoutEffect, useState, useRef } from 'react';
import { Modal, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

import pattern from '@/common/pattern';
import type * as Type from '@/common/interface';
import { SchemaForm, JSONSchema, UISchema, initFormData } from '@/components';

const div = document.createElement('div');
const root = ReactDOM.createRoot(div);

interface IProps {
  title?: string;
  onConfirm?: (formData: any) => void;
}
const useAddUserModal = (props: IProps = {}) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.users.user_modal',
  });

  const { title = t('title'), onConfirm } = props;
  const [visible, setVisibleState] = useState(false);
  const schema: JSONSchema = {
    title: t('title'),
    required: ['display_name', 'email', 'password'],
    properties: {
      display_name: {
        type: 'string',
        title: t('form.fields.display_name.label'),
      },
      email: {
        type: 'string',
        title: t('form.fields.email.label'),
      },
      password: {
        type: 'string',
        title: t('form.fields.password.label'),
      },
    },
  };
  const uiSchema: UISchema = {
    email: {
      'ui:options': {
        type: 'email',
        validator: (value) => {
          if (value && !pattern.email.test(value)) {
            return t('form.fields.email.msg');
          }
          return true;
        },
      },
    },
    password: {
      'ui:options': {
        type: 'password',
      },
    },
  };
  const [formData, setFormData] = useState<Type.FormDataType>(
    initFormData(schema),
  );

  const formRef = useRef<{
    validator: () => Promise<boolean>;
  }>(null);

  const onClose = () => {
    setVisibleState(false);
  };

  const onShow = () => {
    setVisibleState(true);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    event.stopPropagation();
    const isValid = await formRef.current?.validator();

    if (!isValid) {
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

export default useAddUserModal;
