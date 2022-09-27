import { FC, memo } from 'react';
import { ListGroupItem, ButtonGroup, Button } from 'react-bootstrap';
import { useSearchParams, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

const sortBtns = [
  {
    name: 'newest',
  },
  {
    name: 'active',
  },
  {
    name: 'score',
  },
];

interface Props {
  count: number;
  sort: string;
}
const Index: FC<Props> = ({ sort, count = 0 }) => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();

  const { t } = useTranslation('translation', {
    keyPrefix: 'search.sort_btns',
  });

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

  return (
    <ListGroupItem className="d-flex flex-wrap align-items-center justify-content-between divide-line pb-3 border-bottom px-0">
      <h5 className="mb-0">{t('counts', { count, keyPrefix: 'search' })}</h5>
      <ButtonGroup size="sm">
        {sortBtns.map((item) => {
          return (
            <Button
              as="a"
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
    </ListGroupItem>
  );
};

export default memo(Index);
