import { memo, useState, FC } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { marked } from 'marked';
import classNames from 'classnames';

import { Editor, Modal, TextArea } from '@/components';
import { FormDataType } from '@/common/interface';
import { postAnswer } from '@/services';

interface Props {
  visible?: boolean;
  data: {
    /** question  id */
    qid: string;
    answered?: boolean;
  };
  callback?: (obj) => void;
}

const Index: FC<Props> = ({ visible = false, data, callback }) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'question_detail.write_answer',
  });
  const [formData, setFormData] = useState<FormDataType>({
    content: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });
  const [showEditor, setShowEditor] = useState<boolean>(visible);
  const [focusType, setForceType] = useState('');

  const handleSubmit = () => {
    if (!formData.content.value) {
      setFormData({
        content: {
          value: '',
          isInvalid: true,
          errorMsg: t('empty'),
        },
      });
      return;
    }
    postAnswer({
      question_id: data?.qid,
      content: formData.content.value,
      html: marked.parse(formData.content.value),
    }).then((res) => {
      setShowEditor(false);
      setFormData({
        content: {
          value: '',
          isInvalid: false,
          errorMsg: '',
        },
      });
      callback?.(res.info);
    });
  };

  const clickBtn = () => {
    if (data?.answered && !showEditor) {
      Modal.confirm({
        title: t('confirm_title'),
        content: t('confirm_info'),
        confirmText: t('continue'),
        onConfirm: () => {
          setShowEditor(true);
        },
      });
      return;
    }

    if (!showEditor) {
      setShowEditor(true);
      return;
    }

    handleSubmit();
  };
  const handleFocusForTextArea = () => {
    setShowEditor(true);
  };

  return (
    <Form noValidate className="mt-4">
      {(!data.answered || showEditor) && (
        <Form.Group className="mb-3">
          <Form.Label>
            <h5>{t('title')}</h5>
          </Form.Label>
          <Form.Control
            isInvalid={formData.content.isInvalid}
            className="d-none"
          />
          {!showEditor && !data.answered && (
            <div className="d-flex">
              <TextArea
                className="w-100"
                rows={8}
                autoFocus={false}
                onFocus={handleFocusForTextArea}
              />
            </div>
          )}
          {showEditor && (
            <Editor
              className={classNames(
                'form-control p-0',
                focusType === 'answer' && 'focus',
              )}
              value={formData.content.value}
              autoFocus
              onChange={(val) => {
                setFormData({
                  content: {
                    value: val,
                    isInvalid: false,
                    errorMsg: '',
                  },
                });
              }}
              onFocus={() => {
                setForceType('answer');
              }}
              onBlur={() => {
                setForceType('');
              }}
            />
          )}

          <Form.Control.Feedback type="invalid">
            {formData.content.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
      )}

      {data.answered && !showEditor ? (
        <Button onClick={clickBtn}>{t('add_another_answer')}</Button>
      ) : (
        <Button onClick={clickBtn}>{t('btn_name')}</Button>
      )}
    </Form>
  );
};

export default memo(Index);
