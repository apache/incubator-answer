import { FC, memo } from 'react';
import { ListGroupItem } from 'react-bootstrap';

interface Props {
  count?: number;
}

const Index: FC<Props> = ({ count = 10 }) => {
  const list = new Array(count).fill(0).map((v, i) => v + i);
  return (
    <>
      {list.map((v) => (
        <ListGroupItem
          className="py-3 px-0 border-start-0 border-end-0 bg-transparent placeholder-glow"
          key={v}>
          <div className="mb-2">
            <div
              className="placeholder me-2"
              style={{ height: '25px', width: '30px' }}
            />
            <div
              className="h5 mb-0 w-75 placeholder"
              style={{ height: '25px' }}
            />
          </div>
          <div
            className="placeholder w-50 h5 align-top mb-2"
            style={{ height: '21px' }}
          />

          <div
            className="placeholder w-100 d-block align-top mb-2"
            style={{ height: '42px' }}
          />

          <div
            className="placeholder w-25 align-top"
            style={{ height: '24px' }}
          />
        </ListGroupItem>
      ))}
    </>
  );
};

export default memo(Index);
