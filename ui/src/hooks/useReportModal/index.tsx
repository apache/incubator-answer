import { useState, useRef, useEffect, useLayoutEffect } from 'react';
import { Modal, Form, Button, FormCheck } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ReactDOM from 'react-dom/client';

import { useToast } from '@/hooks';
import type * as Type from '@/common/interface';
import { reportList, postReport, closeQuestion, putReport } from '@/services';

interface Params {
  isBackend?: boolean;
  type: Type.ReportType;
  id: string;
  title?: string;
  action: Type.ReportAction;
}

const useReportModal = (callback?: () => void) => {
  const { t } = useTranslation('translation', { keyPrefix: 'report_modal' });
  const toast = useToast();
  const [params, setParams] = useState<Params | null>(null);
  const [isInvalid, setInvalidState] = useState(false);
  const [reportType, setReportType] = useState({
    type: -1,
    haveContent: false,
  });
  const rootRef = useRef<{ root: ReactDOM.Root | null }>({
    root: null,
  });

  const [content, setContent] = useState({
    value: '',
    isInvalid: false,
    errorMsg: '',
  });
  const [show, setShow] = useState(false);
  const [list, setList] = useState<any[]>([]);

  useEffect(() => {
    const div = document.createElement('div');
    rootRef.current.root = ReactDOM.createRoot(div);
  }, []);
  const getList = ({ type, action, isBackend }: Params) => {
    reportList({ type, action, isBackend }).then((res) => {
      setList(res);
      setShow(true);
    });
  };

  const handleRadio = (val) => {
    setInvalidState(false);
    setContent({
      value: '',
      isInvalid: false,
      errorMsg: '',
    });
    setReportType({
      type: val.reason_type,
      haveContent: Boolean(val.content_type),
    });
  };

  const onClose = () => {
    setContent({
      value: '',
      isInvalid: false,
      errorMsg: '',
    });
    setShow(false);
  };

  const handleSubmit = () => {
    if (!params) {
      return;
    }
    if (reportType.type === -1) {
      setInvalidState(true);
      return;
    }

    if (reportType.haveContent && !content.value) {
      setContent({
        value: content.value,
        isInvalid: true,
        errorMsg: t('remark.empty'),
      });
      return;
    }
    if (params.type === 'question' && params.action === 'close') {
      closeQuestion({
        id: params.id,
        close_type: reportType.type,
        close_msg: content.value,
      }).then(() => {
        callback?.();
        onClose();
      });
      return;
    }
    if (!params.isBackend && params.action === 'flag') {
      postReport({
        source: params.type,
        report_type: reportType.type,
        object_id: params.id,
        content: content.value,
      }).then(() => {
        toast.onShow({
          msg: t('flag_success', { keyPrefix: 'toast' }),
          variant: 'warning',
        });
        callback?.();
        onClose();
      });
    }

    if (params.isBackend && params.action === 'review') {
      putReport({
        action: params.type,
        flagged_content: content.value,
        flagged_type: reportType.type,
        id: params.id,
      }).then(() => {
        callback?.();
        onClose();
      });
    }
  };

  const onShow = (obj: Params) => {
    setParams(obj);
    getList(obj);
  };
  let title = '';
  if (typeof params === 'object' && params) {
    title = params.title || t(`${params.action}_title`);
    if (params.action === 'review') {
      title = t(`${params.action}_${params.type}_title`);
    }
  }
  useLayoutEffect(() => {
    rootRef.current.root?.render(
      <Modal show={show} onHide={onClose}>
        <Modal.Header closeButton>
          <Modal.Title as="h5">{title}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            {list.map((item) => {
              return (
                <div key={item?.reason_type}>
                  <Form.Group
                    controlId={`report_${item?.reason_type}`}
                    className={`${
                      item.have_content && reportType === item.type
                        ? 'mb-2'
                        : 'mb-3'
                    }`}>
                    <FormCheck>
                      <FormCheck.Input
                        id={item.reason_type}
                        type="radio"
                        checked={reportType.type === item.reason_type}
                        onChange={() => handleRadio(item)}
                        isInvalid={isInvalid}
                      />
                      <FormCheck.Label htmlFor={item.reason_type}>
                        <span className="fw-bold">{item?.name}</span>
                        <br />
                        <span className="text-secondary">
                          {item?.description}
                        </span>
                      </FormCheck.Label>
                      <Form.Control.Feedback type="invalid">
                        {t('msg.empty')}
                      </Form.Control.Feedback>
                    </FormCheck>
                  </Form.Group>
                  {reportType.haveContent &&
                    reportType.type === item.reason_type && (
                      <Form.Group controlId="content" className="ps-4 mb-3">
                        <Form.Control
                          type="text"
                          as={
                            item.content_type === 'text' ? 'input' : 'textarea'
                          }
                          value={content.value}
                          isInvalid={content.isInvalid}
                          placeholder={item.placeholder}
                          onChange={(e) =>
                            setContent({
                              value: e.target.value,
                              isInvalid: false,
                              errorMsg: '',
                            })
                          }
                        />
                        <Form.Control.Feedback type="invalid">
                          {content.errorMsg}
                        </Form.Control.Feedback>
                      </Form.Group>
                    )}
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

export default useReportModal;
