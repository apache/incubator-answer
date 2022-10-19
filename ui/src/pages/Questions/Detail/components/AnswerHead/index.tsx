import { memo, FC } from 'react';
import { useTranslation } from 'react-i18next';

import { QueryGroup } from '@answer/components';

interface Props {
  count: number;
  order: string;
}

const sortBtns = [
  {
    name: 'score',
    sort: 'default',
  },
  {
    name: 'newest',
    sort: 'updated',
  },
];

const Index: FC<Props> = ({ count = 0, order = 'default' }) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'question_detail.answers',
  });

  return (
    <div
      className="d-flex align-items-center justify-content-between mt-5 mb-3"
      id="answerHeader">
      <h5 className="mb-0">
        {count} {t('title')}
      </h5>
      <QueryGroup
        data={sortBtns}
        currentSort={order === 'updated' ? 'newest' : 'score'}
        i18nKeyPrefix="question_detail.answers"
      />
    </div>
  );
};

export default memo(Index);
