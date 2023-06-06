import { FC, useState, useEffect } from 'react';
import { Form, Table } from 'react-bootstrap';
import { Link, useParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { pathFactory } from '@/router/pathFactory';
import { loggedUserInfoStore } from '@/stores';
import { getTimelineData } from '@/services';
import { Empty } from '@/components';
import * as Type from '@/common/interface';

import HistoryItem from './components/Item';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'timeline' });
  const { qid = '', aid = '', tid = '' } = useParams();
  const { role_id } = loggedUserInfoStore((state) => state.user);
  const [showVotes, setShowVotes] = useState(false);
  const [isLoading, setLoading] = useState(false);
  const [timelineData, setTimelineData] = useState<Type.TimelineRes>();

  const getPageData = (bol: boolean) => {
    setLoading(true);
    getTimelineData({
      object_id: tid || aid || qid,
      show_vote: bol,
    })
      .then((res) => {
        setTimelineData(res);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleSwitch = (bol: boolean) => {
    setShowVotes(bol);
    getPageData(bol);
  };

  useEffect(() => {
    getPageData(false);
  }, []);

  let linkUrl = '';
  let pageTitle = '';
  if (timelineData?.object_info.object_type === 'question') {
    linkUrl = pathFactory.questionLanding(
      timelineData?.object_info.question_id,
      timelineData?.object_info.url_title,
    );
    pageTitle = `${t('title_for_question')} ${timelineData?.object_info.title}`;
  }

  if (timelineData?.object_info.object_type === 'answer') {
    linkUrl = pathFactory.answerLanding({
      questionId: timelineData?.object_info.question_id,
      slugTitle: timelineData?.object_info.url_title,
      answerId: timelineData?.object_info.answer_id,
    });
    pageTitle = `${t('title_for_answer', {
      title: timelineData?.object_info.title,
      author: timelineData?.object_info.display_name,
    })}`;
  }

  if (timelineData?.object_info.object_type === 'tag') {
    linkUrl = `/tags/${
      timelineData?.object_info.main_tag_slug_name ||
      timelineData?.object_info.title
    }`;
    pageTitle = `${t('title_for_tag')} '${timelineData?.object_info.title}'`;
  }

  const revisionList =
    timelineData?.timeline?.filter((item) => item.revision_id > 0) || [];
  usePageTags({
    title: pageTitle,
  });
  return (
    <div className="py-4 mb-5">
      <h5 className="mb-4">
        {timelineData?.object_info.object_type === 'tag'
          ? t('tag_title')
          : t('title')}{' '}
        <Link to={linkUrl}>{timelineData?.object_info?.title}</Link>
      </h5>
      {timelineData?.object_info.object_type !== 'tag' && (
        <Form.Check
          className="mb-4"
          type="switch"
          id="custom-switch"
          label={t('show_votes')}
          checked={showVotes}
          onChange={(e) => handleSwitch(e.target.checked)}
        />
      )}
      <Table hover>
        <thead>
          <tr>
            <th style={{ width: '20%' }}>{t('datetime')}</th>
            <th style={{ width: '15%' }}>{t('type')}</th>
            <th style={{ width: '19%' }}>{t('by')}</th>
            <th>{t('comment')}</th>
          </tr>
        </thead>
        <tbody>
          {timelineData?.timeline?.map((item) => {
            return (
              <HistoryItem
                data={item}
                objectInfo={timelineData?.object_info}
                key={item.activity_id}
                isAdmin={role_id === 2}
                revisionList={revisionList}
              />
            );
          })}
        </tbody>
      </Table>
      {!isLoading && Number(timelineData?.timeline?.length) <= 0 && (
        <Empty>{t('no_data')}</Empty>
      )}
    </div>
  );
};

export default Index;
