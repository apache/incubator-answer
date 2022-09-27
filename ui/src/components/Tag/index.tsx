import React, { memo, FC } from 'react';

import classNames from 'classnames';

interface IProps {
  className?: string;
  children?: React.ReactNode;
  href: string;
}

const Index: FC<IProps> = ({ className = '', children, href }) => {
  href = href.toLowerCase();
  return (
    <a href={href} className={classNames('badge-tag rounded-1', className)}>
      {children}
    </a>
  );
};

export default memo(Index);
