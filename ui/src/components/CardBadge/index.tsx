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
import { Link } from 'react-router-dom';

import classnames from 'classnames';

import { Icon } from '@/components';
import * as Type from '@/common/interface';
import { formatCount } from '@/utils';

import './index.scss';

interface IProps {
  data: Type.BadgeListItem;
  showAwardedCount?: boolean;
  urlSearchParams?: string;
  badgePillType?: 'earned' | 'count';
}

const Index: FC<IProps> = ({
  data,
  badgePillType = 'earned',
  showAwardedCount = false,
  urlSearchParams,
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'badges' });
  return (
    <Link
      className="card text-center badge-card"
      to={`/badges/${data.id}${urlSearchParams ? `?${urlSearchParams}` : ''}`}>
      <Card.Body>
        {Number(data?.earned_count) > 0 && badgePillType === 'earned' && (
          <Badge
            bg="success"
            style={{ position: 'absolute', top: '1rem', right: '1rem' }}>
            {`${t('earned')}${
              Number(data?.earned_count) > 1 ? ` ×${data.earned_count}` : ''
            }`}
          </Badge>
        )}

        {badgePillType === 'count' && Number(data?.earned_count) > 1 && (
          <Badge
            pill
            bg="secondary"
            style={{ position: 'absolute', top: '1rem', right: '1rem' }}>
            ×{data.earned_count}
          </Badge>
        )}
        {data.icon.startsWith('http') ? (
          <img src={data.icon} width={96} height={96} alt={data.name} />
        ) : (
          <Icon
            name={data.icon}
            size="96px"
            className={classnames(
              'lh-1',
              data.level === 1 && 'bronze',
              data.level === 2 && 'silver',
              data.level === 3 && 'gold',
            )}
          />
        )}

        <h6 className="mb-0 mt-3 text-center">{data.name}</h6>
        {showAwardedCount && (
          <div className="small text-secondary mt-2">
            {t('×_awarded', { number: formatCount(data.award_count) })}
          </div>
        )}
      </Card.Body>
    </Link>
  );
};

export default Index;
