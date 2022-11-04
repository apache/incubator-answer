import { useState } from 'react';
import { Container, Row, Col, Card, Button, Spinner } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';

import { PageTitle } from '@/components';
import { upgradSystem } from '@/services';

const Index = () => {
  const { t } = useTranslation('translation', {
    keyPrefix: 'upgrade',
  });
  const [step] = useState(1);
  const [loading, setLoading] = useState(false);

  const handleUpdate = async () => {
    await upgradSystem();
    setLoading(true);
  };
  return (
    <div className="page-wrap2">
      <Container style={{ paddingTop: '74px' }}>
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
                    {loading ? (
                      <Button variant="primary" disabled className="float-end">
                        <Spinner
                          as="span"
                          animation="border"
                          size="sm"
                          role="status"
                          aria-hidden="true"
                        />
                        <span> {t('update_btn')}</span>
                      </Button>
                    ) : (
                      <Button className="float-end" onClick={handleUpdate}>
                        {t('update_btn')}
                      </Button>
                    )}
                  </>
                )}

                {step === 2 && (
                  <>
                    <h5>{t('done_title')}</h5>
                    <p>{t('done_desscription')}</p>
                    <Button className="float-end" href="/">
                      {t('done_btn')}
                    </Button>
                  </>
                )}
              </Card.Body>
            </Card>
          </Col>
        </Row>
      </Container>
    </div>
  );
};

export default Index;
