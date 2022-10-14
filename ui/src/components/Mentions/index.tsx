import React, { useEffect, useRef, useState, FC } from 'react';
import { Dropdown } from 'react-bootstrap';

import * as Types from '@answer/common/interface';

interface IProps {
  children: React.ReactNode;
  pageUsers;
  onSelected: (val: string) => void;
}

const MAX_RECODE = 5;

const Mentions: FC<IProps> = ({ children, pageUsers, onSelected }) => {
  const menuRef = useRef<HTMLDivElement>(null);
  const dropdownRef = useRef<HTMLDivElement>(null);
  const [val, setValue] = useState('');
  const [users, setUsers] = useState<Types.PageUser[]>([]);
  const [cursor, setCursor] = useState(0);
  const [isRequested, setRequestedState] = useState(false);

  const searchUser = () => {
    const element = dropdownRef.current?.children[0];
    const { value, selectionStart = 0 } = element as HTMLTextAreaElement;

    if (value.indexOf('@') < 0) {
      setValue('');
    }
    if (!selectionStart) {
      return;
    }

    const str = value.substring(
      value.substring(0, selectionStart).lastIndexOf('@'),
      selectionStart,
    );

    if (str.substring(str.lastIndexOf(' '), selectionStart).indexOf('@') < 0) {
      return;
    }
    setValue(str.substring(1));

    if (!str.substring(1)) {
      return;
    }
    if (isRequested) {
      return;
    }
    setRequestedState(true);
  };

  useEffect(() => {
    const element = dropdownRef.current?.children[0] as HTMLTextAreaElement;

    if (element) {
      element.addEventListener('input', searchUser);
    }
    return () => {
      element.removeEventListener('input', searchUser);
    };
  }, [dropdownRef]);

  useEffect(() => {
    setUsers(pageUsers);
  }, [pageUsers, val]);

  const handleClick = (item) => {
    const element = dropdownRef.current?.children[0] as HTMLTextAreaElement;

    const { value, selectionStart = 0 } = element;

    if (!selectionStart) {
      return;
    }

    const text = `@${item?.displayName}[${item?.userName}] `;
    onSelected(
      `${value.substring(
        0,
        value.substring(0, selectionStart).lastIndexOf('@'),
      )}${text}${value.substring(selectionStart)}`,
    );
    setUsers([]);
    setValue('');
  };
  const filterData = val
    ? users.filter(
        (item) =>
          item.displayName?.indexOf(val) === 0 ||
          item.userName?.indexOf(val) === 0,
      )
    : [];
  const handleKeyDown = (e) => {
    const { keyCode } = e;

    if (keyCode === 38 && cursor > 0) {
      e.preventDefault();
      setCursor(cursor - 1);
    }
    if (keyCode === 40 && cursor < filterData.length - 1) {
      e.preventDefault();

      setCursor(cursor + 1);
    }
    if (keyCode === 13 && cursor > -1 && cursor <= filterData.length - 1) {
      e.preventDefault();

      const item = filterData[cursor];

      handleClick(item);
      setCursor(0);
    }
  };

  return (
    <Dropdown
      ref={dropdownRef}
      className="mentions-wrap"
      show={filterData.length > 0}
      onKeyDown={handleKeyDown}>
      {children}
      <Dropdown.Menu
        className={filterData.length > 0 ? 'visible' : 'invisible'}
        ref={menuRef}>
        {filterData
          .filter((_, index) => index < MAX_RECODE)
          .map((item, index) => {
            return (
              <Dropdown.Item
                className={`${cursor === index ? 'bg-gray-200' : ''}`}
                key={item.displayName}
                onClick={() => handleClick(item)}>
                <span className="link-dark me-1">{item.displayName}</span>
                <small className="link-secondary">@{item.userName}</small>
              </Dropdown.Item>
            );
          })}
      </Dropdown.Menu>
    </Dropdown>
  );
};

export default Mentions;
