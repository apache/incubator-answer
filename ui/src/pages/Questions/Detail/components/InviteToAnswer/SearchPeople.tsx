import { FC, useState } from 'react';
import { Card, Form, Dropdown } from 'react-bootstrap';

import { Avatar } from '@/components';

const data = [
  {
    name: 'name1',
    value: 1,
  },
  {
    name: 'name2',
    value: 2,
  },
  {
    name: 'name2',
    value: 3,
  },
  {
    name: 'name2',
    value: 4,
  },
  {
    name: 'name2',
    value: 5,
  },
];

const Index: FC = () => {
  const [checkedList, setCheckList] = useState<Array<number>>([]);
  const [currentIndex, setCurrentIndex] = useState(-1);

  const handleItemClick = (item: any) => {
    const index = checkedList.findIndex((v) => v === item.value);
    if (index > -1) {
      checkedList.splice(index, 1);
    } else {
      checkedList.push(item.value);
    }
    setCheckList([...checkedList]);
  };

  const handleKeyDown = (evt) => {
    evt.stopPropagation();

    const { keyCode } = evt;
    console.log('keyCode', keyCode, currentIndex);
    if (keyCode === 38 && currentIndex > 0) {
      setCurrentIndex(currentIndex - 1);
    }
    if (keyCode === 40 && currentIndex < data.length - 1) {
      setCurrentIndex(currentIndex + 1);
    }

    if (keyCode === 13 && currentIndex > -1) {
      evt.preventDefault();
      console.log('enter', currentIndex);
      handleItemClick(data[currentIndex]);
    }
  };

  const handleSelect = (evt) => {
    const { eventKey } = evt;
    console.log('key', eventKey);
  };
  return (
    <Card.Body className="py-3">
      <Dropdown
        autoClose="outside"
        onKeyDown={handleKeyDown}
        onSelect={handleSelect}
        show
        // focusFirstItemOnShow
      >
        <Dropdown.Toggle
          id="dropdown-autoclose-true"
          as={Form.Control}
          placeholder="Search people"
          className="search-people"
          autoFocus
        />
        <Dropdown.Menu show className="w-100 position-relative">
          {data.map((item, index) => {
            return (
              <Dropdown.Item
                key={item.value}
                eventKey={index}
                // active={index === currentIndex}
                onClick={() => handleItemClick(item)}>
                <Form.Check type="checkbox" id={item.value.toString()}>
                  <Form.Check.Input
                    type="checkbox"
                    checked={Boolean(checkedList.find((v) => v === item.value))}
                  />
                  <Form.Check.Label>
                    <Avatar
                      avatar=""
                      size="24"
                      alt="test"
                      className="rounded-1"
                    />
                    <small className="ms-2">{item.name}</small>
                  </Form.Check.Label>
                </Form.Check>
              </Dropdown.Item>
            );
          })}
        </Dropdown.Menu>
      </Dropdown>

      {data.map((item) => {
        return (
          // eslint-disable-next-line jsx-a11y/anchor-is-valid
          <a href="#" key={item.value} onClick={() => handleItemClick(item)}>
            <Form.Check type="checkbox" id={item.value.toString()}>
              <Form.Check.Input
                type="checkbox"
                checked={Boolean(checkedList.find((v) => v === item.value))}
              />
              <Form.Check.Label>
                <Avatar avatar="" size="24" alt="test" className="rounded-1" />
                <small className="ms-2">{item.name}</small>
              </Form.Check.Label>
            </Form.Check>
          </a>
        );
      })}
    </Card.Body>
  );
};

export default Index;
