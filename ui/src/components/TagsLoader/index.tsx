import { FC, memo } from 'react';
import { Col, Card } from 'react-bootstrap';

interface Props {
  count?: number;
}

const Index: FC<Props> = ({ count = 20 }) => {
  const list = new Array(count).fill(0).map((v, i) => v + i);
  return (
    <>
      {list.map((v) => (
        <Col
          key={v}
          xs={12}
          lg={3}
          md={4}
          sm={6}
          className="mb-4 placeholder-glow">
          <Card className="h-100">
            <Card.Body className="d-flex flex-column align-items-start">
              <div
                className="placeholder align-top w-25 mb-3"
                style={{ height: '24px' }}
              />

              <p
                className="placeholder small text-truncate-3 w-100"
                style={{ height: '42px' }}
              />
              <div className="d-flex align-items-center">
                <div
                  className="placeholder me-2"
                  style={{ width: '80px', height: '31px' }}
                />
                <span
                  className="placeholder text-secondary small text-nowrap"
                  style={{ width: '100px', height: '21px' }}
                />
              </div>
            </Card.Body>
          </Card>
        </Col>
      ))}
    </>
  );
};

export default memo(Index);
