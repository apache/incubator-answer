import React from 'react';
import { Container, Row, Col, ListGroup } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useSearchParams } from 'react-router-dom';

import { Pagination, PageTitle } from '@answer/components';

import { Head, SearchHead, SearchItem, Tips, Empty } from './components';

import { useSearch } from '@/services';

const Index = () => {
  const { t } = useTranslation('translation');
  const [searchParams] = useSearchParams();
  const page = searchParams.get('page') || 1;
  const q = searchParams.get('q') || '';
  const order = searchParams.get('order') || 'relevance';

  const { data, isLoading } = useSearch({
    q,
    order,
    page: Number(page),
    size: 20,
  });

  const { count = 0, list = [], extra = null } = data || {};
  let pageTitle = t('search', { keyPrefix: 'page_title' });
  if (q) {
    pageTitle = `${t('posts_containing', { keyPrefix: 'page_title' })} '${q}'`;
  }
  return (
    <>
      <PageTitle title={pageTitle} />
      <Container className="pt-4 mt-2 mb-5">
        <Row className="justify-content-center">
          <Col xxl={7} lg={8} sm={12} className="mb-3">
            <Head data={extra} />

            <ListGroup variant="flush" className="mb-5">
              <SearchHead sort={order} count={count} />

              {list?.map((item) => {
                return <SearchItem key={item.object.id} data={item} />;
              })}
            </ListGroup>

            {!isLoading && !list?.length && <Empty />}

            <div className="d-flex justify-content-center">
              <Pagination
                currentPage={Number(page)}
                pageSize={20}
                totalSize={count}
              />
            </div>
          </Col>
          <Col xxl={3} lg={4} sm={12} className="mt-5 mt-lg-0">
            <Tips />
          </Col>
        </Row>
      </Container>
    </>
  );
};

export default Index;
