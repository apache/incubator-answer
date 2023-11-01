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

import { memo, FC } from 'react';
import { Card } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';

const Index: FC = () => {
  const { t } = useTranslation();
  return (
    <Card>
      <Card.Header>{t('search.tips.title')}</Card.Header>
      <Card.Body className="small ext-secondary">
        <div className="mb-1">
          <Trans i18nKey="search.tips.tag" components={{ 1: <code /> }} />
        </div>
        <div className="mb-1">
          <Trans i18nKey="search.tips.user" components={{ 1: <code /> }} />
        </div>
        <div className="mb-1">
          <Trans i18nKey="search.tips.answer" components={{ 1: <code /> }} />
        </div>
        <div className="mb-1">
          <Trans i18nKey="search.tips.score" components={{ 1: <code /> }} />
        </div>
        <div className="mb-1">
          <Trans i18nKey="search.tips.question" components={{ 1: <code /> }} />
        </div>
        <div>
          <Trans i18nKey="search.tips.is_answer" components={{ 1: <code /> }} />
        </div>
      </Card.Body>
    </Card>
  );
};

export default memo(Index);
