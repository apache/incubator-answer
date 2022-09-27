import { memo, FC } from 'react';
import { Trans } from 'react-i18next';

const Index: FC = () => {
  return (
    <div className="mt-5 text-center">
      <Trans i18nKey="search.empty">
        We couldn't find anything.
        <br />
        Try different or less specific keywords.
      </Trans>
    </div>
  );
};

export default memo(Index);
