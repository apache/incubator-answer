import { FC } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';

import Progress from '../Progress';

interface Props {
  visible: boolean;
  siteUrl: string;
}
const Index: FC<Props> = ({ visible, siteUrl = '' }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'install' });

  if (!visible) return null;
  return (
    <div>
      <h5>{t('ready_title')}</h5>
      <p>
        <Trans i18nKey="install.ready_description">
          If you ever feel like changing more settings, visit
          <a href={`${siteUrl}/users/login`}> admin section</a>; find it in the
          site menu.
        </Trans>
      </p>
      <p>{t('good_luck')}</p>

      <div className="d-flex align-items-center justify-content-between">
        <Progress step={5} />
        <Button href={siteUrl}>{t('done')}</Button>
      </div>
    </div>
  );
};

export default Index;
