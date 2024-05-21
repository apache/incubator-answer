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

import { FC } from 'react';
import { Card, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { formatUptime } from '@/utils';

interface IProps {
  data: Type.AdminDashboard['info'];
}
const SystemInfo: FC<IProps> = ({ data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.dashboard' });

  return (
    <Card className="mb-4">
      <Card.Body>
        <h6 className="mb-3">{t('system_info')}</h6>
        <Row>
          <Col xs={6}>
            <span className="text-secondary me-1">{t('go_version')}</span>
            <strong>{data.go_version}</strong>
          </Col>
          <Col xs={6}>
            <span className="text-secondary me-1">{t('database')}</span>
            <strong>{data.database_version}</strong>
          </Col>
          <Col xs={6}>
            <span className="text-secondary me-1">{t('storage_used')}</span>
            <strong>{data.occupying_storage_space}</strong>
          </Col>
          <Col xs={6}>
            <span className="text-secondary me-1">{t('database_size')}</span>
            <strong>{data.database_size}</strong>
          </Col>
          {data.app_start_time ? (
            <Col xs={6}>
              <span className="text-secondary me-1">{t('uptime')}</span>
              <strong>{formatUptime(data.app_start_time)}</strong>
            </Col>
          ) : null}
        </Row>
      </Card.Body>
    </Card>
  );
};

export default SystemInfo;
