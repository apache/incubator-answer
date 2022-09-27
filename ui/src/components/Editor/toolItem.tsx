import { FC, useContext, useEffect } from 'react';
import { Dropdown, OverlayTrigger, Tooltip, Button } from 'react-bootstrap';

import { EditorContext } from './EditorContext';

interface IProps {
  keyMap?: string[];
  click?: () => void;
  tip?: string;
  className?: string;
  as?: any;
  children?;
  label?: string;
  disable?: boolean;
  isShow?: boolean;
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
    click,
    className,
    as,
    children,
  } = props;

  useEffect(() => {
    if (!keyMap) {
      return;
    }

    keyMap.forEach((key) => {
      editor.addKeyMap({
        [key]: () => {
          if (typeof click === 'function') {
            click();
          }
        },
      });
    });
  }, []);

  const btnRender = () => (
    <OverlayTrigger placement="bottom" overlay={<Tooltip>{tip}</Tooltip>}>
      <Button
        variant="link"
        className={`p-0 b-0 btn-no-border toolbar icon-${label} ${
          disable ? 'disabled' : ''
        } `}
        disabled={disable}
        onClick={(e) => {
          e.preventDefault();
          if (typeof click === 'function') {
            click();
          }
        }}
        onBlur={(e) => {
          e.preventDefault();
          if (typeof click === 'function') {
            click();
          }
        }}
      />
    </OverlayTrigger>
  );

  if (!context) {
    return null;
  }
  return (
    <div className={`toolbar-item-wrap ${className || ''}`}>
      {as === 'dropdown' ? (
        <Dropdown className="h-100 w-100" show={isShow}>
          <Dropdown.Toggle as="div" id="tootlbars-toogle" className="h-100">
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
