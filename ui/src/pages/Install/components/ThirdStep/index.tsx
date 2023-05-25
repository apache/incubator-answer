import { FC } from 'react';
import { Form, Button, FormGroup } from 'react-bootstrap';
import { useTranslation, Trans } from 'react-i18next';

import Progress from '../Progress';

interface Props {
  visible: boolean;
  errorMsg;
  nextCallback: () => void;
}

const Index: FC<Props> = ({ visible, errorMsg, nextCallback }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'install' });

  if (!visible) return null;
  return (
    <div>
      <h5>{t('config_yaml.title')}</h5>

      {errorMsg?.msg?.length > 0 ? (
        <>
          <div className="fmt">
            <p>
              <Trans
                i18nKey="install.config_yaml.desc"
                components={{ 1: <code /> }}
              />
            </p>
          </div>
          <FormGroup className="mb-3">
            <Form.Control
              type="text"
              as="textarea"
              rows={8}
              className="small"
              value={errorMsg?.default_config}
            />
          </FormGroup>
          <div className="mb-3">{t('config_yaml.info')}</div>
        </>
      ) : (
        <div className="mb-3">{t('config_yaml.label')}</div>
      )}

      <div className="d-flex align-items-center justify-content-between">
        <Progress step={3} />
        <Button onClick={nextCallback}>{t('next')}</Button>
      </div>
    </div>
  );
};

export default Index;
