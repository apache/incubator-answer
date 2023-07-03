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
  alt: string;
}

const Index: FC<IProps> = ({
  avatar,
  size,
  className,
  searchStr = '',
  alt,
}) => {
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
    <>
      {/* eslint-disable jsx-a11y/no-noninteractive-element-to-interactive-role,jsx-a11y/control-has-associated-label */}
      <img
        role="button"
        src={url || DefaultAvatar}
        width={size}
        height={size}
        className={classNames(roundedCls, className)}
        alt={alt}
      />
    </>
  );
};

export default memo(Index);
