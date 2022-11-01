import { FC, useState } from 'react';
import { Container, Row, Col, Card, Alert } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';

import {
  FirstStep,
  SecondStep,
  ThirdStep,
  FourthStep,
  Fifth,
} from './components';

import { PageTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'install' });
  const [step] = useState(7);

  return (
    <div style={{ background: '#f5f5f5', minHeight: '100vh' }}>
      <PageTitle title={t('install', { keyPrefix: 'page_title' })} />
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
                {step === 6 && (
                  <div>
                    <h5>{t('warning')}</h5>
                    <p>
                      <Trans i18nKey="install.warning_description">
                        The file <code>config.yaml</code> already exists. If you
                        need to reset any of the configuration items in this
                        file, please delete it first. You may try{' '}
                        <a href="/">installing now</a>.
                      </Trans>
                    </p>
                  </div>
                )}

                {step === 7 && (
                  <div>
                    <h5>{t('installed')}</h5>
                    <p>{t('installed_description')}</p>
                  </div>
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
