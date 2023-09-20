import { memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

let context: IEditorContext;
const Hr = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'hr',
    keyMap: ['Ctrl-R'],
    tip: `${t('hr.text')} (Ctrl+R)`,
  };
  const handleClick = (ctx) => {
    context = ctx;
    const { appendBlock, editor } = context;
    appendBlock('----------\n');
    editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(Hr);
