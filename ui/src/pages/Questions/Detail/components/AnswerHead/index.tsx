import { memo, FC } from 'react';
import { ButtonGroup } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link, useLocation } from 'react-router-dom';

interface Props {
  count: number;
  order: string;
}
const Index: FC<Props> = ({ count = 0, order = 'default' }) => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'question_detail.answers',
  });
  const location = useLocation();
  return (
    <div
      className="d-flex align-items-center justify-content-between mt-5 mb-3"
      id="answerHeader">
      <h5 className="mb-0">
        {count} {t('title')}
      </h5>
      <ButtonGroup size="sm">
        <Link
          to={`${location.pathname}?order=default`}
          className={`btn btn-outline-secondary ${
            order !== 'updated' ? 'active' : ''
          }`}>
          {t('score')}
        </Link>
        <Link
          to={`${location.pathname}?order=updated`}
          className={`btn btn-outline-secondary ${
            order === 'updated' ? 'active' : ''
          }`}>
          {t('newest')}
        </Link>
      </ButtonGroup>
    </div>
  );
};

export default memo(Index);
