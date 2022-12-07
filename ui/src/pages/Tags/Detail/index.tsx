import { FC, useEffect, useState } from 'react';
import { Container, Row, Col, Button } from 'react-bootstrap';
import { useParams, Link, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import * as Type from '@/common/interface';
import { FollowingTags } from '@/components';
import { useTagInfo, useFollow, useQuerySynonymsTags } from '@/services';
import QuestionList from '@/components/QuestionList';
import HotQuestions from '@/components/HotQuestions';
import { escapeRemove } from '@/utils';
import { pathFactory } from '@/router/pathFactory';

const Questions: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'tags' });
  const navigate = useNavigate();
  const routeParams = useParams();
  const curTagName = routeParams.tagName || '';
  const [tagInfo, setTagInfo] = useState<any>({});
  const [tagFollow, setTagFollow] = useState<Type.FollowParams>();
  const { data: tagResp } = useTagInfo({ name: curTagName });
  const { data: followResp } = useFollow(tagFollow);
  const { data: synonymsRes } = useQuerySynonymsTags(tagInfo?.tag_id);
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
        navigate(pathFactory.tagLanding(info.main_tag_slug_name), {
          replace: true,
        });
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
  if (tagInfo?.display_name) {
    pageTitle = `'${tagInfo.display_name}' ${t('questions', {
      keyPrefix: 'page_title',
    })}`;
  }
  const keywords: string[] = [];
  if (tagInfo?.slug_name) {
    keywords.push(tagInfo.slug_name);
  }
  synonymsRes?.synonyms.forEach((o) => {
    keywords.push(o.slug_name);
  });
  usePageTags({
    title: pageTitle,
    description: tagInfo?.description,
    keywords: keywords.join(','),
  });
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col xxl={7} lg={8} sm={12}>
          <div className="tag-box mb-5">
            <h3 className="mb-3">
              <Link
                to={pathFactory.tagLanding(tagInfo.slug_name)}
                replace
                className="link-dark">
                {tagInfo.display_name}
              </Link>
            </h3>

            <p className="text-break">
              {escapeRemove(tagInfo.excerpt) || t('no_desc')}
              <Link to={pathFactory.tagInfo(curTagName)} className="ms-1">
                [{t('more')}]
              </Link>
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
  );
};

export default Questions;
