/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import { memo } from 'react';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

let context: IEditorContext;
const Table = () => {
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
  const handleClick = (ctx) => {
    context = ctx;
    const { editor } = context;
    let table = '\n';

    table += makeHeader(2, tableData);
    table += makeBody(2, 2, tableData);
    editor?.replaceSelection(table);
    editor?.focus();
  };

  return <ToolItem {...item} onClick={handleClick} />;
};

export default memo(Table);
