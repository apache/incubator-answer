import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import classname from 'classnames';

import { Icon } from '@/components';

interface Props {
  data: {
    votes: number;
    answers: number;
    views: number;
  };
  showVotes?: boolean;
  showAnswers?: boolean;
  showViews?: boolean;
  showAccepted?: boolean;
  isAccepted?: boolean;
  className?: string;
}
const Index: FC<Props> = ({
  data,
  showVotes = true,
  showAnswers = true,
  showViews = true,
  isAccepted = false,
  showAccepted = false,
  className = '',
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'counts' });

  return (
    <div
      className={classname('d-flex align-items-center flex-wrap', className)}>
      {showVotes && (
        <div className="d-flex align-items-center flex-shrink-0">
          <Icon name="hand-thumbs-up-fill me-1" />
          <span>
            {data.votes} {t('votes')}
          </span>
        </div>
      )}

      {showAccepted && (
        <div className="d-flex align-items-center ms-3 text-success flex-shrink-0">
          <Icon name="check-circle-fill me-1" />
          <span>{t('accepted')}</span>
        </div>
      )}

      {showAnswers && (
        <div
          className={`d-flex align-items-center ms-3 flex-shrink-0 ${
            isAccepted ? 'text-success' : ''
          }`}>
          {isAccepted ? (
            <Icon name="check-circle-fill me-1" />
          ) : (
            <Icon name="chat-square-text-fill me-1" />
          )}
          <span>
            {data.answers} {t('answers')}
          </span>
        </div>
      )}
      {showViews && (
        <span className="summary-stat ms-3 flex-shrink-0">
          <Icon name="eye-fill" />
          <em className="fst-normal ms-1">
            {data.views} {t('views')}
          </em>
        </span>
      )}
    </div>
  );
};

export default memo(Index);
