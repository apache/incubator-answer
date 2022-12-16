import { FC, memo } from 'react';
import { ListGroup, ListGroupItem } from 'react-bootstrap';

import { pathFactory } from '@/router/pathFactory';
import { FormatTime } from '@/components';

interface Props {
  visible: boolean;
  data;
}

const Index: FC<Props> = ({ visible, data }) => {
  if (!visible || !data?.length) {
    return null;
  }
  return (
    <ListGroup variant="flush">
      {data.map((item) => {
        return (
          <ListGroupItem className="py-3 px-0" key={item.comment_id}>
            <a
              className="text-break"
              href={
                item.object_type === 'question'
                  ? pathFactory.questionLanding(item.object_id, item.url_title)
                  : pathFactory.answerLanding({
                      questionId: item.question_id,
                      slugTitle: item.url_title,
                      answerId: item.object_id,
                    })
              }>
              {item.title}
            </a>
            <div
              className="fs-14 mb-2 last-p text-break text-truncate-2"
              dangerouslySetInnerHTML={{
                __html: item.content,
              }}
            />

            <FormatTime
              time={item.created_at}
              className="fs-14 text-secondary"
            />
          </ListGroupItem>
        );
      })}
    </ListGroup>
  );
};

export default memo(Index);
