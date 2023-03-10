import { memo, FC } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import { useGetStartUseOauthConnector } from '@/services';
import { base64ToSvg } from '@/utils';

interface Props {
  className?: string;
}
const Index: FC<Props> = ({ className }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'plugins.oauth' });
  const { data } = useGetStartUseOauthConnector();

  if (!data?.length) return null;
  return (
    <div className={classnames('d-grid gap-2', className)}>
      {data?.map((item) => {
        return (
          <Button variant="outline-secondary" href={item.link} key={item.name}>
            <span
              dangerouslySetInnerHTML={{
                __html: base64ToSvg(item.icon),
              }}
            />

            <span>{t('connect', { auth_name: item.name })}</span>
          </Button>
        );
      })}
    </div>
  );
};

export default memo(Index);
