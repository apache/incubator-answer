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

import { FC, useEffect, useState, memo } from 'react';
import { Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import ToolItem from '../toolItem';
import { IEditorContext } from '../types';

const Chart: FC<IEditorContext> = ({ editor }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'editor' });

  const headerList = [
    {
      label: t('chart.flow_chart'),
      tpl: `graph TD
      A[Christmas] -->|Get money| B(Go shopping)
      B --> C{Let me think}
      C -->|One| D[Laptop]
      C -->|Two| E[iPhone]
      C -->|Three| F[fa:fa-car Car]`,
    },
    {
      label: t('chart.sequence_diagram'),
      tpl: `sequenceDiagram
      Alice->>+John: Hello John, how are you?
      Alice->>+John: John, can you hear me?
      John-->>-Alice: Hi Alice, I can hear you!
      John-->>-Alice: I feel great!
              `,
    },
    {
      label: t('chart.state_diagram'),
      tpl: `stateDiagram-v2
      [*] --> Still
      Still --> [*]
      Still --> Moving
      Moving --> Still
      Moving --> Crash
      Crash --> [*]
              `,
    },
    {
      label: t('chart.class_diagram'),
      tpl: `classDiagram
      Animal <|-- Duck
      Animal <|-- Fish
      Animal <|-- Zebra
      Animal : +int age
      Animal : +String gender
      Animal: +isMammal()
      Animal: +mate()
      class Duck{
        +String beakColor
        +swim()
        +quack()
      }
      class Fish{
        -int sizeInFeet
        -canEat()
      }
      class Zebra{
        +bool is_wild
        +run()
      }
              `,
    },
    {
      label: t('chart.pie_chart'),
      tpl: `pie title Pets adopted by volunteers
      "Dogs" : 386
      "Cats" : 85
      "Rats" : 15
              `,
    },
    {
      label: t('chart.gantt_chart'),
      tpl: `gantt
      title A Gantt Diagram
      dateFormat  YYYY-MM-DD
      section Section
      A task           :a1, 2014-01-01, 30d
      Another task     :after a1  , 20d
      section Another
      Task in sec      :2014-01-12  , 12d
      another task      : 24d
              `,
    },
    {
      label: t('chart.entity_relationship_diagram'),
      tpl: `erDiagram
      CUSTOMER }|..|{ DELIVERY-ADDRESS : has
      CUSTOMER ||--o{ ORDER : places
      CUSTOMER ||--o{ INVOICE : "liable for"
      DELIVERY-ADDRESS ||--o{ ORDER : receives
      INVOICE ||--|{ ORDER : covers
      ORDER ||--|{ ORDER-ITEM : includes
      PRODUCT-CATEGORY ||--|{ PRODUCT : contains
      PRODUCT ||--o{ ORDER-ITEM : "ordered in"
        `,
    },
  ];
  const item = {
    label: 'chart',
    tip: `${t('chart.text')}`,
  };
  const [isShow, setShowState] = useState(false);
  const [isLocked, setLockState] = useState(false);

  useEffect(() => {
    if (!editor) {
      return;
    }
    editor.on('focus', () => {
      setShowState(false);
    });
  }, []);

  const click = (tpl) => {
    const { ch } = editor.getCursor();

    editor.replaceSelection(`${ch ? '\n' : ''}\`\`\`mermaid\n${tpl}\n\`\`\`\n`);
  };

  const onAddHeader = () => {
    setShowState(!isShow);
  };
  const handleMouseEnter = () => {
    if (isLocked) {
      return;
    }
    setLockState(true);
  };

  const handleMouseLeave = () => {
    setLockState(false);
  };
  return (
    <ToolItem
      as="dropdown"
      {...item}
      onClick={onAddHeader}
      onBlur={onAddHeader}>
      <Dropdown.Menu
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}>
        {headerList.map((header) => {
          return (
            <Dropdown.Item
              key={header.label}
              onClick={(e) => {
                e.preventDefault();
                click(header.tpl);
              }}>
              {header.label}
            </Dropdown.Item>
          );
        })}
      </Dropdown.Menu>
    </ToolItem>
  );
};

export default memo(Chart);
