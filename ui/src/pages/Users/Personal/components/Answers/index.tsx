import { FC, memo } from 'react';
import { ListGroup, ListGroupItem } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { FormatTime, Tag, Counts } from '@/components';
import { pathFactory } from '@/router/pathFactory';

interface Props {
  visible: boolean;
  data: any[];
}
const Index: FC<Props> = ({ visible, data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });
  if (!visible || !data?.length) {
    return null;
  }
  return (
    <ListGroup className="rounded-0">
      {data.map((item) => {
        return (
          <ListGroupItem
            className="py-3 px-0 bg-transparent border-start-0 border-end-0"
            key={item.answer_id}>
            <h6 className="mb-2">
              <a
                href={pathFactory.answerLanding({
                  questionId: item.question_id,
                  slugTitle: item.question_info?.url_title,
                  answerId: item.answer_id,
                })}
                className="text-break">
                {item.question_info?.title}
              </a>
            </h6>
            <div className="d-flex align-items-center small text-secondary mb-2">
              <FormatTime
                time={item.create_time}
                className="me-3"
                preFix={t('answered')}
              />

              <Counts
                data={{ votes: item?.vote_count, views: 0, answers: 0 }}
                showAnswers={false}
                showViews={false}
                showAccepted={item.accepted === 2}
              />
            </div>
            <div>
              {item.question_info?.tags?.map((tag) => {
                return <Tag key={tag.slug_name} className="me-1" data={tag} />;
              })}
            </div>
          </ListGroupItem>
        );
      })}
    </ListGroup>
  );
};

export default memo(Index);
