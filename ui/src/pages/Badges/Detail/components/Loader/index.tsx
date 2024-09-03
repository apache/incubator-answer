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
import { Col } from 'react-bootstrap';

interface Props {
  count?: number;
}

const Index: FC<Props> = ({ count = 12 }) => {
  const list = new Array(count).fill(0).map((v, i) => v + i);
  return (
    <>
      {list.map((v) => (
        <Col sm={12} md={6} lg={3} key={v} className="mb-4 placeholder-glow">
          <div className="small mb-1 placeholder" style={{ width: '100px' }} />
          <div className="d-flex align-items-center">
            <div
              style={{ width: '40px', height: '40px' }}
              className="placeholder rounded flex-shrink-0"
            />
            <div className="small ms-2">
              <div className="placeholder lh-1" style={{ width: '80px' }} />
              <div
                className="text-secondary placeholder"
                style={{ width: '150px' }}
              />
            </div>
          </div>
          <div className="mt-1 d-block placeholder" />
          <div className="mt-1 d-block placeholder" />
        </Col>
      ))}
    </>
  );
};

export default memo(Index);
