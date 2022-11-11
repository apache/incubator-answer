import { FC, useEffect, useState } from 'react';
import { Container, Row, Col, Button } from 'react-bootstrap';
import { useParams, Link, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import * as Type from '@/common/interface';
import { PageTitle, FollowingTags } from '@/components';
import { useTagInfo, useFollow } from '@/services';
import QuestionList from '@/components/QuestionList';
import HotQuestions from '@/components/HotQuestions';
import { escapeRemove } from '@/utils';

const Questions: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'tags' });
  const navigate = useNavigate();
  const routeParams = useParams();
  const curTagName = routeParams.tagName;
  const [tagInfo, setTagInfo] = useState<any>({});
  const [tagFollow, setTagFollow] = useState<Type.FollowParams>();
  const { data: tagResp } = useTagInfo({ name: curTagName });
  const { data: followResp } = useFollow(tagFollow);

  const toggleFollow = () => {
    setTagFollow({
      is_cancel: tagInfo.is_follower,
      object_id: tagInfo.tag_id,
    });
  };
  useEffect(() => {
    if (tagResp) {
      const info = { ...tagResp };
      if (info.main_tag_slug_name) {
        navigate(`/tags/${info.main_tag_slug_name}`, { replace: true });
        return;
      }
      if (followResp) {
        info.is_follower = followResp.is_followed;
      }

      if (info.excerpt) {
        info.excerpt =
          info.excerpt.length > 256
            ? [...info.excerpt].slice(0, 256).join('')
            : info.excerpt;
      }

      setTagInfo(info);
    }
  }, [tagResp, followResp]);
  let pageTitle = '';
  if (tagInfo) {
    pageTitle = `'${tagInfo.display_name}' ${t('questions', {
      keyPrefix: 'page_title',
    })}`;
  }
  return (
    <>
      <PageTitle title={pageTitle} />
      <Container className="pt-4 mt-2 mb-5">
        <Row className="justify-content-center">
          <Col xxl={7} lg={8} sm={12}>
            <div className="tag-box mb-5">
              <h3 className="mb-3">
                <Link
                  to={`/tags/${tagInfo?.slug_name}`}
                  replace
                  className="link-dark">
                  {tagInfo.display_name}
                </Link>
              </h3>

              <p className="text-break">
                {escapeRemove(tagInfo.excerpt) || t('no_description')}
                <Link to={`/tags/${curTagName}/info`}> [{t('more')}]</Link>
              </p>

              <div className="box-ft">
                {tagInfo.is_follower ? (
                  <Button variant="primary" onClick={() => toggleFollow()}>
                    {t('button_following')}
                  </Button>
                ) : (
                  <Button
                    variant="outline-primary"
                    onClick={() => toggleFollow()}>
                    {t('button_follow')}
                  </Button>
                )}
              </div>
            </div>
            <QuestionList source="tag" />
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
