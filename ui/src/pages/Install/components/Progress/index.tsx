import { FC, memo } from 'react';
import { ProgressBar } from 'react-bootstrap';

interface IProps {
  step: number;
}

const Index: FC<IProps> = ({ step }) => {
  return (
    <div className="d-flex align-items-center small text-secondary">
      <ProgressBar
        now={(step / 5) * 100}
        variant="success"
        style={{ width: '200px' }}
        className="me-2"
      />
      <span>{step}/5</span>
    </div>
  );
};

export default memo(Index);
