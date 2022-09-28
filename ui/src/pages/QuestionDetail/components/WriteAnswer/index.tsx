import { memo, useState, FC } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { marked } from 'marked';

import { Editor, Modal } from '@answer/components';
import { FormDataType } from '@answer/common/interface';
import { postAnswer } from '@answer/api';

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
        content: t('confirm_info'),
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

  return (
    <Form noValidate className="mt-4">
      {showEditor && (
        <Form.Group className="mb-3">
          <Form.Label>
            <h5>{t('title')}</h5>
          </Form.Label>
          <Form.Control
            isInvalid={formData.content.isInvalid}
            className="d-none"
          />
          <Editor
            value={formData.content.value}
            onChange={(val) => {
              setFormData({
                content: {
                  value: val,
                  isInvalid: false,
                  errorMsg: '',
                },
              });
            }}
          />

          <Form.Control.Feedback type="invalid">
            {formData.content.errorMsg}
          </Form.Control.Feedback>
        </Form.Group>
      )}

      <Button onClick={clickBtn}>{t('btn_name')}</Button>
    </Form>
  );
};

export default memo(Index);
