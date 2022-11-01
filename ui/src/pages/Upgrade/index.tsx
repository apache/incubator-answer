import { useState } from 'react';
import { Container, Row, Col, Card, Button } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';

import { PageTitle } from '@/components';

const Index = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'upgrade',
  });
  const [step, setStep] = useState(1);

  const handleUpdate = () => {
    setStep(2);
  };
  return (
    <Container className="page-wrap2" style={{ paddingTop: '74px' }}>
      <PageTitle title={t('upgrade', { keyPrefix: 'page_title' })} />
      <Row className="justify-content-center">
        <Col lg={6}>
          <h2 className="text-center mb-4">{t('title')}</h2>
          <Card>
            <Card.Body>
              {step === 1 && (
                <>
                  <h5>{t('update_title')}</h5>
                  <Trans
                    i18nKey="upgrade.update_description"
                    components={{ 1: <p /> }}
                  />
                  <Button className="float-end" onClick={handleUpdate}>
                    {t('update_btn')}
                  </Button>
                </>
              )}

              {step === 2 && (
                <>
                  <h5>{t('done_title')}</h5>
                  <p>{t('done_desscription')}</p>
                  <Button className="float-end">{t('done_btn')}</Button>
                </>
              )}
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
};

export default Index;
