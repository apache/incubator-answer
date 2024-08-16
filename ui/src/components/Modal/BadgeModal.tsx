import { FC } from 'react';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';

import classNames from 'classnames';

import type * as Type from '@/common/interface';
import { loggedUserInfoStore } from '@/stores';
import { readNotification, useQueryNotificationStatus } from '@/services';
import Icon from '../Icon';

import Modal from './Modal';

interface BadgeModalProps {
  badge?: Type.NotificationBadgeAward | null;
  visible: boolean;
}
const BadgeModal: FC<BadgeModalProps> = ({ badge, visible }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'badges.modal' });
  const { user } = loggedUserInfoStore();
  const navigate = useNavigate();
  const { data } = useQueryNotificationStatus();

  const handleCancel = async () => {
    if (!data) return;
    await readNotification(badge?.notification_id);
  };
  const handleConfirm = async () => {
    await readNotification(badge?.notification_id);

    const url = `/badges/${badge?.badge_id}?username=${user.username}`;
    navigate(url);
  };

  return (
    <Modal
      title={t('title')}
      visible={visible}
      onCancel={handleCancel}
      onConfirm={handleConfirm}
      cancelText={t('close')}
      cancelBtnVariant="link"
      confirmText={t('confirm')}
      confirmBtnVariant="primary"
      scrollable={false}>
      {badge && (
        <div className="text-center">
          {badge.icon?.startsWith('http') ? (
            <img src={badge.icon} width={96} height={96} alt={badge.name} />
          ) : (
            <Icon
              name={badge.icon}
              size="96px"
              className={classNames(
                'lh-1',
                badge.level === 1 && 'bronze',
                badge.level === 2 && 'silver',
                badge.level === 3 && 'gold',
              )}
            />
          )}
          <h5 className="mt-3">{badge?.name}</h5>
          <p>{t('content')}</p>
        </div>
      )}
    </Modal>
  );
};

export default BadgeModal;
