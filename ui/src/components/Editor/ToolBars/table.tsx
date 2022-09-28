import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const Table: FC<IEditorContext> = ({ editor }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const item = {
    label: 'table',
    tip: t('table.text'),
  };
  const tableData = [
    [`${t('table.heading')} A`],
    [`${t('table.heading')} B`],
    [`${t('table.cell')} 1`],
    [`${t('table.cell')} 2`],
    [`${t('table.cell')} 3`],
    [`${t('table.cell')} 4`],
  ];

  const makeHeader = (col, data) => {
    let header = '|';
    let border = '|';
    let index = 0;

    while (col) {
      if (data) {
        header += ` ${data[index]} |`;
        index += 1;
      } else {
        header += '  |';
      }

      border += ' ----- |';

      col -= 1;
    }

    return `${header}\n${border}\n`;
  };

  const makeBody = (col, row, data) => {
    let body = '';
    let index = col;

    for (let irow = 0; irow < row; irow += 1) {
      body += '|';

      for (let icol = 0; icol < col; icol += 1) {
        if (data) {
          body += ` ${data[index]} |`;
          index += 1;
        } else {
          body += '  |';
        }
      }

      body += '\n';
    }

    body = body.replace(/\n$/g, '');

    return body;
  };
  const handleClick = () => {
    let table = '\n';

    table += makeHeader(2, tableData);
    table += makeBody(2, 2, tableData);
    editor?.replaceSelection(table);
    editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(Table);
