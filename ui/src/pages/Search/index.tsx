import { Row, Col, ListGroup } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useSearchParams } from 'react-router-dom';
import { useEffect, useState } from 'react';

import { usePageTags, useCaptchaModal } from '@/hooks';
import { Pagination } from '@/components';
import { getSearchResult } from '@/services';
import type { SearchParams, SearchRes } from '@/common/interface';

import {
  Head,
  SearchHead,
  SearchItem,
  Tips,
  Empty,
  ListLoader,
} from './components';

const Index = () => {
  const { t } = useTranslation('translation');
  const [searchParams] = useSearchParams();
  const page = searchParams.get('page') || 1;
  const q = searchParams.get('q') || '';
  const order = searchParams.get('order') || 'active';
  const [isLoading, setIsLoading] = useState(false);
  const [data, setData] = useState<SearchRes>({
    count: 0,
    list: [],
    extra: null,
  });
  const { count = 0, list = [], extra = null } = data || {};

  const searchCaptcha = useCaptchaModal('search');

  const doSearch = () => {
    setIsLoading(true);
    const params: SearchParams = {
      q,
      order,
      page: Number(page),
      size: 20,
    };

    const captcha = searchCaptcha.getCaptcha();
    if (captcha?.verify) {
      params.captcha_id = captcha.captcha_id;
      params.captcha_code = captcha.captcha_code;
    }

    getSearchResult(params)
      .then((resp) => {
        searchCaptcha.close();
        setData(resp);
      })
      .catch((err) => {
        if (err.isError) {
          searchCaptcha.handleCaptchaError(err.list);
        }
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  useEffect(() => {
    searchCaptcha.check(() => {
      doSearch();
    });
  }, [q, order, page]);

  let pageTitle = t('search', { keyPrefix: 'page_title' });
  if (q) {
    pageTitle = `${t('posts_containing', { keyPrefix: 'page_title' })} '${q}'`;
  }
  usePageTags({
    title: pageTitle,
  });

  return (
    <Row className="pt-4 mb-5">
      <Col className="page-main flex-auto">
        <Head data={extra} />
        <SearchHead sort={order} count={count} />
        <ListGroup className="rounded-0 mb-5">
          {isLoading ? (
            <ListLoader />
          ) : (
            list?.map((item) => {
              return <SearchItem key={item.object.id} data={item} />;
            })
          )}
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
      <Col className="page-right-side mt-4 mt-xl-0">
        <Tips />
      </Col>
    </Row>
  );
};

export default Index;
