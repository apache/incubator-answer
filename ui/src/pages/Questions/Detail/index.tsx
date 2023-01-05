import { useEffect, useState } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import {
  useParams,
  useSearchParams,
  useNavigate,
  useLocation,
} from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import Pattern from '@/common/pattern';
import { Pagination } from '@/components';
import { loggedUserInfoStore, toastStore } from '@/stores';
import { scrollTop, bgFadeOut } from '@/utils';
import { usePageTags, usePageUsers } from '@/hooks';
import type {
  ListResult,
  QuestionDetailRes,
  AnswerItem,
} from '@/common/interface';
import { questionDetail, getAnswers } from '@/services';

import {
  Question,
  Answer,
  AnswerHead,
  RelatedQuestions,
  WriteAnswer,
  Alert,
  ContentLoader,
} from './components';

import './index.scss';

const Index = () => {
  const navigate = useNavigate();
  const { t } = useTranslation('translation');
  const { qid = '', slugPermalink = '' } = useParams();
  /**
   * Note: Compatible with Permalink
   */
  let { aid = '' } = useParams();
  if (!aid && Pattern.isAnswerId.test(slugPermalink)) {
    aid = slugPermalink;
  }

  const [urlSearch] = useSearchParams();
  const page = Number(urlSearch.get('page') || 0);
  const order = urlSearch.get('order') || '';
  const [question, setQuestion] = useState<QuestionDetailRes | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [answers, setAnswers] = useState<ListResult<AnswerItem>>({
    count: -1,
    list: [],
  });
  const { setUsers } = usePageUsers();
  const userInfo = loggedUserInfoStore((state) => state.user);
  const isAuthor = userInfo?.username === question?.user_info?.username;
  const isLogged = Boolean(userInfo?.access_token);
  const { state: locationState } = useLocation();

  useEffect(() => {
    if (locationState?.isReview) {
      toastStore.getState().show({
        msg: t('review', { keyPrefix: 'toast' }),
        variant: 'warning',
      });
    }
  }, [locationState]);

  const requestAnswers = async () => {
    const res = await getAnswers({
      order: order === 'updated' ? order : 'default',
      question_id: qid,
      page: 1,
      page_size: 999,
    });
    if (res) {
      setAnswers(res);
      if (page > 0 || order) {
        // scroll into view;
        const element = document.getElementById('answerHeader');
        scrollTop(element);
        bgFadeOut(element);
      }

      res.list.forEach((item) => {
        setUsers([
          {
            displayName: item.user_info.display_name,
            userName: item.user_info.username,
          },
          {
            displayName: item?.update_user_info?.display_name,
            userName: item?.update_user_info?.username,
          },
        ]);
      });
    }
  };

  const getDetail = async () => {
    setIsLoading(true);
    try {
      const res = await questionDetail(qid);
      if (res) {
        setUsers([
          res.user_info,
          res?.update_user_info,
          res?.last_answered_user_info,
        ]);
        setQuestion(res);
      }
      setIsLoading(false);
    } catch (e) {
      setIsLoading(false);
    }
  };

  const initPage = (type: string) => {
    if (type === 'delete_question') {
      setTimeout(() => {
        navigate(-1);
      }, 1000);
      return;
    }

    if (type === 'default') {
      window.scrollTo(0, 0);
      getDetail();
      return;
    }
    requestAnswers();
  };

  const writeAnswerCallback = (obj: AnswerItem) => {
    setAnswers({
      count: answers.count + 1,
      list: [...answers.list, obj],
    });

    if (question) {
      setQuestion({
        ...question,
        answered: true,
      });
    }
  };

  useEffect(() => {
    if (!qid) {
      return;
    }
    getDetail();
    requestAnswers();
  }, [qid]);

  useEffect(() => {
    if (page || order) {
      requestAnswers();
    }
  }, [page, order]);
  usePageTags({
    title: question?.title,
    description: question?.description,
    keywords: question?.tags.map((_) => _.slug_name).join(','),
  });
  return (
    <Container className="pt-4 mt-2 mb-5 questionDetailPage">
      <Row className="justify-content-center">
        <Col xxl={7} lg={8} sm={12} className="mb-5 mb-md-0">
          {question?.operation?.operation_type && (
            <Alert data={question.operation} />
          )}
          {isLoading ? (
            <ContentLoader />
          ) : (
            <Question
              data={question}
              initPage={initPage}
              hasAnswer={answers.count > 0}
              isLogged={isLogged}
            />
          )}
          {!isLoading && answers.count > 0 && (
            <>
              <AnswerHead count={answers.count} order={order} />
              {answers?.list?.map((item) => {
                return (
                  <Answer
                    aid={aid}
                    key={item?.id}
                    data={item}
                    questionTitle={question?.title || ''}
                    slugTitle={question?.url_title}
                    isAuthor={isAuthor}
                    callback={initPage}
                    isLogged={isLogged}
                  />
                );
              })}
            </>
          )}

          {!isLoading && Math.ceil(answers.count / 15) > 1 && (
            <div className="d-flex justify-content-center answer-item pt-4">
              <Pagination
                currentPage={Number(page || 1)}
                pageSize={15}
                totalSize={answers?.count || 0}
              />
            </div>
          )}

          {!isLoading && !question?.operation?.operation_type && (
            <WriteAnswer
              visible={answers.count === 0}
              data={{
                qid,
                answered: question?.answered,
              }}
              callback={writeAnswerCallback}
            />
          )}
        </Col>
        <Col xxl={3} lg={4} sm={12} className="mt-5 mt-lg-0">
          <RelatedQuestions id={question?.id || ''} />
        </Col>
      </Row>
    </Container>
  );
};

export default Index;
