import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import { QueryGroup } from '@answer/components';

const sortBtns = ['newest', 'score'];

interface Props {
  tabName: string;
  count: number;
  sort: string;
  visible: boolean;
}
const Index: FC<Props> = ({
  tabName = 'answers',
  visible,
  sort,
  count = 0,
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });

  if (!visible) {
    return null;
  }

  return (
    <div className="d-flex  align-items-center justify-content-between pb-3 border-bottom">
      <h5 className="mb-0">
        {count} {t(tabName)}
      </h5>
      {(tabName === 'answers' || tabName === 'questions') && (
        <QueryGroup
          data={sortBtns}
          currentSort={sort}
          i18nkeyPrefix="personal"
        />
      )}
    </div>
  );
};

export default memo(Index);
