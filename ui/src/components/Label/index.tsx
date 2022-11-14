import React, { memo, FC } from 'react';

import classNames from 'classnames';

import { labelStyle } from '@/utils';

interface IProps {
  className?: string;
  children?: React.ReactNode;
  color: string;
}

const Index: FC<IProps> = ({ className = '', children, color }) => {
  // hover
  const [hover, setHover] = React.useState(false);
  return (
    <span
      className={classNames('badge-label rounded-1', className)}
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
      style={labelStyle(color, hover)}>
      {children}
    </span>
  );
};

export default memo(Index);
