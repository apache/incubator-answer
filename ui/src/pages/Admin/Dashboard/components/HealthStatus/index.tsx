import { FC } from 'react';
import { Card, Row, Col, Badge } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { Link } from 'react-router-dom';

import type * as Type from '@answer/common/interface';

interface IProps {
  data: Type.AdminDashboard['info'];
}

const HealthStatus: FC<IProps> = ({ data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.dashboard' });

  return (
    <Card className="mb-4">
      <Card.Body>
        <h6 className="mb-3">{t('site_health_status')}</h6>
        <Row>
          <Col xs={6} className="mb-1 d-flex align-items-center">
            <span className="text-secondary me-1">{t('version')}</span>
            <strong>90</strong>
            <Badge pill bg="warning" text="dark" className="ms-1">
              {t('update_to')} {data.app_version}
            </Badge>
          </Col>
          <Col xs={6} className="mb-1">
            <span className="text-secondary me-1">{t('https')}</span>
            <strong>{data.https ? t('yes') : t('yes')}</strong>
          </Col>
          <Col xs={6} className="mb-1">
            <span className="text-secondary me-1">{t('uploading_files')}</span>
            <strong>
              {data.uploading_files ? t('allowed') : t('not_allowed')}
            </strong>
          </Col>
          <Col xs={6}>
            <span className="text-secondary me-1">{t('smtp')}</span>
            <strong>{data.smtp ? t('enabled') : t('disabled')}</strong>
            <Link to="/admin/smtp" className="ms-2">
              {t('config')}
            </Link>
          </Col>
          <Col xs={6}>
            <span className="text-secondary me-1">{t('timezone')}</span>
            <strong>{data.time_zone}</strong>
          </Col>
        </Row>
      </Card.Body>
    </Card>
  );
};

export default HealthStatus;
