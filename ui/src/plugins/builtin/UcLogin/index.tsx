import { memo, FC } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import pluginKit, { PluginInfo } from '@/utils/pluginKit';
import { SvgIcon } from '@/components';
import { userCenterStore } from '@/stores';
import './i18n';

import info from './info.yaml';

interface Props {
  className?: classnames.Argument;
}

const pluginInfo: PluginInfo = {
  slug_name: info.slug_name,
};

const Index: FC<Props> = ({ className }) => {
  const { t } = useTranslation(pluginKit.getTransNs(), {
    keyPrefix: pluginKit.getTransKeyPrefix(pluginInfo),
  });
  const ucAgent = userCenterStore().agent;
  const ucLoginRedirect =
    ucAgent?.enabled && ucAgent?.agent_info?.login_redirect_url;

  if (ucLoginRedirect) {
    return (
      <Button
        className={classnames('w-100', className)}
        variant="outline-secondary"
        href={ucAgent?.agent_info.login_redirect_url}>
        <SvgIcon base64={ucAgent?.agent_info.icon} />
        <span>
          {t('connect', { auth_name: ucAgent?.agent_info.display_name })}
        </span>
      </Button>
    );
  }
  return null;
};
export default {
  info: pluginInfo,
  component: memo(Index),
};
