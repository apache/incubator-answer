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

import React, { memo, FC, useState, useEffect } from 'react';
import { Card } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';

import QrCode from 'qrcode';

import { userCenterStore } from '@/stores';
import { guard, getUaType, floppyNavigation } from '@/utils';
import { USER_AGENT_NAMES } from '@/common/constants';

import { getLoginConf, checkLoginResult } from './service';

let checkTimer: NodeJS.Timeout;
const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'user_center' });
  const navigate = useNavigate();
  const ucAgent = userCenterStore().agent;
  const agentName = ucAgent?.agent_info?.name || '';
  const [qrcodeDataUrl, setQrCodeDataUrl] = useState('');
  const handleLoginResult = (key: string) => {
    if (!key) {
      return;
    }
    checkLoginResult(key).then((res) => {
      if (res.is_login) {
        guard.handleLoginWithToken(res.token, navigate);
        return;
      }
      clearTimeout(checkTimer);
      checkTimer = setTimeout(() => {
        handleLoginResult(key);
      }, 2000);
    });
  };
  const handleQrCode = (targetUrl: string) => {
    if (!targetUrl) {
      return;
    }
    QrCode.toDataURL(targetUrl, { width: 240, margin: 0 }, (err, url) => {
      if (err) {
        return;
      }
      setQrCodeDataUrl(url);
    });
  };

  useEffect(() => {
    if (!agentName) {
      return;
    }
    getLoginConf().then((res) => {
      if (getUaType() === USER_AGENT_NAMES.WeCom) {
        floppyNavigation.navigate(res?.redirect_url, {
          handler: 'replace',
        });
      } else {
        handleQrCode(res?.redirect_url);
        handleLoginResult(res?.key);
      }
    });
  }, [agentName]);
  useEffect(() => {
    return () => {
      clearTimeout(checkTimer);
    };
  }, []);

  if (getUaType() !== USER_AGENT_NAMES.WeCom) {
    return (
      <Card className="text-center">
        <Card.Body>
          <Card.Title as="h3" className="mb-3">
            {ucAgent?.agent_info?.display_name} {t('login')}
          </Card.Title>
          {qrcodeDataUrl ? (
            <>
              <img
                className="w-100"
                style={{ maxWidth: '240px' }}
                src={qrcodeDataUrl}
                alt={agentName}
              />
              <div className="text-secondary mt-3">
                {t('qrcode_login_tip', {
                  agentName: ucAgent?.agent_info?.display_name,
                })}
              </div>
            </>
          ) : null}
        </Card.Body>
      </Card>
    );
  }
  return null;
};

export default memo(Index);
