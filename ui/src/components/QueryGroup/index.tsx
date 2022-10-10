import { FC, memo } from 'react';
import { ButtonGroup, Button } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

interface Props {
  data: string[] | Array<{ name: string; sort: string }>;
  i18nkeyPrefix: string;
  currentSort: string;
  sortKey?: string;
  className?: string;
}

const Index: FC<Props> = ({
  data,
  currentSort = '',
  sortKey = 'order',
  i18nkeyPrefix = '',
  className = '',
}) => {
  const [searchParams, setUrlSearchParams] = useSearchParams();

  const { t } = useTranslation('translation', {
    keyPrefix: i18nkeyPrefix,
  });

  const handleParams = (order): string => {
    searchParams.delete('page');
    searchParams.set(sortKey, order);
    const searchStr = searchParams.toString();
    return `?${searchStr}`;
  };

  const handleClick = (e, type) => {
    e.preventDefault();
    const str = handleParams(type);
    setUrlSearchParams(str);
  };

  return (
    <ButtonGroup size="sm">
      {data.map((btn) => {
        const key = typeof btn === 'string' ? btn : btn.sort;
        const name = typeof btn === 'string' ? btn : btn.name;
        return (
          <Button
            as="a"
            key={key}
            variant="outline-secondary"
            active={currentSort === name}
            className={`text-capitalize ${className}`}
            href={handleParams(key)}
            onClick={(evt) => handleClick(evt, key)}>
            {t(name)}
          </Button>
        );
      })}
    </ButtonGroup>
  );
};

export default memo(Index);
