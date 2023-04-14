import React, { memo, FC, useState, useEffect } from 'react';
import { Card } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import QrCode from 'qrcode';

import { userCenterStore } from '@/stores';
import { guard } from '@/utils';

import { getLoginConf, checkLoginResult } from './wecom.service';

let checkTimer: NodeJS.Timeout;
const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'plugins' });
  const ucAgent = userCenterStore().agent;
  const agentName = ucAgent?.agent_info?.name || '';
  const [qrcodeDataUrl, setQrCodeDataUrl] = useState('');
  const handleLoginResult = (key: string) => {
    if (!key) {
      return;
    }
    checkLoginResult(key).then((res) => {
      if (res.is_login) {
        guard.handleLoginWithToken(res.token);
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
    QrCode.toDataURL(targetUrl, { width: 240 }, (err, url) => {
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
      handleQrCode(res?.redirect_url);
      handleLoginResult(res?.key);
    });
  }, [agentName]);
  useEffect(() => {
    return () => {
      clearTimeout(checkTimer);
    };
  }, []);
  if (/WeCom/i.test(agentName)) {
    return (
      <Card className="text-center">
        <Card.Body>
          <Card.Title as="h3">
            {agentName} {t('login')}
          </Card.Title>
          {qrcodeDataUrl ? (
            <>
              <img width={240} height={240} src={qrcodeDataUrl} alt="" />
              <div className="text-secondary">
                {t('qrcode_login_tip', { agentName })}
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
