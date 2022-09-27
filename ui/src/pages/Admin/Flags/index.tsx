import React, { FC } from 'react';
import { ButtonGroup, Button, Form, Table, Stack } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import {
  FormatTime,
  BaseUserCard,
  Empty,
  Pagination,
} from '@answer/components';
import { useReportModal } from '@answer/hooks';
import * as Type from '@answer/services/types';
import { useFlagSearch } from '@answer/services/flag-admin.api';

import '../index.scss';

const flagFilterKeys: Type.FlagStatus[] = ['pending', 'completed'];
const flagTypeKeys: Type.FlagType[] = ['all', 'question', 'answer', 'comment'];

const Flags: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.flags' });
  const [urlSearchParams, setUrlSearchParams] = useSearchParams();
  const curFilter = urlSearchParams.get('status') || flagFilterKeys[0];
  const curType = urlSearchParams.get('type') || flagTypeKeys[0];
  const pageSize = 20;
  const curPage = Number(urlSearchParams.get('page')) || 1;
  const { data: listData, isLoading } = useFlagSearch({
    page_size: pageSize,
    page: curPage,
    status: curFilter as Type.FlagStatus,
    object_type: curType as Type.FlagType,
  });
  const reportModal = useReportModal();

  const count = listData?.count || 0;
  const onFilterChange = (filter) => {
    if (filter === curFilter) {
      return;
    }
    urlSearchParams.set('page', '1');
    urlSearchParams.set('status', filter);
    setUrlSearchParams(urlSearchParams);
  };
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
        <ButtonGroup size="sm">
          {flagFilterKeys.map((k) => {
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
        <Form.Select
          value={curType}
          onChange={onTypeChange}
          size="sm"
          style={{ width: '12.25rem' }}>
          {flagTypeKeys.map((li) => {
            return (
              <option value={li} key={li}>
                {li}
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
            <th style={{ width: '20%' }}>{t('action')}</th>
          </tr>
        </thead>
        <tbody className="align-middle">
          {listData?.list?.map((li) => {
            return (
              <tr key={li.id}>
                <td>
                  <Stack>
                    <small className="text-secondary">
                      Flagged {li.object_type}
                    </small>
                    <BaseUserCard data={li.reported_user} className="mt-2" />
                    <a
                      href={`/questions/${li.question_id}`}
                      target="_blank"
                      className="text-wrap text-break mt-2"
                      rel="noreferrer">
                      {li.title}
                    </a>
                    <small className="text-break text-wrap word">
                      {li.excerpt}
                    </small>
                  </Stack>
                </td>
                <td>
                  <Stack>
                    <FormatTime
                      time={li.created_at}
                      className="fs-14 text-secondary"
                    />
                    <BaseUserCard data={li.report_user} className="mt-2 mb-2" />
                    {li.flaged_reason ? (
                      <small>{li.flaged_content}</small>
                    ) : (
                      <small>
                        {li.reason?.name}
                        <br />
                        <span className="text-secondary">{li.content}</span>
                      </small>
                    )}
                  </Stack>
                </td>
                <td>
                  <Button variant="link" onClick={() => handleReview(li)}>
                    {t('review')}
                  </Button>
                </td>
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
          pageSize={pageSize}
        />
      </div>
    </>
  );
};

export default Flags;
