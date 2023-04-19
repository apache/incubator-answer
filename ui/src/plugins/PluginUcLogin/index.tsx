import { memo, FC } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import { SvgIcon } from '@/components';
import { userCenterStore } from '@/stores';

interface Props {
  className?: classnames.Argument;
}
const Index: FC<Props> = ({ className }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'plugins.oauth' });
  const ucAgent = userCenterStore().agent;
  const agentName = ucAgent?.agent_info?.name || '';
  const ucLoginRedirect =
    ucAgent?.enabled && ucAgent?.agent_info?.login_redirect_url;

  if (ucLoginRedirect) {
    return (
      <Button
        className={classnames('w-100', className)}
        variant="outline-secondary"
        href={ucAgent?.agent_info.login_redirect_url}>
        <SvgIcon base64={ucAgent?.agent_info.icon} />
        <span>{t('connect', { auth_name: agentName })}</span>
      </Button>
    );
  }
  return null;
};

export default memo(Index);
