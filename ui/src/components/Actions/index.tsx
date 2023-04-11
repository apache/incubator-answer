import { memo, FC, useState, useEffect } from 'react';
import { Button, ButtonGroup } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { Icon } from '@/components';
import { loggedUserInfoStore } from '@/stores';
import { useToast } from '@/hooks';
import { tryNormalLogged } from '@/utils/guard';
import { bookmark, postVote } from '@/services';

interface Props {
  className?: string;
  source: 'question' | 'answer';
  data: {
    id: string;
    votesCount: number;
    isLike: boolean;
    isHate: boolean;
    hideCollect?: boolean;
    collected: boolean;
    collectCount: number;
    username: string;
  };
}

const Index: FC<Props> = ({ className, data, source }) => {
  const [votes, setVotes] = useState(0);
  const [like, setLike] = useState(false);
  const [hate, setHated] = useState(false);
  const [bookmarkState, setBookmark] = useState({
    state: data?.collected,
    count: data?.collectCount,
  });
  const { username = '' } = loggedUserInfoStore((state) => state.user);
  const toast = useToast();
  const { t } = useTranslation();
  useEffect(() => {
    if (data) {
      setVotes(data.votesCount);
      setLike(data.isLike);
      setHated(data.isHate);
      setBookmark({
        state: data?.collected,
        count: data?.collectCount,
      });
    }
  }, []);

  const handleVote = (type: 'up' | 'down') => {
    if (!tryNormalLogged(true)) {
      return;
    }

    if (data.username === username) {
      toast.onShow({
        msg: t('cannot_vote_for_self'),
        variant: 'danger',
      });
      return;
    }
    const isCancel = (type === 'up' && like) || (type === 'down' && hate);
    postVote(
      {
        object_id: data?.id,
        is_cancel: isCancel,
      },
      type,
    )
      .then((res) => {
        setVotes(res.votes);
        setLike(res.vote_status === 'vote_up');
        setHated(res.vote_status === 'vote_down');
      })
      .catch((err) => {
        const errMsg = err?.value;
        if (errMsg) {
          toast.onShow({
            msg: errMsg,
            variant: 'danger',
          });
        }
      });
  };

  const handleBookmark = () => {
    if (!tryNormalLogged(true)) {
      return;
    }
    bookmark({
      group_id: '0',
      object_id: data?.id,
    }).then((res) => {
      setBookmark({
        state: res.switch,
        count: res.object_collection_count,
      });
    });
  };

  return (
    <div className={classNames(className)}>
      <ButtonGroup>
        <Button
          title={
            source === 'question'
              ? t('question_detail.question_useful')
              : t('question_detail.answer_useful')
          }
          variant="outline-secondary"
          active={like}
          onClick={() => handleVote('up')}>
          <Icon name="hand-thumbs-up-fill" />
        </Button>
        <Button variant="outline-secondary" className="opacity-100" disabled>
          {votes}
        </Button>
        <Button
          title={
            source === 'question'
              ? t('question_detail.question_un_useful')
              : t('question_detail.answer_un_useful')
          }
          variant="outline-secondary"
          active={hate}
          onClick={() => handleVote('down')}>
          <Icon name="hand-thumbs-down-fill" />
        </Button>
      </ButtonGroup>
      {!data?.hideCollect && (
        <Button
          variant="outline-secondary ms-3"
          active={bookmarkState.state}
          onClick={handleBookmark}>
          <Icon name="bookmark-fill" />
          <span style={{ paddingLeft: '10px' }}>{bookmarkState.count}</span>
        </Button>
      )}
    </div>
  );
};

export default memo(Index);
