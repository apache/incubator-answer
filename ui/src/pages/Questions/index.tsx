import { FC } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { useMatch, Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { FollowingTags } from '@/components';
import QuestionList from '@/components/QuestionList';
import HotQuestions from '@/components/HotQuestions';
import { siteInfoStore, loggedUserInfoStore } from '@/stores';

const Questions: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'question' });
  const { user: loggedUser } = loggedUserInfoStore((_) => _);
  const isIndexPage = useMatch('/');
  let pageTitle = t('questions', { keyPrefix: 'page_title' });
  let slogan = '';
  const { siteInfo } = siteInfoStore();
  if (isIndexPage) {
    pageTitle = `${siteInfo.name}`;
    slogan = `${siteInfo.short_description}`;
  }

  usePageTags({ title: pageTitle, subtitle: slogan });
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col xxl={7} lg={8} sm={12}>
          <QuestionList source="questions" />
        </Col>
        <Col xxl={3} lg={4} sm={12} className="mt-5 mt-lg-0">
          {!loggedUser.access_token && (
            <div className="card mb-4">
              <div className="card-body">
                <h5 className="card-title">
                  {t('page_title', {
                    keyPrefix: 'login',
                    site_name: siteInfo.name,
                  })}
                </h5>
                <p className="card-text">{siteInfo.description}</p>
                <Link to="/users/login" className="card-link btn btn-primary">
                  {t('login', { keyPrefix: 'btns' })}
                </Link>
                <Link to="/users/register" className="card-link">
                  {t('signup', { keyPrefix: 'btns' })}
                </Link>
              </div>
            </div>
          )}
          {loggedUser.access_token && <FollowingTags />}
          <HotQuestions />
        </Col>
      </Row>
    </Container>
  );
};

export default Questions;
