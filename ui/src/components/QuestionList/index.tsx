import { FC } from 'react';
import { ListGroup } from 'react-bootstrap';
import { NavLink, useParams, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { pathFactory } from '@/router/pathFactory';
import type * as Type from '@/common/interface';
import {
  Tag,
  Pagination,
  FormatTime,
  Empty,
  BaseUserCard,
  QueryGroup,
  QuestionListLoader,
  Counts,
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
      <ListGroup className="rounded-0">
        {isLoading ? (
          <QuestionListLoader />
        ) : (
          listData?.list?.map((li) => {
            return (
              <ListGroup.Item
                key={li.id}
                className="bg-transparent py-3 px-0 border-start-0 border-end-0">
                <h5 className="text-wrap text-break">
                  <NavLink
                    to={pathFactory.questionLanding(li.id, li.url_title)}
                    className="link-dark">
                    {li.title}
                    {li.status === 2 ? ` [${t('closed')}]` : ''}
                  </NavLink>
                </h5>
                <div className="d-flex flex-column flex-md-row align-items-md-center fs-14 mb-2 text-secondary">
                  <div className="d-flex">
                    <BaseUserCard
                      data={li.operator}
                      showAvatar={false}
                      className="me-1"
                    />
                    •
                    <FormatTime
                      time={li.operated_at}
                      className="text-secondary ms-1"
                      preFix={t(li.operation_type)}
                    />
                  </div>
                  <Counts
                    data={{
                      votes: li.vote_count,
                      answers: li.answer_count,
                      views: li.view_count,
                      top: li.top,
                    }}
                    showTop
                    isAccepted={li.accepted_answer_id >= 1}
                    className="ms-0 ms-md-3 mt-2 mt-md-0"
                  />
                </div>
                <div className="question-tags m-n1">
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
          })
        )}
      </ListGroup>
      {count <= 0 && !isLoading && <Empty />}
      <div className="mt-4 mb-2 d-flex justify-content-center">
        <Pagination
          currentPage={curPage}
          totalSize={count}
          pageSize={pageSize}
          pathname={source === 'questions' ? '/questions' : ''}
        />
      </div>
    </div>
  );
};
export default QuestionList;
