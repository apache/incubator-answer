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

import { FC, memo, useEffect, useState } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Link, useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { loggedUserInfoStore } from '@/stores';
import { changeEmailVerify } from '@/services';
import { WelcomeTitle } from '@/components';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'account_result' });
  const [searchParams] = useSearchParams();
  const [step, setStep] = useState('loading');

  const updateUser = loggedUserInfoStore((state) => state.update);

  useEffect(() => {
    const code = searchParams.get('code');
    if (code) {
      // do
      changeEmailVerify({ code })
        .then((res) => {
          setStep('success');
          if (res?.access_token) {
            // update user info
            updateUser(res);
          }
        })
        .catch(() => {
          setStep('invalid');
        });
    }
  }, []);
  usePageTags({
    title: t('confirm_email', { keyPrefix: 'page_title' }),
  });
  return (
    <Container className="pt-4 mt-2 mb-5">
      <Row className="justify-content-center">
        <Col lg={6}>
          <WelcomeTitle className="mt-3 mb-5" />
          {step === 'success' && (
            <>
              <p className="text-center">{t('confirm_new_email')}</p>
              <div className="text-center">
                <Link to="/">{t('link')}</Link>
              </div>
            </>
          )}

          {step === 'invalid' && (
            <p className="text-center">{t('confirm_new_email_invalid')}</p>
          )}
        </Col>
      </Row>
    </Container>
  );
};

export default memo(Index);
