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
    <Alert className="mb-4" variant="info">
      <div>
        {data.operation_msg.indexOf('http') > -1 ? (
          <p>
            {data.operation_description}{' '}
            <a href={data.operation_msg} style={{ color: '#055160' }}>
              <strong>{t('question_detail.show_exist')}</strong>
            </a>
          </p>
        ) : (
          <p>
            {data.operation_msg
              ? data.operation_msg
              : data.operation_description}
          </p>
        )}
        <div className="fs-14">
          {t('question_detail.closed_in')}{' '}
          <time
            dateTime={dayjs.unix(data.operation_time).toISOString()}
            title={dayjs
              .unix(data.operation_time)
              .format(t('dates.long_date_with_time'))}>
            {dayjs
              .unix(data.operation_time)
              .format(t('dates.long_date_with_year'))}
          </time>
          .
        </div>
      </div>
    </Alert>
  );
};

export default memo(Index);
