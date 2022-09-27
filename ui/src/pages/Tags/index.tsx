import { useState } from 'react';
import {
  Container,
  Row,
  Col,
  Card,
  ButtonGroup,
  Button,
  Form,
} from 'react-bootstrap';
import { useSearchParams, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { useQueryTags, following } from '@answer/services/api';
import { Tag, Pagination, PageTitle } from '@answer/components';
import { formatCount } from '@answer/utils';

const Tags = () => {
  const [urlSearch] = useSearchParams();
  const { t } = useTranslation('translation', { keyPrefix: 'tags' });
  const [searchTag, setSearchTag] = useState('');
  const navigate = useNavigate();

  const page = Number(urlSearch.get('page')) || 1;
  const sort = urlSearch.get('sort');

  const pageSize = 20;
  const { data: tags, mutate } = useQueryTags({
    page,
    page_size: pageSize,
    ...(searchTag ? { slug_name: searchTag } : {}),
    ...(sort ? { query_cond: sort } : {}),
  });

  const handleChange = (e) => {
    setSearchTag(e.target.value);
  };

  const handleSort = (param) => {
    navigate(`/tags?sort=${param}`);
  };

  const handleFollow = (tag) => {
    following({
      object_id: tag.tag_id,
      is_cancel: tag.is_follower,
    }).then(() => {
      mutate();
    });
  };
  return (
    <>
      <PageTitle title={t('tags', { keyPrefix: 'page_title' })} />
      <Container className="py-3 my-3">
        <Row className="mb-4 d-flex justify-content-center">
          <Col lg={10}>
            <h3 className="mb-4">{t('title')}</h3>
            <div className="d-flex justify-content-between align-items-center">
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
              <ButtonGroup size="sm">
                <Button
                  variant={
                    !sort || sort === 'popular'
                      ? 'secondary'
                      : 'outline-secondary'
                  }
                  onClick={() => handleSort('popular')}>
                  {t('sort_buttons.popular')}
                </Button>
                <Button
                  variant={sort === 'name' ? 'secondary' : 'outline-secondary'}
                  onClick={() => handleSort('name')}>
                  {t('sort_buttons.name')}
                </Button>
                <Button
                  className="text-capitalize"
                  variant={
                    sort === 'newest' ? 'secondary' : 'outline-secondary'
                  }
                  onClick={() => handleSort('newest')}>
                  {t('sort_buttons.newest')}
                </Button>
              </ButtonGroup>
            </div>
          </Col>

          <Col className="mt-4" lg={10}>
            <Row>
              {tags?.list?.map((tag) => (
                <Col key={tag.slug_name} lg={3} md={4} className="mb-4">
                  <Card className="h-100">
                    <Card.Body className="d-flex flex-column align-items-start">
                      <Tag className="mb-3" href={`/tags/${tag.slug_name}`}>
                        {tag.slug_name}
                      </Tag>
                      <p className="fs-14 flex-fill text-break text-wrap text-truncate-4">
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
              ))}
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
    </>
  );
};

export default Tags;
