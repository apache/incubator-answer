import { FC } from 'react';
import { Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import { useDashBoard } from '@/services';

import {
  AnswerLinks,
  HealthStatus,
  Statistics,
  SystemInfo,
} from './components';

const Dashboard: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.dashboard' });
  const { data } = useDashBoard();

  if (!data) {
    return null;
  }

  return (
    <>
      <h3 className="text-capitalize">{t('title')}</h3>
      <p className="mt-4">{t('welcome')}</p>
      <Row>
        <Col lg={6}>
          <Statistics data={data.info} />
        </Col>
        <Col lg={6}>
          <HealthStatus data={data.info} />
        </Col>
        <Col lg={6}>
          <SystemInfo data={data.info} />
        </Col>
        <Col lg={6}>
          <AnswerLinks />
        </Col>
      </Row>
      {process.env.REACT_APP_VERSION && (
        <p className="mt-4">
          {`${t('version')} `}
          {process.env.REACT_APP_VERSION}
        </p>
      )}
    </>
  );
};
export default Dashboard;
