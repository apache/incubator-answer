import { FC, memo } from 'react';
import { ListGroupItem } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { QueryGroup } from '@answer/components';

const sortBtns = ['newest', 'active', 'score'];

interface Props {
  count: number;
  sort: string;
}
const Index: FC<Props> = ({ sort, count = 0 }) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'search.sort_btns',
  });

  return (
    <ListGroupItem className="d-flex flex-wrap align-items-center justify-content-between divide-line pb-3 border-bottom px-0">
      <h5 className="mb-0">{t('counts', { count, keyPrefix: 'search' })}</h5>
      <QueryGroup
        data={sortBtns}
        currentSort={sort}
        sortKey="order"
        i18nkeyPrefix="search.sort_btns"
      />
    </ListGroupItem>
  );
};

export default memo(Index);
