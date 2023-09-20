import { FC, useState } from 'react';
import { Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import 'katex/dist/katex.min.css';

import icon from './icon.svg';
import { useRenderFormula } from './hooks';

interface FormulaProps {
  editor;
  previewElement: HTMLElement;
}

const Formula: FC<FormulaProps> = ({ editor, previewElement }) => {
  useRenderFormula(previewElement);
  const { t } = useTranslation('plugin', {
    keyPrefix: 'formula',
  });
  const [isLocked, setLockState] = useState(false);

  const handleMouseEnter = () => {
    if (isLocked) {
      return;
    }
    setLockState(true);
  };

  const handleMouseLeave = () => {
    setLockState(false);
  };
  const formulaList = [
    {
      type: 'line',
      label: t('options.inline'),
    },
    {
      type: 'block',
      label: t('options.block'),
    },
  ];

  const handleClick = (type: string, label: string) => {
    if (!editor) {
      return;
    }
    const { wrapText } = editor;
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
  };
  return (
    <div className="toolbar-item-wrap">
      <Dropdown className="p-0 b-0 btn-no-border btn btn-link" title="chart">
        <Dropdown.Toggle
          type="button"
          as="button"
          className="p-0 b-0 btn-no-border btn btn-link">
          <img src={icon} alt="formula" />
        </Dropdown.Toggle>
        <Dropdown.Menu
          onMouseEnter={handleMouseEnter}
          onMouseLeave={handleMouseLeave}>
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
      </Dropdown>
    </div>
  );
};

export default Formula;
