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

import React, { FC, useLayoutEffect, useState } from 'react';
import { Button, ButtonProps, Spinner } from 'react-bootstrap';

import { request } from '@/utils';
import type { UIAction, FormKit } from '../types';
import { useToast } from '@/hooks';

interface Props {
  fieldName: string;
  text: string;
  action: UIAction | undefined;
  formKit: FormKit;
  readOnly: boolean;
  variant?: ButtonProps['variant'];
  size?: ButtonProps['size'];
}
const Index: FC<Props> = ({
  fieldName,
  action,
  formKit,
  text = '',
  readOnly = false,
  variant = 'primary',
  size,
}) => {
  const Toast = useToast();
  const [isLoading, setLoading] = useState(false);
  const handleToast = (msg, type: 'success' | 'danger' = 'success') => {
    const tm = action?.on_complete?.toast_return_message;
    if (tm === false || !msg) {
      return;
    }
    Toast.onShow({
      msg,
      variant: type,
    });
  };
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const handleCallback = (resp) => {
    const callback = action?.on_complete;
    if (callback?.refresh_form_config) {
      formKit.refreshConfig();
    }
  };
  const handleAction = () => {
    if (!action) {
      return;
    }
    setLoading(true);
    request
      .request({
        method: action.method,
        url: action.url,
        timeout: 0,
      })
      .then((resp) => {
        if ('message' in resp) {
          handleToast(resp.message, 'success');
        }
        handleCallback(resp);
      })
      .catch((ex) => {
        if (ex && 'msg' in ex) {
          handleToast(ex.msg, 'danger');
        }
      })
      .finally(() => {
        setLoading(false);
      });
  };
  useLayoutEffect(() => {
    if (action?.loading?.state === 'pending') {
      setLoading(true);
    }
  }, []);
  const loadingText = action?.loading?.text || text;
  const disabled = isLoading || readOnly;

  return (
    <div className="d-flex">
      <Button
        name={fieldName}
        onClick={handleAction}
        disabled={disabled}
        size={size}
        variant={variant}>
        {isLoading ? (
          <>
            <Spinner
              className="align-middle me-2"
              animation="border"
              size="sm"
              variant={variant}
            />
            {loadingText}
          </>
        ) : (
          text
        )}
      </Button>
    </div>
  );
};

export default Index;
