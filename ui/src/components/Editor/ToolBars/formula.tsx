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
  const [isLocked, setLockState] = useState(false);

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
    if (isLocked) {
      return;
    }
    setShowState(!isShow);
  };

  const handleMouseEnter = () => {
    setLockState(true);
  };

  const handleMouseLeave = () => {
    setLockState(false);
  };
  return (
    <ToolItem
      as="dropdown"
      {...item}
      isShow={isShow}
      onClick={onAddFormula}
      onBlur={onAddFormula}>
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
    </ToolItem>
  );
};

export default memo(Formula);
