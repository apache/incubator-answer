import { FC, memo } from 'react';
import { ListGroup, ListGroupItem } from 'react-bootstrap';

import { FormatTime } from '@/components';

interface Props {
  visible: boolean;
  data: any[];
}

const Index: FC<Props> = ({ visible, data }) => {
  if (!visible || !data?.length) {
    return null;
  }
  return (
    <ListGroup variant="flush">
      {data.map((item) => {
        return (
          <ListGroupItem className="d-flex py-3 px-0" key={item.object_id}>
            <div
              className="me-3 text-end text-secondary"
              style={{ width: '80px' }}>
              {item.vote_type}
            </div>
            <div>
              <a
                className="text-break"
                href={`/questions/${
                  item.object_type === 'question'
                    ? item.question_id
                    : `${item.question_id}/${item.answer_id}`
                }`}>
                {item.title}
              </a>
              <div className="d-flex align-items-center fs-14 text-secondary">
                <span>{item.object_type}</span>

                <span className="split-dot" />
                <FormatTime time={item.created_at} className="me-4" />
              </div>
            </div>
          </ListGroupItem>
        );
      })}
    </ListGroup>
  );
};

export default memo(Index);
