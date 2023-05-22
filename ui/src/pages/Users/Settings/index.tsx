import { FC, memo } from 'react';
import { Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Outlet } from 'react-router-dom';

import { usePageTags } from '@/hooks';

import Nav from './components/Nav';

import './index.scss';

const Index: FC = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'settings.profile',
  });

  usePageTags({
    title: t('settings', { keyPrefix: 'page_title' }),
  });
  return (
    <Row className="mt-4 mb-5 pb-5">
      <Col className="settings-nav mb-4">
        <Nav />
      </Col>
      <Col className="settings-main">
        <Outlet />
      </Col>
    </Row>
  );
};

export default memo(Index);
