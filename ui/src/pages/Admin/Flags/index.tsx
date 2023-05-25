import React, { FC } from 'react';
import { Button, Form, Table, Stack } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import {
  FormatTime,
  BaseUserCard,
  Empty,
  Pagination,
  QueryGroup,
} from '@/components';
import { useReportModal } from '@/hooks';
import * as Type from '@/common/interface';
import { useFlagSearch } from '@/services';
import { escapeRemove } from '@/utils';
import { pathFactory } from '@/router/pathFactory';

const flagFilterKeys: Type.FlagStatus[] = ['pending', 'completed'];
const flagTypeKeys: Type.FlagType[] = ['all', 'question', 'answer', 'comment'];

const Flags: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.flags' });
  const [urlSearchParams, setUrlSearchParams] = useSearchParams();
  const curFilter = urlSearchParams.get('status') || flagFilterKeys[0];
  const curType = urlSearchParams.get('type') || flagTypeKeys[0];
  const PAGE_SIZE = 20;
  const curPage = Number(urlSearchParams.get('page')) || 1;
  const {
    data: listData,
    isLoading,
    mutate: refreshList,
  } = useFlagSearch({
    page_size: PAGE_SIZE,
    page: curPage,
    status: curFilter as Type.FlagStatus,
    object_type: curType as Type.FlagType,
  });
  const reportModal = useReportModal(refreshList);

  const count = listData?.count || 0;

  const onTypeChange = (evt) => {
    urlSearchParams.set('type', evt.target.value);
    setUrlSearchParams(urlSearchParams);
  };

  const handleReview = ({ id, object_type }) => {
    reportModal.onShow({
      id,
      type: object_type,
      isBackend: true,
      action: 'review',
    });
  };

  return (
    <>
      <h3 className="mb-4">{t('title')}</h3>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <QueryGroup
          data={flagFilterKeys}
          currentSort={curFilter}
          sortKey="status"
          i18nKeyPrefix="admin.flags"
        />

        <Form.Select
          value={curType}
          onChange={onTypeChange}
          size="sm"
          style={{ width: '12.25rem' }}>
          {flagTypeKeys.map((li) => {
            return (
              <option value={li} key={li}>
                {t(li, { keyPrefix: 'btns' })}
              </option>
            );
          })}
        </Form.Select>
      </div>
      <Table>
        <thead>
          <tr>
            <th>{t('flagged')}</th>
            <th style={{ width: '20%' }}>{t('created')}</th>
            {curFilter !== 'completed' ? (
              <th style={{ width: '20%' }}>{t('action')}</th>
            ) : null}
          </tr>
        </thead>
        <tbody className="align-middle">
          {listData?.list?.map((li) => {
            return (
              <tr key={li.id}>
                <td>
                  <Stack>
                    <small className="text-secondary">
                      {t('flagged_type', {
                        type: t(li.object_type, { keyPrefix: 'btns' }),
                      })}
                    </small>
                    <BaseUserCard
                      data={li.reported_user}
                      className="mt-2 small"
                    />
                    <a
                      href={pathFactory.questionLanding(
                        li.question_id,
                        li.url_title,
                      )}
                      target="_blank"
                      className="text-wrap text-break mt-2"
                      rel="noreferrer">
                      {li.title}
                    </a>
                    <small className="text-break text-wrap word">
                      {escapeRemove(li.excerpt)}
                    </small>
                  </Stack>
                </td>
                <td>
                  <Stack>
                    <FormatTime
                      time={li.created_at}
                      className="small text-secondary"
                    />
                    <BaseUserCard
                      data={li.report_user}
                      className="mt-2 mb-2 small"
                    />
                    {li.flagged_reason ? (
                      <small>{li.flagged_content}</small>
                    ) : (
                      <small>
                        {li.reason?.name}
                        <br />
                        <span className="text-secondary">{li.content}</span>
                      </small>
                    )}
                  </Stack>
                </td>
                {curFilter !== 'completed' ? (
                  <td>
                    <Button variant="link" onClick={() => handleReview(li)}>
                      {t('review')}
                    </Button>
                  </td>
                ) : null}
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

export default Flags;
