import React, { memo, FC } from 'react';

import classNames from 'classnames';

import { Tag } from '@/common/interface';

interface IProps {
  data: Tag;
  href?: string;
  className?: string;
}

const Index: FC<IProps> = ({ className = '', href, data }) => {
  href =
    href || `/tags/${data.main_tag_slug_name || data.slug_name}`.toLowerCase();
  return (
    <a
      href={href}
      className={classNames(
        'badge-tag rounded-1',
        data.reserved && 'badge-tag-reserved',
        data.recommend && 'badge-tag-required',
        className,
      )}>
      {data.slug_name}
    </a>
  );
};

export default memo(Index);
