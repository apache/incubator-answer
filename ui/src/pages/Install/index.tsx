import { FC, useState, useEffect } from 'react';
import { Container, Row, Col, Card, Alert } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';

import type { FormDataType } from '@/common/interface';
import { PageTitle } from '@/components';
import {
  dbCheck,
  installInit,
  installBaseInfo,
  checkConfigFileExists,
} from '@/services';
import { Storage } from '@/utils';
import { CURRENT_LANG_STORAGE_KEY } from '@/common/constants';

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
  const [loading, setLoading] = useState(true);
  const [errorData, setErrorData] = useState<{ [propName: string]: any }>({
    msg: '',
  });
  const [tableExist, setTableExist] = useState(false);

  const [formData, setFormData] = useState<FormDataType>({
    lang: {
      value: '',
      isInvalid: false,
      errorMsg: '',
    },
    db_type: {
      value: 'mysql',
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
    site_url: {
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

  const handleErr = (data) => {
    window.scrollTo(0, 0);
    setErrorData(data);
  };

  const handleNext = async () => {
    setErrorData({
      msg: '',
    });
    setStep((pre) => pre + 1);
  };

  const submitDatabaseForm = () => {
    const params = {
      lang: formData.lang.value,
      db_type: formData.db_type.value,
      db_username: formData.db_username.value,
      db_password: formData.db_password.value,
      db_host: formData.db_host.value,
      db_name: formData.db_name.value,
      db_file: formData.db_file.value,
    };
    dbCheck(params)
      .then(() => {
        handleNext();
      })
      .catch((err) => {
        console.log(err);
        handleErr(err);
      });
  };

  const checkInstall = () => {
    const params = {
      lang: formData.lang.value,
      db_type: formData.db_type.value,
      db_username: formData.db_username.value,
      db_password: formData.db_password.value,
      db_host: formData.db_host.value,
      db_name: formData.db_name.value,
      db_file: formData.db_file.value,
    };
    installInit(params)
      .then(() => {
        handleNext();
      })
      .catch((err) => {
        handleErr(err);
      });
  };

  const submitSiteConfig = () => {
    const params = {
      site_name: formData.site_name.value,
      contact_email: formData.contact_email.value,
      admin_name: formData.admin_name.value,
      admin_password: formData.admin_password.value,
      admin_email: formData.admin_email.value,
    };
    installBaseInfo(params)
      .then(() => {
        handleNext();
      })
      .catch((err) => {
        handleErr(err);
      });
  };

  const handleStep = () => {
    if (step === 1) {
      Storage.set(CURRENT_LANG_STORAGE_KEY, formData.lang.value);
      handleNext();
    }
    if (step === 2) {
      submitDatabaseForm();
    }
    if (step === 3) {
      checkInstall();
    }
    if (step === 4) {
      submitSiteConfig();
    }
    if (step > 4) {
      handleNext();
    }
  };

  const handleInstallNow = (e) => {
    e.preventDefault();
    if (tableExist) {
      setStep(7);
    } else {
      setStep(4);
    }
  };

  const configYmlCheck = () => {
    checkConfigFileExists()
      .then((res) => {
        setTableExist(res?.db_table_exist);
        if (res && res.config_file_exist) {
          setStep(5);
        }
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    configYmlCheck();
  }, []);

  if (loading) {
    return <div />;
  }

  return (
    <div className="page-wrap2">
      <PageTitle title={t('install', { keyPrefix: 'page_title' })} />
      <Container style={{ paddingTop: '74px' }}>
        <Row className="justify-content-center">
          <Col lg={6}>
            <h2 className="mb-4 text-center">{t('title')}</h2>
            <Card>
              <Card.Body>
                {errorData?.msg && (
                  <Alert variant="danger">{errorData?.msg}</Alert>
                )}

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

                <ThirdStep
                  visible={step === 3}
                  nextCallback={handleStep}
                  errorMsg={errorData}
                />

                <FourthStep
                  visible={step === 4}
                  data={formData}
                  changeCallback={handleChange}
                  nextCallback={handleStep}
                />

                <Fifth visible={step === 5} siteUrl={formData.site_url.value} />
                {step === 6 && (
                  <div>
                    <h5>{t('warning')}</h5>
                    <p>
                      <Trans i18nKey="install.warning_description">
                        The file <code>config.yaml</code> already exists. If you
                        need to reset any of the configuration items in this
                        file, please delete it first. You may try{' '}
                        <a href="###" onClick={(e) => handleInstallNow(e)}>
                          installing now
                        </a>
                        .
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
