import { FC } from 'react';
import {
  ButtonGroup,
  Button,
  Form,
  Table,
  Stack,
  Badge,
} from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import {
  FormatTime,
  Icon,
  Pagination,
  Modal,
  BaseUserCard,
  Empty,
} from '@answer/components';
import { ADMIN_LIST_STATUS } from '@answer/common/constants';
import { useEditStatusModal } from '@answer/hooks';
import { useAnswerSearch, changeAnswerStatus } from '@answer/api';
import * as Type from '@answer/common/interface';

import '../index.scss';

const answerFilterItems: Type.AdminContentsFilterBy[] = ['normal', 'deleted'];

const Answers: FC = () => {
  const [urlSearchParams, setUrlSearchParams] = useSearchParams();
  const curFilter = urlSearchParams.get('status') || answerFilterItems[0];
  const PAGE_SIZE = 20;
  const curPage = Number(urlSearchParams.get('page')) || 1;
  const { t } = useTranslation('translation', { keyPrefix: 'admin.answers' });

  const {
    data: listData,
    isLoading,
    mutate: refreshList,
  } = useAnswerSearch({
    page_size: PAGE_SIZE,
    page: curPage,
    status: curFilter as Type.AdminContentsFilterBy,
  });
  const count = listData?.count || 0;
  const onFilterChange = (filter) => {
    if (filter === curFilter) {
      return;
    }
    urlSearchParams.set('page', '1');
    urlSearchParams.set('status', filter);
    setUrlSearchParams(urlSearchParams);
  };

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
          item.adopted === 2
            ? t('answer_accepted', { keyPrefix: 'delete' })
            : `<p>${t('other', { keyPrefix: 'delete' })}</p>`,
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

  return (
    <>
      <h3 className="mb-4">{t('page_title')}</h3>
      <div className="d-flex justify-content-between align-items-center mb-3">
        <ButtonGroup size="sm">
          {answerFilterItems.map((li) => {
            return (
              <Button
                key={li}
                size="sm"
                className="text-capitalize"
                onClick={() => onFilterChange(li)}
                variant={curFilter === li ? 'secondary' : 'outline-secondary'}>
                {t(li)}
              </Button>
            );
          })}
        </ButtonGroup>
        <Form.Control
          size="sm"
          type="input"
          placeholder="Filter by title"
          className="d-none"
          style={{ width: '12.25rem' }}
        />
      </div>
      <Table>
        <thead>
          <tr>
            <th style={{ width: '45%' }}>{t('post')}</th>
            <th>{t('votes')}</th>
            <th style={{ width: '20%' }}>{t('created')}</th>
            <th>{t('status')}</th>
            {curFilter !== 'deleted' && <th>{t('action')}</th>}
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
                        href={`/questions/${li.question_id}/${li.id}`}
                        target="_blank"
                        className="text-break text-wrap"
                        rel="noreferrer">
                        {li.question_info.title}
                      </a>
                      {li.adopted === 2 && (
                        <Icon
                          name="check-circle-fill"
                          className="ms-2 text-success"
                        />
                      )}
                    </Stack>
                    <div
                      dangerouslySetInnerHTML={{
                        __html: li.description,
                      }}
                      className="last-p text-truncate-2 fs-14"
                    />
                  </Stack>
                </td>
                <td>{li.vote_count}</td>
                <td>
                  <Stack>
                    <BaseUserCard data={li.user_info} />

                    <FormatTime
                      className="fs-14 text-secondary"
                      time={li.create_time}
                    />
                  </Stack>
                </td>
                <td>
                  <Badge bg={ADMIN_LIST_STATUS[curFilter]?.variant}>
                    {t(ADMIN_LIST_STATUS[curFilter]?.name)}
                  </Badge>
                </td>
                {curFilter !== 'deleted' && (
                  <td>
                    <Button variant="link" onClick={() => handleChange(li.id)}>
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
