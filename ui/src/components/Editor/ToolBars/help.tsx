import { memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';

const Help = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });

  const item = {
    label: 'help',
    tip: t('help.text'),
  };
  const handleClick = () => {
    window.open('https://commonmark.org/help/');
  };

  return <ToolItem {...item} click={handleClick} />;
};

export default memo(Help);
