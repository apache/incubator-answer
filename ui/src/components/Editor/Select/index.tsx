import { FC, useEffect, useState } from 'react';
import { Dropdown, FormControl } from 'react-bootstrap';

interface IProps {
  options;
  value?;
  onChange?;
  placeholder?;
  onSelect?;
}
const Select: FC<IProps> = ({
  options = [],
  value = '',
  onChange,
  placeholder = '',
  onSelect,
}) => {
  const [isFocus, setFocusState] = useState(false);
  const [cursor, setCursor] = useState(0);

  useEffect(() => {
    setCursor(0);
  }, [value]);
  const handleKeyDown = (e) => {
    const { keyCode } = e;

    if (keyCode === 38 && cursor > 0) {
      e.preventDefault();
      setCursor(cursor - 1);
    }
    if (keyCode === 40 && cursor < options.length - 1) {
      e.preventDefault();

      setCursor(cursor + 1);
    }
    if (keyCode === 13 && cursor > -1 && cursor <= options.length - 1) {
      const lang = options.filter((opt) =>
        value ? opt.indexOf(value) === 0 : true,
      )[cursor];

      setFocusState(false);
      onSelect(lang);
    }
  };

  const result = options.filter((opt) =>
    value ? opt.indexOf(value) === 0 : true,
  );

  return (
    <div className="position-relative" onKeyDown={handleKeyDown}>
      <FormControl
        type="search"
        value={value}
        placeholder={placeholder}
        onChange={(e) => {
          setFocusState(true);
          if (onChange instanceof Function) {
            onChange(e);
          }
        }}
      />
      {result.length > 0 && (
        <Dropdown.Menu
          show={value && isFocus}
          className="border py-2 rounded w-100"
          style={{ overflowY: 'auto', maxHeight: '250px' }}>
          {result.map((opt, index) => {
            return (
              <Dropdown.Item
                key={opt}
                className={`${cursor === index ? 'active' : ''}`}
                onClick={(e) => {
                  e.preventDefault();
                  setFocusState(false);
                  onSelect(opt);
                }}>
                {opt}
              </Dropdown.Item>
            );
          })}
        </Dropdown.Menu>
      )}
    </div>
  );
};

export default Select;
