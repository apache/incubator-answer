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
import { Form, Table, Stack } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { QueryGroup } from '@/components';
import * as Type from '@/common/interface';

import Action from './components/Action';

const BadgeFilterKeys: Type.BadgeFilterBy[] = ['all', 'active', 'inactive'];

// const bgMap = {
//   normal: 'text-bg-success',
//   suspended: 'text-bg-danger',
//   deleted: 'text-bg-danger',
//   inactive: 'text-bg-secondary',
// };

const Users: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.badges' });

  const [urlSearchParams, setUrlSearchParams] = useSearchParams();
  const curFilter = urlSearchParams.get('filter') || BadgeFilterKeys[0];
  const curQuery = urlSearchParams.get('query') || '';

  const handleFilter = (e) => {
    urlSearchParams.set('query', e.target.value);
    urlSearchParams.delete('page');
    setUrlSearchParams(urlSearchParams);
  };

  return (
    <>
      <h3 className="mb-4">{t('title')}</h3>
      <div className="d-flex flex-wrap justify-content-between align-items-center mb-3">
        <Stack direction="horizontal" gap={3}>
          <QueryGroup
            data={BadgeFilterKeys}
            currentSort={curFilter}
            sortKey="filter"
            i18nKeyPrefix="admin.badges"
          />
        </Stack>

        <Form.Control
          size="sm"
          type="search"
          value={curQuery}
          onChange={handleFilter}
          placeholder={t('filter.placeholder')}
          style={{ width: '12.25rem' }}
          className="mt-3 mt-sm-0"
        />
      </div>
      <Table responsive="md">
        <thead>
          <tr>
            <th>{t('name')}</th>
            <th>{t('group')}</th>
            <th>{t('awards')}</th>

            <th>{t('status')}</th>

            <th className="text-end">{t('action')}</th>
          </tr>
        </thead>
        <tbody className="align-middle">
          <tr>
            <td className="d-flex align-items-center">
              <img
                src="https://s3-alpha-sig.figma.com/img/b6d9/2c6b/dfa4017fd4654f72c13bfc406377416a?Expires=1723420800&Key-Pair-Id=APKAQ4GOSFWCVNEHN3O4&Signature=lp0b2MuP5yiv7IB4yEuhFk--W9D5CWsud77ftgjtdFrSIPPsIxcnZxz-RyLl40euIysaQLVVWwvYqJP75wLnccCQ1XbzwKfU1sOj3Z52jMTLMZ5PGwYL~dnx0sUJVv3khew7Xe8FiebLTwK4yV62jlW2RYq~HvK3s3RL5z9ZrkSnZUIWOC1nD~RTlsS9K3-hJ9GHwSCA9i0VupM5qHBMgxZDstTy6MO5VnACCCD1865NsKpkLM770wlVXP7XARVl5AhcRFYr0J8VTjccwg3dRHvOUUy4sM0wOqRctX7dgQfp1V-bc49RNcO0CkHifof2hn4oLyaC4fUd6GBrthOpXg__"
                className="rounded-circle bg-white me-2"
                width="32px"
                height="32px"
                alt="badge"
              />
              <div>
                <div className="text-primary">Nice Question</div>
                <div className="text-small">Question score of 10 or more.</div>
              </div>
            </td>

            <td>Community Badges</td>
            <td className="text-primary">200</td>
            <td>Active</td>
            <Action badgeData={{}} />
          </tr>
        </tbody>
      </Table>
      {/* {Number(data?.count) <= 0 && !isLoading && <Empty />} */}
      {/* <div className="mt-4 mb-2 d-flex justify-content-center">
        <Pagination
          currentPage={curPage}
          totalSize={data?.count || 0}
          pageSize={PAGE_SIZE}
        />
      </div> */}
    </>
  );
};

export default Users;
