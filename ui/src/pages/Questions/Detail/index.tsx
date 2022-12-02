import { useEffect, useState } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { useParams, useSearchParams, useNavigate } from 'react-router-dom';

import { Pagination, PageTitle } from '@/components';
import { loggedUserInfoStore } from '@/stores';
import { scrollTop } from '@/utils';
import { usePageUsers } from '@/hooks';
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
} from './components';

import './index.scss';

const Index = () => {
  const navigate = useNavigate();
  const { qid = '', aid = '' } = useParams();
  const [urlSearch] = useSearchParams();
  const page = Number(urlSearch.get('page') || 0);
  const order = urlSearch.get('order') || '';
  const [question, setQuestion] = useState<QuestionDetailRes | null>(null);
  const [answers, setAnswers] = useState<ListResult<AnswerItem>>({
    count: -1,
    list: [],
  });
  const { setUsers } = usePageUsers();
  const userInfo = loggedUserInfoStore((state) => state.user);
  const isAuthor = userInfo?.username === question?.user_info?.username;
  const isLogged = Boolean(userInfo?.access_token);
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
    const res = await questionDetail(qid);
    if (res) {
      // undo
      setUsers([
        res.user_info,
        res?.update_user_info,
        res?.last_answered_user_info,
      ]);
      setQuestion(res);
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

  return (
    <>
      <PageTitle title={question?.title} />
      <Container className="pt-4 mt-2 mb-5 questionDetailPage">
        <Row className="justify-content-center">
          <Col xxl={7} lg={8} sm={12} className="mb-5 mb-md-0">
            {question?.operation?.operation_type && (
              <Alert data={question.operation} />
            )}
            <Question
              data={question}
              initPage={initPage}
              hasAnswer={answers.count > 0}
              isLogged={isLogged}
            />
            {answers.count > 0 && (
              <>
                <AnswerHead count={answers.count} order={order} />
                {answers?.list?.map((item) => {
                  return (
                    <Answer
                      aid={aid}
                      key={item?.id}
                      data={item}
                      questionTitle={question?.title || ''}
                      isAuthor={isAuthor}
                      callback={initPage}
                      isLogged={isLogged}
                    />
                  );
                })}
              </>
            )}

            {Math.ceil(answers.count / 15) > 1 && (
              <div className="d-flex justify-content-center answer-item pt-4">
                <Pagination
                  currentPage={Number(page || 1)}
                  pageSize={15}
                  totalSize={answers?.count || 0}
                />
              </div>
            )}

            {!question?.operation?.operation_type && (
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
    </>
  );
};

export default Index;
