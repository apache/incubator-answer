import { FC, memo } from 'react';
import { ButtonGroup, Button } from 'react-bootstrap';
import { useSearchParams, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

const sortBtns = [
  {
    name: 'newest',
  },
  {
    name: 'score',
  },
];

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
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const { t } = useTranslation('translation', { keyPrefix: 'personal' });

  const handleParams = (order): string => {
    const basePath = window.location.pathname;
    searchParams.delete('page');
    searchParams.set('order', order);
    const searchStr = searchParams.toString();
    return `${basePath}?${searchStr}`;
  };

  const handleClick = (e, type) => {
    e.preventDefault();
    const str = handleParams(type);
    navigate(str);
  };

  if (!visible) {
    return null;
  }

  return (
    <div className="d-flex  align-items-center justify-content-between pb-3 border-bottom">
      <h5 className="mb-0">
        {count} {t(tabName)}
      </h5>
      {(tabName === 'answers' || tabName === 'questions') && (
        <ButtonGroup size="sm">
          {sortBtns.map((item) => {
            return (
              <Button
                variant="outline-secondary"
                active={sort === item.name}
                href={handleParams(item.name)}
                key={item.name}
                onClick={(e) => handleClick(e, item.name)}>
                {t(item.name)}
              </Button>
            );
          })}
        </ButtonGroup>
      )}
    </div>
  );
};

export default memo(Index);
