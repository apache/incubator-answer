import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const Bold: FC<IEditorContext> = ({ editor, wrapText }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'bold',
    keyMap: ['Ctrl-B'],
    tip: `${t('bold.text')} (Ctrl+B)`,
  };
  const DEFAULTTEXT = t('bold.text');

  const handleClick = () => {
    wrapText('**', '**', DEFAULTTEXT);
    editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(Bold);
