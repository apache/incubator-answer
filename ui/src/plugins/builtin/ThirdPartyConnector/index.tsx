import { memo, FC } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import { PluginInfo } from '@/utils/pluginKit';
import { getTransNs, getTransKeyPrefix } from '@/utils/pluginKit/utils';
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
  const { t } = useTranslation(getTransNs(), {
    keyPrefix: getTransKeyPrefix(pluginInfo),
  });

  const { data } = useGetStartUseOauthConnector();

  if (!data?.length) return null;
  return (
    <div className={classnames('d-grid gap-2', className)}>
      {data?.map((item) => {
        return (
          <Button variant="outline-secondary" href={item.link} key={item.name}>
            <SvgIcon base64={item.icon} svgClassName="btnSvg" />
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
