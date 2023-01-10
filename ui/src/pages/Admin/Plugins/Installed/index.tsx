import { FC } from 'react';
import { Table, Dropdown, Stack } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { Pagination, Empty, QueryGroup, Icon } from '@/components';
import * as Type from '@/common/interface';
import { useQueryUsers } from '@/services';

const InstalledPluginsFilterKeys: Type.InstalledPluginsFilterBy[] = [
  'all',
  'active',
  'inactive',
  'outdated',
];

const bgMap = {
  normal: 'text-bg-success',
  suspended: 'text-bg-danger',
  deleted: 'text-bg-danger',
  inactive: 'text-bg-secondary',
};

const PAGE_SIZE = 10;
const Users: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.installed_plugins',
  });

  const [urlSearchParams] = useSearchParams();
  const curFilter =
    urlSearchParams.get('filter') || InstalledPluginsFilterKeys[0];
  const curPage = Number(urlSearchParams.get('page') || '1');
  const curQuery = urlSearchParams.get('query') || '';
  const { data, isLoading } = useQueryUsers({
    page: curPage,
    page_size: PAGE_SIZE,
    query: curQuery,
    ...(curFilter === 'all'
      ? {}
      : curFilter === 'staff'
      ? { staff: true }
      : { status: curFilter }),
  });

  const handleAction = (type, plugin) => {
    console.log(type, plugin);
  };

  return (
    <>
      <h3 className="mb-4">{t('title')}</h3>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <Stack direction="horizontal" gap={3}>
          <QueryGroup
            data={InstalledPluginsFilterKeys}
            currentSort={curFilter}
            sortKey="filter"
            i18nKeyPrefix="admin.installed_plugins.filter"
          />
        </Stack>
      </div>
      <Table>
        <thead>
          <tr>
            <th style={{ width: '12%' }}>{t('name')}</th>
            <th style={{ width: '20%' }}>{t('version')}</th>
            <th style={{ width: '12%' }}>{t('status')}</th>
            {curFilter !== 'deleted' ? (
              <th style={{ width: '8%' }} className="text-end">
                {t('action')}
              </th>
            ) : null}
          </tr>
        </thead>
        <tbody className="align-middle">
          {data?.list.map((plugin) => {
            return (
              <tr key={plugin.user_id}>
                <td>
                  <div>Twitter Logins</div>
                  <div className="text-muted text-small">
                    Enable login with Twitter
                  </div>
                </td>
                <td className="text-break">{plugin.version}</td>
                <td>
                  <span className={classNames('badge', bgMap[plugin.status])}>
                    {t(`filter.${plugin.status}`)}
                  </span>
                </td>
                {curFilter !== 'deleted' ? (
                  <td className="text-end">
                    <Dropdown>
                      <Dropdown.Toggle variant="link" className="no-toggle">
                        <Icon name="three-dots-vertical" />
                      </Dropdown.Toggle>
                      <Dropdown.Menu>
                        <Dropdown.Item
                          onClick={() => handleAction('deactivate', plugin)}>
                          {t('deactivate')}
                        </Dropdown.Item>
                        <Dropdown.Item
                          onClick={() => handleAction('settings', plugin)}>
                          {t('settings')}
                        </Dropdown.Item>
                      </Dropdown.Menu>
                    </Dropdown>
                  </td>
                ) : null}
              </tr>
            );
          })}
        </tbody>
      </Table>
      {Number(data?.count) <= 0 && !isLoading && <Empty />}
      <div className="mt-4 mb-2 d-flex justify-content-center">
        <Pagination
          currentPage={curPage}
          totalSize={data?.count || 0}
          pageSize={PAGE_SIZE}
        />
      </div>
    </>
  );
};

export default Users;
