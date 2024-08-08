/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import { Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link, useParams, useSearchParams } from 'react-router-dom';

// import classnames from 'classnames';

import { Avatar, FormatTime, Pagination, Empty } from '@/components';
import { usePageTags, useSkeletonControl } from '@/hooks';
// import { formatCount } from '@/utils';
import { useGetBadgeInfo, useBadgeDetailList } from '@/services';

import BadgeDetail from './components/Badge';
import Loader from './components/Loader';
import HeaderLoader from './components/HeaderLoader';

const Index = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'badges' });

  const { badge_id = '' } = useParams();
  const [urlSearchParams] = useSearchParams();

  const page = Number(urlSearchParams.get('page')) || 1;
  const pageSize = 30;
  const { data: badgeInfo, isLoading: isHeaderLoading } =
    useGetBadgeInfo(badge_id);
  const { data: badges, isLoading: isDetailLoading } = useBadgeDetailList({
    badge_id,
    page,
    page_size: pageSize,
  });

  const { isSkeletonShow } = useSkeletonControl(isDetailLoading);

  usePageTags({
    title: badgeInfo?.name || '',
  });

  if (badgeInfo === undefined) {
    return null;
  }

  console.log(badges);

  return (
    <div className="pt-4 mb-5">
      <h3 className="mb-4">{t('title')}</h3>
      {isHeaderLoading ? <HeaderLoader /> : <BadgeDetail data={badgeInfo} />}
      <Row>
        <Loader />
        {isSkeletonShow ? (
          <Loader />
        ) : (
          badges?.list?.map((item) => {
            return (
              <Col sm={12} md={6} lg={3} key={item.id} className="mb-4">
                <FormatTime
                  time={1722397094672}
                  preFix={t('awarded')}
                  className="small mb-1 d-block"
                />
                <div className="d-flex align-items-center">
                  <Link to="/user">
                    <Avatar size="40px" avatar="" alt="" />
                  </Link>
                  <div className="small ms-2">
                    <Link
                      to="/user"
                      className="lh-1 name-ellipsis"
                      style={{ maxWidth: '200px' }}>
                      username
                    </Link>
                    <div className="text-secondary">
                      980 {t('x_reputation', { keyPrefix: 'personal' })}
                    </div>
                  </div>
                </div>
                <Link to="/question" className="mt-1 d-block">
                  How to `go test` all tests in my project?
                </Link>
              </Col>
            );
          })
        )}
      </Row>
      {Number(badges?.count) <= 0 && !isDetailLoading && <Empty />}
      <div className="d-flex justify-content-center">
        <Pagination
          currentPage={page}
          pageSize={pageSize}
          totalSize={badges?.count || 0}
        />
      </div>
    </div>
  );
};

export default Index;
