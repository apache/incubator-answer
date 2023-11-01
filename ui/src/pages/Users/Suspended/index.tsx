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

import { useTranslation } from 'react-i18next';
import { Button } from 'react-bootstrap';

import { siteInfoStore } from '@/stores';
import { usePageTags } from '@/hooks';

const Suspended = () => {
  const { contact_email = '' } = siteInfoStore((state) => state.siteInfo);
  const { t } = useTranslation('translation', { keyPrefix: 'suspended' });
  usePageTags({
    title: t('account_suspended', { keyPrefix: 'page_title' }),
  });

  return (
    <div className="d-flex flex-column align-items-center mt-5 pt-3">
      <h3 className="mb-5">{t('title')}</h3>
      <p className="text-center">
        {t('forever')}
        <br />
        {t('end')}
      </p>
      <Button href={`mailto:${contact_email}`} variant="link">
        {t('contact_us')}
      </Button>
    </div>
  );
};

export default Suspended;
