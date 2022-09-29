import { memo, FC } from 'react';
import { Link } from 'react-router-dom';

import { Avatar } from '@answer/components';

interface Props {
  data: any;
  showAvatar?: boolean;
  avatarSize?: string;
  className?: string;
}

const Index: FC<Props> = ({
  data,
  showAvatar = true,
  avatarSize = '20px',
  className = 'fs-14',
}) => {
  return (
    <div className={`text-secondary ${className}`}>
      {data.status !== 'deleted' ? (
        <Link to={`/users/${data?.username}`}>
          {showAvatar && (
            <Avatar avatar={data?.avatar} size={avatarSize} className="me-1" />
          )}
          <span className="me-1 text-break">{data?.display_name}</span>
        </Link>
      ) : (
        <>
          {showAvatar && (
            <Avatar avatar={data?.avatar} size={avatarSize} className="me-1" />
          )}
          <span className="me-1 text-break">{data?.display_name}</span>
        </>
      )}

      <span className="fw-bold">{data?.rank}</span>
    </div>
  );
};

export default memo(Index);
