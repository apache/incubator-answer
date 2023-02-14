import { useLayoutEffect, useState } from 'react';
import { Modal, Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

const div = document.createElement('div');
const root = ReactDOM.createRoot(div);

const MAX_LENGTH = 35;
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
    } else if (displayName.value.length > MAX_LENGTH) {
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
    } else if (slugName.value.length > MAX_LENGTH) {
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
