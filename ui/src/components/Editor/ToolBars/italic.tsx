import { memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

let context: IEditorContext;
const Italic = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'italic',
    keyMap: ['Ctrl-I'],
    tip: `${t('italic.text')} (Ctrl+I)`,
  };
  const DEFAULTTEXT = t('italic.text');

  const handleClick = (ctx) => {
    context = ctx;
    const { editor, wrapText } = context;
    wrapText('*', '*', DEFAULTTEXT);
    editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(Italic);
