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

import { useNavigate } from 'react-router-dom';

import { floppyNavigation } from '@/utils';
import Storage from '@/utils/storage';
import { RouteAlias } from '@/router/alias';
import { REDIRECT_PATH_STORAGE_KEY } from '@/common/constants';

const Index = () => {
  const navigate = useNavigate();

  const loginRedirect = () => {
    const redirect = Storage.get(REDIRECT_PATH_STORAGE_KEY) || RouteAlias.home;
    Storage.remove(REDIRECT_PATH_STORAGE_KEY);
    floppyNavigation.navigate(redirect, {
      handler: navigate,
      options: {
        replace: true,
      },
    });
  };

  return { loginRedirect };
};

export default Index;
