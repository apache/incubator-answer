import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const Outdent: FC<IEditorContext> = ({ editor, replaceLines }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'outdent',
    keyMap: ['Shift-Tab'],
    tip: t('outdent.text'),
  };
  const handleClick = () => {
    replaceLines((line) => {
      line = line.replace(/^(\s{0,})/, (_1, $1) => {
        return $1.length > 4 ? $1.substring(4) : '';
      });
      return line;
    });
    editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(Outdent);
