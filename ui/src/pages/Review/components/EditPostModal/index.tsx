import { FC, useState } from 'react';
import { Modal, Button, Form } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { modifyQuestion } from '@/services';
import { useCaptchaModal } from '@/hooks';
import { Editor, TagSelector } from '@/components';
import { handleFormError } from '@/utils';
import type * as Type from '@/common/interface';

import './index.scss';

interface Props {
  visible: boolean;
  handleClose: () => void;
}

interface FormDataItem {
  title: Type.FormValue<string>;
  tags: Type.FormValue<Type.Tag[]>;
  content: Type.FormValue<string>;
}

const Index: FC<Props> = ({ visible = false, handleClose }) => {
  const initFormData = {
    title: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    tags: {
      value: [],
      isInvalid: false,
      errorMsg: '',
    },
    content: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  };
  const { t } = useTranslation('translation', { keyPrefix: 'ask' });
  const [formData, setFormData] = useState<FormDataItem>(initFormData);
  const [focusEditor, setFocusEditor] = useState(false);

  const editCaptcha = useCaptchaModal('edit');

  const handleInput = (data: Partial<FormDataItem>) => {
    setFormData({
      ...formData,
      ...data,
    });
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    event.stopPropagation();

    const params: Type.QuestionParams = {
      title: formData.title.value,
      content: formData.content.value,
      tags: formData.tags.value,
    };

    editCaptcha.check(() => {
      const ep = {
        ...params,
        id: '',
        edit_summary: '',
      };
      const imgCode = editCaptcha.getCaptcha();
      if (imgCode.verify) {
        ep.captcha_code = imgCode.captcha_code;
        ep.captcha_id = imgCode.captcha_id;
      }
      modifyQuestion(ep)
        .then(async (res) => {
          await editCaptcha.close();
          console.log('res', res);
          // navigate(pathFactory.questionLanding(qid, res?.url_title), {
          //   state: { isReview: res?.wait_for_review },
          // });
        })
        .catch((err) => {
          if (err.isError) {
            editCaptcha.handleCaptchaError(err.list);
            const data = handleFormError(err, formData);
            setFormData({ ...data });
          }
        });
    });
  };
  return (
    <Modal
      show={visible}
      onHide={handleClose}
      className="w-100"
      dialogClassName="edit-post-modal">
      <Modal.Header closeButton>
        <Modal.Title>Edit post</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form noValidate onSubmit={handleSubmit}>
          <Form.Group controlId="title" className="mb-3">
            <Form.Label>{t('form.fields.title.label')}</Form.Label>
            <Form.Control
              type="text"
              value={formData.title.value}
              isInvalid={formData.title.isInvalid}
              onChange={(e) => {
                handleInput({
                  title: {
                    value: e.target.value,
                    isInvalid: false,
                    errorMsg: '',
                  },
                });
              }}
              placeholder={t('form.fields.title.placeholder')}
              autoFocus
              contentEditable
            />

            <Form.Control.Feedback type="invalid">
              {formData.title.errorMsg}
            </Form.Control.Feedback>
          </Form.Group>

          <Form.Group controlId="body">
            <Form.Label>{t('form.fields.body.label')}</Form.Label>
            <Form.Control
              defaultValue={formData.content.value}
              isInvalid={formData.content.isInvalid}
              hidden
            />
            <Editor
              value={formData.content.value}
              onChange={(value) => {
                handleInput({
                  content: { value, errorMsg: '', isInvalid: false },
                });
              }}
              className={classNames(
                'form-control p-0',
                focusEditor ? 'focus' : '',
              )}
              onFocus={() => {
                setFocusEditor(true);
              }}
              onBlur={() => {
                setFocusEditor(false);
              }}
            />
            <Form.Control.Feedback type="invalid">
              {formData.content.errorMsg}
            </Form.Control.Feedback>
          </Form.Group>
          <Form.Group controlId="tags" className="my-3">
            <Form.Label>{t('form.fields.tags.label')}</Form.Label>
            <Form.Control
              defaultValue={JSON.stringify(formData.tags.value)}
              isInvalid={formData.tags.isInvalid}
              hidden
            />
            <TagSelector
              value={formData.tags.value}
              onChange={(value) => {
                handleInput({
                  tags: { value, errorMsg: '', isInvalid: false },
                });
              }}
              showRequiredTag
              maxTagLength={5}
            />
            <Form.Control.Feedback type="invalid">
              {formData.tags.errorMsg}
            </Form.Control.Feedback>
          </Form.Group>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleClose}>
          {t('close', { keyPrefix: 'btns' })}
        </Button>
        <Button variant="primary" onClick={handleClose}>
          {t('submit', { keyPrefix: 'btns' })}
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default Index;
