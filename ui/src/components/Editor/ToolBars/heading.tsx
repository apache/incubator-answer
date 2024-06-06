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

import { useState, memo } from 'react';
import { Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

let context: IEditorContext;
const Heading = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });
  const headerList = [
    {
      text: `<h1 class="mb-0 h3">${t('heading.options.h1')}</h1>`,
      level: 1,
      label: t('heading.options.h1'),
    },
    {
      text: `<h2 class="mb-0 h4">${t('heading.options.h2')}</h2>`,
      level: 2,
      label: t('heading.options.h2'),
    },
    {
      text: `<h3 class="mb-0 h5">${t('heading.options.h3')}</h3>`,
      level: 3,
      label: t('heading.options.h3'),
    },
    {
      text: `<h4 class="mb-0 h6">${t('heading.options.h4')}</h4>`,
      level: 4,
      label: t('heading.options.h4'),
    },
    {
      text: `<h5 class="mb-0 small">${t('heading.options.h5')}</h5>`,
      level: 5,
      label: t('heading.options.h5'),
    },
    {
      text: `<h6 class="mb-0 fs-12">${t('heading.options.h6')}</h6>`,
      level: 6,
      label: t('heading.options.h6'),
    },
  ];
  const item = {
    label: 'type',
    keyMap: ['Ctrl-h'],
    tip: `${t('heading.text')} (Ctrl+h)`,
  };
  const [isShow, setShowState] = useState(false);
  const [isLocked, setLockState] = useState(false);

  const handleClick = (level = 2, label = '大标题') => {
    const { replaceLines } = context;

    replaceLines((line) => {
      line = line.trim().replace(/^#*/, '').trim();
      line = `${'#'.repeat(level)} ${line || label}`;
      return line;
    }, level + 1);
    setShowState(false);
  };
  const onAddHeader = (ctx) => {
    context = ctx;
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
      onClick={onAddHeader}
      onBlur={onAddHeader}>
      <Dropdown.Menu
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}>
        {headerList.map((header) => {
          return (
            <Dropdown.Item
              key={header.text}
              onClick={(e) => {
                e.preventDefault();
                e.stopPropagation();
                handleClick(header.level, header.label);
              }}
              dangerouslySetInnerHTML={{ __html: header.text }}
            />
          );
        })}
      </Dropdown.Menu>
    </ToolItem>
  );
};

export default memo(Heading);
