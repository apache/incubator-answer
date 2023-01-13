import { memo, FC } from 'react';
import { Button } from 'react-bootstrap';

import classnames from 'classnames';

import { useGetStartUseOauthConnector } from '@/services';

interface Props {
  className?: string;
}
const Index: FC<Props> = ({ className }) => {
  const { data } = useGetStartUseOauthConnector();

  if (!data?.length) return null;
  return (
    <div className={classnames('d-grid gap-2', className)}>
      {data?.map((item) => {
        return (
          <Button variant="outline-secondary" href={item.link} key={item.name}>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              className="btnSvg"
              viewBox="0 0 24 24">
              <path d={item.icon} />
            </svg>
            <span> {item.name}</span>
          </Button>
        );
      })}
    </div>
  );
};

export default memo(Index);
