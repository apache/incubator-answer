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

import { useCallback } from 'react';
import {
  useBeforeUnload,
  unstable_usePrompt as usePrompt,
} from 'react-router-dom';
import { useTranslation } from 'react-i18next';

// https://gist.github.com/chaance/2f3c14ec2351a175024f62fd6ba64aa6
// The link above is an example of implementing usePrompt with useBlocker.
interface PromptProps {
  when: boolean;
  beforeUnload?: boolean;
}

const usePromptWithUnload = ({
  when = false,
  beforeUnload = true,
}: PromptProps) => {
  const { t } = useTranslation('translation', { keyPrefix: 'prompt' });

  usePrompt({
    when,
    message: `${t('leave_page')} ${t('changes_not_save')}`,
  });

  useBeforeUnload(
    useCallback(
      (event) => {
        if (beforeUnload && when) {
          const msg = t('changes_not_save');
          event.preventDefault();
          event.returnValue = msg;
        }
      },
      [when, beforeUnload],
    ),
    { capture: true },
  );
};

export default usePromptWithUnload;
