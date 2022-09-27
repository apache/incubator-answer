import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const UL: FC<IEditorContext> = ({ editor, replaceLines }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'unorderedList',
    keyMap: ['Ctrl-U'],
    tip: `${t('unordered_list.text')} (Ctrl+U)`,
  };

  const handleClick = () => {
    if (!editor) {
      return;
    }
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

  return <ToolItem {...item} click={handleClick} />;
};

export default memo(UL);
