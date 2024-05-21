/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import { memo, FC, useState, useEffect } from 'react';
import { Button, ButtonGroup } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { Icon } from '@/components';
import { loggedUserInfoStore } from '@/stores';
import { useToast } from '@/hooks';
import { useCaptchaPlugin } from '@/utils/pluginKit';
import { tryNormalLogged } from '@/utils/guard';
import { bookmark, postVote } from '@/services';
import * as Types from '@/common/interface';

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
  const vCaptcha = useCaptchaPlugin('vote');

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

  const submitVote = (type) => {
    const isCancel = (type === 'up' && like) || (type === 'down' && hate);
    const imgCode: Types.ImgCodeReq = {
      captcha_id: undefined,
      captcha_code: undefined,
    };
    vCaptcha?.resolveCaptchaReq?.(imgCode);

    postVote(
      {
        object_id: data?.id,
        is_cancel: isCancel,
        ...imgCode,
      },
      type,
    )
      .then(async (res) => {
        await vCaptcha?.close();
        setVotes(res.votes);
        setLike(res.vote_status === 'vote_up');
        setHated(res.vote_status === 'vote_down');
      })
      .catch((err) => {
        if (err?.isError) {
          vCaptcha?.handleCaptchaError(err.list);
        }
        const errMsg = err?.value;
        if (errMsg) {
          toast.onShow({
            msg: errMsg,
            variant: 'danger',
          });
        }
      });
  };

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

    if (!vCaptcha) {
      submitVote(type);
      return;
    }

    vCaptcha.check(() => {
      submitVote(type);
    });
  };

  const handleBookmark = () => {
    if (!tryNormalLogged(true)) {
      return;
    }
    bookmark({
      group_id: '0',
      object_id: data?.id,
      bookmark: !bookmarkState.state,
    }).then((res) => {
      setBookmark({
        state: !bookmarkState.state,
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
          title={t('question_detail.question_bookmark')}
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
