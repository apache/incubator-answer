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
    <ListGroup className="rounded-0">
      {data.map((item) => {
        return (
          <ListGroupItem
            className="py-3 px-0 bg-transparent border-start-0 border-end-0"
            key={item.comment_id}>
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
            <div
              className="small mb-2 last-p text-break text-truncate-2"
              dangerouslySetInnerHTML={{
                __html: item.content,
              }}
            />

            <FormatTime
              time={item.created_at}
              className="small text-secondary"
            />
          </ListGroupItem>
        );
      })}
    </ListGroup>
  );
};

export default memo(Index);
