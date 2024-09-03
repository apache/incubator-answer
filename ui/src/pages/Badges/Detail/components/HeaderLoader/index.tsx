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

import { Card } from 'react-bootstrap';

const Index = () => {
  return (
    <Card className="mb-4 placeholder-glow">
      <Card.Body className="d-block d-sm-flex">
        <div
          className="placeholder me-3 flex-shrink-0"
          style={{ width: '96px', height: '96px' }}
        />

        <div className="w-100 mt-3 mt-sm-0">
          <div className="placeholder h5 w-25" />
          <div className="placeholder w-100" />
          <div className="placeholder w-75" />

          <div className="placeholder mt-2 w-50" />

          <div className="small mt-2">
            <span className="placeholder" style={{ width: '80px' }} />

            <span className="placeholder ms-2" style={{ width: '80px' }} />
          </div>
        </div>
      </Card.Body>
    </Card>
  );
};

export default Index;
