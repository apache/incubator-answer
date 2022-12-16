import { FC, memo } from 'react';
import { ListGroup, ListGroupItem } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { pathFactory } from '@/router/pathFactory';
import { Icon } from '@/components';

interface Props {
  data: any[];
  type: 'answer' | 'question';
}
const Index: FC<Props> = ({ data, type }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });
  return (
    <ListGroup variant="flush" className="mb-4">
      {data?.map((item) => {
        return (
          <ListGroupItem
            className="p-0 border-0 mb-2"
            key={type === 'answer' ? item.answer_id : item.question_id}>
            <a
              href={
                type === 'answer'
                  ? pathFactory.answerLanding({
                      questionId: item.question_id,
                      slugTitle: item.question_info?.url_title,
                      answerId: item.answer_id,
                    })
                  : pathFactory.questionLanding(
                      item.question_id,
                      item.url_title,
                    )
              }>
              {type === 'answer' ? item.question_info.title : item.title}
            </a>
            <div className="d-inline-block text-secondary ms-3 fs-14">
              <Icon name="hand-thumbs-up-fill" />
              <span> {item.vote_count}</span>
            </div>
            {type === 'question' && (
              <div
                className={`d-inline-block text-secondary ms-3 fs-14 ${
                  Number(item.accepted_answer_id) > 0 ? 'text-success' : ''
                }`}>
                {Number(item.accepted_answer_id) > 0 ? (
                  <Icon name="check-circle-fill" />
                ) : (
                  <Icon name="chat-square-text-fill" />
                )}

                <span> {item.answer_count}</span>
              </div>
            )}

            {type === 'answer' && item.adopted === 2 && (
              <div className="d-inline-block text-success ms-3 fs-14">
                <Icon name="check-circle-fill" />
                <span> {t('accepted')}</span>
              </div>
            )}
          </ListGroupItem>
        );
      })}
    </ListGroup>
  );
};

export default memo(Index);
