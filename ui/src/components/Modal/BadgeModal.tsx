/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

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
  const { data, mutate } = useQueryNotificationStatus();

  const handle = async () => {
    if (!data) return;
    await readNotification(badge?.notification_id);
    await mutate({
      ...data,
      badge_award: null,
    });
  };
  const handleCancel = async () => {
    await handle();
  };
  const handleConfirm = async () => {
    await handle();

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
