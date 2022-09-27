import { FC, memo } from 'react';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';
import dayjs from 'dayjs';

interface Props {
  time: number;
  className?: string;
  preFix?: string;
}

const Index: FC<Props> = ({ time, preFix, className }) => {
  const { t } = useTranslation();
  const formatTime = (from) => {
    const now = Math.floor(dayjs().valueOf() / 1000);
    const between = now > from ? now - from : 0;

    if (between <= 1) {
      return t('dates.now');
    }
    if (between > 1 && between < 60) {
      return t('dates.x_seconds_ago', { count: between });
    }

    if (between >= 60 && between < 3600) {
      const min = Math.floor(between / 60);
      return t('dates.x_minutes_ago', { count: min });
    }
    if (between >= 3600 && between < 3600 * 24) {
      const h = Math.floor(between / 3600);
      return t('dates.x_hours_ago', { count: h });
    }

    if (
      between >= 3600 * 24 &&
      between < 3600 * 24 * 366 &&
      dayjs.unix(from).format('YYYY') === dayjs.unix(now).format('YYYY')
    ) {
      return dayjs.unix(from).format(t('dates.long_date'));
    }

    return dayjs.unix(from).format(t('dates.long_date_with_year'));
  };

  if (!time) {
    return null;
  }

  return (
    <time
      className={classNames('', className)}
      dateTime={dayjs.unix(time).toISOString()}
      title={dayjs.unix(time).format(t('dates.long_date_with_time'))}>
      {preFix ? `${preFix} ` : ''}
      {formatTime(time)}
    </time>
  );
};

export default memo(Index);
