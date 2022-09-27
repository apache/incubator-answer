import { memo, FC } from 'react';
import { Link } from 'react-router-dom';

import { Avatar } from '@answer/components';

interface Props {
  data: any;
  avatarSize?: string;
  className?: string;
}

const Index: FC<Props> = ({
  data,
  avatarSize = '20px',
  className = 'fs-14',
}) => {
  return (
    <div className={`text-secondary ${className}`}>
      <Link to={`/users/${data?.username}`}>
        <Avatar avatar={data?.avatar} size={avatarSize} className="me-1" />
      </Link>
      <Link to={`/users/${data?.username}`} className="me-1 text-break">
        {data?.display_name}
      </Link>
      <span className="fw-bold">{data?.rank}</span>
    </div>
  );
};

export default memo(Index);
