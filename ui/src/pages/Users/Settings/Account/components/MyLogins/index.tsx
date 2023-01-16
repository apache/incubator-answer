import { memo } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { Modal } from '@/components';
import { useOauthConnectorInfoByUser, userOauthUnbind } from '@/services';
import { useToast } from '@/hooks';

const Index = () => {
  const { data, mutate } = useOauthConnectorInfoByUser();
  const toast = useToast();

  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.my_logins',
  });

  const deleteLogins = (e, item) => {
    if (!item.binding) {
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
      <small className="form-text mt-0">{t('lable')}</small>

      <div className="mt-3">
        {data?.map((item) => {
          return (
            <Button
              variant={item.binding ? 'outline-danger' : 'outline-secondary'}
              href={item.link}
              onClick={(e) => deleteLogins(e, item)}
              key={item.name}>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                className="btnSvg"
                viewBox="0 0 24 24">
                <path d={item.icon} />
              </svg>
              <span> {item.name}</span>
            </Button>
          );
        })}
      </div>
    </div>
  );
};

export default memo(Index);
