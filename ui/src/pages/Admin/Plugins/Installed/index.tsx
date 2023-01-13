import { FC } from 'react';
import { Table, Dropdown, Stack } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { Empty, QueryGroup, Icon } from '@/components';
import * as Type from '@/common/interface';
import { useQueryPlugins, updatePluginStatus } from '@/services';

const InstalledPluginsFilterKeys: Type.InstalledPluginsFilterBy[] = [
  'all',
  'active',
  'inactive',
];

const bgMap = {
  active: 'text-bg-success',
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
  const { data, isLoading, mutate } = useQueryPlugins({
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
    if (type === 'deactivate') {
      updatePluginStatus({
        enabled: false,
        plugin_slug_name: plugin.slug_name,
      }).then(() => {
        mutate();
      });
    }
    if (type === 'activate') {
      updatePluginStatus({
        enabled: true,
        plugin_slug_name: plugin.slug_name,
      }).then(() => {
        mutate();
      });
    }
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
          {data?.map((plugin) => {
            return (
              <tr key={plugin.slug_name}>
                <td>
                  <div>{plugin.name}</div>
                  <div className="text-muted text-small">
                    {plugin.description}
                  </div>
                </td>
                <td className="text-break">{plugin.version}</td>
                <td>
                  <span
                    className={classNames(
                      'badge',
                      bgMap[plugin.enabled ? 'active' : 'inactive'],
                    )}>
                    {t(`filter.${plugin.enabled ? 'active' : 'inactive'}`)}
                  </span>
                </td>
                {curFilter !== 'deleted' ? (
                  <td className="text-end">
                    <Dropdown>
                      <Dropdown.Toggle variant="link" className="no-toggle">
                        <Icon name="three-dots-vertical" />
                      </Dropdown.Toggle>
                      <Dropdown.Menu>
                        {plugin.enabled ? (
                          <Dropdown.Item
                            onClick={() => handleAction('deactivate', plugin)}>
                            {t('deactivate')}
                          </Dropdown.Item>
                        ) : (
                          <Dropdown.Item
                            onClick={() => handleAction('activate', plugin)}>
                            {t('activate')}
                          </Dropdown.Item>
                        )}

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
      {Number(data?.length) <= 0 && !isLoading && <Empty />}
    </>
  );
};

export default Users;
