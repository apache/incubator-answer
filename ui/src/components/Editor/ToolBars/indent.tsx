import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const Indent: FC<IEditorContext> = ({ editor, replaceLines }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'indent',
    tip: t('indent.text'),
  };
  const handleClick = () => {
    replaceLines((line) => {
      line = `    ${line}`;
      return line;
    });
    editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(Indent);
