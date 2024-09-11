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

import { FC, useEffect, useState } from 'react';
import { Form, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type { LangsType, FormValue, FormDataType } from '@/common/interface';
import Progress from '../Progress';
import { getInstallLangOptions } from '@/services';
import { setupInstallLanguage } from '@/utils/localize';
import { CURRENT_LANG_STORAGE_KEY } from '@/common/constants';
import { Storage } from '@/utils';

interface Props {
  data: FormValue;
  changeCallback: (value: FormDataType) => void;
  nextCallback: () => void;
  visible: boolean;
}
const Index: FC<Props> = ({ visible, data, changeCallback, nextCallback }) => {
  const { t } = useTranslation('translation', { keyPrefix: 'install' });

  const [langs, setLangs] = useState<LangsType[]>();

  const getLangs = async () => {
    const res: LangsType[] = await getInstallLangOptions();
    const currentLang = Storage.get(CURRENT_LANG_STORAGE_KEY);
    const selectedLang = currentLang || res[0].value;

    setLangs(res);
    setupInstallLanguage(selectedLang);

    changeCallback({
      lang: {
        value: selectedLang,
        isInvalid: false,
        errorMsg: '',
      },
    });
  };

  const handleSubmit = () => {
    nextCallback();
  };

  useEffect(() => {
    getLangs();
  }, []);

  if (!visible) return null;
  return (
    <Form noValidate onSubmit={handleSubmit}>
      <Form.Group controlId="lang" className="mb-3">
        <Form.Label>{t('lang.label')}</Form.Label>
        <Form.Select
          value={data.value}
          isInvalid={data.isInvalid}
          onChange={(e) => {
            setupInstallLanguage(e.target.value);
            changeCallback({
              lang: {
                value: e.target.value,
                isInvalid: false,
                errorMsg: '',
              },
            });
          }}>
          {langs?.map((item) => {
            return (
              <option value={item.value} key={item.value}>
                {item.label}
              </option>
            );
          })}
        </Form.Select>
      </Form.Group>

      <div className="d-flex align-items-center justify-content-between">
        <Progress step={1} />
        <Button type="submit">{t('next')}</Button>
      </div>
    </Form>
  );
};

export default Index;
