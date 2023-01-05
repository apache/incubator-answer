import { memo, FC, useEffect, useRef } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link, useSearchParams } from 'react-router-dom';

import {
  Actions,
  Operate,
  UserCard,
  Icon,
  Comment,
  FormatTime,
  htmlRender,
} from '@/components';
import { scrollTop, bgFadeOut } from '@/utils';
import { AnswerItem } from '@/common/interface';
import { acceptanceAnswer } from '@/services';

interface Props {
  data: AnswerItem;
  /** router answer id */
  aid?: string;
  /** is author */
  isAuthor: boolean;
  questionTitle: string;
  slugTitle: string;
  isLogged: boolean;
  callback: (type: string) => void;
}
const Index: FC<Props> = ({
  aid,
  data,
  isAuthor,
  isLogged,
  questionTitle = '',
  slugTitle,
  callback,
}) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'question_detail',
  });
  const [searchParams] = useSearchParams();
  const answerRef = useRef<HTMLDivElement>(null);
  const acceptAnswer = () => {
    acceptanceAnswer({
      question_id: data.question_id,
      answer_id: data.accepted === 2 ? '0' : data.id,
    }).then(() => {
      callback?.('');
    });
  };

  useEffect(() => {
    if (!answerRef?.current) {
      return;
    }
    if (aid === data.id) {
      setTimeout(() => {
        const element = answerRef.current;
        scrollTop(element);
      }, 100);
    }
    htmlRender(answerRef.current.querySelector('.fmt'));
    if (!searchParams.get('commentId')) {
      bgFadeOut(answerRef.current);
    }
  }, [data.id, answerRef.current]);
  if (!data?.id) {
    return null;
  }
  return (
    <div id={data.id} ref={answerRef} className="answer-item py-4">
      <article
        dangerouslySetInnerHTML={{ __html: data?.html }}
        className="fmt text-break text-wrap"
      />
      <div className="d-flex align-items-center mt-4">
        <Actions
          data={{
            id: data?.id,
            isHate: data?.vote_status === 'vote_down',
            isLike: data?.vote_status === 'vote_up',
            votesCount: data?.vote_count,
            hideCollect: true,
            collected: data?.collected,
            collectCount: 0,
            username: data?.user_info?.username,
          }}
        />

        {data?.accepted === 2 && (
          <Button
            disabled={!isAuthor}
            variant="outline-success"
            className="ms-3 active opacity-100 bg-success text-white"
            onClick={acceptAnswer}>
            <Icon name="check-circle-fill" className="me-2" />
            <span>{t('answers.btn_accepted')}</span>
          </Button>
        )}

        {isAuthor && data.accepted === 1 && (
          <Button
            variant="outline-success"
            className="ms-3"
            onClick={acceptAnswer}>
            <Icon name="check-circle-fill" className="me-2" />
            <span>{t('answers.btn_accept')}</span>
          </Button>
        )}
      </div>

      <div className="d-block d-md-flex flex-wrap mt-4 mb-3">
        <div className="mb-3 mb-md-0 me-4 flex-grow-1">
          <Operate
            qid={data.question_id}
            aid={data.id}
            memberActions={data?.member_actions}
            type="answer"
            isAccepted={data.accepted === 2}
            title={questionTitle}
            slugTitle={slugTitle}
            callback={callback}
          />
        </div>
        <div className="mb-3 mb-md-0 me-4" style={{ minWidth: '196px' }}>
          {data.update_user_info &&
          data.update_user_info?.username !== data.user_info?.username ? (
            <UserCard
              data={data?.update_user_info}
              time={Number(data.update_time)}
              preFix={t('edit')}
              isLogged={isLogged}
              timelinePath={`/posts/${data.question_id}/${data.id}/timeline`}
            />
          ) : isLogged ? (
            <Link to={`/posts/${data.question_id}/${data.id}/timeline`}>
              <FormatTime
                time={Number(data.update_time)}
                preFix={t('edit')}
                className="link-secondary fs-14"
              />
            </Link>
          ) : (
            <FormatTime
              time={Number(data.update_time)}
              preFix={t('edit')}
              className="text-secondary fs-14"
            />
          )}
        </div>
        <div style={{ minWidth: '196px' }}>
          <UserCard
            data={data?.user_info}
            time={Number(data.create_time)}
            preFix={t('answered')}
            isLogged={isLogged}
            timelinePath={`/posts/${data.question_id}/${data.id}/timeline`}
          />
        </div>
      </div>

      <Comment
        objectId={data.id}
        mode="answer"
        commentId={searchParams.get('commentId')}
      />
    </div>
  );
};

export default memo(Index);
