import { FC, memo, useEffect, useState } from 'react';
import { Button, OverlayTrigger, Tooltip } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { Icon } from '@/components';
import { queryReactions, updateReaction } from '@/services';
import { tryNormalLogged } from '@/utils/guard';

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
  const [reactions, setReactions] = useState<Record<string, string[]>>();
  const { t } = useTranslation('translation');

  useEffect(() => {
    queryReactions(objectId).then((res) => {
      setReactions(res);
    });
  }, []);

  const handleSubmit = (params: { object_id: string; emoji: string }) => {
    if (!tryNormalLogged(true)) {
      return;
    }
    updateReaction({ ...params, type: 'toggle' }).then((res) => {
      setReactions(res);
    });
  };

  const convertToTooltip = (names: string[]) => {
    const n: number = Math.min(5, names.length);
    let ret = '';
    for (let i = 0; i < n; i += 1) {
      if (i === n - 1) {
        ret += names[i];
      } else {
        ret += `${names[i]}, `;
      }
    }
    if (names.length > 5) {
      ret += t('reaction.tooltip', { count: names.length - 5 });
    }
    return ret;
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
    <div className="d-block d-md-flex flex-wrap mt-4 mb-3">
      {showAddCommentBtn && (
        <Button
          className="rounded-pill btn-no-border answer-reaction-btn bg-light"
          size="sm"
          onClick={handleClickComment}>
          <Icon name="chat-text-fill" />
          <span className="ms-1">{t('comment.btn_add_comment')}</span>
        </Button>
      )}

      <OverlayTrigger trigger="click" placement="top" overlay={renderTooltip}>
        <Button
          size="sm"
          className="rounded-pill ms-2 answer-reaction-btn bg-light btn-no-border">
          <Icon name="emoji-smile-fill" />
          <span className="ms-1">+</span>
        </Button>
      </OverlayTrigger>

      {reactions &&
        emojiMap.map((emoji) => {
          if (!reactions[emoji.name] || reactions[emoji.name].length === 0) {
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
                    {convertToTooltip(reactions[emoji.name])}
                  </div>
                </Tooltip>
              }>
              <Button
                title="hahah"
                className="rounded-pill ms-2 answer-reaction-btn bg-light btn-no-border"
                size="sm"
                onClick={() =>
                  handleSubmit({ object_id: objectId, emoji: emoji.name })
                }>
                <Icon name={emoji.icon} className={emoji.className} />
                <span className="ms-1">{reactions[emoji.name].length}</span>
              </Button>
            </OverlayTrigger>
          );
        })}
    </div>
  );
};

export default memo(Index);
