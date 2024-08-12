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

import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';
import { Row, Col } from 'react-bootstrap';

// import * as Type from '@/common/interface';
import { CardBadge } from '@/components';
import { useGetRecentAwardBadges } from '@/services';
import TopList from '../TopList';

interface Props {
  username: string;
  visible: boolean;
  introduction: string;
  data;
}
const Index: FC<Props> = ({ visible, introduction, data, username }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });
  const { data: recentBadges } = useGetRecentAwardBadges(
    visible ? username : null,
  );
  if (!visible) {
    return null;
  }
  return (
    <div>
      <h5 className="mb-3">{t('about_me')}</h5>
      {introduction ? (
        <div
          className="mb-5 text-break fmt"
          dangerouslySetInnerHTML={{ __html: introduction }}
        />
      ) : (
        <div className="text-center py-5 mb-5">{t('about_me_empty')}</div>
      )}

      <Row className="mb-4">
        <Col sm={12} md={6} className="mb-4">
          {data?.answer?.length > 0 && (
            <>
              <h5 className="mb-3">{t('top_answers')}</h5>
              <TopList data={data?.answer} type="answer" />
            </>
          )}
        </Col>
        <Col sm={12} md={6}>
          {data?.question?.length > 0 && (
            <>
              <h5 className="mb-3">{t('top_questions')}</h5>
              <TopList data={data?.question} type="question" />
            </>
          )}
        </Col>
      </Row>

      {Number(recentBadges?.count) > 0 && (
        <div className="mb-5">
          <h5 className="mb-3">{t('recent_badges')}</h5>
          <div className="d-flex flex-wrap" style={{ margin: '-12px' }}>
            {recentBadges?.list?.map((item) => {
              return (
                <CardBadge
                  data={item}
                  urlSearchParams={`username=${username}`}
                  badgePillType="count"
                />
              );
            })}
          </div>
        </div>
      )}
    </div>
  );
};

export default memo(Index);
