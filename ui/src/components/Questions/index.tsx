import { FC } from 'react';
import { Row, Col, ButtonGroup, Button, ListGroup } from 'react-bootstrap';
import { NavLink, useParams, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { useQuestionList } from '@answer/api';
import type * as Type from '@answer/common/interface';
import {
  Icon,
  Tag,
  Pagination,
  FormatTime,
  Empty,
  BaseUserCard,
} from '@answer/components';

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
      <>
        <BaseUserCard
          data={q.last_answered_user_info}
          showAvatar={false}
          className="me-1"
        />
        •
        <FormatTime
          time={q.update_time}
          className="text-secondary mx-1"
          preFix={t('answered')}
        />
      </>
    );
  }

  if (q.edit_time > q.update_time) {
    // question modified
    return (
      <>
        <BaseUserCard
          data={q.update_user_info}
          showAvatar={false}
          className="me-1"
        />
        •
        <FormatTime
          time={q.edit_time}
          className="text-secondary mx-1"
          preFix={t('modified')}
        />
      </>
    );
  }

  // default: asked
  return (
    <>
      <BaseUserCard data={q.user_info} showAvatar={false} className="me-1" />
      •
      <FormatTime
        time={q.create_time}
        preFix={t('asked')}
        className="text-secondary mx-1"
      />
    </>
  );
};

const QuestionList: FC<Props> = ({ source }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'question' });
  const { tagName = '' } = useParams();
  const [urlSearchParams, setUrlSearchParams] = useSearchParams();
  const curOrder = urlSearchParams.get('order') || QuestionOrderKeys[0];
  const curPage = Number(urlSearchParams.get('page')) || 1;
  const pageSize = 20;
  const reqParams: Type.QueryQuestionsReq = {
    page_size: pageSize,
    page: curPage,
    order: curOrder as Type.QuestionOrderBy,
    tags: [tagName],
  };

  if (source === 'questions') {
    delete reqParams.tags;
  }
  const { data: listData, isLoading } = useQuestionList(reqParams);
  const count = listData?.count || 0;
  const onOrderChange = (evt, order) => {
    evt.preventDefault();
    if (order === curOrder) {
      return;
    }
    urlSearchParams.set('page', '1');
    urlSearchParams.set('order', order);
    setUrlSearchParams(urlSearchParams);
  };

  return (
    <div>
      <Row className="mb-3">
        <Col className="d-flex align-items-center">
          <h5 className="fs-5 text-nowrap mb-3 mb-md-0">
            {source === 'questions'
              ? t('all_questions')
              : t('x_questions', { count })}
          </h5>
        </Col>
        <Col>
          <ButtonGroup size="sm">
            {QuestionOrderKeys.map((k) => {
              return (
                <Button
                  as="a"
                  key={k}
                  className="text-capitalize"
                  href={`?page=1&order=${k}`}
                  onClick={(evt) => onOrderChange(evt, k)}
                  variant={curOrder === k ? 'secondary' : 'outline-secondary'}>
                  {t(k)}
                </Button>
              );
            })}
          </ButtonGroup>
        </Col>
      </Row>
      <ListGroup variant="flush" className="border-top border-bottom-0">
        {listData?.list?.map((li) => {
          return (
            <ListGroup.Item key={li.id} className="border-bottom py-3 px-0">
              <h5 className="text-wrap text-break">
                <NavLink to={`/questions/${li.id}`} className="text-body">
                  {li.title}
                </NavLink>
              </h5>
              <div className="d-flex align-items-center fs-14 text-secondary">
                <QuestionLastUpdate q={li} />
                <span className="ms-3">
                  <Icon name="hand-thumbs-up-fill" />
                  <em className="fst-normal mx-1">{li.vote_count}</em>
                </span>
                <span className="ms-3">
                  {li.accepted_answer_id >= 1 ? (
                    <Icon name="check-circle-fill" className="text-success" />
                  ) : (
                    <Icon name="chat-square-text-fill" />
                  )}
                  <em className="fst-normal mx-1">{li.answer_count}</em>
                </span>
                <span className="summary-stat ms-3">
                  <Icon name="eye-fill" />
                  <em className="fst-normal mx-1">{li.view_count}</em>
                </span>
              </div>
              <div className="question-tags">
                {Array.isArray(li.tags)
                  ? li.tags.map((tag) => {
                      return (
                        <Tag
                          key={tag.slug_name}
                          className="me-2 mt-2"
                          href={`/tags/${
                            tag.main_tag_slug_name || tag.slug_name
                          }`}>
                          {tag.slug_name}
                        </Tag>
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
