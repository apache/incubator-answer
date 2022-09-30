import { memo, FC, useState } from 'react';
import { Alert, Col } from 'react-bootstrap';

interface Props {
  data;
}
const Index: FC<Props> = ({ data }) => {
  const [show, setShow] = useState(Boolean(data));

  return (
    <Col lg={10} className="mb-3">
      <Alert
        variant="info"
        show={show}
        dismissible
        onClose={() => {
          setShow(false);
        }}>
        <div dangerouslySetInnerHTML={{ __html: data }} />
      </Alert>
    </Col>
  );
};

export default memo(Index);
