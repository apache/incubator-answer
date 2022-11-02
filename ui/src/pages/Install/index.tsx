import { FC, useState, useEffect } from 'react';
import { Container, Row, Col, Card, Alert } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';

import type { FormDataType } from '@/common/interface';
import { Storage } from '@/utils';
import { PageTitle } from '@/components';

import {
  FirstStep,
  SecondStep,
  ThirdStep,
  FourthStep,
  Fifth,
} from './components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'install' });
  const [step, setStep] = useState(1);
  const [showError] = useState(false);

  const [formData, setFormData] = useState<FormDataType>({
    lang: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    db_type: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    db_username: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    db_password: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    db_host: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    db_name: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    db_file: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },

    site_name: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    contact_email: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    admin_name: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    admin_password: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    admin_email: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
  });

  const handleChange = (params: FormDataType) => {
    console.log(params);
    setFormData({ ...formData, ...params });
  };

  const handleStep = () => {
    setStep((pre) => pre + 1);
  };

  // const handleSubmit = () => {
  //   const params = {
  //     lang: formData.lang.value,
  //     db_type: formData.db_type.value,
  //     db_username: formData.db_username.value,
  //     db_password: formData.db_password.value,
  //     db_host: formData.db_host.value,
  //     db_name: formData.db_name.value,
  //     db_file: formData.db_file.value,
  //     site_name: formData.site_name.value,
  //     contact_email: formData.contact_email.value,
  //     admin_name: formData.admin_name.value,
  //     admin_password: formData.admin_password.value,
  //     admin_email: formData.admin_email.value,
  //   };

  //   console.log(params);
  // };

  useEffect(() => {
    console.log('step===', Storage.get('INSTALL_STEP'));
  }, []);

  return (
    <div className="page-wrap2">
      <PageTitle title={t('install', { keyPrefix: 'page_title' })} />
      <Container style={{ paddingTop: '74px' }}>
        <Row className="justify-content-center">
          <Col lg={6}>
            <h2 className="mb-4 text-center">{t('title')}</h2>
            <Card>
              <Card.Body>
                {showError && <Alert variant="danger"> show error msg </Alert>}

                <FirstStep
                  visible={step === 1}
                  data={formData.lang}
                  changeCallback={handleChange}
                  nextCallback={handleStep}
                />

                <SecondStep
                  visible={step === 2}
                  data={formData}
                  changeCallback={handleChange}
                  nextCallback={handleStep}
                />

                <ThirdStep visible={step === 3} nextCallback={handleStep} />

                <FourthStep
                  visible={step === 4}
                  data={formData}
                  changeCallback={handleChange}
                  nextCallback={handleStep}
                />

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
