import { ListGroup } from 'react-bootstrap';
import { Link } from 'react-router-dom';

import classNames from 'classnames';

import './index.scss';

const Achievements = ({ data, handleReadNotification }) => {
  if (!data) {
    return null;
  }
  return (
    <ListGroup
      className="border-top border-bottom achievement-wrap"
      variant="flush">
      {data.map((item) => {
        let url = '';
        switch (item.object_info.object_type) {
          case 'question':
            url = `/questions/${item.object_info.object_id}`;
            break;
          case 'answer':
            url = `/questions/${item.object_info?.object_map?.question}/${item.object_info.object_id}`;
            break;
          default:
            url = '';
        }
        return (
          <ListGroup.Item
            className={classNames('d-flex', !item.is_read && 'warning')}>
            {item.rank > 0 && (
              <div className="text-success num text-end">{`+${item.rank}`}</div>
            )}
            {item.rank === 0 && <div className="num text-end">{item.rank}</div>}
            {item.rank < 0 && (
              <div className="text-danger num text-end">{`${item.rank}`}</div>
            )}
            <div className="d-flex flex-column ms-3 flex-fill">
              <Link to={url} onClick={() => handleReadNotification(item.id)}>
                {item.object_info.title}
              </Link>
              <span className="text-secondary">
                {item.object_info.object_type}
              </span>
            </div>
          </ListGroup.Item>
        );
      })}
    </ListGroup>
  );
};

export default Achievements;
