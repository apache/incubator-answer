import { FC } from 'react';
import { Card, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type * as Type from '@/common/interface';

interface IProps {
  data: Type.AdminDashboard['info'];
}
const Statistics: FC<IProps> = ({ data }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.dashboard' });

  return (
    <Card className="mb-4">
      <Card.Body>
        <h6 className="mb-3">{t('site_statistics')}</h6>
        <Row>
          <Col xs={6} className="mb-1">
            <span className="text-secondary me-1">{t('questions')}</span>
            <strong>{data.question_count}</strong>
          </Col>
          <Col xs={6} className="mb-1">
            <span className="text-secondary me-1">{t('answers')}</span>
            <strong>{data.answer_count}</strong>
          </Col>
          <Col xs={6} className="mb-1">
            <span className="text-secondary me-1">{t('comments')}</span>
            <strong>{data.comment_count}</strong>
          </Col>
          <Col xs={6} className="mb-1">
            <span className="text-secondary me-1">{t('votes')}</span>
            <strong>{data.vote_count}</strong>
          </Col>
          <Col xs={6}>
            <span className="text-secondary me-1">{t('active_users')}</span>
            <strong>{data.user_count}</strong>
          </Col>
          <Col xs={6}>
            <span className="text-secondary me-1">{t('flags')}</span>
            <strong>{data.report_count}</strong>
            <a href="###" className="ms-2">
              {t('review')}
            </a>
          </Col>
        </Row>
      </Card.Body>
    </Card>
  );
};

export default Statistics;
