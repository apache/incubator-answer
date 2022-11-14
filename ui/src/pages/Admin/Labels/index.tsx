import { FC, useState } from 'react';
import { Button, Table } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { Pagination, BaseUserCard, Empty } from '@/components';
import * as Type from '@/common/interface';
import { useChangeModal } from '@/hooks';
import { useQueryUsers } from '@/services';

import CreateForm from './Form';

import '../index.scss';

const UserFilterKeys: Type.UserFilterBy[] = [
  'all',
  'inactive',
  'suspended',
  'deleted',
];

const PAGE_SIZE = 10;
const Users: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.labels' });

  const [urlSearchParams] = useSearchParams();
  const curFilter = urlSearchParams.get('filter') || UserFilterKeys[0];
  const curPage = Number(urlSearchParams.get('page') || '1');
  const curQuery = urlSearchParams.get('query') || '';

  const [isCreate, setCreateState] = useState(true);
  const {
    data,
    isLoading,
    mutate: refreshUsers,
  } = useQueryUsers({
    page: curPage,
    page_size: PAGE_SIZE,
    query: curQuery,
    ...(curFilter === 'all' ? {} : { status: curFilter }),
  });
  const changeModal = useChangeModal({
    callback: refreshUsers,
  });

  const handleClick = ({ user_id, status }) => {
    changeModal.onShow({
      id: user_id,
      type: status,
    });
  };

  if (isCreate) {
    return <CreateForm onClose={() => setCreateState(false)} />;
  }

  return (
    <>
      <h3 className="mb-4">{t('title')}</h3>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <Button
          variant="outline-secondary"
          size="sm"
          onClick={() => setCreateState(true)}>
          {t('new_label')}
        </Button>
      </div>
      <Table>
        <thead>
          <tr>
            <th>{t('name')}</th>
            <th style={{ width: '12%' }}>{t('color')}</th>
            <th style={{ width: '20%' }}>{t('description')}</th>

            {curFilter !== 'deleted' ? (
              <th style={{ width: '10%' }}>{t('action')}</th>
            ) : null}
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
                    avatarSearchStr="s=48"
                  />
                </td>
                <td>{user.rank}</td>
                <td className="text-break">{user.e_mail}</td>

                {curFilter !== 'deleted' ? (
                  <td>
                    {user.status !== 'deleted' && (
                      <Button
                        className="p-0 btn-no-border"
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
