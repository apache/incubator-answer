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
import { Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Outlet } from 'react-router-dom';

import { usePageTags } from '@/hooks';

import Nav from './components/Nav';

import './index.scss';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.profile',
  });

  usePageTags({
    title: t('settings', { keyPrefix: 'page_title' }),
  });
  return (
    <Row className="mt-4 mb-5 pb-5">
      <Col className="settings-nav mb-4">
        <Nav />
      </Col>
      <Col className="settings-main">
        <Outlet />
      </Col>
    </Row>
  );
};

export default memo(Index);
