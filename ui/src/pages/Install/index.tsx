import { FC, useState } from 'react';
import { Container, Row, Col, Card, Alert } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import {
  FirstStep,
  SecondStep,
  ThirdStep,
  FourthStep,
  Fifth,
} from './components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'install' });
  const [step] = useState(5);

  return (
    <div style={{ background: '#f5f5f5', minHeight: '100vh' }}>
      <Container style={{ paddingTop: '74px' }}>
        <Row className="justify-content-center">
          <Col lg={6}>
            <h2 className="mb-4 text-center">{t('title')}</h2>
            <Card>
              <Card.Body>
                <Alert variant="danger"> show error msg </Alert>
                <FirstStep visible={step === 1} />

                <SecondStep visible={step === 2} />

                <ThirdStep visible={step === 3} />

                <FourthStep visible={step === 4} />

                <Fifth visible={step === 5} />
              </Card.Body>
            </Card>
          </Col>
        </Row>
      </Container>
    </div>
  );
};

export default Index;
