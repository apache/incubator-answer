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

import { memo } from 'react';
import { Container } from 'react-bootstrap';
import { useSearchParams } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import { userCenterStore } from '@/stores';
import { USER_AGENT_NAMES } from '@/common/constants';
import { usePageTags } from '@/hooks';

import WeCom from './components/WeCom';

const Index = () => {
  const { t } = useTranslation('translation');
  const [searchParam] = useSearchParams();
  const { agent: ucAgent } = userCenterStore();
  let agentName = ucAgent?.agent_info?.name || '';
  if (searchParam.get('agent_name')) {
    agentName = searchParam.get('agent_name') || '';
  }
  usePageTags({
    title: t('login', { keyPrefix: 'page_title' }),
  });
  return (
    <Container>
      {USER_AGENT_NAMES.WeCom.toLowerCase() === agentName.toLowerCase() ? (
        <WeCom />
      ) : null}
    </Container>
  );
};

export default memo(Index);
