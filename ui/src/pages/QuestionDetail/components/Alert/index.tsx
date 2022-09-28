import { memo, FC } from 'react';
import { Alert } from 'react-bootstrap';

interface Props {
  data;
}
const Index: FC<Props> = ({ data }) => {
  return (
    <Alert className="mb-4" variant="info">
      <div>
        <strong>{data.operation_msg} </strong>
        {data.operation_description}
      </div>
    </Alert>
  );
};

export default memo(Index);
