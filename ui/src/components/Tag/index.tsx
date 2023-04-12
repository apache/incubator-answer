import React, { memo, FC } from 'react';
import { Link } from 'react-router-dom';

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
  href ||= pathFactory.tagLanding(data?.slug_name);

  return (
    <Link
      to={href}
      className={classNames(
        'badge-tag rounded-1',
        data.reserved && 'badge-tag-reserved',
        data.recommend && 'badge-tag-required',
        className,
      )}>
      <span className={textClassName}>
        {data.display_name || data.slug_name}
      </span>
    </Link>
  );
};

export default memo(Index);
