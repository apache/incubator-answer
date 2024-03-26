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

import { FC } from 'react';
import { InputGroup } from 'react-bootstrap';

import type { FormKit, InputGroupOptions } from '../types';

import Button from './Button';

interface Props {
  formKitWithContext: FormKit;
  uiOpt: InputGroupOptions;
  prefixText?: string;
  suffixText?: string;
  children: React.ReactNode;
}

const InputGroupBtn = ({
  formKitWithContext,
  uiOpt,
}: {
  formKitWithContext: FormKit;
  uiOpt:
    | InputGroupOptions['prefixBtnOptions']
    | InputGroupOptions['suffixBtnOptions'];
}) => {
  return (
    <Button
      fieldName="1"
      text={String(uiOpt?.text)}
      iconName={uiOpt?.iconName ? uiOpt?.iconName : ''}
      action={uiOpt?.action ? uiOpt?.action : undefined}
      actionType="click"
      clickCallback={uiOpt?.clickCallback ? uiOpt?.clickCallback : undefined}
      formKit={formKitWithContext}
      variant={uiOpt?.variant ? uiOpt.variant : undefined}
      size={uiOpt?.size ? uiOpt?.size : undefined}
      title={uiOpt?.title ? uiOpt?.title : ''}
      nowrap
      readOnly={false}
    />
  );
};

const Index: FC<Props> = ({
  formKitWithContext,
  uiOpt,
  prefixText = null,
  suffixText = null,
  children,
}) => {
  return (
    <InputGroup>
      {prefixText && <InputGroup.Text>{prefixText}</InputGroup.Text>}
      {uiOpt && 'prefixBtnOptions' in uiOpt && (
        <InputGroupBtn
          uiOpt={uiOpt.prefixBtnOptions}
          formKitWithContext={formKitWithContext}
        />
      )}
      {children}
      {uiOpt && 'suffixBtnOptions' in uiOpt && (
        <InputGroupBtn
          uiOpt={uiOpt.suffixBtnOptions}
          formKitWithContext={formKitWithContext}
        />
      )}
      {suffixText ? <InputGroup.Text>{suffixText}</InputGroup.Text> : null}
    </InputGroup>
  );
};

export default Index;
