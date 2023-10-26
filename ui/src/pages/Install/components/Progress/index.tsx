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

import { FC, memo } from 'react';
import { ProgressBar } from 'react-bootstrap';

interface IProps {
  step: number;
}

const Index: FC<IProps> = ({ step }) => {
  return (
    <div className="d-flex align-items-center small text-secondary">
      <ProgressBar
        now={(step / 5) * 100}
        variant="success"
        style={{ width: '200px' }}
        className="me-2"
      />
      <span>{step}/5</span>
    </div>
  );
};

export default memo(Index);
