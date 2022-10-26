import { memo, FC } from 'react';
import { Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { Modal } from '@answer/components';
import { useReportModal, useToast } from '@answer/hooks';
import { deleteQuestion, deleteAnswer } from '@answer/api';
import { isLogin } from '@answer/utils';
import Share from '../Share';

interface IProps {
  type: 'answer' | 'question';
  qid: string;
  aid?: string;
  title: string;
  hasAnswer?: boolean;
  isAccepted: boolean;
  callback: (type: string) => void;
  memberActions;
}
const Index: FC<IProps> = ({
  type,
  qid,
  aid = '',
  title,
  isAccepted = false,
  hasAnswer = false,
  memberActions = [],
  callback,
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'delete' });
  const toast = useToast();
  const reportModal = useReportModal();
  const closeModal = useReportModal();
  const editUrl =
    type === 'answer' ? `/posts/${qid}/${aid}/edit` : `/posts/${qid}/edit`;

  const handleReport = () => {
    reportModal.onShow({
      type,
      id: type === 'answer' ? aid : qid,
      action: 'flag',
    });
  };

  const handleClose = () => {
    closeModal.onShow({
      type,
      id: qid,
      action: 'close',
    });
  };

  const handleDelete = () => {
    if (type === 'question') {
      Modal.confirm({
        title: t('title'),
        content: hasAnswer ? `<p>${t('question')}</p>` : `<p>${t('other')}</p>`,
        cancelBtnVariant: 'link',
        confirmBtnVariant: 'danger',
        confirmText: t('delete', { keyPrefix: 'btns' }),
        onConfirm: () => {
          deleteQuestion({
            id: qid,
          }).then(() => {
            toast.onShow({
              msg: t('tip_question_deleted'),
              variant: 'success',
            });
            callback?.('delete_question');
          });
        },
      });
    }

    if (type === 'answer' && aid) {
      Modal.confirm({
        title: t('title'),
        content: isAccepted ? t('answer_accepted') : `<p>${t('other')}</p>`,
        cancelBtnVariant: 'link',
        confirmBtnVariant: 'danger',
        confirmText: t('delete', { keyPrefix: 'btns' }),
        onConfirm: () => {
          deleteAnswer({
            id: aid,
          }).then(() => {
            // refersh page
            toast.onShow({
              msg: t('tip_answer_deleted'),
              variant: 'success',
            });
            callback?.('all');
          });
        },
      });
    }
  };

  const handleAction = (action) => {
    if (!isLogin(true)) {
      return;
    }
    if (action === 'delete') {
      handleDelete();
    }

    if (action === 'report') {
      handleReport();
    }

    if (action === 'close') {
      handleClose();
    }
  };

  return (
    <div className="d-flex align-items-center">
      <Share type={type} qid={qid} aid={aid} title={title} />
      {memberActions?.map((item) => {
        if (item.action === 'edit') {
          return (
            <Link
              key={item.action}
              to={editUrl}
              className="link-secondary p-0 fs-14 me-3"
              style={{ lineHeight: '23px' }}>
              {t(`btn_${item.action}`)}
            </Link>
          );
        }
        return (
          <Button
            key={item.action}
            variant="link"
            className="link-secondary p-0 fs-14 me-3"
            onClick={() => handleAction(item.action)}>
            {t(`btn_${item.action}`)}
          </Button>
        );
      })}
    </div>
  );
};

export default memo(Index);
