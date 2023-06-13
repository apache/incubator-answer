import { memo, FC } from 'react';
import { useTranslation } from 'react-i18next';
import { Alert } from 'react-bootstrap';

import pluginKit, { PluginInfo } from '@/utils/pluginKit';
import './i18n';

import info from './info.yaml';

const pluginInfo: PluginInfo = {
  slug_name: info.slug_name,
};

const Index: FC = () => {
  const { t } = useTranslation(pluginKit.getTransNs(), {
    keyPrefix: pluginKit.getTransKeyPrefix(pluginInfo),
  });

  return <Alert variant="info">{t('msg')}</Alert>;
};
export default {
  info: pluginInfo,
  component: memo(Index),
};
