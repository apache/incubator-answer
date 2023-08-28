import { memo, FC } from 'react';
import { Link } from 'react-router-dom';

import { Avatar } from '@/components';
import { formatCount } from '@/utils';

interface Props {
  data: any;
  showAvatar?: boolean;
  avatarSize?: string;
  showReputation?: boolean;
  avatarSearchStr?: string;
  className?: string;
  avatarClass?: string;
  nameMaxWidth?: string;
}

const Index: FC<Props> = ({
  data,
  showAvatar = true,
  avatarClass = '',
  avatarSize = '20px',
  className = 'small',
  avatarSearchStr = 's=48',
  showReputation = true,
  nameMaxWidth = '300px',
}) => {
  return (
    <div className={`d-flex align-items-center  text-secondary ${className}`}>
      {data?.status !== 'deleted' ? (
        <Link
          to={`/users/${data?.username}`}
          className="d-flex align-items-center">
          {showAvatar && (
            <Avatar
              avatar={data?.avatar}
              size={avatarSize}
              className={`me-1 ${avatarClass}`}
              searchStr={avatarSearchStr}
              alt={data?.display_name}
            />
          )}
          <span
            className="me-1 name-ellipsis"
            style={{ maxWidth: nameMaxWidth }}>
            {data?.display_name}
          </span>
        </Link>
      ) : (
        <>
          {showAvatar && (
            <Avatar
              avatar={data?.avatar}
              size={avatarSize}
              className={`me-1 ${avatarClass}`}
              searchStr={avatarSearchStr}
              alt={data?.display_name}
            />
          )}
          <span className="me-1 name-ellipsis">{data?.display_name}</span>
        </>
      )}

      {showReputation && (
        <span className="fw-bold" title="Reputation">
          {formatCount(data?.rank)}
        </span>
      )}
    </div>
  );
};

export default memo(Index);
