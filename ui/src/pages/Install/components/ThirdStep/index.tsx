/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

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
