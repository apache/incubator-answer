import { memo } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { Modal } from '@/components';
import { useOauthConnectorInfoByUser, userOauthUnbind } from '@/services';
import { useToast } from '@/hooks';
import { base64ToSvg } from '@/utils';
import Storage from '@/utils/storage';
import { REDIRECT_PATH_STORAGE_KEY } from '@/common/constants';

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
      Storage.set(REDIRECT_PATH_STORAGE_KEY, window.location.pathname);
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
                    __html: base64ToSvg(item.icon),
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
