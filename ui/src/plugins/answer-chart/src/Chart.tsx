import { FC, useState } from 'react';
import { Dropdown } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import icon from './icon.svg';
import { useRenderChart } from './hooks';

interface ChartProps {
  editor;
  previewElement: HTMLElement;
}

const Chart: FC<ChartProps> = ({ editor, previewElement }) => {
  useRenderChart(previewElement);
  const { t } = useTranslation('plugin', {
    keyPrefix: 'chart',
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
  const headerList = [
    {
      label: t('flow_chart'),
      tpl: `graph TD
      A[Christmas] -->|Get money| B(Go shopping)
      B --> C{Let me think}
      C -->|One| D[Laptop]
      C -->|Two| E[iPhone]
      C -->|Three| F[fa:fa-car Car]`,
    },
    {
      label: t('sequence_diagram'),
      tpl: `sequenceDiagram
      Alice->>+John: Hello John, how are you?
      Alice->>+John: John, can you hear me?
      John-->>-Alice: Hi Alice, I can hear you!
      John-->>-Alice: I feel great!
              `,
    },
    {
      label: t('state_diagram'),
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
      label: t('class_diagram'),
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
      label: t('pie_chart'),
      tpl: `pie title Pets adopted by volunteers
      "Dogs" : 386
      "Cats" : 85
      "Rats" : 15
              `,
    },
    {
      label: t('gantt_chart'),
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
      label: t('entity_relationship_diagram'),
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

  const handleChange = (tpl: string) => {
    const { ch } = editor.getCursor();

    editor.replaceSelection(`${ch ? '\n' : ''}\`\`\`mermaid\n${tpl}\n\`\`\`\n`);
  };

  return (
    <div className="toolbar-item-wrap">
      <Dropdown className="p-0 b-0 btn-no-border btn btn-link" title="chart">
        <Dropdown.Toggle
          type="button"
          as="button"
          className="p-0 b-0 btn-no-border btn btn-link">
          <img src={icon} alt="chart" />
        </Dropdown.Toggle>
        <Dropdown.Menu
          onMouseEnter={handleMouseEnter}
          onMouseLeave={handleMouseLeave}>
          {headerList.map((header) => {
            return (
              <Dropdown.Item
                key={header.label}
                onClick={(e) => {
                  e.preventDefault();
                  handleChange(header.tpl);
                }}>
                {header.label}
              </Dropdown.Item>
            );
          })}
        </Dropdown.Menu>
      </Dropdown>
    </div>
  );
};

export default Chart;
