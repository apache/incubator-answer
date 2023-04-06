import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import { QueryGroup } from '@/components';

const sortBtns = ['active', 'newest', 'relevance', 'score'];

interface Props {
  count: number;
  sort: string;
}
const Index: FC<Props> = ({ sort, count = 0 }) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'search.sort_btns',
  });

  return (
    <div className="d-flex flex-wrap align-items-center justify-content-between pt-2 pb-3">
      <h5 className="mb-0">{t('counts', { count, keyPrefix: 'search' })}</h5>
      <QueryGroup
        data={sortBtns}
        currentSort={sort}
        sortKey="order"
        i18nKeyPrefix="search.sort_btns"
      />
    </div>
  );
};

export default memo(Index);
