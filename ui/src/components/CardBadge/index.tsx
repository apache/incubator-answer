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
import { useTranslation } from 'react-i18next';
import { FC } from 'react';
import { Card, Badge } from 'react-bootstrap';

import { formatCount } from '@/utils';

import './index.scss';

interface IProps {
  data: any;
  badgePill: boolean;
}

const Index: FC<IProps> = ({ data, badgePill }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'badges' });
  console.log(data);
  return (
    <Card className="text-center badge-card">
      <Card.Body>
        <Badge pill={badgePill} bg="success" className="label">
          0
        </Badge>
        <img src="" width={96} height={96} alt="" />
        <h6 className="mb-0 mt-3 text-center">Nice Question</h6>
        <div className="small text-secondary">
          {t('x_awarded', { number: formatCount(16) })}
        </div>
      </Card.Body>
    </Card>
  );
};

export default Index;
