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

import { FC, memo, useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { useSearchParams, useNavigate } from 'react-router-dom';

import { usePageTags } from '@/hooks';
import { loggedUserInfoStore } from '@/stores';
import { activateAccount } from '@/services';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'page_title' });
  const [searchParams] = useSearchParams();
  const updateUser = loggedUserInfoStore((state) => state.update);
  const navigate = useNavigate();
  useEffect(() => {
    const code = searchParams.get('code');

    if (code) {
      activateAccount(encodeURIComponent(code)).then((res) => {
        updateUser(res);
        setTimeout(() => {
          navigate('/users/account-activation/success', { replace: true });
        }, 0);
      });
    } else {
      navigate('/', { replace: true });
    }
  }, []);
  usePageTags({
    title: t('account_activation'),
  });
  return null;
};

export default memo(Index);
