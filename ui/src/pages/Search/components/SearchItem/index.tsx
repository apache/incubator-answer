import { memo, FC } from 'react';
import { ListGroupItem, Badge } from 'react-bootstrap';

import { Icon, Tag, FormatTime } from '@answer/components';
import type { SearchResItem } from '@answer/services/types';
import { formatCount } from '@answer/utils';

interface Props {
  data: SearchResItem;
}
const Index: FC<Props> = ({ data }) => {
  if (!data?.object_type) {
    return null;
  }
  return (
    <ListGroupItem className="py-3 px-0">
      <div className="mb-2 clearfix">
        <Badge
          bg="dark"
          className="me-2 float-start"
          style={{ marginTop: '2px' }}>
          {data.object_type === 'question' ? 'Q' : 'A'}
        </Badge>
        <a
          className="h5 mb-0 text-body text-break"
          href={`/questions/${data.object.id}`}>
          {data.object.title}
        </a>
      </div>
      <div className="d-flex flex-wrap align-items-center fs-14 text-secondary mb-2">
        <a href={`/users/${data.object?.user_info?.username}`}>
          {data.object?.user_info?.display_name}
        </a>
        {data.object?.user_info?.rank > 0 && (
          <span className="fw-bold ms-1">
            {formatCount(data.object.user_info.rank)}
          </span>
        )}
        <span className="split-dot" />
        <FormatTime
          time={data.object?.created_at}
          className="me-3"
          preFix={data.object_type === 'question' ? 'asked' : 'answered'}
        />
        <div className="d-flex align-items-center my-2 my-sm-0">
          <div className="d-flex align-items-center me-3">
            <Icon name="hand-thumbs-up-fill me-1" />
            <span> {data.object?.vote_count}</span>
          </div>
          <div
            className={`d-flex align-items-center ${
              data.object?.accepted ? 'text-success' : ''
            }`}>
            {data.object?.accepted ? (
              <Icon name="check-circle-fill me-1" />
            ) : (
              <Icon name="chat-square-text-fill me-1" />
            )}
            <span>{data.object?.answer_count}</span>
          </div>
          {/* <div className="d-flex align-items-center">
            <Icon name="eye-fill fs-6 me-1" />
            <span> 0</span>
          </div> */}
        </div>
      </div>

      {data.object?.excerpt && (
        <p
          className="fs-14 text-truncate-2 mb-2 last-p text-break"
          dangerouslySetInnerHTML={{ __html: data.object.excerpt }}
        />
      )}

      {data.object?.tags?.map((item) => {
        return (
          <Tag href={`/tags/${item.slug_name}`} className="me-1">
            {item.slug_name}
          </Tag>
        );
      })}
    </ListGroupItem>
  );
};

export default memo(Index);
