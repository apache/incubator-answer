import { memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

let context: IEditorContext;
const UL = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'unorderedList',
    keyMap: ['Ctrl-U'],
    tip: `${t('unordered_list.text')} (Ctrl+U)`,
  };

  const handleClick = (ctx) => {
    context = ctx;
    const { editor, replaceLines } = context;

    replaceLines((line) => {
      const FIND_UL_RX = /^(\s{0,})(-|\*)\s/;

      if (line.match(FIND_UL_RX)) {
        line = line.replace(FIND_UL_RX, '');
      } else {
        line = `* ${line}`;
      }
      return line;
    });
    editor.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(UL);
