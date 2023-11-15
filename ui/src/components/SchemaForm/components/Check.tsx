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

import React, { FC } from 'react';
import { Form, Stack } from 'react-bootstrap';

import type * as Type from '@/common/interface';

interface Props {
  type: 'radio' | 'checkbox';
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  enumValues: (string | boolean | number)[];
  enumNames: string[];
  formData: Type.FormDataType;
  readOnly?: boolean;
}
const Index: FC<Props> = ({
  type = 'radio',
  fieldName,
  onChange,
  enumValues,
  enumNames,
  formData,
  readOnly = false,
}) => {
  const fieldObject = formData[fieldName];
  const handleCheck = (
    evt: React.ChangeEvent<HTMLInputElement>,
    index: number,
  ) => {
    const { name, checked } = evt.currentTarget;
    enumValues[index] = checked;

    const state = {
      ...formData,
      [name]: {
        ...formData[name],
        value: enumValues,
        isInvalid: false,
      },
    };
    if (typeof onChange === 'function') {
      onChange(state);
    }
  };
  return (
    <Stack direction="horizontal">
      {enumValues?.map((item, index) => {
        return (
          <Form.Check
            key={String(item)}
            inline
            type={type}
            name={fieldName}
            id={`${fieldName}-${enumNames?.[index]}`}
            label={enumNames?.[index]}
            checked={fieldObject?.value?.[index] || false}
            feedback={fieldObject?.errorMsg}
            feedbackType="invalid"
            isInvalid={fieldObject?.isInvalid}
            disabled={readOnly}
            onChange={(evt) => handleCheck(evt, index)}
          />
        );
      })}
    </Stack>
  );
};

export default Index;
