import { memo, FC } from 'react';
import { Link } from 'react-router-dom';

import classnames from 'classnames';

import { Avatar, FormatTime } from '@/components';
import { formatCount } from '@/utils';

interface Props {
  data: any;
  time: number;
  preFix: string;
  isLogged: boolean;
  timelinePath: string;
  className?: string;
}

const Index: FC<Props> = ({
  data,
  time,
  preFix,
  isLogged,
  timelinePath,
  className = '',
}) => {
  return (
    <div className={classnames('d-flex', className)}>
      {data?.status !== 'deleted' ? (
        <Link to={`/users/${data?.username}`}>
          <Avatar
            avatar={data?.avatar}
            size="40px"
            className="me-2 d-none d-md-block"
            searchStr="s=96"
          />

          <Avatar
            avatar={data?.avatar}
            size="24px"
            className="me-2 d-block d-md-none"
            searchStr="s=48"
          />
        </Link>
      ) : (
        <>
          <Avatar
            avatar={data?.avatar}
            size="40px"
            className="me-2 d-none d-md-block"
            searchStr="s=96"
          />

          <Avatar
            avatar={data?.avatar}
            size="24px"
            className="me-2 d-block d-md-none"
            searchStr="s=48"
          />
        </>
      )}
      <div className="small text-secondary d-flex flex-row flex-md-column align-items-center align-items-md-start">
        <div className="me-1 me-md-0">
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
        {time &&
          (isLogged ? (
            <Link to={timelinePath}>
              <FormatTime
                time={time}
                preFix={preFix}
                className="link-secondary"
              />
            </Link>
          ) : (
            <FormatTime time={time} preFix={preFix} />
          ))}
      </div>
    </div>
  );
};

export default memo(Index);
