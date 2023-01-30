import { Card, Row, Col } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

const AnswerLinks = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'admin.dashboard' });

  return (
    <Card className="mb-4">
      <Card.Body>
        <h6 className="mb-3">{t('answer_links')}</h6>
        <Row>
          <Col xs={6}>
            <a href="https://answer.dev/docs" target="_blank" rel="noreferrer">
              {t('documents')}
            </a>
          </Col>
          <Col xs={6}>
            <a href="https://meta.answer.dev" target="_blank" rel="noreferrer">
              {t('support')}
            </a>
          </Col>
        </Row>
      </Card.Body>
    </Card>
  );
};

export default AnswerLinks;
