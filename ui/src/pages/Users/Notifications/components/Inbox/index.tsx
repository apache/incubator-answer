import { ListGroup } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import classNames from 'classnames';
import { isEmpty } from 'lodash';

import { FormatTime, Empty } from '@/components';

const Inbox = ({ data, handleReadNotification }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'notifications' });
  if (!data) {
    return null;
  }
  if (isEmpty(data)) {
    return <Empty />;
  }
  return (
    <ListGroup className="rounded-0">
      {data.map((item) => {
        const { comment, question, answer } =
          item?.object_info?.object_map || {};
        let url = '';
        switch (item.object_info.object_type) {
          case 'question':
            url = `/questions/${item.object_info.object_id}`;
            break;
          case 'answer':
            url = `/questions/${question}/${item.object_info.object_id}`;
            break;
          case 'comment':
            url = `/questions/${question}/${answer}?commentId=${comment}`;
            break;
          default:
            url = '';
        }
        return (
          <ListGroup.Item
            key={item.id}
            className={classNames(
              'py-3 border-start-0 border-end-0',
              !item.is_read && 'warning',
            )}>
            <div>
              {item.user_info && item.user_info.status !== 'deleted' ? (
                <Link to={`/users/${item.user_info.username}`}>
                  {item.user_info.display_name}{' '}
                </Link>
              ) : (
                // someone for anonymous user display
                <span>{item.user_info?.display_name || t('someone')} </span>
              )}
              {item.notification_action}{' '}
              <Link to={url} onClick={() => handleReadNotification(item.id)}>
                {item.object_info.title}
              </Link>
            </div>
            <div className="text-secondary small">
              <FormatTime time={item.update_time} />
            </div>
          </ListGroup.Item>
        );
      })}
    </ListGroup>
  );
};

export default Inbox;
