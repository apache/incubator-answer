import { memo, FC, useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { Row, Col, Button } from 'react-bootstrap';

import {
  Tag,
  Actions,
  Operate,
  UserCard,
  Comment,
  FormatTime,
} from '@answer/components';
import { formatCount } from '@answer/utils';
import { following } from '@answer/services/api';

interface Props {
  data: any;
  hasAnswer: boolean;
  initPage: (type: string) => void;
}

const Index: FC<Props> = ({ data, initPage, hasAnswer }) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'question_detail',
  });
  const [followed, setFollowed] = useState(data?.is_followed);

  const handleFollow = (e) => {
    e.preventDefault();
    following({
      object_id: data?.id,
      is_cancel: followed,
    }).then((res) => {
      setFollowed(res.is_followed);
    });
  };

  useEffect(() => {
    if (data) {
      setFollowed(data?.is_followed);
    }
  }, [data]);

  if (!data?.id) {
    return null;
  }
  return (
    <div>
      <h1 className="fs-3 mb-3 text-wrap text-break">
        <Link className="text-body" reloadDocument to={`/questions/${data.id}`}>
          {data.title}
        </Link>
      </h1>
      <div className="d-flex align-items-center fs-14 mb-3 text-secondary">
        <Button
          variant="link"
          size="sm"
          className="p-0 me-3 btn-no-border"
          onClick={(e) => handleFollow(e)}>
          {followed ? 'Following' : 'Follow'}
        </Button>

        <FormatTime
          time={data.create_time}
          preFix={t('Asked')}
          className="me-3"
        />

        <FormatTime
          time={data.update_time}
          preFix={t('update')}
          className="me-3"
        />
        {data?.view_count > 0 && (
          <div>
            {t('Views')} {formatCount(data.view_count)}
          </div>
        )}
      </div>
      <div className="mb-3">
        {data?.tags?.map((item: any) => {
          return (
            <Tag
              className="me-1"
              href={`/tags/${item.main_tag_slug_name || item.slug_name}`}
              key={item.slug_name}>
              {item.slug_name}
            </Tag>
          );
        })}
      </div>
      <article
        dangerouslySetInnerHTML={{ __html: data?.html }}
        className="fmt text-break text-wrap"
      />

      <Actions
        className="mt-4"
        data={{
          id: data?.id,
          isHate: data?.vote_status === 'vote_down',
          isLike: data?.vote_status === 'vote_up',
          votesCount: data?.vote_count,
          collected: data?.collected,
          collectCount: data?.collection_count,
          username: data.user_info?.username,
        }}
      />

      <Row className="mt-4 mb-3">
        <Col lg={5}>
          <Operate
            qid={data?.id}
            type="question"
            memberActions={data?.member_actions}
            title={data.title}
            hasAnswer={hasAnswer}
            isAccepted={Boolean(data?.accepted_answer_id)}
            callback={initPage}
          />
        </Col>
        <Col lg={3}>
          {data.update_user_info?.username !== data.user_info?.username ? (
            <UserCard
              data={data?.user_info}
              time={data.edit_time}
              preFix={t('edit')}
            />
          ) : (
            <FormatTime
              time={data.edit_time}
              preFix={t('edit')}
              className="text-secondary fs-14"
            />
          )}
        </Col>
        <Col lg={3}>
          <UserCard
            data={data?.user_info}
            time={data.create_time}
            preFix={t('asked')}
          />
        </Col>
      </Row>

      <Comment objectId={data?.id} mode="question" />
    </div>
  );
};

export default memo(Index);
