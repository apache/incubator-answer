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
import { Card, Badge } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classnames from 'classnames';

import * as Type from '@/common/interface';
import { Icon } from '@/components';
import { formatCount } from '@/utils';

interface IProps {
  data: Type.BadgeInfo;
}

const Index: FC<IProps> = ({ data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'badges' });

  if (!data?.id) {
    return null;
  }

  return (
    <Card className="mb-4">
      <Card.Body className="d-flex">
        {data.icon?.startsWith('http') ? (
          <img
            src={data.icon}
            width={96}
            height={96}
            alt={data.name}
            className="me-3"
          />
        ) : (
          <Icon
            name={data?.icon}
            size="96px"
            className={classnames(
              'lh-1 me-3',
              data?.level === 1 && 'bronze',
              data?.level === 2 && 'silver',
              data?.level === 3 && 'gold',
            )}
          />
        )}
        <div>
          <h5>{data.name}</h5>
          <div dangerouslySetInnerHTML={{ __html: data.description || '' }} />

          {!data.is_single && (
            <div className="mt-2">{t('can_earn_multiple')}</div>
          )}

          {(data.award_count > 0 || data.earned_count > 0) && (
            <div className="small mt-2">
              {data.award_count > 0 && (
                <span className="text-secondary me-2">
                  {t('×_awarded', { number: formatCount(data.award_count) })}
                </span>
              )}

              {data.earned_count > 1 && (
                <Badge bg="success">
                  {t('earned_×', { number: data.earned_count })}
                </Badge>
              )}
              {data.earned_count === 1 && (
                <Badge bg="success">{t('earned')}</Badge>
              )}
            </div>
          )}
        </div>
      </Card.Body>
    </Card>
  );
};

export default Index;
