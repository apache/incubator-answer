import { FC } from 'react';
import { Button, Form, Table, Stack } from 'react-bootstrap';
import { Link, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import {
  FormatTime,
  Icon,
  Pagination,
  Modal,
  BaseUserCard,
  Empty,
  QueryGroup,
} from '@/components';
import { ADMIN_LIST_STATUS } from '@/common/constants';
import { useEditStatusModal, useReportModal } from '@/hooks';
import * as Type from '@/common/interface';
import { useQuestionSearch, changeQuestionStatus } from '@/services';
import { pathFactory } from '@/router/pathFactory';

const questionFilterItems: Type.AdminContentsFilterBy[] = [
  'normal',
  'closed',
  'deleted',
];

const PAGE_SIZE = 20;
const Questions: FC = () => {
  const [urlSearchParams, setUrlSearchParams] = useSearchParams();
  const curFilter = urlSearchParams.get('status') || questionFilterItems[0];
  const curPage = Number(urlSearchParams.get('page')) || 1;
  const curQuery = urlSearchParams.get('query') || '';
  const { t } = useTranslation('translation', { keyPrefix: 'admin.questions' });

  const {
    data: listData,
    isLoading,
    mutate: refreshList,
  } = useQuestionSearch({
    page_size: PAGE_SIZE,
    page: curPage,
    status: curFilter as Type.AdminContentsFilterBy,
    query: curQuery,
  });
  const count = listData?.count || 0;

  const closeModal = useReportModal(refreshList);

  const handleCallback = (id, type) => {
    if (type === 'normal') {
      changeQuestionStatus(id, 'available').then(() => {
        refreshList();
      });
    }
    if (type === 'closed') {
      closeModal.onShow({
        type: 'question',
        id,
        action: 'close',
      });
    }
    if (type === 'deleted') {
      const item = listData?.list?.filter((v) => v.id === id)?.[0];
      Modal.confirm({
        title: t('title', { keyPrefix: 'delete' }),
        content:
          item.answer_count > 0
            ? t('question', { keyPrefix: 'delete' })
            : t('other', { keyPrefix: 'delete' }),
        cancelBtnVariant: 'link',
        confirmBtnVariant: 'danger',
        confirmText: t('delete', { keyPrefix: 'btns' }),
        onConfirm: () => {
          changeQuestionStatus(id, 'deleted').then(() => {
            refreshList();
          });
        },
      });
    }
  };

  const changeModal = useEditStatusModal({
    editType: 'question',
    callback: handleCallback,
  });

  const handleChange = (itemId) => {
    changeModal.onShow({
      id: itemId,
      type: curFilter,
    });
  };

  const handleFilter = (e) => {
    urlSearchParams.set('query', e.target.value);
    urlSearchParams.delete('page');
    setUrlSearchParams(urlSearchParams);
  };
  return (
    <>
      <h3 className="mb-4">{t('page_title')}</h3>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <QueryGroup
          data={questionFilterItems}
          currentSort={curFilter}
          sortKey="status"
          i18nKeyPrefix="admin.questions"
        />

        <Form.Control
          value={curQuery}
          size="sm"
          type="input"
          placeholder={t('filter.placeholder')}
          onChange={handleFilter}
          style={{ width: '12.25rem' }}
        />
      </div>
      <Table>
        <thead>
          <tr>
            <th>{t('post')}</th>
            <th style={{ width: '8%' }}>{t('votes')}</th>
            <th style={{ width: '8%' }}>{t('answers')}</th>
            <th style={{ width: '20%' }}>{t('created')}</th>
            <th style={{ width: '9%' }}>{t('status')}</th>
            {curFilter !== 'deleted' && (
              <th style={{ width: '10%' }}>{t('action')}</th>
            )}
          </tr>
        </thead>
        <tbody className="align-middle">
          {listData?.list?.map((li) => {
            return (
              <tr key={li.id}>
                <td>
                  <a
                    href={pathFactory.questionLanding(li.id, li.url_title)}
                    target="_blank"
                    className="text-break text-wrap"
                    rel="noreferrer">
                    {li.title}
                  </a>
                  {li.accepted_answer_id > 0 && (
                    <Icon
                      name="check-circle-fill"
                      className="ms-2 text-success"
                    />
                  )}
                </td>
                <td>{li.vote_count}</td>
                <td>
                  <Link
                    to={`/admin/answers?questionId=${li.id}`}
                    rel="noreferrer">
                    {li.answer_count}
                  </Link>
                </td>
                <td>
                  <Stack>
                    <BaseUserCard data={li.user_info} />
                    <FormatTime
                      className="small text-secondary"
                      time={li.create_time}
                    />
                  </Stack>
                </td>
                <td>
                  <span
                    className={classNames(
                      'badge',
                      ADMIN_LIST_STATUS[curFilter]?.variant,
                    )}>
                    {t(ADMIN_LIST_STATUS[curFilter]?.name)}
                  </span>
                </td>
                {curFilter !== 'deleted' && (
                  <td>
                    <Button
                      variant="link"
                      className="p-0 btn-no-border"
                      onClick={() => handleChange(li.id)}>
                      {t('change')}
                    </Button>
                  </td>
                )}
              </tr>
            );
          })}
        </tbody>
      </Table>
      {Number(count) <= 0 && !isLoading && <Empty />}
      <div className="mt-4 mb-2 d-flex justify-content-center">
        <Pagination
          currentPage={curPage}
          totalSize={count}
          pageSize={PAGE_SIZE}
        />
      </div>
    </>
  );
};

export default Questions;
