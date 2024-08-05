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

import { Card, Badge, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import { Avatar, FormatTime } from '@/components';
import { usePageTags } from '@/hooks';
import { formatCount } from '@/utils';

const Index = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'badges' });

  usePageTags({
    title: t('title'),
  });

  return (
    <div className="pt-4 mb-5">
      <h3 className="mb-4">{t('title')}</h3>
      <Card className="mb-4">
        <Card.Body className="d-flex">
          <img src="" alt="" width={96} height={96} />
          <div>
            <h5>Support Expert</h5>
            <div className="mb-2">
              This badge is granted for achieving the expert level of our
              community support programme. This certifies that the recipient has
              demonstrated a high level of skills and abilities to manage and
              support multiple communities/instances.
            </div>
            <div className="mb-2">{t('can_earn_multiple')}</div>
            <div className="small">
              <span className="text-secondary">
                {t('x_awarded', { number: formatCount(16) })}
              </span>
              <Badge bg="success" className="ms-2">
                {t('earned_x', { number: 2 })}
              </Badge>
            </div>
          </div>
        </Card.Body>
      </Card>
      <Row>
        {[0, 1, 2, 3, 4, 5, 6].map((item) => {
          return (
            <Col sm={12} md={6} lg={3} key={item} className="mb-4">
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
        })}
      </Row>
    </div>
  );
};

export default Index;
