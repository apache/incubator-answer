import { FC } from 'react';
import { Row, Col } from 'react-bootstrap';
import { useParams, useSearchParams, Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { useQuestionLink } from '@/services';
import * as Type from '@/common/interface';
import {
  QuestionList,
  CustomSidebar,
  HotQuestions,
  FollowingTags,
} from '@/components';
import { userCenter, floppyNavigation } from '@/utils';
import { QUESTION_ORDER_KEYS } from '@/components/QuestionList';
import {
  loggedUserInfoStore,
  siteInfoStore,
  loginSettingStore,
} from '@/stores';

const LinkedQuestions: FC = () => {
  const { qid } = useParams<{ qid: string }>();
  const { t } = useTranslation('translation', { keyPrefix: 'linked_question' });
  const { t: t2 } = useTranslation('translation');
  const { user: loggedUser } = loggedUserInfoStore((_) => _);
  const [urlSearchParams] = useSearchParams();
  const curPage = Number(urlSearchParams.get('page')) || 1;
  const curOrder = (urlSearchParams.get('order') ||
    QUESTION_ORDER_KEYS[0]) as Type.QuestionOrderBy;
  const pageSize = 10;
  const { siteInfo } = siteInfoStore();
  const { data: listData, isLoading: listLoading } = useQuestionLink({
    question_id: qid || '',
    page: curPage,
    page_size: pageSize,
  });
  const { login: loginSetting } = loginSettingStore();

  usePageTags({
    title: t('title'),
  });

  return (
    <Row className="pt-4 mb-5">
      <Col className="page-main flex-auto">
        <QuestionList
          source="linked"
          data={listData}
          order={curOrder}
          orderList={
            loggedUser.username
              ? QUESTION_ORDER_KEYS
              : QUESTION_ORDER_KEYS.filter((key) => key !== 'recommend')
          }
          isLoading={listLoading}
        />
      </Col>
      <Col className="page-right-side mt-4 mt-xl-0">
        <CustomSidebar />
        {!loggedUser.username && (
          <div className="card mb-4">
            <div className="card-body">
              <h5 className="card-title">
                {t2('website_welcome', {
                  site_name: siteInfo.name,
                })}
              </h5>
              <p className="card-text">{siteInfo.description}</p>
              <Link
                to={userCenter.getLoginUrl()}
                className="btn btn-primary"
                onClick={floppyNavigation.handleRouteLinkClick}>
                {t('login', { keyPrefix: 'btns' })}
              </Link>
              {loginSetting.allow_new_registrations ? (
                <Link
                  to={userCenter.getSignUpUrl()}
                  className="btn btn-link ms-2"
                  onClick={floppyNavigation.handleRouteLinkClick}>
                  {t('signup', { keyPrefix: 'btns' })}
                </Link>
              ) : null}
            </div>
          </div>
        )}
        {loggedUser.access_token && <FollowingTags />}
        <HotQuestions />
      </Col>
    </Row>
  );
};

export default LinkedQuestions;
