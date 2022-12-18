import { FC } from 'react';
import { Container, Row, Col, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useParams, useSearchParams } from 'react-router-dom';

import { usePageTags } from '@/hooks';
import { Pagination, FormatTime, Empty } from '@/components';
import { loggedUserInfoStore } from '@/stores';
import {
  usePersonalInfoByName,
  usePersonalTop,
  usePersonalListByTabName,
} from '@/services';

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
  let pageTitle = '';
  if (userInfo) {
    pageTitle = `${userInfo.info.display_name} (${userInfo.info.username})`;
  }
  const { count = 0, list = [] } = listData?.[tabName] || {};
  usePageTags({
    title: pageTitle,
  });
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        {userInfo?.info?.status !== 'normal' && userInfo?.info?.status_msg && (
          <Alert data={userInfo?.info.status_msg} />
        )}
        <Col xxl={7} lg={8} sm={12}>
          <UserInfo data={userInfo?.info} />
        </Col>
        <Col
          xxl={3}
          lg={4}
          sm={12}
          className="d-flex justify-content-start justify-content-md-end">
          {isSelf && (
            <div className="mb-3">
              <Button
                variant="outline-secondary"
                href="/users/settings/profile"
                className="btn">
                {t('edit_profile')}
              </Button>
            </div>
          )}
        </Col>
      </Row>

      <Row className="justify-content-center">
        <Col xxl={10}>
          <NavBar tabName={tabName} slug={username} isSelf={isSelf} />
        </Col>
        <Col xxl={7} lg={8} sm={12}>
          <Overview
            visible={tabName === 'overview'}
            introduction={userInfo?.info?.bio_html}
            data={topData}
          />
          <ListHead
            count={tabName === 'reputation' ? userInfo?.info?.rank : count}
            sort={order}
            visible={tabName !== 'overview'}
            tabName={tabName}
          />
          <Answers data={list} visible={tabName === 'answers'} />
          <DefaultList
            data={list}
            tabName={tabName}
            visible={tabName === 'questions' || tabName === 'bookmarks'}
          />
          <Reputation data={list} visible={tabName === 'reputation'} />
          <Comments data={list} visible={tabName === 'comments'} />
          <Votes data={list} visible={tabName === 'votes'} />
          {!list?.length && !isLoading && <Empty />}

          {count > 0 && (
            <div className="d-flex justify-content-center border-top py-4">
              <Pagination
                pageSize={30}
                totalSize={count || 0}
                currentPage={Number(page)}
              />
            </div>
          )}
        </Col>
        <Col xxl={3} lg={4} sm={12} className="mt-5 mt-lg-0">
          <h5 className="mb-3">{t('stats')}</h5>
          {userInfo?.info && (
            <>
              <FormatTime time={1671290521} preFix={t('last_login')} />
              <div className="text-secondary">
                <FormatTime
                  time={userInfo.info.created_at}
                  preFix={t('joined')}
                />
              </div>
              <div className="text-secondary">
                <FormatTime
                  time={userInfo.info.last_login_date}
                  preFix={t('last_login')}
                />
              </div>
            </>
          )}
        </Col>
      </Row>
    </Container>
  );
};
export default Personal;
