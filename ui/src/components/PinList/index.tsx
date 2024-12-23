import { FC } from 'react';
import { ListGroup, Stack, Card } from 'react-bootstrap';
import { NavLink } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { Counts } from '@/components';
import { pathFactory } from '@/router/pathFactory';

interface IProps {
  data: any[];
}

const PinList: FC<IProps> = ({ data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'question' });
  if (!data?.length) return null;

  return (
    <ListGroup.Item className="py-3 px-0 border-start-0 border-end-0">
      <Stack
        direction="horizontal"
        gap={3}
        className="overflow-x-auto align-items-stretch">
        {data.map((item) => {
          return (
            <Card
              key={item.id}
              style={{
                minWidth: '238px',
                width: `${100 / data.length}%`,
              }}>
              <Card.Body>
                <h6 className="text-wrap text-break">
                  <NavLink
                    to={pathFactory.questionLanding(item.id, item.url_title)}
                    className="link-dark text-truncate-2">
                    {item.title}
                    {item.status === 2 ? ` [${t('closed')}]` : ''}
                  </NavLink>
                </h6>

                <Counts
                  data={{
                    votes: item.vote_count,
                    answers: item.answer_count,
                    views: item.view_count,
                  }}
                  isAccepted={item.accepted_answer_id >= 1}
                  showViews={false}
                  className="mt-2 mt-md-0 small text-secondary"
                />
              </Card.Body>
            </Card>
          );
        })}
      </Stack>
    </ListGroup.Item>
  );
};

export default PinList;
