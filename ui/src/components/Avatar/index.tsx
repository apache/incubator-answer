import { memo, FC } from 'react';

import classNames from 'classnames';

import DefaultAvatar from '@/assets/images/default-avatar.svg';

interface IProps {
  /** avatar url */
  avatar: string | { type: string; gravatar: string; custom: string };
  size: string;
  searchStr?: string;
  className?: string;
}

const Index: FC<IProps> = ({ avatar, size, className, searchStr = '' }) => {
  let url = '';
  if (typeof avatar === 'string') {
    if (avatar.length > 1) {
      url = `${avatar}?${searchStr}`;
    }
  } else if (avatar?.type !== 'default') {
    url = `${avatar[avatar.type]}?${searchStr}`;
  }

  return (
    <img
      src={url || DefaultAvatar}
      width={size}
      height={size}
      className={classNames('rounded', className)}
      alt=""
    />
  );
};

export default memo(Index);
