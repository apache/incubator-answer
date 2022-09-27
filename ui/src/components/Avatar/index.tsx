import { memo, FC } from 'react';

import classNames from 'classnames';

import DefaultAvatar from '@/assets/images/default-avatar.svg';

interface IProps {
  /** avatar url */
  avatar: string;
  size: string;
  className?: string;
}

const Index: FC<IProps> = ({ avatar, size, className }) => {
  return (
    <img
      src={avatar || DefaultAvatar}
      width={size}
      height={size}
      className={classNames('rounded', className)}
      alt=""
    />
  );
};

export default memo(Index);
