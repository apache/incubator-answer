import { memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

let context: IEditorContext;
const BlockQuote = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });

  const item = {
    label: 'blockquote',
    keyMap: ['Ctrl-Q'],
    tip: `${t('blockquote.text')} (Ctrl+Q)`,
  };

  const handleClick = (ctx) => {
    context = ctx;
    context.replaceLines((line) => {
      const FIND_BLOCKQUOTE_RX = /^>\s+?/g;

      if (line === `> ${t('blockquote.text')}`) {
        line = '';
      } else if (line.match(FIND_BLOCKQUOTE_RX)) {
        line = line.replace(FIND_BLOCKQUOTE_RX, '');
      } else {
        line = `> ${line || t('blockquote.text')}`;
      }
      return line;
    }, 2);
    context.editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(BlockQuote);
