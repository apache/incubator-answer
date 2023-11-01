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

import React, { useState } from 'react';
import { Container, Col } from 'react-bootstrap';
import { Trans, useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';

import SendEmail from './components/sendEmail';

const Index: React.FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'account_forgot' });
  const [step, setStep] = useState(1);
  const [email, setEmail] = useState('');

  const callback = (param: number, mail: string) => {
    setStep(param);
    setEmail(mail);
  };
  usePageTags({
    title: t('account_recovery', { keyPrefix: 'page_title' }),
  });
  return (
    <Container style={{ paddingTop: '4rem', paddingBottom: '6rem' }}>
      <h3 className="text-center mb-5">{t('page_title')}</h3>
      {step === 1 && (
        <Col className="mx-auto" md={6} lg={4} xl={3}>
          <SendEmail visible={step === 1} callback={callback} />
        </Col>
      )}
      {step === 2 && (
        <Col className="mx-auto px-4" md={6}>
          <div className="text-center">
            <p>
              <Trans
                i18nKey="account_forgot.send_success"
                values={{ mail: email }}
                components={{ bold: <strong /> }}
              />
            </p>
          </div>
        </Col>
      )}
    </Container>
  );
};

export default React.memo(Index);
