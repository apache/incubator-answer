import { FC, useState, memo } from 'react';
import { Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const Formula: FC<IEditorContext> = ({ editor, wrapText }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const formulaList = [
    {
      type: 'line',
      label: t('formula.options.inline'),
    },
    {
      type: 'block',
      label: t('formula.options.block'),
    },
  ];
  const item = {
    label: 'formula',
    tip: t('formula.text'),
  };
  const [isShow, setShowState] = useState(false);

  const handleClick = (type, label) => {
    if (!editor) {
      return;
    }
    if (type === 'line') {
      wrapText('\\\\( ', ' \\\\)', label);
    } else {
      const cursor = editor.getCursor();

      wrapText('\n$$\n', '\n$$\n', label);

      editor.setSelection(
        { line: cursor.line + 2, ch: 0 },
        { line: cursor.line + 2, ch: label.length },
      );
    }
    editor?.focus();
    setShowState(false);
  };
  const onAddFormula = () => {
    setShowState(!isShow);
  };

  return (
    <ToolItem as="dropdown" {...item} isShow={isShow} click={onAddFormula}>
      <Dropdown.Menu>
        {formulaList.map((formula) => {
          return (
            <Dropdown.Item
              key={formula.label}
              onClick={(e) => {
                e.preventDefault();
                handleClick(formula.type, formula.label);
              }}>
              {formula.label}
            </Dropdown.Item>
          );
        })}
      </Dropdown.Menu>
    </ToolItem>
  );
};

export default memo(Formula);
