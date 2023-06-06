import { memo, FC } from 'react';
import { ListGroupItem } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { pathFactory } from '@/router/pathFactory';
import { Tag, FormatTime, BaseUserCard, Counts } from '@/components';
import type { SearchResItem } from '@/common/interface';
import { escapeRemove } from '@/utils';

interface Props {
  data: SearchResItem;
}
const Index: FC<Props> = ({ data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'question' });
  if (!data?.object_type) {
    return null;
  }
  let itemUrl = pathFactory.questionLanding(
    data.object.id,
    data.object.url_title,
  );
  if (data.object_type === 'answer' && data.object.question_id) {
    itemUrl = pathFactory.answerLanding({
      questionId: data.object.question_id,
      slugTitle: data.object.url_title,
      answerId: data.object.id,
    });
  }

  return (
    <ListGroupItem className="py-3 px-0 border-start-0 border-end-0 bg-transparent">
      <div className="mb-2 clearfix">
        <span
          className="float-start me-2 badge text-bg-dark"
          style={{ marginTop: '2px' }}>
          {data.object_type === 'question' ? 'Q' : 'A'}
        </span>
        <a className="h5 mb-0 link-dark text-break" href={itemUrl}>
          {data.object.title}
          {data.object.status === 'closed'
            ? ` [${t('closed', { keyPrefix: 'question' })}]`
            : null}
        </a>
      </div>
      <div className="d-flex flex-wrap align-items-center small text-secondary mb-2">
        <BaseUserCard data={data.object?.user_info} showAvatar={false} />

        <span className="split-dot" />
        <FormatTime
          time={data.object?.created_at}
          className="me-3"
          preFix={data.object_type === 'question' ? 'asked' : 'answered'}
        />

        <Counts
          className="my-2 my-sm-0"
          showViews={false}
          isAccepted={data.object?.accepted}
          data={{
            votes: data.object?.vote_count,
            answers: data.object?.answer_count,
            views: 0,
          }}
        />
      </div>

      {data.object?.excerpt && (
        <p className="small text-truncate-2 mb-2 last-p text-break">
          {escapeRemove(data.object.excerpt)}
        </p>
      )}

      {data.object?.tags?.map((item) => {
        return <Tag key={item.slug_name} className="me-1" data={item} />;
      })}
    </ListGroupItem>
  );
};

export default memo(Index);
