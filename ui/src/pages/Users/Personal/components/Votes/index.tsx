import { FC, memo } from 'react';
import { ListGroup, ListGroupItem } from 'react-bootstrap';

import { pathFactory } from '@/router/pathFactory';
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
    <ListGroup className="rounded-0">
      {data.map((item) => {
        return (
          <ListGroupItem
            className="d-flex py-3 px-0 bg-transparent border-start-0 border-end-0"
            key={item.object_id}>
            <div
              className="me-3 text-end text-secondary flex-shrink-0"
              style={{ width: '80px' }}>
              {item.vote_type}
            </div>
            <div>
              <a
                className="text-break"
                href={
                  item.object_type === 'question'
                    ? pathFactory.questionLanding(
                        item.question_id,
                        item.url_title,
                      )
                    : pathFactory.answerLanding({
                        questionId: item.question_id,
                        slugTitle: item.url_title,
                        answerId: item.answer_id,
                      })
                }>
                {item.title}
              </a>
              <div className="d-flex align-items-center small text-secondary">
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
