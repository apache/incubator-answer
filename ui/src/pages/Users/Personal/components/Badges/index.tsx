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

import * as Type from '@/common/interface';
import { CardBadge } from '@/components';

interface IProps {
  data: Type.BadgeListItem[];
  username: string;
  visible: boolean;
}

const Index: FC<IProps> = ({ data, visible, username }) => {
  if (!visible) {
    return null;
  }
  return (
    <div className="d-flex flex-wrap" style={{ margin: '-12px' }}>
      {data.map((item) => {
        return (
          <CardBadge
            data={item}
            urlSearchParams={`username=${username}`}
            badgePillType="count"
          />
        );
      })}
    </div>
  );
};

export default Index;
