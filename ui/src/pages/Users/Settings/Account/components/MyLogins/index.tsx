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

import { memo } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { Modal } from '@/components';
import { useOauthConnectorInfoByUser, userOauthUnbind } from '@/services';
import { useToast } from '@/hooks';
import { base64ToSvg } from '@/utils';
import Storage from '@/utils/storage';
import { REDIRECT_PATH_STORAGE_KEY } from '@/common/constants';
import { REACT_BASE_PATH } from '@/router/alias';

const Index = () => {
  const { data, mutate } = useOauthConnectorInfoByUser();
  const toast = useToast();

  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.my_logins',
  });

  const { t: t2 } = useTranslation('translation', {
    keyPrefix: 'oauth',
  });

  const deleteLogins = (e, item) => {
    if (!item.binding) {
      Storage.set(
        REDIRECT_PATH_STORAGE_KEY,
        window.location.pathname.replace(REACT_BASE_PATH, ''),
      );
      return;
    }
    e.preventDefault();
    Modal.confirm({
      title: t('modal_title'),
      content: t('modal_content'),
      confirmBtnVariant: 'danger',
      confirmText: t('modal_confirm_btn'),
      onConfirm: () => {
        userOauthUnbind({ external_id: item.external_id }).then(() => {
          mutate();
          toast.onShow({
            msg: t('remove_success'),
            variant: 'success',
          });
        });
      },
    });
  };

  if (!data?.length) return null;
  return (
    <div className="mt-5">
      <div className="form-label">{t('title')}</div>
      <small className="form-text mt-0">{t('label')}</small>

      <div className="d-grid gap-2 mt-3">
        {data?.map((item) => {
          return (
            <div key={item.name}>
              <Button
                variant={item.binding ? 'outline-danger' : 'outline-secondary'}
                href={item.link}
                onClick={(e) => deleteLogins(e, item)}>
                <span
                  dangerouslySetInnerHTML={{
                    __html: base64ToSvg(item.icon, 'btnSvg me-2'),
                  }}
                />
                <span>
                  {t2(item.binding ? 'remove' : 'connect', {
                    auth_name: item.name,
                  })}
                </span>
              </Button>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default memo(Index);
