import { useEffect, useState } from 'react';
import { Row, Col } from 'react-bootstrap';
import {
  useParams,
  useSearchParams,
  useNavigate,
  useLocation,
} from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { Pagination, CustomSidebar } from '@/components';
import { loggedUserInfoStore, toastStore } from '@/stores';
import { scrollToElementTop, scrollToDocTop } from '@/utils';
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
  InviteToAnswer,
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
  if (!aid && slugPermalink) {
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
  const isAdmin = userInfo?.role_id === 2;
  const isModerator = userInfo?.role_id === 3;
  const isLogged = Boolean(userInfo?.access_token);
  const loggedUserRank = userInfo?.rank;
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
      res.list = res.list?.filter((v) => {
        // delete answers only show to author and admin and has search params aid
        if (v.status === 10) {
          if (
            (v?.user_info?.username === userInfo?.username || isAdmin) &&
            aid === v.id
          ) {
            return v;
          }
          return null;
        }

        return v;
      });

      setAnswers({ ...res, count: res.list.length });
      if (page > 0 || order) {
        // scroll into view;
        const element = document.getElementById('answerHeader');
        scrollToElementTop(element);
      }

      res.list.forEach((item) => {
        setUsers([
          {
            displayName: item.user_info?.display_name,
            userName: item.user_info?.username,
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
          {
            id: res.user_info?.id,
            displayName: res.user_info?.display_name,
            userName: res.user_info?.username,
            avatar_url: res.user_info?.avatar,
          },
          {
            id: res?.update_user_info?.id,
            displayName: res?.update_user_info?.display_name,
            userName: res?.update_user_info?.username,
            avatar_url: res?.update_user_info?.avatar,
          },
          {
            id: res?.last_answered_user_info?.id,
            displayName: res?.last_answered_user_info?.display_name,
            userName: res?.last_answered_user_info?.username,
            avatar_url: res?.last_answered_user_info?.avatar,
          },
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
        navigate('/', { replace: true });
      }, 1000);
      return;
    }
    if (type === 'default') {
      scrollToDocTop();
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
    scrollToDocTop();
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

  const showInviteToAnswer = question?.id;
  let canInvitePeople = false;
  if (showInviteToAnswer && Array.isArray(question.extends_actions)) {
    const inviteAct = question.extends_actions.find((op) => {
      return op.action === 'invite_other_to_answer';
    });
    if (inviteAct) {
      canInvitePeople = true;
    }
  }

  return (
    <Row className="questionDetailPage pt-4 mb-5">
      <Col className="page-main flex-auto">
        {question?.operation?.level && <Alert data={question.operation} />}
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
                  canAccept={isAuthor || isAdmin || isModerator}
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

        {!isLoading &&
          Number(question?.status) !== 2 &&
          !question?.operation?.type && (
            <WriteAnswer
              data={{
                qid,
                answered: question?.answered,
                loggedUserRank,
              }}
              callback={writeAnswerCallback}
            />
          )}
      </Col>
      <Col className="page-right-side mt-4 mt-xl-0">
        <CustomSidebar />
        <RelatedQuestions id={question?.id || ''} />
        {showInviteToAnswer ? (
          <InviteToAnswer
            questionId={question.id}
            readOnly={!canInvitePeople}
          />
        ) : null}
      </Col>
    </Row>
  );
};

export default Index;
