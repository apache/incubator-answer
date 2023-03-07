import { useState } from 'react';
import {
  Container,
  Row,
  Col,
  Card,
  Button,
  Form,
  Stack,
} from 'react-bootstrap';
import { useSearchParams, Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { Tag, Pagination, QueryGroup, TagsLoader } from '@/components';
import { formatCount } from '@/utils';
import { tryNormalLogged } from '@/utils/guard';
import { useQueryTags, following } from '@/services';
import { loggedUserInfoStore } from '@/stores';

const sortBtns = ['popular', 'name', 'newest'];

const Tags = () => {
  const [urlSearch] = useSearchParams();
  const { t } = useTranslation('translation', { keyPrefix: 'tags' });
  const [searchTag, setSearchTag] = useState('');
  const { role_id } = loggedUserInfoStore((_) => _.user);

  const page = Number(urlSearch.get('page')) || 1;
  const sort = urlSearch.get('sort');

  const pageSize = 20;
  const {
    data: tags,
    mutate,
    isLoading,
  } = useQueryTags({
    page,
    page_size: pageSize,
    ...(searchTag ? { slug_name: searchTag } : {}),
    ...(sort ? { query_cond: sort } : {}),
  });

  const handleChange = (e) => {
    setSearchTag(e.target.value);
  };

  const handleFollow = (tag) => {
    if (!tryNormalLogged(true)) {
      return;
    }
    following({
      object_id: tag.tag_id,
      is_cancel: tag.is_follower,
    }).then(() => {
      mutate();
    });
  };
  usePageTags({
    title: t('tags', { keyPrefix: 'page_title' }),
  });
  return (
    <Container className="py-3 my-3">
      <Row className="mb-4 d-flex justify-content-center">
        <Col xxl={10} sm={12}>
          <h3 className="mb-4">{t('title')}</h3>
          <div className="d-flex justify-content-between align-items-center flex-wrap">
            <Stack direction="horizontal" gap={3}>
              <Form>
                <Form.Group controlId="formBasicEmail">
                  <Form.Control
                    value={searchTag}
                    placeholder={t('search_placeholder')}
                    type="text"
                    onChange={handleChange}
                    size="sm"
                  />
                </Form.Group>
              </Form>
              {role_id === 2 || role_id === 3 ? (
                <Link
                  className="btn btn-outline-primary btn-sm"
                  to="/tags/create">
                  {t('title', { keyPrefix: 'tag_modal' })}
                </Link>
              ) : null}
            </Stack>
            <QueryGroup
              data={sortBtns}
              currentSort={sort || 'popular'}
              sortKey="sort"
              i18nKeyPrefix="tags.sort_buttons"
            />
          </div>
        </Col>

        <Col className="mt-4" xxl={10} sm={12}>
          <Row>
            {isLoading ? (
              <TagsLoader />
            ) : (
              tags?.list?.map((tag) => (
                <Col
                  key={tag.slug_name}
                  xs={12}
                  lg={3}
                  md={4}
                  sm={6}
                  className="mb-4">
                  <Card className="h-100">
                    <Card.Body className="d-flex flex-column align-items-start">
                      <Tag className="mb-3" data={tag} />

                      <p className="fs-14 flex-fill text-break text-wrap text-truncate-3">
                        {tag.original_text}
                      </p>
                      <div className="d-flex align-items-center">
                        <Button
                          className={`me-2 ${tag.is_follower ? 'active' : ''}`}
                          variant="outline-primary"
                          size="sm"
                          onClick={() => handleFollow(tag)}>
                          {tag.is_follower
                            ? t('button_following')
                            : t('button_follow')}
                        </Button>
                        <span className="text-secondary fs-14 text-nowrap">
                          {formatCount(tag.question_count)} {t('tag_label')}
                        </span>
                      </div>
                    </Card.Body>
                  </Card>
                </Col>
              ))
            )}
          </Row>
          <div className="d-flex justify-content-center">
            <Pagination
              currentPage={page}
              totalSize={tags?.count || 0}
              pageSize={pageSize}
            />
          </div>
        </Col>
      </Row>
    </Container>
  );
};

export default Tags;
