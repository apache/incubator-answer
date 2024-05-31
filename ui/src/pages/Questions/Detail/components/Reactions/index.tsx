import { FC, memo, useEffect, useState } from 'react';
import { Button, OverlayTrigger, Tooltip } from 'react-bootstrap';
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
      setReactions(res.reaction_summary);
    });
  }, []);

  const handleSubmit = (params: { object_id: string; emoji: string }) => {
    if (!tryNormalLogged(true)) {
      setReactIsActive(false);
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

  const renderTooltip = (props) => (
    <Tooltip id="reaction-button-tooltip" {...props} bsPrefix="tooltip">
      <div className="d-block d-md-flex flex-wrap m-0 p-0">
        {emojiMap.map((d) => (
          <Button
            key={d.icon}
            variant="light"
            size="sm"
            onClick={() =>
              handleSubmit({ object_id: objectId, emoji: d.name })
            }>
            <Icon name={d.icon} className={d.className} />
          </Button>
        ))}
      </div>
    </Tooltip>
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
        overlay={renderTooltip}
        show={reactIsActive}
        onToggle={(show) => setReactIsActive(show)}>
        <Button
          size="sm"
          active={reactIsActive}
          className="smile-btn rounded-pill link-secondary"
          variant="light">
          <Icon name="emoji-smile-fill" />
          <span className="ms-1">+</span>
        </Button>
      </OverlayTrigger>

      {reactions &&
        emojiMap.map((emoji) => {
          if (!reactions[emoji.name] || reactions[emoji.name].count === 0) {
            return null;
          }
          return (
            <OverlayTrigger
              key={emoji.name}
              placement="top"
              overlay={
                <Tooltip>
                  <div className="text-start">
                    <b>{t(`reaction.${emoji.name}`)}</b> <br />{' '}
                    {reactions[emoji.name].tooltip}
                  </div>
                </Tooltip>
              }>
              <Button
                title={emoji.name}
                className="rounded-pill ms-2 link-secondary align-items-center"
                variant="light"
                size="sm"
                onClick={() =>
                  handleSubmit({ object_id: objectId, emoji: emoji.name })
                }>
                <Icon name={emoji.icon} className={emoji.className} />
                <span className="ms-1 lh-1">{reactions[emoji.name].count}</span>
              </Button>
            </OverlayTrigger>
          );
        })}
    </div>
  );
};

export default memo(Index);
