/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import { FC, useState, useEffect } from 'react';
import { Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useParams, useSearchParams, Link } from 'react-router-dom';

import {
  usePersonalInfoByName,
  usePersonalTop,
  usePersonalListByTabName,
  questionDetail,
  getAnswers,
} from '@/services';
import { usePageTags } from '@/hooks';
import { Pagination, FormatTime, Empty } from '@/components';
import { loggedUserInfoStore } from '@/stores';
import type { UserInfoRes } from '@/common/interface';

import {
  UserInfo,
  NavBar,
  Overview,
  Alert,
  ListHead,
  DefaultList,
  Reputation,
  Comments,
  Answers,
  Votes,
} from './components';

const Personal: FC = () => {
  const { tabName = 'overview', username = '' } = useParams();
  const [searchParams] = useSearchParams();
  const page = searchParams.get('page') || 1;
  const order = searchParams.get('order') || 'newest';
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });
  const sessionUser = loggedUserInfoStore((state) => state.user);
  const isSelf = sessionUser?.username === username;
  const currentUserId = sessionUser?.id;
  const roleId = sessionUser?.role_id;
  const { data: userInfo } = usePersonalInfoByName(username);
  const { data: topData } = usePersonalTop(username, tabName);

  const { data: listData, isLoading = true } = usePersonalListByTabName(
    {
      username,
      page: Number(page),
      page_size: 30,
      order,
    },
    tabName,
  );
  const getFullQuestionInfo = async (list) => {
    const promises = list.map((item) => {
      const questionId = item.question_id;
      return questionDetail(questionId)
        .then((questionInfo) => ({ ...item, questionInfo }))
        .catch((error) => {
          console.error(`error_ID: ${questionId}`, error);
          return null;
        });
    });
    const completedList = await Promise.all(promises);
    return completedList;
  };

  const getAnswersUserId = async (list) => {
    const promises = list.map((item) =>
      getAnswers({
        order: order === 'updated' ? order : 'default',
        question_id: item.question_id,
        page: 1,
        page_size: 999,
      })
        .then((res) => ({
          questionId: item.question_id,
          user_ids: res.list.map((answer) => answer.user_info.id),
        }))
        .catch((error) => {
          console.error(
            `Error fetching answers for question ID: ${item.question_id}`,
            error,
          );
          return { questionId: item.question_id, user_ids: [] };
        }),
    );

    const results = await Promise.all(promises);
    const answersMap = new Map();
    results.forEach(({ questionId, user_ids }) =>
      answersMap.set(questionId, user_ids),
    );
    return answersMap;
  };
  const getFiltered = (completedList, answersMap) => {
    const filtered = completedList.filter((item) => {
      if (!item) return false;
      const userId = item.questionInfo?.user_info?.id;
      const showValue = item.questionInfo?.show;
      const userAnswered = answersMap
        .get(item.question_id)
        ?.includes(currentUserId);
      return showValue !== 2 || userId === currentUserId || userAnswered;
    });
    return filtered;
  };
  const { list = [] } = listData?.[tabName] || {};
  const top = topData ?? { answer: [], question: [] };
  console.log('top data', topData);
  console.log('list data', list);
  const [filteredList, setFilteredList] = useState<any[]>([]);
  const [filteredTop, setFilteredTop] = useState<{
    answer: any[];
    question: any[];
  }>({ answer: [], question: [] });
  useEffect(() => {
    const fetchTop = async () => {
      if (roleId === 2) {
        setFilteredTop(top);
        return;
      }
      const completedAnswers = await getFullQuestionInfo(top.answer);
      const completedQuestion = await getFullQuestionInfo(top.question);
      const answersAnswerMap = await getAnswersUserId(completedAnswers);
      const questionsAnswerMap = await getAnswersUserId(completedQuestion);
      const filteredAnswer = getFiltered(completedAnswers, answersAnswerMap);
      const filteredQuestion = getFiltered(
        completedQuestion,
        questionsAnswerMap,
      );
      const answerQuestionIds = filteredAnswer.map((answer) => {
        return answer.question_id;
      });
      const questionQuestionIds = filteredQuestion.map((question) => {
        return question.question_id;
      });
      const filtered = {
        answer: top.answer.filter((answer) =>
          answerQuestionIds.includes(answer.question_id),
        ),
        question: top.question.filter((question) =>
          questionQuestionIds.includes(question.question_id),
        ),
      };
      setFilteredTop(filtered);
    };
    if (top.answer.length || top.question.length) {
      fetchTop();
    }
  }, [top, currentUserId, roleId]);
  useEffect(() => {
    const fetchQuestionDetails = async () => {
      if (roleId === 2) {
        setFilteredList(list);
        return;
      }
      let completedList;
      if (tabName === 'bookmarks') {
        completedList = list.map((item) => ({ ...item, questionInfo: item }));
      } else {
        completedList = await getFullQuestionInfo(list);
      }
      const answersMap = await getAnswersUserId(completedList);
      const filtered = getFiltered(completedList, answersMap);
      setFilteredList(filtered);
    };

    if (list.length) {
      fetchQuestionDetails();
    }
  }, [list, currentUserId, tabName, roleId]);

  const count = filteredList.length;

  let pageTitle = '';
  if (userInfo?.username) {
    pageTitle = `${userInfo?.display_name} (${userInfo?.username})`;
  }
  usePageTags({
    title: pageTitle,
  });

  return (
    <div className="pt-4 mb-5">
      <Row>
        {userInfo?.status !== 'normal' && userInfo?.status_msg && (
          <Alert data={userInfo?.status_msg} />
        )}
        <Col className="page-main flex-auto">
          <UserInfo data={userInfo as UserInfoRes} />
        </Col>
        <Col
          xxl={3}
          lg={4}
          sm={12}
          className="page-right-side mt-4 mt-xl-0 d-flex justify-content-start justify-content-md-end">
          {isSelf && (
            <div className="mb-3">
              <Link
                className="btn btn-outline-secondary"
                to="/users/settings/profile">
                {t('edit_profile')}
              </Link>
            </div>
          )}
        </Col>
      </Row>
      <NavBar tabName={tabName} slug={username} isSelf={isSelf} />
      <Row>
        <Col className="page-main flex-auto">
          <Overview
            visible={tabName === 'overview'}
            introduction={userInfo?.bio_html || ''}
            data={filteredTop}
          />
          <ListHead
            count={tabName === 'reputation' ? Number(userInfo?.rank) : count}
            sort={order}
            visible={tabName !== 'overview'}
            tabName={tabName}
          />
          <Answers data={filteredList} visible={tabName === 'answers'} />
          <DefaultList
            data={filteredList}
            tabName={tabName}
            visible={tabName === 'questions' || tabName === 'bookmarks'}
          />
          <Reputation data={filteredList} visible={tabName === 'reputation'} />
          <Comments data={filteredList} visible={tabName === 'comments'} />
          <Votes data={filteredList} visible={tabName === 'votes'} />
          {!list?.length && !isLoading && <Empty />}

          {count > 0 && (
            <div className="d-flex justify-content-center py-4">
              <Pagination
                pageSize={30}
                totalSize={count || 0}
                currentPage={Number(page)}
              />
            </div>
          )}
        </Col>
        <Col className="page-right-side mt-4 mt-xl-0">
          <h5 className="mb-3">{t('stats')}</h5>
          {userInfo?.created_at && (
            <>
              <div className="text-secondary">
                <FormatTime time={userInfo.created_at} preFix={t('joined')} />
              </div>
              <div className="text-secondary">
                <FormatTime
                  time={userInfo.last_login_date}
                  preFix={t('last_login')}
                />
              </div>
            </>
          )}
        </Col>
      </Row>
    </div>
  );
};
export default Personal;
