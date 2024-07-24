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
import { ButtonGroup, Button } from 'react-bootstrap';

import classNames from 'classnames';

import { Icon, UploadImg } from '@/components';
import { UploadType } from '@/common/interface';

interface Props {
  type: UploadType;
  value: string;
  onChange: (value: string) => void;
  acceptType?: string;
  readOnly?: boolean;
  imgClassNames?: classNames.Argument;
}

const Index: FC<Props> = ({
  type = 'post',
  value,
  onChange,
  acceptType,
  readOnly = false,
  imgClassNames = '',
}) => {
  const onUpload = (imgPath: string) => {
    onChange(imgPath);
  };

  const onRemove = () => {
    onChange('');
  };
  return (
    <div className="d-flex">
      <div className="bg-gray-300 upload-img-wrap me-2 d-flex align-items-center justify-content-center">
        <img
          className={classNames(imgClassNames)}
          src={value}
          alt=""
          style={{ maxWidth: '100%', maxHeight: '100%' }}
        />
      </div>
      <ButtonGroup vertical className="fit-content">
        <UploadImg
          type={type}
          uploadCallback={onUpload}
          className="mb-0"
          disabled={readOnly}
          acceptType={acceptType}>
          <Icon name="cloud-upload" />
        </UploadImg>

        <Button
          disabled={readOnly}
          variant="outline-secondary"
          onClick={onRemove}>
          <Icon name="trash" />
        </Button>
      </ButtonGroup>
    </div>
  );
};

export default Index;
