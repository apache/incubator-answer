import { FC, memo } from 'react';
import { Pagination } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useSearchParams, useNavigate, useLocation } from 'react-router-dom';

interface Props {
  currentPage: number;
  pageSize: number;
  totalSize: number;
  pathname?: string;
}

interface PageItemProps {
  page: number;
  currentPage: number;
  path: string;
}

const pageArr = [
  {
    href: '1',
    page: 1,
  },
  {
    href: '#!',
    page: 2,
  },
  {
    href: '#!',
    page: 3,
  },
  {
    href: '#!',
    page: 4,
  },
  {
    href: '#!',
    page: 5,
  },
];

const PageItem = ({ page, currentPage, path }: PageItemProps) => {
  const navigate = useNavigate();
  return (
    <Pagination.Item
      active={currentPage === page}
      href={path}
      onClick={(e) => {
        e.preventDefault();
        navigate(path);
        window.scrollTo(0, 0);
      }}>
      {page}
    </Pagination.Item>
  );
};

const Index: FC<Props> = ({
  currentPage = 1,
  pageSize = 15,
  totalSize = 0,
  pathname = '',
}) => {
  const { t } = useTranslation('translation', { keyPrefix: 'pagination' });
  const location = useLocation();
  if (!pathname) {
    pathname = location.pathname;
  }
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const totalPage = Math.ceil(totalSize / pageSize);
  const realPage = currentPage > totalPage ? totalPage : currentPage;

  const mapPage = pageArr.filter((i) => i.page <= totalPage);

  if (totalPage <= 1) {
    return null;
  }

  const handleParams = (pageNum): string => {
    searchParams.set('page', String(pageNum));
    const searchStr = searchParams.toString();
    return `${pathname}?${searchStr}`;
  };
  return (
    <Pagination size="sm" className="d-inline-flex mb-0">
      {currentPage > 1 && (
        <Pagination.Prev
          href={handleParams(currentPage - 1)}
          onClick={(e) => {
            e.preventDefault();
            navigate(handleParams(currentPage - 1));
            window.scrollTo(0, 0);
          }}>
          {t('prev')}
        </Pagination.Prev>
      )}
      {currentPage >= 1 && currentPage <= 4 && (
        <>
          {mapPage.map((item) => {
            return (
              <PageItem
                key={item.page}
                page={item.page}
                currentPage={currentPage}
                path={handleParams(item.page)}
              />
            );
          })}
        </>
      )}
      {currentPage === 4 && totalPage > 6 && (
        <PageItem
          key="6"
          page={6}
          currentPage={currentPage}
          path={handleParams(6)}
        />
      )}

      {currentPage > 4 && (
        <>
          <PageItem
            key="first"
            page={1}
            currentPage={currentPage}
            path={handleParams(1)}
          />

          <Pagination.Ellipsis disabled />
        </>
      )}
      {currentPage >= 5 && (
        <>
          <PageItem
            key="prev2"
            page={realPage - 2}
            currentPage={currentPage}
            path={handleParams(realPage - 2)}
          />
          <PageItem
            key="prev1"
            page={realPage - 1}
            currentPage={currentPage}
            path={handleParams(realPage - 1)}
          />
        </>
      )}

      {currentPage > totalPage && (
        <PageItem
          key={realPage}
          page={realPage}
          currentPage={currentPage}
          path={handleParams(realPage)}
        />
      )}

      {currentPage >= 5 &&
        totalPage >= currentPage &&
        new Array(
          totalPage <= 3
            ? totalPage - currentPage + 1
            : Math.min(totalPage - currentPage + 1, 3),
        )
          .fill('')
          .map((v, i) => {
            return (
              <PageItem
                key={`${currentPage + i}`}
                page={currentPage + i}
                currentPage={currentPage}
                path={handleParams(currentPage + i)}
              />
            );
          })}
      {totalPage > 5 && realPage + 2 < totalPage && (
        <Pagination.Ellipsis disabled />
      )}

      {totalPage > 0 && currentPage < totalPage && (
        <Pagination.Next
          disabled={currentPage === totalPage}
          href={handleParams(currentPage + 1)}
          onClick={(e) => {
            e.preventDefault();
            navigate(handleParams(currentPage + 1));
            window.scrollTo(0, 0);
          }}>
          {t('next')}
        </Pagination.Next>
      )}
    </Pagination>
  );
};

export default memo(Index);
