import { memo, FC } from 'react';
import { Link } from 'react-router-dom';

import { Avatar, FormatTime } from '@answer/components';

interface Props {
  data: any;
  time: number;
  preFix: string;
}

const Index: FC<Props> = ({ data, time, preFix }) => {
  return (
    <div className="d-flex">
      <Link to={`/users/${data?.username}`}>
        <Avatar avatar={data?.avatar} size="40px" className="me-2" />
      </Link>
      <div className="fs-14 text-secondary">
        <div>
          <Link to={`/users/${data?.username}`} className="me-1 text-break">
            {data?.display_name}
          </Link>
          <span className="fw-bold">{data?.rank}</span>
        </div>
        {time && <FormatTime time={time} preFix={preFix} />}
      </div>
    </div>
  );
};

export default memo(Index);
