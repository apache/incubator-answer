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

import { FC, memo, useEffect } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Link, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { unsubscribe } from '@/services';
import { usePageTags } from '@/hooks';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'unsubscribe' });
  usePageTags({
    title: t('page_title'),
  });
  const [searchParams] = useSearchParams();
  const code = searchParams.get('code');
  useEffect(() => {
    if (code) {
      unsubscribe(code);
    }
  }, [code]);
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col lg={6}>
          <h3 className="text-center mt-3 mb-5">{t('success_title')}</h3>
          <p className="text-center">{t('success_desc')}</p>
          <div className="text-center">
            <Link to="/users/settings/notify">{t('link')}</Link>
          </div>
        </Col>
      </Row>
    </Container>
  );
};

export default memo(Index);
