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
import { Container, Row, Col } from 'react-bootstrap';
import { Link, useLocation } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { WelcomeTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'account_result' });
  const location = useLocation();
  usePageTags({
    title: t('account_activation', { keyPrefix: 'page_title' }),
  });
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col lg={6}>
          <WelcomeTitle className="mt-3 mb-5" />
          {location.pathname?.includes('success') && (
            <>
              <p className="text-center">{t('success')}</p>
              <div className="text-center">
                <Link to="/">{t('link')}</Link>
              </div>
            </>
          )}

          {location.pathname?.includes('failed') && (
            <p className="text-center">{t('invalid')}</p>
          )}
        </Col>
      </Row>
    </Container>
  );
};

export default memo(Index);
