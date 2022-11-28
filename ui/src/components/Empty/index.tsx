import { FC, memo } from 'react';
import { Trans } from 'react-i18next';

const Index: FC = () => {
  return (
    <div className="text-center py-5">
      <Trans i18nKey="personal.list_empty">
        We couldn't find anything. <br /> Try different or less specific
        keywords.
      </Trans>
    </div>
  );
};

export default memo(Index);
