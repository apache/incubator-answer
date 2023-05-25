import { memo, FC } from 'react';
import { Alert } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import dayjs from 'dayjs';

interface Props {
  data;
}
const Index: FC<Props> = ({ data }) => {
  const { t } = useTranslation();
  return (
    <Alert className="mb-4" variant={data.level}>
      {data.level === 'info' ? (
        <div>
          {data.msg.indexOf('http') > -1 ? (
            <p>
              {data.description}{' '}
              <a href={data.msg} style={{ color: '#055160' }}>
                <strong>{t('question_detail.show_exist')}</strong>
              </a>
            </p>
          ) : (
            <p>{data.msg ? data.msg : data.description}</p>
          )}
          <div className="small">
            {t('question_detail.closed_in')}{' '}
            <time
              dateTime={dayjs.unix(data.time).tz().toISOString()}
              title={dayjs
                .unix(data.time)
                .tz()
                .format(t('dates.long_date_with_time'))}>
              {dayjs
                .unix(data.time)
                .tz()
                .format(t('dates.long_date_with_year'))}
            </time>
            .
          </div>
        </div>
      ) : (
        data.msg
      )}
    </Alert>
  );
};

export default memo(Index);
