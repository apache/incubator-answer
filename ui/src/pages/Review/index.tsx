import { FC, useEffect, useState } from 'react';
import { Row, Col, Alert, Stack, Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { BaseUserCard, FormatTime, Empty, DiffContent } from '@/components';
import { getReviewList, revisionAudit } from '@/services';
import { pathFactory } from '@/router/pathFactory';
import { scrollToDocTop } from '@/utils';
import type * as Type from '@/common/interface';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_review' });
  const [isLoading, setIsLoading] = useState(false);
  const [noTasks, setNoTasks] = useState(false);
  const [page, setPage] = useState(1);
  const [reviewResp, setReviewResp] = useState<Type.ReviewResp>();
  const ro = reviewResp?.list[0];
  const { info, type, unreviewed_info } = ro || {
    info: null,
    type: '',
    unreviewed_info: null,
  };
  const resolveNextOne = (resp, pageNumber) => {
    const { count, list = [] } = resp;
    // auto rollback
    if (!list.length && count && page !== 1) {
      pageNumber = 1;
      setPage(pageNumber);
      // eslint-disable-next-line @typescript-eslint/no-use-before-define
      queryNextOne(pageNumber);
      return;
    }
    if (pageNumber !== page) {
      setPage(pageNumber);
    }
    setReviewResp(resp);
    if (!list.length) {
      setNoTasks(true);
    }
    setTimeout(() => {
      scrollToDocTop();
    }, 150);
  };
  const queryNextOne = (pageNumber) => {
    getReviewList(pageNumber)
      .then((resp) => {
        resolveNextOne(resp, pageNumber);
      })
      .catch((ex) => {
        console.log('ex: ', ex);
      });
  };
  const reviewInfo = unreviewed_info?.content;
  const handlingSkip = () => {
    queryNextOne(page + 1);
  };
  const handlingApprove = () => {
    if (!unreviewed_info) {
      return;
    }
    setIsLoading(true);
    revisionAudit(unreviewed_info.id, 'approve')
      .then(() => {
        queryNextOne(page);
      })
      .catch((ex) => {
        console.log('ex: ', ex);
      })
      .finally(() => {
        setIsLoading(false);
      });
  };
  const handlingReject = () => {
    if (!unreviewed_info) {
      return;
    }
    setIsLoading(true);
    revisionAudit(unreviewed_info.id, 'reject')
      .then(() => {
        queryNextOne(page);
      })
      .catch((ex) => {
        console.log('ex: ', ex);
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  let itemLink = '';
  let itemTitle = '';
  let editBadge = '';
  let editSummary = unreviewed_info?.reason;
  const editor = unreviewed_info?.user_info;
  const editTime = unreviewed_info?.create_at;
  if (type === 'question') {
    itemLink = pathFactory.questionLanding(info?.object_id, info?.url_title);
    itemTitle = info?.title;
    editBadge = t('question_edit');
    editSummary ||= t('edit_question');
  } else if (type === 'answer') {
    itemLink = pathFactory.answerLanding({
      // @ts-ignore
      questionId: unreviewed_info.content.question_id,
      slugTitle: info?.url_title,
      answerId: unreviewed_info.object_id,
    });
    itemTitle = info?.title;
    editBadge = t('answer_edit');
    editSummary ||= t('edit_answer');
  } else if (type === 'tag') {
    const tagInfo = unreviewed_info.content as Type.Tag;
    itemLink = pathFactory.tagLanding(tagInfo.slug_name);
    itemTitle = tagInfo.display_name;
    editBadge = t('tag_edit');
    editSummary ||= t('edit_tag');
  }
  useEffect(() => {
    queryNextOne(page);
  }, []);
  usePageTags({
    title: t('review'),
  });
  return (
    <Row className="pt-4 mb-5">
      <h3 className="mb-4">{t('review')}</h3>
      <Col className="page-main flex-auto">
        {!noTasks && ro && (
          <>
            <Alert variant="secondary">
              <Stack className="align-items-start">
                <span className="badge text-bg-secondary mb-2">
                  {editBadge}
                </span>
                <Link to={itemLink} target="_blank">
                  {itemTitle}
                </Link>
                <p className="mb-0">
                  {t('edit_summary')}: {editSummary}
                </p>
              </Stack>
              <Stack
                direction="horizontal"
                gap={1}
                className="align-items-baseline mt-2">
                <BaseUserCard data={editor} avatarSize="24" />
                {editTime && (
                  <FormatTime
                    time={editTime}
                    className="small text-secondary"
                    preFix={t('proposed')}
                  />
                )}
              </Stack>
            </Alert>
            {type === 'question' &&
              info &&
              reviewInfo &&
              'content' in reviewInfo && (
                <DiffContent
                  className="mt-2"
                  objectType={type}
                  oldData={{
                    title: info.title,
                    original_text: info.content,
                    tags: info.tags,
                  }}
                  newData={{
                    title: reviewInfo.title,
                    original_text: reviewInfo.content,
                    tags: reviewInfo.tags,
                  }}
                />
              )}
            {type === 'answer' &&
              info &&
              reviewInfo &&
              'content' in reviewInfo && (
                <DiffContent
                  className="mt-2"
                  objectType={type}
                  newData={{
                    original_text: reviewInfo.content,
                  }}
                  oldData={{
                    original_text: info.content,
                  }}
                />
              )}
            {type === 'tag' && info && reviewInfo && (
              <DiffContent
                className="mt-2"
                objectType={type}
                newData={{
                  original_text: reviewInfo.original_text,
                }}
                oldData={{
                  original_text: info.content,
                }}
                opts={{ showTitle: false, showTagUrlSlug: false }}
              />
            )}
            <Stack direction="horizontal" gap={2} className="mt-4">
              <Button
                variant="outline-primary"
                disabled={isLoading}
                onClick={handlingApprove}>
                {t('approve', { keyPrefix: 'btns' })}
              </Button>
              <Button
                variant="outline-primary"
                disabled={isLoading}
                onClick={handlingReject}>
                {t('reject', { keyPrefix: 'btns' })}
              </Button>
              <Button
                variant="outline-primary"
                disabled={isLoading}
                onClick={handlingSkip}>
                {t('skip', { keyPrefix: 'btns' })}
              </Button>
            </Stack>
          </>
        )}
        {noTasks && <Empty>{t('empty')}</Empty>}
      </Col>

      <Col className="page-right-side mt-4 mt-xl-0" />
    </Row>
  );
};

export default Index;
