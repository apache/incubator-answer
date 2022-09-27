import { FC } from 'react';

import classNames from 'classnames';

interface IProps {
  type?: 'br' | 'bi';
  /** icon name */
  name: string;
  className?: string;
  size?: string;
  onClick?: () => void;
}
const Icon: FC<IProps> = ({ type = 'br', name, className, size, onClick }) => {
  return (
    <i
      className={classNames(type, `bi-${name}`, className)}
      style={{ ...(size && { fontSize: size }) }}
      onClick={onClick}
      onKeyDown={onClick}
    />
  );
};

export default Icon;
