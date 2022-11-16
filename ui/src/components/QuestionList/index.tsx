import { FC } from 'react';
import { ListGroup } from 'react-bootstrap';
import { NavLink, useParams, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import {
  Icon,
  Tag,
  Pagination,
  FormatTime,
  Empty,
  BaseUserCard,
  QueryGroup,
} from '@/components';
import { useQuestionList } from '@/services';

const QuestionOrderKeys: Type.QuestionOrderBy[] = [
  'newest',
  'active',
  'frequent',
  'score',
  'unanswered',
];

interface Props {
  source: 'questions' | 'tag';
}

const QuestionLastUpdate = ({ q }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'question' });
  if (q.update_time > q.edit_time) {
    // question answered
    return (
      <div className="d-flex">
        <BaseUserCard
          data={q.last_answered_user_info}
          showAvatar={false}
          className="me-1"
        />
        •
        <FormatTime
          time={q.update_time}
          className="text-secondary ms-1"
          preFix={t('answered')}
        />
      </div>
    );
  }

  if (q.edit_time > q.update_time) {
    // question modified
    return (
      <div className="d-flex">
        <BaseUserCard
          data={q.update_user_info}
          showAvatar={false}
          className="me-1"
        />
        •
        <FormatTime
          time={q.edit_time}
          className="text-secondary ms-1"
          preFix={t('modified')}
        />
      </div>
    );
  }

  // default: asked
  return (
    <div className="d-flex">
      <BaseUserCard data={q.user_info} showAvatar={false} className="me-1" />
      •
      <FormatTime
        time={q.create_time}
        preFix={t('asked')}
        className="text-secondary ms-1"
      />
    </div>
  );
};

const QuestionList: FC<Props> = ({ source }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'question' });
  const { tagName = '' } = useParams();
  const [urlSearchParams] = useSearchParams();
  const curOrder = urlSearchParams.get('order') || QuestionOrderKeys[0];
  const curPage = Number(urlSearchParams.get('page')) || 1;
  const pageSize = 20;
  const reqParams: Type.QueryQuestionsReq = {
    page_size: pageSize,
    page: curPage,
    order: curOrder as Type.QuestionOrderBy,
    tag: tagName,
  };

  if (source === 'questions') {
    delete reqParams.tag;
  }
  const { data: listData, isLoading } = useQuestionList(reqParams);
  const count = listData?.count || 0;

  return (
    <div>
      <div className="mb-3 d-flex flex-wrap justify-content-between">
        <h5 className="fs-5 text-nowrap mb-3 mb-md-0">
          {source === 'questions'
            ? t('all_questions')
            : t('x_questions', { count })}
        </h5>
        <QueryGroup
          data={QuestionOrderKeys}
          currentSort={curOrder}
          pathname={source === 'questions' ? '/questions' : ''}
          i18nKeyPrefix="question"
        />
      </div>
      <ListGroup variant="flush" className="border-top border-bottom-0">
        {listData?.list?.map((li) => {
          return (
            <ListGroup.Item
              key={li.id}
              className="border-bottom pt-3 pb-2 px-0">
              <h5 className="text-wrap text-break">
                <NavLink to={`/questions/${li.id}`} className="link-dark">
                  {li.title}
                  {li.status === 2 ? ` [${t('closed')}]` : ''}
                </NavLink>
              </h5>
              <div className="d-flex flex-column flex-md-row align-items-md-center fs-14 text-secondary">
                <QuestionLastUpdate q={li} />
                <div className="ms-0 ms-md-3 mt-2 mt-md-0">
                  <span>
                    <Icon name="hand-thumbs-up-fill" />
                    <em className="fst-normal ms-1">{li.vote_count}</em>
                  </span>
                  <span
                    className={`ms-3 ${
                      li.accepted_answer_id >= 1 ? 'text-success' : ''
                    }`}>
                    <Icon
                      name={
                        li.accepted_answer_id >= 1
                          ? 'check-circle-fill'
                          : 'chat-square-text-fill'
                      }
                    />
                    <em className="fst-normal ms-1">{li.answer_count}</em>
                  </span>
                  <span className="summary-stat ms-3">
                    <Icon name="eye-fill" />
                    <em className="fst-normal ms-1">{li.view_count}</em>
                  </span>
                </div>
              </div>
              <div className="question-tags mx-n1 mt-2">
                {Array.isArray(li.tags)
                  ? li.tags.map((tag) => {
                      return (
                        <Tag key={tag.slug_name} className="m-1" data={tag} />
                      );
                    })
                  : null}
              </div>
            </ListGroup.Item>
          );
        })}
      </ListGroup>
      {count <= 0 && !isLoading && <Empty />}
      <div className="mt-4 mb-2 d-flex justify-content-center">
        <Pagination
          currentPage={curPage}
          totalSize={count}
          pageSize={pageSize}
          pathname="/questions"
        />
      </div>
    </div>
  );
};
export default QuestionList;
