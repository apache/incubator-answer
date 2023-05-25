import { FC } from 'react';
import { Table, Dropdown, Stack } from 'react-bootstrap';
import { useSearchParams, useNavigate } from 'react-router-dom';
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

const Users: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'admin.installed_plugins',
  });
  const navigate = useNavigate();
  const [urlSearchParams] = useSearchParams();
  const curFilter =
    urlSearchParams.get('filter') || InstalledPluginsFilterKeys[0];
  const {
    data,
    isLoading,
    mutate: updatePlugins,
  } = useQueryPlugins({
    status: curFilter === 'all' ? undefined : curFilter,
  });
  const emitPluginChange = (type) => {
    window.postMessage({
      msgType: type,
    });
  };
  const handleStatus = (plugin) => {
    updatePluginStatus({
      enabled: !plugin.enabled,
      plugin_slug_name: plugin.slug_name,
    }).then(() => {
      updatePlugins();
      if (plugin.have_config) {
        emitPluginChange('refreshConfigurablePlugins');
      }
    });
  };
  const handleSettings = (plugin) => {
    const url = `/admin/${plugin.slug_name}`;
    navigate(url);
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
            <th>{t('name')}</th>
            <th style={{ width: '17%' }}>{t('version')}</th>
            <th style={{ width: '11%' }}>{t('status')}</th>
            {curFilter !== 'deleted' ? (
              <th style={{ width: '11%' }} className="text-end">
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
                  <div>
                    {plugin.link ? (
                      <a href={plugin.link} target="_blank" rel="noreferrer">
                        {plugin.name}
                      </a>
                    ) : (
                      plugin.name
                    )}
                  </div>
                  <div className="small">{plugin.description}</div>
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
                          <Dropdown.Item onClick={() => handleStatus(plugin)}>
                            {t('deactivate')}
                          </Dropdown.Item>
                        ) : (
                          <Dropdown.Item onClick={() => handleStatus(plugin)}>
                            {t('activate')}
                          </Dropdown.Item>
                        )}
                        {plugin.enabled && plugin.have_config && (
                          <Dropdown.Item onClick={() => handleSettings(plugin)}>
                            {t('settings')}
                          </Dropdown.Item>
                        )}
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
