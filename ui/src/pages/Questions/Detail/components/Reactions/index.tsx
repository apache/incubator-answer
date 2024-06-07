import { FC, memo, useEffect, useState } from 'react';
import { Button, OverlayTrigger, Popover, Tooltip } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';

import { Icon } from '@/components';
import { queryReactions, updateReaction } from '@/services';
import { tryNormalLogged } from '@/utils/guard';
import { ReactionItem } from '@/common/interface';

interface Props {
  objectId: string;
  showAddCommentBtn?: boolean;
  handleClickComment: () => void;
}

const emojiMap = [
  {
    name: 'heart',
    icon: 'heart-fill',
    className: 'text-danger',
  },
  {
    name: 'smile',
    icon: 'emoji-laughing-fill',
    className: 'text-warning',
  },
  {
    name: 'frown',
    icon: 'emoji-frown-fill',
    className: 'text-warning',
  },
];

const Index: FC<Props> = ({
  objectId,
  showAddCommentBtn,
  handleClickComment,
}) => {
  const [reactions, setReactions] = useState<Record<string, ReactionItem>>();
  const [reactIsActive, setReactIsActive] = useState<boolean>(false);
  const { t } = useTranslation('translation');

  useEffect(() => {
    queryReactions(objectId).then((res) => {
      setReactions(res?.reaction_summary);
    });
  }, []);

  const handleSubmit = (params: { object_id: string; emoji: string }) => {
    if (!tryNormalLogged(true)) {
      return;
    }
    updateReaction({
      ...params,
      reaction:
        reactions &&
        reactions[params.emoji] &&
        reactions[params.emoji].is_active
          ? 'deactivate'
          : 'activate',
    }).then((res) => {
      setReactions(res.reaction_summary);
      setReactIsActive(false);
    });
  };

  const renderPopover = (props) => (
    <Popover id="reaction-button-tooltip" {...props}>
      <Popover.Body className="d-block d-md-flex flex-wrap p-1">
        {emojiMap.map((d, index) => (
          <Button
            aria-label={
              reactions?.[d.name]?.is_active
                ? t('reaction.undo_emoji', { emoji: d.name })
                : t(`reaction.${d.name}`)
            }
            key={d.icon}
            variant="light"
            active={reactions?.[d.name]?.is_active}
            className={`${index !== 0 ? 'ms-1' : ''}`}
            size="sm"
            onClick={() =>
              handleSubmit({ object_id: objectId, emoji: d.name })
            }>
            <Icon name={d.icon} className={d.className} />
          </Button>
        ))}
      </Popover.Body>
    </Popover>
  );

  return (
    <div
      className={classNames('d-block d-md-flex flex-wrap', {
        'mb-3': !showAddCommentBtn,
      })}>
      {showAddCommentBtn && (
        <Button
          className="rounded-pill me-2 link-secondary"
          variant="light"
          size="sm"
          onClick={handleClickComment}>
          <Icon name="chat-text-fill" />
          <span className="ms-1">{t('comment.btn_add_comment')}</span>
        </Button>
      )}

      <OverlayTrigger
        trigger="click"
        placement="top"
        overlay={renderPopover}
        show={reactIsActive}
        onToggle={(show) => setReactIsActive(show)}>
        <Button
          size="sm"
          aria-label={t('reaction.btn_label')}
          aria-haspopup="true"
          active={reactIsActive}
          className="smile-btn rounded-pill link-secondary"
          variant="light">
          <Icon name="emoji-smile-fill" />
          <span className="ms-1">+</span>
        </Button>
      </OverlayTrigger>

      {reactions &&
        Object.keys(reactions).map((emoji) => {
          if (!reactions[emoji] || reactions[emoji]?.count <= 0) {
            return null;
          }
          return (
            <OverlayTrigger
              key={emoji}
              placement="top"
              overlay={
                <Tooltip>
                  <div className="text-start">
                    <b>{t(`reaction.${emoji}`)}</b> <br />{' '}
                    {reactions[emoji].tooltip}
                  </div>
                </Tooltip>
              }>
              <Button
                className="rounded-pill ms-2 link-secondary d-flex align-items-center"
                aria-label={
                  reactions?.[emoji]?.is_active
                    ? t('reaction.unreact_emoji', { emoji })
                    : t('reaction.react_emoji', { emoji })
                }
                aria-pressed="true"
                variant="light"
                active={reactions[emoji].is_active}
                size="sm"
                onClick={() => handleSubmit({ object_id: objectId, emoji })}>
                <Icon
                  name={String(emojiMap.find((v) => v.name === emoji)?.icon)}
                  className={emojiMap.find((v) => v.name === emoji)?.className}
                />
                <span className="ms-1 lh-1">{reactions[emoji].count}</span>
              </Button>
            </OverlayTrigger>
          );
        })}
    </div>
  );
};

export default memo(Index);
