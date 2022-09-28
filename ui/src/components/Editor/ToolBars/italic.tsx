import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const Italic: FC<IEditorContext> = ({ editor, wrapText }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'italic',
    keyMap: ['Ctrl-I'],
    tip: `${t('italic.text')} (Ctrl+I)`,
  };
  const DEFAULTTEXT = t('italic.text');

  const handleClick = () => {
    wrapText('*', '*', DEFAULTTEXT);
    editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(Italic);
