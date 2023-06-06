import { memo, FC } from 'react';
import { Card, ListGroup } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { Icon } from '@/components';
import { useSimilarQuestion } from '@/services';
import { loggedUserInfoStore } from '@/stores';
import { pathFactory } from '@/router/pathFactory';

interface Props {
  id: string;
}
const Index: FC<Props> = ({ id }) => {
  const { user } = loggedUserInfoStore();
  const { t } = useTranslation('translation', {
    keyPrefix: 'related_question',
  });

  const { data, isLoading } = useSimilarQuestion({
    question_id: id,
    page_size: 5,
  });

  if (isLoading) {
    return null;
  }

  return (
    <Card>
      <Card.Header>{t('title')}</Card.Header>
      <ListGroup variant="flush">
        {data?.list?.map((item) => {
          return (
            <ListGroup.Item
              action
              key={item.id}
              as={Link}
              to={pathFactory.questionLanding(item.id, item.url_title)}>
              <div className="link-dark">{item.title}</div>
              {item.answer_count > 0 && (
                <div
                  className={`mt-1 small me-2 ${
                    item.accepted_answer_id > 0
                      ? 'link-success'
                      : 'link-secondary'
                  }`}>
                  <Icon
                    name={
                      item.accepted_answer_id > 0
                        ? 'check-circle-fill'
                        : 'chat-square-text-fill'
                    }
                    className="me-1"
                  />

                  <span>
                    {item.answer_count} {t('answers')}
                  </span>
                </div>
              )}
            </ListGroup.Item>
          );
        })}
      </ListGroup>
      {user?.username ? (
        <Card.Footer className="bg-white">
          <Link to="/questions/ask">{t('btn')}</Link>
        </Card.Footer>
      ) : null}
    </Card>
  );
};

export default memo(Index);
