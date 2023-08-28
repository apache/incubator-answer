import { FC, useContext, useEffect } from 'react';
import { Dropdown, Button } from 'react-bootstrap';

import { EditorContext } from './EditorContext';

interface IProps {
  keyMap?: string[];
  onClick?: () => void;
  tip?: string;
  className?: string;
  as?: any;
  children?;
  label?: string;
  disable?: boolean;
  isShow?: boolean;
  onBlur?: () => void;
}
const ToolItem: FC<IProps> = (props) => {
  const context = useContext(EditorContext);

  const { editor } = context;
  const {
    label,
    tip,
    disable = false,
    isShow,
    keyMap,
    onClick,
    className,
    as,
    children,
    onBlur,
  } = props;

  useEffect(() => {
    if (!keyMap) {
      return;
    }

    keyMap.forEach((key) => {
      editor.addKeyMap({
        [key]: () => {
          if (typeof onClick === 'function') {
            onClick();
          }
        },
      });
    });
  }, []);

  const btnRender = () => (
    <Button
      variant="link"
      title={tip}
      className={`p-0 b-0 btn-no-border toolbar icon-${label} ${
        disable ? 'disabled' : ''
      } `}
      disabled={disable}
      tabIndex={-1}
      onClick={(e) => {
        e.preventDefault();
        if (typeof onClick === 'function') {
          onClick();
        }
      }}
      onBlur={(e) => {
        e.preventDefault();
        if (typeof onBlur === 'function') {
          onBlur();
        }
      }}
    />
  );

  if (!context) {
    return null;
  }
  return (
    <div className={`toolbar-item-wrap ${className || ''}`}>
      {as === 'dropdown' ? (
        <Dropdown className="h-100 w-100" show={isShow}>
          <Dropdown.Toggle as="div" className="h-100">
            {btnRender()}
          </Dropdown.Toggle>
          {children}
        </Dropdown>
      ) : (
        <>
          {btnRender()}
          {children}
        </>
      )}
    </div>
  );
};

export default ToolItem;
