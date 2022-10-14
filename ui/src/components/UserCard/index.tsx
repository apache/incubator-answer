import { memo, FC } from 'react';
import { Link } from 'react-router-dom';

import { Avatar, FormatTime } from '@answer/components';

import { formatCount } from '@/utils';

interface Props {
  data: any;
  time: number;
  preFix: string;
}

const Index: FC<Props> = ({ data, time, preFix }) => {
  return (
    <div className="d-flex">
      {data?.status !== 'deleted' ? (
        <Link to={`/users/${data?.username}`}>
          <Avatar avatar={data?.avatar} size="40px" className="me-2" />
        </Link>
      ) : (
        <Avatar avatar={data?.avatar} size="40px" className="me-2" />
      )}
      <div className="fs-14 text-secondary">
        <div>
          {data?.status !== 'deleted' ? (
            <Link to={`/users/${data?.username}`} className="me-1 text-break">
              {data?.display_name}
            </Link>
          ) : (
            <span className="me-1 text-break">{data?.display_name}</span>
          )}
          <span className="fw-bold" title="Reputation">
            {formatCount(data?.rank)}
          </span>
        </div>
        {time && <FormatTime time={time} preFix={preFix} />}
      </div>
    </div>
  );
};

export default memo(Index);
