import { FC } from 'react';
import { Button, Form, Table, Stack } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
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
import { useEditStatusModal } from '@/hooks';
import * as Type from '@/common/interface';
import { useAnswerSearch, changeAnswerStatus } from '@/services';
import { escapeRemove } from '@/utils';
import { pathFactory } from '@/router/pathFactory';

const answerFilterItems: Type.AdminContentsFilterBy[] = ['normal', 'deleted'];

const Answers: FC = () => {
  const [urlSearchParams, setUrlSearchParams] = useSearchParams();
  const curFilter = urlSearchParams.get('status') || answerFilterItems[0];
  const PAGE_SIZE = 20;
  const curPage = Number(urlSearchParams.get('page')) || 1;
  const curQuery = urlSearchParams.get('query') || '';
  const questionId = urlSearchParams.get('questionId') || '';
  const { t } = useTranslation('translation', { keyPrefix: 'admin.answers' });

  const {
    data: listData,
    isLoading,
    mutate: refreshList,
  } = useAnswerSearch({
    page_size: PAGE_SIZE,
    page: curPage,
    status: curFilter as Type.AdminContentsFilterBy,
    query: curQuery,
    question_id: questionId,
  });
  const count = listData?.count || 0;

  const handleCallback = (id, type) => {
    if (type === 'normal') {
      changeAnswerStatus(id, 'available').then(() => {
        refreshList();
      });
    }
    if (type === 'deleted') {
      const item = listData?.list?.filter((v) => v.id === id)?.[0];
      Modal.confirm({
        title: t('title', { keyPrefix: 'delete' }),
        content:
          item.accepted === 2
            ? t('answer_accepted', { keyPrefix: 'delete' })
            : t('other', { keyPrefix: 'delete' }),
        cancelBtnVariant: 'link',
        confirmBtnVariant: 'danger',
        confirmText: t('delete', { keyPrefix: 'btns' }),
        onConfirm: () => {
          changeAnswerStatus(id, 'deleted').then(() => {
            refreshList();
          });
        },
      });
    }
  };

  const changeModal = useEditStatusModal({
    editType: 'answer',
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
          data={answerFilterItems}
          currentSort={curFilter}
          sortKey="status"
          i18nKeyPrefix="admin.answers"
        />

        <Form.Control
          value={curQuery}
          onChange={handleFilter}
          size="sm"
          type="input"
          placeholder={t('filter.placeholder')}
          style={{ width: '12.25rem' }}
        />
      </div>
      <Table responsive>
        <thead>
          <tr>
            <th>{t('post')}</th>
            <th style={{ width: '11%' }}>{t('votes')}</th>
            <th style={{ width: '14%' }}>{t('created')}</th>
            <th style={{ width: '11%' }}>{t('status')}</th>
            {curFilter !== 'deleted' && (
              <th style={{ width: '11%' }}>{t('action')}</th>
            )}
          </tr>
        </thead>
        <tbody className="align-middle">
          {listData?.list?.map((li) => {
            return (
              <tr key={li.id}>
                <td>
                  <Stack>
                    <Stack direction="horizontal" gap={2}>
                      <a
                        href={pathFactory.answerLanding({
                          questionId: li.question_id,
                          slugTitle: li.question_info.url_title,
                          answerId: li.id,
                        })}
                        target="_blank"
                        className="text-break text-wrap"
                        rel="noreferrer">
                        {li.question_info.title}
                      </a>
                      {li.accepted === 2 && (
                        <Icon
                          name="check-circle-fill"
                          className="ms-2 text-success"
                        />
                      )}
                    </Stack>
                    <div
                      className="text-truncate-2 small"
                      style={{ maxWidth: '30rem' }}>
                      {escapeRemove(li.description)}
                    </div>
                  </Stack>
                </td>
                <td>{li.vote_count}</td>
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

export default Answers;
