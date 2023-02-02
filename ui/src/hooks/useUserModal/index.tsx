import { useLayoutEffect, useState, useRef } from 'react';
import { Modal, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

import pattern from '@/common/pattern';
import type * as Type from '@/common/interface';
import { SchemaForm, JSONSchema, UISchema, initFormData } from '@/components';
import { handleFormError } from '@/utils';

const div = document.createElement('div');
const root = ReactDOM.createRoot(div);

interface IProps {
  title?: string;
  onConfirm?: (formData: any) => Promise<any>;
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
    display_name: {
      'ui:options': {
        validator: (value) => {
          const MIN_LENGTH = 4;
          const MAX_LENGTH = 30;

          if (value.length < MIN_LENGTH || value.length > MAX_LENGTH) {
            return t('form.fields.display_name.msg');
          }
          return true;
        },
      },
    },
    email: {
      'ui:options': {
        inputType: 'email',
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
        inputType: 'password',
        validator: (value) => {
          const MIN_LENGTH = 8;
          const MAX_LENGTH = 32;
          if (value.length < MIN_LENGTH || value.length > MAX_LENGTH) {
            return t('form.fields.password.msg');
          }
          return true;
        },
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
        display_name: formData.display_name.value,
        email: formData.email.value,
        password: formData.password.value,
      })
        .then(() => {
          setFormData({
            display_name: {
              value: '',
              isInvalid: false,
              errorMsg: '',
            },
            email: {
              value: '',
              isInvalid: false,
              errorMsg: '',
            },
            password: {
              value: '',
              isInvalid: false,
              errorMsg: '',
            },
          });
          onClose();
        })
        .catch((err) => {
          if (err.isError) {
            const data = handleFormError(err, formData);
            setFormData({ ...data });
          }
        });
    }
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
