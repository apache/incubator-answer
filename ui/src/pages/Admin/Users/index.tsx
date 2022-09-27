import React, { FC, useState } from 'react';
import { ButtonGroup, Button, Form, Table, Badge } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { useQueryUsers } from '@answer/services/question-admin.api';
import {
  Pagination,
  FormatTime,
  BaseUserCard,
  Empty,
} from '@answer/components';
import * as Type from '@answer/services/types';
import { useChangeModal } from '@answer/hooks';

import '../index.scss';

const UserFilterKeys: Type.UserFilterBy[] = [
  'all',
  'inactive',
  'suspended',
  'deleted',
];

const bgMap = {
  normal: 'success',
  suspended: 'danger',
  deleted: 'danger',
  inactive: 'secondary',
};

const PAGE_SIZE = 10;
const Users: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.users' });
  const [userName, setUserName] = useState('');

  const [urlSearchParams, setUrlSearchParams] = useSearchParams();
  const curFilter = urlSearchParams.get('filter') || UserFilterKeys[0];
  const curPage = Number(urlSearchParams.get('page') || '1');
  const {
    data,
    isLoading,
    mutate: refreshUsers,
  } = useQueryUsers({
    page: curPage,
    page_size: PAGE_SIZE,
    ...(userName ? { username: userName } : {}),
    ...(curFilter === 'all' ? {} : { status: curFilter }),
  });
  const changeModal = useChangeModal({
    callback: refreshUsers,
  });

  const onFilterChange = (filter) => {
    if (filter === urlSearchParams.get('filter')) {
      return;
    }
    urlSearchParams.set('page', '1');
    urlSearchParams.set('filter', filter);
    setUrlSearchParams(urlSearchParams);
  };

  const handleClick = ({ user_id, status }) => {
    changeModal.onShow({
      id: user_id,
      type: status,
    });
  };

  return (
    <>
      <h3 className="mb-4">{t('title')}</h3>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <ButtonGroup size="sm">
          {UserFilterKeys.map((k) => {
            return (
              <Button
                key={k}
                size="sm"
                className="text-capitalize"
                onClick={() => onFilterChange(k)}
                variant={curFilter === k ? 'secondary' : 'outline-secondary'}>
                {t(k)}
              </Button>
            );
          })}
        </ButtonGroup>
        <Form.Control
          className="d-none"
          size="sm"
          value={userName}
          onChange={(e) => setUserName(e.target.value)}
          placeholder="Filter by name"
          style={{ width: '12.25rem' }}
        />
      </div>
      <Table>
        <thead>
          <tr>
            <th style={{ width: '30%' }}>{t('name')}</th>
            <th>{t('reputation')}</th>
            <th style={{ width: '20%' }}>{t('email')}</th>
            <th className="text-nowrap" style={{ width: '20%' }}>
              {t('created_at')}
            </th>
            {(curFilter === 'deleted' || curFilter === 'suspended') && (
              <th className="text-nowrap" style={{ width: '15%' }}>
                {curFilter === 'deleted' ? t('delete_at') : t('suspend_at')}
              </th>
            )}

            <th>{t('status')}</th>
            {curFilter !== 'deleted' ? <th>{t('action')}</th> : null}
          </tr>
        </thead>
        <tbody className="align-middle">
          {data?.list.map((user) => {
            return (
              <tr key={user.user_id}>
                <td>
                  <BaseUserCard
                    data={user}
                    className="fs-6"
                    avatarSize="24px"
                  />
                </td>
                <td>{user.rank}</td>
                <td className="text-break">{user.e_mail}</td>
                <td>
                  <FormatTime time={user.created_at} />
                </td>
                {curFilter === 'suspended' && (
                  <td className="text-nowrap">
                    <FormatTime time={user.suspended_at} />
                  </td>
                )}
                {curFilter === 'deleted' && (
                  <td className="text-nowrap">
                    <FormatTime time={user.deleted_at} />
                  </td>
                )}
                <td>
                  <Badge bg={bgMap[user.status]}>{t(user.status)}</Badge>
                </td>
                {curFilter !== 'deleted' ? (
                  <td>
                    {user.status !== 'deleted' && (
                      <Button
                        className="px-2"
                        variant="link"
                        onClick={() => handleClick(user)}>
                        {t('change')}
                      </Button>
                    )}
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
