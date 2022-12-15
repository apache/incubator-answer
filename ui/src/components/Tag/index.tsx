import React, { memo, FC } from 'react';

import classNames from 'classnames';

import { Tag } from '@/common/interface';
import { pathFactory } from '@/router/pathFactory';

interface IProps {
  data: Tag;
  href?: string;
  className?: string;
  textClassName?: string;
}

const Index: FC<IProps> = ({
  data,
  href,
  className = '',
  textClassName = '',
}) => {
  href ||= pathFactory.tagLanding(data);

  return (
    <a
      href={href}
      className={classNames(
        'badge-tag rounded-1',
        data.reserved && 'badge-tag-reserved',
        data.recommend && 'badge-tag-required',
        className,
      )}>
      <span className={textClassName}>{data.slug_name}</span>
    </a>
  );
};

export default memo(Index);
