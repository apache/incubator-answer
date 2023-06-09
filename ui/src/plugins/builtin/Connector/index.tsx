import { memo, FC } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import pluginKit, { PluginInfo } from '@/utils/pluginKit';
import { SvgIcon } from '@/components';

import info from './info.yaml';
import { useGetStartUseOauthConnector } from './services';
import './i18n';

const pluginInfo: PluginInfo = {
  slug_name: info.slug_name,
  type: info.type,
};
interface Props {
  className?: string;
}
const Index: FC<Props> = ({ className }) => {
  const { t } = useTranslation(pluginKit.getTransNs(), {
    keyPrefix: pluginKit.getTransKeyPrefix(pluginInfo),
  });

  const { data } = useGetStartUseOauthConnector();

  if (!data?.length) return null;
  return (
    <div className={classnames('d-grid gap-2', className)}>
      {data?.map((item) => {
        return (
          <Button variant="outline-secondary" href={item.link} key={item.name}>
            <SvgIcon base64={item.icon} />
            <span>{t('connect', { auth_name: item.name })}</span>
          </Button>
        );
      })}
    </div>
  );
};

export default {
  info: pluginInfo,
  component: memo(Index),
};
