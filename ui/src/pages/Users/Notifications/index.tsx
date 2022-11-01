import { useState, useEffect } from 'react';
import { Container, Row, Col, ButtonGroup, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useParams, useNavigate } from 'react-router-dom';

import { PageTitle } from '@answer/components';

import Inbox from './components/Inbox';
import Achievements from './components/Achievements';

import {
  useQueryNotifications,
  clearUnreadNotification,
  clearNotificationStatus,
  readNotification,
} from '@/services';

const PAGE_SIZE = 10;

const Notifications = () => {
  const [page, setPage] = useState(1);
  const [notificationData, setNotificationData] = useState<any>([]);
  const { t } = useTranslation('translation', { keyPrefix: 'notifications' });
  const { type = 'inbox' } = useParams();

  const { data, mutate } = useQueryNotifications({
    type,
    page,
    page_size: PAGE_SIZE,
  });

  useEffect(() => {
    clearNotificationStatus(type);
  }, []);

  useEffect(() => {
    if (!data) {
      return;
    }
    if (page > 1) {
      setNotificationData([...notificationData, ...(data?.list || [])]);
    } else {
      setNotificationData(data?.list);
    }
  }, [data]);
  const navigate = useNavigate();

  const handleTypeChange = (evt, val) => {
    evt.preventDefault();
    if (type === val) {
      return;
    }
    setPage(1);
    setNotificationData([]);
    navigate(`/users/notifications/${val}`);
  };

  const handleLoadMore = () => {
    setPage(page + 1);
  };

  const handleUnreadNotification = async () => {
    await clearUnreadNotification(type);
    mutate();
  };

  const handleReadNotification = (id) => {
    readNotification(id);
  };

  return (
    <>
      <PageTitle title={t('notifications', { keyPrefix: 'page_title' })} />
      <Container className="pt-4 mt-2 mb-5">
        <Row className="justify-content-center">
          <Col xxl={7} lg={8} sm={12}>
            <h3 className="mb-4">{t('title')}</h3>
            <div className="d-flex justify-content-between mb-3">
              <ButtonGroup size="sm">
                <Button
                  as="a"
                  href="/users/notifications/inbox"
                  variant="outline-secondary"
                  active={type === 'inbox'}
                  onClick={(evt) => handleTypeChange(evt, 'inbox')}>
                  {t('inbox')}
                </Button>
                <Button
                  as="a"
                  href="/users/notifications/achievement"
                  variant="outline-secondary"
                  active={type === 'achievement'}
                  onClick={(evt) => handleTypeChange(evt, 'achievement')}>
                  {t('achievement')}
                </Button>
              </ButtonGroup>
              <Button
                size="sm"
                variant="outline-secondary"
                onClick={handleUnreadNotification}>
                {t('all_read')}
              </Button>
            </div>
            {type === 'inbox' && (
              <Inbox
                data={notificationData}
                handleReadNotification={handleReadNotification}
              />
            )}
            {type === 'achievement' && (
              <Achievements
                data={notificationData}
                handleReadNotification={handleReadNotification}
              />
            )}
            {(data?.count || 0) > PAGE_SIZE * page && (
              <div className="d-flex justify-content-center align-items-center py-3">
                <Button
                  variant="link"
                  className="btn-no-border"
                  onClick={handleLoadMore}>
                  {t('show_more')}
                </Button>
              </div>
            )}
          </Col>
          <Col xxl={3} lg={4} sm={12} className="mt-5 mt-lg-0" />
        </Row>
      </Container>
    </>
  );
};

export default Notifications;
