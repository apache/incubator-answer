import React, { memo, FC } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { SvgIcon } from '@/components';
import { userCenterStore } from '@/stores';

import WeComLogin from './WeCom';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'plugins.oauth' });
  const ucAgent = userCenterStore().agent;
  const agentName = ucAgent?.agent_info?.name || '';
  const ucLoginRedirect =
    ucAgent?.enabled && ucAgent?.agent_info?.login_redirect_url;
  if (/WeCom/i.test(agentName)) {
    return <WeComLogin />;
  }
  if (ucLoginRedirect) {
    return (
      <Button
        className="w-100"
        variant="outline-secondary"
        href={ucAgent?.agent_info.login_redirect_url}>
        <SvgIcon base64={ucAgent?.agent_info.icon} />
        <span>{t('connect', { auth_name: ucAgent?.agent_info.name })}</span>
      </Button>
    );
  }
  return null;
};

export default memo(Index);
