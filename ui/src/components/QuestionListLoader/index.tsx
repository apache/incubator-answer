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
          className="bg-transparent py-3 px-0 border-start-0 border-end-0 placeholder-glow"
          key={v}>
          <div
            className="placeholder w-100 h5 align-top"
            style={{ height: '24px' }}
          />

          <div
            className="placeholder w-75 d-block align-top mb-2"
            style={{ height: '21px' }}
          />

          <div
            className="placeholder w-50 align-top"
            style={{ height: '24px' }}
          />
        </ListGroupItem>
      ))}
    </>
  );
};

export default memo(Index);
