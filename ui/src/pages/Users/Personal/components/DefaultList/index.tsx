import { FC, memo } from 'react';
import { ListGroup, ListGroupItem } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { Icon, FormatTime, Tag, BaseUserCard } from '@answer/components';

interface Props {
  visible: boolean;
  tabName: string;
  data: any[];
}

const Index: FC<Props> = ({ visible, tabName, data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });
  if (!visible) {
    return null;
  }
  return (
    <ListGroup variant="flush">
      {data.map((item) => {
        return (
          <ListGroupItem className="py-3 px-0" key={item.question_id}>
            <h6 className="mb-2">
              <a
                className="text-break"
                href={`/questions/${
                  tabName === 'questions' ? item.question_id : item.id
                }`}>
                {item.title}
                {tabName === 'questions' && item.status === 'closed'
                  ? ` [${t('closed', { keyPrefix: 'question' })}]`
                  : null}
              </a>
            </h6>
            <div className="d-flex align-items-center fs-14 text-secondary mb-2">
              {tabName === 'bookmarks' && (
                <>
                  <BaseUserCard data={item.user_info} showAvatar={false} />
                  <span className="split-dot" />
                </>
              )}
              <FormatTime
                time={item.create_time}
                className="me-3"
                preFix={t('asked')}
              />

              <div className="d-flex align-items-center me-3">
                <Icon name="hand-thumbs-up-fill me-1" />
                <span>{item.vote_count}</span>
              </div>

              {tabName !== 'answers' && (
                <div
                  className={`d-flex align-items-center me-3 ${
                    Number(item.accepted_answer_id) > 0 ? 'text-success' : ''
                  }`}>
                  {Number(item.accepted_answer_id) > 0 ? (
                    <Icon name="check-circle-fill me-1" />
                  ) : (
                    <Icon name="chat-square-text-fill me-1" />
                  )}
                  <span>{item.answer_count}</span>
                </div>
              )}

              <div className="d-flex align-items-center me-3">
                <Icon name="eye-fill me-1" />
                <span>{item.view_count}</span>
              </div>
            </div>
            <div>
              {item.tags?.map((tag) => {
                return (
                  <Tag
                    href={`/t/${tag.main_tag_slug_name || tag.slug_name}`}
                    className="me-1"
                    key={tag.slug_name}>
                    {tag.slug_name}
                  </Tag>
                );
              })}
            </div>
          </ListGroupItem>
        );
      })}
    </ListGroup>
  );
};

export default memo(Index);
