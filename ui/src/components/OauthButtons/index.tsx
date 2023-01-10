import { memo, FC } from 'react';
import { Button } from 'react-bootstrap';

import classnames from 'classnames';

import { Icon } from '@/components';

interface Props {
  // data: any[]; // should use oauth plugin schemes
  className?: string;
}
const Index: FC<Props> = ({ className }) => {
  return (
    <div className={classnames('d-grid gap-2', className)}>
      <Button
        variant="outline-secondary"
        href="https://github.com/login/oauth/authorize?client_id=8cb9d4760cfd71c24de9&edirect_uri=http://10.0.20.88:8080/answer/api/v1/connector/redirect/github">
        <Icon name="github" className="me-2" />
        <span>Connect with Github</span>
      </Button>

      <Button variant="outline-secondary">
        <Icon name="twitter" className="me-2" />
        <span>Connect with Google</span>
      </Button>
    </div>
  );
};

export default memo(Index);
