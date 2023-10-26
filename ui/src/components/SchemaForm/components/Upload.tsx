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
import { Form } from 'react-bootstrap';

import classNames from 'classnames';

import type * as Type from '@/common/interface';
import BrandUpload from '@/components/BrandUpload';

interface Props {
  type: Type.UploadType | undefined;
  acceptType: string | undefined;
  fieldName: string;
  onChange?: (fd: Type.FormDataType) => void;
  formData: Type.FormDataType;
  readOnly?: boolean;
  imgClassNames?: classNames.Argument;
}
const Index: FC<Props> = ({
  type = 'avatar',
  acceptType = '',
  fieldName,
  onChange,
  formData,
  readOnly = false,
  imgClassNames = '',
}) => {
  const fieldObject = formData[fieldName];
  const handleChange = (name: string, value: string) => {
    const state = {
      ...formData,
      [name]: {
        ...formData[name],
        value,
      },
    };
    if (typeof onChange === 'function') {
      onChange(state);
    }
  };
  return (
    <>
      <BrandUpload
        type={type}
        acceptType={acceptType}
        value={fieldObject?.value}
        readOnly={readOnly}
        onChange={(value) => handleChange(fieldName, value)}
        imgClassNames={imgClassNames}
      />
      <Form.Control
        name={fieldName}
        className="d-none"
        isInvalid={fieldObject?.isInvalid}
      />
    </>
  );
};

export default Index;
