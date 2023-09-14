import { Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { Icon, Modal } from '@/components';
import { changeAnswerStatus } from '@/services';

const AnswerActions = ({ itemData, curFilter, refreshList }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'delete' });

  const handleAction = (type) => {
    console.log(type);
    if (type === 'delete') {
      Modal.confirm({
        title: t('title'),
        content: itemData.accepted === 2 ? t('answer_accepted') : t('other'),
        cancelBtnVariant: 'link',
        confirmBtnVariant: 'danger',
        confirmText: t('delete', { keyPrefix: 'btns' }),
        onConfirm: () => {
          changeAnswerStatus(itemData.id, 'deleted').then(() => {
            refreshList();
          });
        },
      });
    }

    if (type === 'undelete') {
      Modal.confirm({
        title: t('undelete_title'),
        content: t('undelete_desc'),
        cancelBtnVariant: 'link',
        confirmBtnVariant: 'danger',
        confirmText: t('undelete', { keyPrefix: 'btns' }),
        onConfirm: () => {
          changeAnswerStatus(itemData.id, 'available').then(() => {
            refreshList();
          });
        },
      });
    }
  };

  return (
    <Dropdown>
      <Dropdown.Toggle variant="link" className="no-toggle p-0">
        <Icon
          name="three-dots-vertical"
          title={t('action', { keyPrefix: 'admin.answers' })}
        />
      </Dropdown.Toggle>
      <Dropdown.Menu>
        {curFilter === 'deleted' ? (
          <Dropdown.Item onClick={() => handleAction('undelete')}>
            {t('undelete', { keyPrefix: 'btns' })}
          </Dropdown.Item>
        ) : (
          <Dropdown.Item onClick={() => handleAction('delete')}>
            {t('delete', { keyPrefix: 'btns' })}
          </Dropdown.Item>
        )}
      </Dropdown.Menu>
    </Dropdown>
  );
};

export default AnswerActions;
