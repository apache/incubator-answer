import { FC } from 'react';
import { Card, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';
import { formatUptime } from '@/utils';

interface IProps {
  data: Type.AdminDashboard['info'];
}
const SystemInfo: FC<IProps> = ({ data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.dashboard' });

  return (
    <Card className="mb-4">
      <Card.Body>
        <h6 className="mb-3">{t('system_info')}</h6>
        <Row>
          <Col xs={6}>
            <span className="text-secondary me-1">{t('storage_used')}</span>
            <strong>{data.occupying_storage_space}</strong>
          </Col>
          {data.app_start_time ? (
            <Col xs={6}>
              <span className="text-secondary me-1">{t('uptime')}</span>
              <strong>{formatUptime(data.app_start_time)}</strong>
            </Col>
          ) : null}
        </Row>
      </Card.Body>
    </Card>
  );
};

export default SystemInfo;
