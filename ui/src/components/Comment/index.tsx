import { useState, useEffect, useCallback } from 'react';
import { Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import classNames from 'classnames';
import { unionBy } from 'lodash';

import * as Types from '@/common/interface';
import { Modal } from '@/components';
import { usePageUsers, useReportModal } from '@/hooks';
import {
  matchedUsers,
  parseUserInfo,
  scrollToElementTop,
  bgFadeOut,
} from '@/utils';
import { tryNormalLogged } from '@/utils/guard';
import {
  useQueryComments,
  addComment,
  deleteComment,
  updateComment,
  postVote,
} from '@/services';

import { Form, ActionBar, Reply } from './components';

import './index.scss';

const Comment = ({ objectId, mode, commentId }) => {
  const pageUsers = usePageUsers();
  const [pageIndex, setPageIndex] = useState(0);
  const [visibleComment, setVisibleComment] = useState(false);
  const pageSize = pageIndex === 0 ? 3 : 15;
  const { data, mutate } = useQueryComments({
    object_id: objectId,
    comment_id: commentId,
    page: pageIndex,
    page_size: pageSize,
  });
  const [comments, setComments] = useState<any>([]);

  const reportModal = useReportModal();

  const { t } = useTranslation('translation', { keyPrefix: 'comment' });
  const scrollCallback = useCallback((el, co) => {
    if (pageIndex === 0 && co.comment_id === commentId) {
      setTimeout(() => {
        scrollToElementTop(el);
        bgFadeOut(el);
      }, 100);
    }
  }, []);

  useEffect(() => {
    if (!data) {
      return;
    }
    if (pageIndex === 1 || pageIndex === 0) {
      setComments(data?.list);
    } else {
      setComments([...comments, ...data.list]);
    }
    const user: Types.PageUser[] = [];
    data.list.forEach((item) => {
      user.push({
        id: item.user_id,
        displayName: item.user_display_name,
        userName: item.username,
      });
      user.push({
        id: item.reply_comment_id,
        displayName: item.reply_user_display_name,
        userName: item.username,
      });
    });
    pageUsers.setUsers(user);
  }, [data]);

  const handleReply = (id) => {
    if (!tryNormalLogged(true)) {
      return;
    }
    setComments(
      comments.map((item) => {
        if (item.comment_id === id) {
          item.showReply = !item.showReply;
        }
        return item;
      }),
    );
  };

  const handleEdit = (id) => {
    setComments(
      comments.map((item) => {
        if (item.comment_id === id) {
          item.showEdit = !item.showEdit;
        }
        return item;
      }),
    );
  };

  const handleSendReply = (item) => {
    const users = matchedUsers(item.value);
    const userNames = unionBy(users.map((user) => user.userName));
    const commentMarkDown = parseUserInfo(item.value);

    const params = {
      object_id: objectId,
      original_text: commentMarkDown,
      mention_username_list: userNames,
      ...(item.type === 'reply'
        ? {
            reply_comment_id: item.comment_id,
          }
        : {}),
    };

    if (item.type === 'edit') {
      return updateComment({
        ...params,
        comment_id: item.comment_id,
      }).then((res) => {
        setComments(
          comments.map((comment) => {
            if (comment.comment_id === item.comment_id) {
              comment.showEdit = false;
              comment.parsed_text = res.parsed_text;
              comment.original_text = res.original_text;
            }
            return comment;
          }),
        );
      });
    }
    return addComment(params).then((res) => {
      if (item.type === 'reply') {
        const index = comments.findIndex(
          (comment) => comment.comment_id === item.comment_id,
        );
        comments[index].showReply = false;
        comments.splice(index + 1, 0, res);
        setComments([...comments]);
      } else {
        setComments([
          ...comments.map((comment) => {
            if (comment.comment_id === item.comment_id) {
              comment.showReply = false;
            }
            return comment;
          }),
          res,
        ]);
      }

      setVisibleComment(false);
    });
  };

  const handleDelete = (id) => {
    Modal.confirm({
      title: t('title', { keyPrefix: 'delete' }),
      content: t('other', { keyPrefix: 'delete' }),
      confirmBtnVariant: 'danger',
      confirmText: t('delete', { keyPrefix: 'btns' }),
      onConfirm: () => {
        deleteComment(id).then(() => {
          if (pageIndex === 0) {
            mutate();
          }
          setComments(comments.filter((item) => item.comment_id !== id));
        });
      },
    });
  };

  const handleVote = (id, is_cancel) => {
    if (!tryNormalLogged(true)) {
      return;
    }

    postVote(
      {
        object_id: id,
        is_cancel,
      },
      'up',
    ).then(() => {
      setComments(
        comments.map((item) => {
          if (item.comment_id === id) {
            item.vote_count = is_cancel
              ? item.vote_count - 1
              : item.vote_count + 1;
            item.is_vote = !is_cancel;
          }
          return item;
        }),
      );
    });
  };

  const handleAction = ({ action }, item) => {
    if (!tryNormalLogged(true)) {
      return;
    }
    if (action === 'report') {
      reportModal.onShow({
        id: item.comment_id,
        type: 'comment',
        action: 'flag',
      });
    } else if (action === 'delete') {
      handleDelete(item.comment_id);
    } else if (action === 'edit') {
      handleEdit(item.comment_id);
    }
  };

  const handleCancel = (id) => {
    setComments(
      comments.map((item) => {
        if (item.comment_id === id) {
          item.showReply = false;
          item.showEdit = false;
        }
        return item;
      }),
    );
  };
  return (
    <div className="comments-wrap">
      {comments.map((item, index) => {
        return (
          <div
            key={item.comment_id}
            ref={(el) => {
              scrollCallback(el, item);
            }}
            className={classNames(
              'border-bottom py-2 comment-item',
              index === 0 && 'border-top',
            )}>
            {item.showEdit ? (
              <Form
                className="mt-2"
                value={item.original_text}
                type="edit"
                mode={mode}
                onSendReply={(value) =>
                  handleSendReply({ ...item, value, type: 'edit' })
                }
                onCancel={() => handleCancel(item.comment_id)}
              />
            ) : (
              <div className="d-block">
                {item.reply_user_display_name && (
                  <Link to="." className="fs-14 me-1 text-nowrap">
                    @{item.reply_user_display_name}
                  </Link>
                )}

                <div
                  className="fmt fs-14 text-break text-wrap"
                  dangerouslySetInnerHTML={{ __html: item.parsed_text }}
                />
              </div>
            )}

            {item.showReply ? (
              <Reply
                userName={item.user_display_name}
                mode={mode}
                onSendReply={(value) =>
                  handleSendReply({ ...item, value, type: 'reply' })
                }
                onCancel={() => handleCancel(item.comment_id)}
              />
            ) : null}
            {item.showEdit || item.showReply ? null : (
              <ActionBar
                nickName={item.user_display_name}
                username={item.username}
                createdAt={item.created_at}
                voteCount={item.vote_count}
                isVote={item.is_vote}
                memberActions={item.member_actions}
                userStatus={item.user_status}
                onReply={() => {
                  handleReply(item.comment_id);
                }}
                onAction={(action) => handleAction(action, item)}
                onVote={(e) => {
                  e.preventDefault();
                  handleVote(item.comment_id, item.is_vote);
                }}
              />
            )}
          </div>
        );
      })}

      <div className="mt-2">
        <Button
          variant="link"
          className="p-0 fs-14 btn-no-border"
          onClick={() => {
            if (tryNormalLogged(true)) {
              setVisibleComment(!visibleComment);
            }
          }}>
          {t('btn_add_comment')}
        </Button>
        {data && (pageIndex || 1) < Math.ceil((data?.count || 0) / pageSize) && (
          <Button
            variant="link"
            className="p-0 fs-14 ms-3 btn-no-border"
            onClick={() => {
              setPageIndex(pageIndex + 1);
            }}>
            {t('show_more')}
          </Button>
        )}
      </div>

      {visibleComment && (
        <Form
          mode={mode}
          className="mt-2"
          onSendReply={(value) => handleSendReply({ value, type: 'comment' })}
          onCancel={() => setVisibleComment(!visibleComment)}
        />
      )}
    </div>
  );
};

export default Comment;
