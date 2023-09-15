import { memo, FC } from 'react';
import { useTranslation } from 'react-i18next';
import { Alert } from 'react-bootstrap';

import { PluginInfo } from '@/utils/pluginKit';
import { getTransNs, getTransKeyPrefix } from '@/utils/pluginKit/utils';

import './i18n';

import info from './info.yaml';

const pluginInfo: PluginInfo = {
  slug_name: info.slug_name,
  type: info.type,
};

const Index: FC = () => {
  const { t } = useTranslation(getTransNs(), {
    keyPrefix: getTransKeyPrefix(pluginInfo),
  });

  return <Alert variant="info">{t('msg')}</Alert>;
};
export default {
  info: pluginInfo,
  component: memo(Index),
};
