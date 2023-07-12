import { memo, FC } from 'react';

import classNames from 'classnames';

import DefaultAvatar from '@/assets/images/default-avatar.svg';

interface IProps {
  /** avatar url */
  avatar: string | { type: string; gravatar: string; custom: string };
  /** size 48 96 128 256 */
  size: string;
  searchStr?: string;
  className?: string;
}

const Index: FC<IProps> = ({ avatar, size, className, searchStr = '' }) => {
  let url = '';
  if (typeof avatar === 'string') {
    if (avatar.length > 1) {
      url = `${avatar}?${searchStr}${
        avatar?.includes('gravatar') ? '&d=identicon' : ''
      }`;
    }
  } else if (avatar?.type === 'gravatar' && avatar.gravatar) {
    url = `${avatar.gravatar}?${searchStr}&d=identicon`;
  } else if (avatar?.type === 'custom' && avatar.custom) {
    url = `${avatar.custom}?${searchStr}`;
  }

  const roundedCls =
    className && className.indexOf('rounded') !== -1 ? '' : 'rounded';

  return (
    <img
      src={url || DefaultAvatar}
      width={size}
      height={size}
      className={classNames(roundedCls, className)}
      alt=""
    />
  );
};

export default memo(Index);
