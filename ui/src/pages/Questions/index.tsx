import { FC } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { useMatch } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { PageTitle, FollowingTags } from '@answer/components';

import QuestionList from '@/components/QuestionList';
import HotQuestions from '@/components/HotQuestions';
import { siteInfoStore } from '@/stores';

const Questions: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'question' });

  const isIndexPage = useMatch('/');
  let pageTitle = t('questions', { keyPrefix: 'page_title' });
  let slogan = '';
  const { siteInfo } = siteInfoStore();
  if (isIndexPage) {
    pageTitle = `${siteInfo.name}`;
    slogan = `${siteInfo.short_description}`;
  }

  return (
    <>
      <PageTitle title={pageTitle} suffix={slogan} />
      <Container className="pt-4 mt-2 mb-5">
        <Row className="justify-content-center">
          <Col xxl={7} lg={8} sm={12}>
            <QuestionList source="questions" />
          </Col>
          <Col xxl={3} lg={4} sm={12} className="mt-5 mt-lg-0">
            <FollowingTags />
            <HotQuestions />
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default Questions;
