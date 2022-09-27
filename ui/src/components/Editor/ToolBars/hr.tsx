import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const Hr: FC<IEditorContext> = ({ editor, appendBlock }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'hr',
    keyMap: ['Ctrl-R'],
    tip: `${t('hr.text')} (Ctrl+R)`,
  };
  const handleClick = () => {
    appendBlock('----------\n');
    editor?.focus();
  };

  return <ToolItem {...item} click={handleClick} />;
};

export default memo(Hr);
